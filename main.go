package main

import (
	"fmt"
	"go/build"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Use(apiGateway)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/hello", func(c echo.Context) error {
		proxyHeaders := make(map[string]string)

		for k, v := range c.Request().Header {
			proxyHeaders[k] = v[0]
		}

		requestBody := map[string]string{}
		requestBody["name"] = c.FormValue("BODY")

		return c.JSON(http.StatusOK, map[string]string{
			"data": "SIP TENAN",
		})
		// return c.JSON(http.StatusOK, events.APIGatewayProxyResponse{
		// 	StatusCode:      http.StatusOK,
		// 	Headers:         proxyHeaders,
		// 	Body:            map[string],
		// 	IsBase64Encoded: false,
		// })
	})

	e.POST("/hello/:id", func(c echo.Context) error {
		proxyHeaders := make(map[string]string)

		for k, v := range c.Request().Header {
			proxyHeaders[k] = v[0]
		}

		body := map[string]string{}
		body["name"] = c.FormValue("name")
		body["id"] = c.Param("id")
		body["limit"] = c.QueryParam("limit")

		return c.JSON(http.StatusOK, map[string]map[string]string{
			"data": body,
		})
		// return c.JSON(http.StatusOK, events.APIGatewayProxyResponse{
		// 	StatusCode:      http.StatusOK,
		// 	Headers:         proxyHeaders,
		// 	Body:            map[string],
		// 	IsBase64Encoded: false,
		// })
	})

	e.GET("/photo", func(c echo.Context) error {
		return c.File("image1.jpg")
	})

	e.POST("/photo", func(c echo.Context) error {
		file, err := c.FormFile("image")

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// // Destination
		// ex, err := os.Executable()
		// if err != nil {
		// 	panic(err)
		// }
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}
		dst, err := os.Create(gopath + "/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.File(gopath + "/" + file.Filename)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func apiGateway(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// apigw := events.APIGatewayProxyRequest{}
		// if err := c.Bind(apigw); err != nil {
		// 	c.Error(err)
		// }

		fmt.Printf("%+v\n")
		return next(c)
	}
}
