package main

import (
	"context"
	"fmt"
	"os"

	"github.com/222dis/gauth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT")
	mongodb := os.Getenv("MONGODB_URI")

	if port == "" {
		port = "3000"
	}

	if mongodb == "" { //https://www.mongodb.com/docs/drivers/go/current/quick-start/
		mongodb = "mongodb://root:root@localhost:27018/"
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Static("/", "./public")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodb))
	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error { //app.Get("/:value?", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User
		coll := client.Database("food").Collection("fruits")

		results, error := coll.Find(context.TODO(), bson.M{})
		if error != nil {
			panic(error)
		}

		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"users": users,
		})

	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User

		c.BodyParser(&user)

		coll := client.Database("food").Collection("fruits")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
		}})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	app.Listen(":" + port)
	fmt.Println("Server on port 3000")

}
