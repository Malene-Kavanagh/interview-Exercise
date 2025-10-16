package main

import (
	"encoding/json"
	"net"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Intro struct {
	Message string `json:"message"`
	Time    int64  `json:"timestamp"`
}

// func Convert(t time.Time) string {
//     return t.UTC().Format(time.RFC3339)
// }

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//app := fiber.New()  // default config
	app := fiber.New(fiber.Config{})
	app.Get("/", func(c *fiber.Ctx) error {
		//current time
		//t := time.Now()
		t := time.Now().UTC().Unix()

		intro := Intro{
			Message: "My name is Malene Kavanagh",
			Time:    t,
		}

		fine, err := json.Marshal(intro)
		if err != nil {
			panic(err)
		}

		// to show what json.Marshal is sending
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		//c.SendString(intro)
		//return c.JSON(fine)
		return c.SendString(string(fine)) //error check
		//return c.Send(fine)
		//^  this gives garbled output
		//^accidentally made json.Marshal useless
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
	//log.Fatal(app.Listen(":" + port))
	//port is hardcoded to be 80
	//log.Fatal(app.Listen(":80"))
	app.Listen(":80")
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
