package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", Page)

	e.GET("/articles/:article", ArticlePage)
	e.Logger.Fatal(e.Start(":1323"))
}
