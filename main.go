package main

import (
	"encoding/json"
	"fmt"
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
	app := fiber.New()
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
		fmt.Println(string(fine))
		// to show what json.Marshal is sending
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		//c.SendString(intro)
		//return c.JSON(fine)
		return c.Send(fine)
		//^  this gives garbled output
		//^accidentally made json.Marshal useless
	})

	//app.Listen(":3000")
	app.Listen(":8080") // for cloud run
}
