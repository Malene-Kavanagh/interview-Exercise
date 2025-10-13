package main

import (
	"encoding/json"
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
		//Unix allows it to be computer time instead of human readable
		t := time.Now().UTC().Unix()

		intro := Intro{
			Message: "My name is Malene Kavanagh",
			Time:    t,
		}

		fine, err := json.Marshal(intro)
		if err != nil {
			panic(err)
		}
		// TEST print
		//fmt.Println(string(fine))
		// to show what json.Marshal is sending
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		//Sends the encoded json to the browser
		return c.Send(fine)
	})
	//FIX: Needs to be sent to the cloud browser
	//     Docker needs to be implemented with server
	app.Listen(":3000")
}
