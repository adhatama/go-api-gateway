package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

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

		// Request
		reqBody := []byte{}
		if c.Request().Body != nil { // Read
			reqBody, _ = ioutil.ReadAll(c.Request().Body)
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset

		return c.JSON(http.StatusOK, events.APIGatewayProxyResponse{
			StatusCode:      http.StatusOK,
			Headers:         proxyHeaders,
			Body:            "HELLO WORLD",
			IsBase64Encoded: false,
		})
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
