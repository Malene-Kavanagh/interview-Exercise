package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Intro struct {
	Message string    `json:"message"`
	Time    time.Time `json:"timestamp"`
}

// func Convert(t time.Time) string {
//     return t.UTC().Format(time.RFC3339)
// }

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		//current time
		//t := time.Now()
		t := time.Now().UTC()
		// t := Convert(time.Now())

		// func Convert(t time.Time) UTC
		//format time to JSON
		//timeStamp := t.MarshalJSON()
		intro := Intro{
			Message: "My name is Malene Kavanagh",
			Time:    t,
		}

		fmt.Println(intro)
		// test output to show the info being sent
		fine, err := json.Marshal(intro)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(fine))
		// to show what json.Marshal is sending
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		//c.SendString(intro)
		//return c.JSON(fine)
		return c.JSON(intro)
	})

	app.Listen(":3000")
}
