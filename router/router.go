package router

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Set(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}

	// tag := e.Group("/tag")
	// tag.POST("", controller.Creat)
	// tag.POST("/user", controller.Creat)
	// tag.POST("/product", controller.Insert)
	// tag.GET("", controller.Get)
	// tag.GET("/:id", controller.GetOne)
	// tag.PUT("/:id", controller.Update)
	// tag.DELETE("/:id", controller.Delete)
	// tag.GET("/search", controller.Search)
}
