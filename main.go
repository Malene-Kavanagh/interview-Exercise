package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Intro struct {
	Message string `json:"message"`
	Time    int64  `json:"timestamp"`
} //it works
// demo4
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // default port 80 if not specified
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		//current time in unix timestamp
		p := time.Now().UTC()
		t := p.UnixMilli()
		//works fine

		intro := Intro{
			Message: "My name is Malene Kavanagh",
			Time:    t,
		}

		fine, err := json.Marshal(intro)
		if err != nil {
			panic(err)
		}
		var minibuf bytes.Buffer

		err2 := json.Compact(&minibuf, fine)
		if err2 != nil {
			panic(err2)
		}
		//set content type to application/json
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		//return c.JSON(intro)
		return c.Send(minibuf.Bytes())
	})

	app.Get("/port", func(c *fiber.Ctx) error {
		if p := c.Get("Forwarded-Port"); p != "" {
			return c.SendString(p)
		}
		if p := extractPort(c.Get("Forwarded-Host")); p != "" {
			return c.SendString(p)
		}
		return c.SendString(extractPort(string(c.Context().Request.Header.Peek("Host"))))
	})
	log.Fatal(app.Listen(":" + port))
}

func extractPort(hostport string) string {
	//takes a string of host:port and returns port
	if hostport == "" || !strings.Contains(hostport, ":") {
		return ""
	} // no port empty string return
	if _, p, err := net.SplitHostPort(hostport); err == nil && p != "" {
		return p
	} //splits host and port
	// Some proxies may send multiple hosts, take the first
	if idx := strings.Index(hostport, ","); idx > 0 {
		hostport = strings.TrimSpace(hostport[:idx])
		// underscore ignores error
		if _, p, err := net.SplitHostPort(hostport); err == nil && p != "" {
			return p
		}
	}
	return ""
	//if none worked return empty string
}
