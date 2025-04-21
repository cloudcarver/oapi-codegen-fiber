package apigen

import "github.com/gofiber/fiber/v2"

type Validator interface { 
    PreValidate(*fiber.Ctx) error
    
    PostValidate(*fiber.Ctx) error

    CheckOperationID(c *fiber.Ctx, operationID string) error

    OwnCluster(c *fiber.Ctx, userId int32, clusterId int32) error
 
    GetUserID(c *fiber.Ctx) int32
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
	if err := x.PreValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	operationID := "GetUser"
	 
	if err := x.CheckOperationID(c, operationID); err != nil {
	    return c.Status(fiber.StatusForbidden).SendString(err.Error())
	} 
	if err := x.OwnCluster(c, x.GetUserID(c), clusterId); err != nil {
	    return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}  
	if err := x.PostValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
    return x.Handler.GetUser(c, clusterId)
}

// (POST /sign-in)
func (x *XMiddleware) SignIn(c *fiber.Ctx) error {
    
	if err := x.PreValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	
	
	if err := x.PostValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
    return x.Handler.SignIn(c)
}

// (GET /test0)
func (x *XMiddleware) Test0(c *fiber.Ctx) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	if err := x.PreValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	
	  
	if err := x.PostValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
    return x.Handler.Test0(c)
}

// (POST /test0)
func (x *XMiddleware) Test0Post(c *fiber.Ctx) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	if err := x.PreValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	
	    
	if err := x.PostValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
    return x.Handler.Test0Post(c)
}

// (GET /test2)
func (x *XMiddleware) GetTest2(c *fiber.Ctx) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	if err := x.PreValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	
	  
	if err := x.PostValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
    return x.Handler.GetTest2(c)
}

// (GET /test3)
func (x *XMiddleware) GetTest3(c *fiber.Ctx) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	if err := x.PreValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	
	  
	if err := x.PostValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
    return x.Handler.GetTest3(c)
}

// (GET /user/{id})
func (x *XMiddleware) GetUserId(c *fiber.Ctx, id string) error {
    if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header is required")
	} 
	if err := x.PreValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
	operationID := "GetUserId"
	 
	if err := x.CheckOperationID(c, operationID); err != nil {
	    return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}  
	if err := x.PostValidate(c); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}
    return x.Handler.GetUserId(c, id)
}

