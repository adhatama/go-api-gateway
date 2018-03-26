package main

import (
	"fmt"
	"net/http"

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
