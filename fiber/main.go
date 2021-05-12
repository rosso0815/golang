// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()

// 	app.Prefork = true // enable prefork

// 	app.Get("/", func(c *fiber.Ctx) {
// 		c.Send(fmt.Sprintf("Hi, I'm worker #%v", os.Getpid()))
// 		// => Hi, I'm worker #16858
// 		// => Hi, I'm worker #16877
// 		// => Hi, I'm worker #16895
// 	})

// 	app.Listen(8080)
// }

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Get("/", hello)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
