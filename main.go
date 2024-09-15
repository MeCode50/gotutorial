package main 

import ("fmt"
		"log"
 "github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello world")
	app := fiber.Nwe()
	log.Fatal(app.Listen(":4000"))
}