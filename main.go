package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})

	app.Get("/", healthCheck)

	app.Post("/question", addQuestion)

	app.Post("/pdf", addPDF)

	log.Fatal(app.Listen(":3000"))
}

func healthCheck(c *fiber.Ctx) error {
	log.Println("Health Check Performed")
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}

func addQuestion(c *fiber.Ctx) error {
	log.Println("Question Performed")

	body := c.Body()

	var question Question
	err := json.Unmarshal(body, &question)

	if err != nil {
		log.Println(err)
	}

	return c.JSON(fiber.Map{
		"body": &question,
	})
}

func addPDF(c *fiber.Ctx) error {

	log.Println("PDF Performed")

	body := c.Body()

	var customer Customer
	err := json.Unmarshal(body, &customer)

	fileName := customer.FirstName + "_" + customer.LastName + "_" + customer.TransactionId

	generatePDF("pdfs", fileName)

	if err != nil {
		log.Println(err)
	}

	return c.JSON(fiber.Map{
		"body":    &customer,
		"message": "PDF Generated with the filename: " + fileName + ".pdf",
	})

}
