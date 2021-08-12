package main

import (
	"context"
	"fmt"
	"imo-server/db"
	"imo-server/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
  app := fiber.New()

  fmt.Println("Trying to connect db")
  db.InitDatabase()

  app.Get("/test", func (c *fiber.Ctx) error {
    collection := db.Collection("shows")

    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
      return c.Status(500).SendString(fmt.Sprintf("%v", err)) 
    }

    var shows []models.Show
    if err = cursor.All(context.Background(), &shows); err != nil {
      return c.Status(500).SendString(fmt.Sprintf("%v", err)) 
    }

    return c.JSON(shows)
  })

  app.Get("/hello-world", func (c *fiber.Ctx) error {
    return c.SendString("Hello World!")
  })

  fmt.Println("Listening on port 9990")
  app.Listen(":9990")
}