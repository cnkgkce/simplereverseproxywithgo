package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Proxy interface{
	Accept(key string) bool
	Proxy(c *fiber.Ctx) error
}

var proxies = []Proxy{
	NewLimitProxy("cenk",3,3*time.Minute),
}
  
func ProxyHandler(c *fiber.Ctx) error{

	for _,value := range proxies{
		if value.Accept(c.Params("key")){
			return value.Proxy(c)
		}
	}

	c.Response().SetStatusCode(404)
	return nil
	
}