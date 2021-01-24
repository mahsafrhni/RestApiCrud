package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)
import "First/Handler"

func main() {
	e := echo.New()
	e.POST("/customers", Handler.Create)
	e.PUT("/customers/:id", Handler.Edit)
	e.DELETE("/customers/:id", Handler.Delete)
	e.GET("/report/:month", Handler.Report)
	e.GET("/customers", Handler.Information)
	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}
}
