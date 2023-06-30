package main

import (
	"fmt"
	"main/controler"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	// Routes
	e.GET("/", controler.GetAllItem)
	e.GET("/:id", controler.GetDetailItem)
	e.POST("/", controler.PostNewItem)
	e.PUT("/:id", controler.ChangeItem)
	e.DELETE("/:id", controler.DeleteItem)
	fmt.Println("something")
	// Start server at localhost:1323
	e.Logger.Fatal(e.Start(":1232"))
}
