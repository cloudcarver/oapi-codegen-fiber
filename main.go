package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"net/url"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

var (
	specPath    string
	outPath     string
	packageName string
	version     bool
)

func main() {
	flag.StringVar(&specPath, "path", "", "Path to the OpenAPI spec file")
	flag.StringVar(&outPath, "out", "", "Path to the output file")
	flag.StringVar(&packageName, "package", "main", "Package name")
	flag.BoolVar(&version, "version", false, "show version")
	flag.Parse()

	if version {
		fmt.Println("v0.3.0")
		return
	}

	if specPath == "" {
		panic("specPath is required")
	}

	if specPath == "" {
		panic("out is required")
	}

	if len(packageName) == 0 {
		packageName = "fiberx"
	}

	doc, err := getSchema(specPath)
	if err != nil {
		panic(err)
	}

	middlewareCode, err := generateMiddleware(&SecurityTmplVar{
		PackageName:          packageName,
		SecurityRequirements: parseSecurityRequirements(doc),
	})
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(outPath, []byte(middlewareCode), 0644); err != nil {
		panic(err)
	}
}

func getSchema(specPath string) (*openapi3.T, error) {
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile(specPath)
	if err != nil {
		return nil, err
	}
	if err := loader.ResolveRefsIn(doc, &url.URL{}); err != nil {
		return nil, err
	}
	return doc, nil
}

type SecurityTmplVar struct {
	PackageName          string
	SecurityRequirements []SecurityRequirement
}

type SecurityRequirement struct {
	Path   string
	Method string
	Rules  map[string][]string
}

func parseSecurityRequirements(doc *openapi3.T) []SecurityRequirement {
	var baseURL string
	if len(doc.Servers) == 0 {
		baseURL = ""
	} else {
		baseURL = doc.Servers[0].URL
	}

	var rtn []SecurityRequirement
	for path, pathItem := range doc.Paths.Map() {
		for method, operation := range pathItem.Operations() {
			if operation.Security == nil {
				continue
			}
			rq := SecurityRequirement{
				Path:   replaceURLParameters(baseURL + path),
				Method: toCamelCase(method),
				Rules:  map[string][]string{},
			}
			for _, se := range *operation.Security {
				for typ, rules := range se {
					rq.Rules[typ] = rules
				}
			}
			rtn = append(rtn, rq)
		}
	}
	return rtn
}

const securityMiddlewareTemplate = `package {{.PackageName}} 

import "github.com/gofiber/fiber/v2"

type AuthFunc func(c *fiber.Ctx, rules ...string) error

func RegisterAuthFunc(app *fiber.App, f AuthFunc) {
	{{range .SecurityRequirements}}
	app.{{.Method}}("{{.Path}}", func(c *fiber.Ctx) error { {{range $key, $value := .Rules}}{{if eq $key "BearerAuth"}}
		if c.Get("Authorization") == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		} {{if eq (len $value) 0}}
		if err := f(c); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		{{else}}
		rules := []string{
			{{range $value}}"{{.}}", {{end}}
		}
		if err := f(c, rules...); err != nil {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}{{end}}{{end}}{{end}}
		return c.Next()
	}){{end}}
}
`

func generateMiddleware(v *SecurityTmplVar) (string, error) {
	tmpl, err := template.New("securityMiddleware").Parse(securityMiddlewareTemplate)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, v); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func toCamelCase(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
}

func replaceURLParameters(path string) string {
	return strings.ReplaceAll(strings.ReplaceAll(path, "{", ":"), "}", "")
}
