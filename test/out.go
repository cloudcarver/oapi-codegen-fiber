package apigen

import "github.com/gofiber/fiber/v2"

type Validator interface { 
    OwnCluster(c *fiber.Ctx, userId int32, clusterId int32) error
 
    GetUserID(c *fiber.Ctx, ) int32

}


type XMiddleware struct {
	Handler ServerInterface
	Validator
}

func NewXMiddleware(handler ServerInterface, validator Validator) ServerInterface {
	return &XMiddleware{Handler: handler, Validator: validator}
}


// (GET /clusters/{clusterId})
func (x *XMiddleware) GetUser(c *fiber.Ctx, clusterId int32) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	c.Locals("operationID", "GetUser")
	 
	if err := x.OwnCluster(c, x.GetUserID(c), clusterId); err != nil {
	    return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}  
    return x.Handler.GetUser(c, clusterId)
}

// (GET /test0)
func (x *XMiddleware) Test0(c *fiber.Ctx) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	c.Locals("operationID", "Test0")
	  
    return x.Handler.Test0(c)
}

// (POST /test0)
func (x *XMiddleware) Test0Post(c *fiber.Ctx) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	c.Locals("operationID", "Test0Post")
	    
    return x.Handler.Test0Post(c)
}

// (GET /test2)
func (x *XMiddleware) GetTest2(c *fiber.Ctx) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	c.Locals("operationID", "GetTest2")
	  
    return x.Handler.GetTest2(c)
}

// (GET /test3)
func (x *XMiddleware) GetTest3(c *fiber.Ctx) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	c.Locals("operationID", "GetTest3")
	  
    return x.Handler.GetTest3(c)
}

// (GET /user/{id})
func (x *XMiddleware) GetUserId(c *fiber.Ctx, id string) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	c.Locals("operationID", "GetUserId")
	  
    return x.Handler.GetUserId(c, id)
}

