package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber"
	"github.com/igudgz/desafioMoneri/src/config"
	"github.com/igudgz/desafioMoneri/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBook(c *fiber.Ctx) {
	collection, err := config.GetMongoDbCollection("library", "books")
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)

		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)
}
func GetBookAuthor(c *fiber.Ctx) {
	collection, err := config.GetMongoDbCollection("library", "books")
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	var filter bson.M = bson.M{}
	if c.Params("author") != "" {
		author := c.Params("author")
		filter = bson.M{"author": author}
	}
	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)
}

func GetBookTitle(c *fiber.Ctx) {
	collection, err := config.GetMongoDbCollection("library", "books")
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	var filter bson.M = bson.M{}
	if c.Params("title") != "" {
		title := c.Params("title")

		filter = bson.M{"title": title}
	}
	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)

}
func CreateBook(c *fiber.Ctx) {
	collection, err := config.GetMongoDbCollection("library", "books")
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}
	var book models.Book
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	book.Create_at = time.Now().In(loc)
	book.Update_at = time.Now().In(loc)
	json.Unmarshal([]byte(c.Body()), &book)

	res, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}
func UpdateBook(c *fiber.Ctx) {
	collection, err := config.GetMongoDbCollection("library", "books")
	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}
	var book models.Book
	book.Update_at = time.Now()
	json.Unmarshal([]byte(c.Body()), &book)

	update := bson.M{
		"$set": book,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}
func DeleteBook(c *fiber.Ctx) {
	collection, err := config.GetMongoDbCollection("library", "books")

	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500).Send([]byte(err.Error()))
		return
	}

	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)
}
