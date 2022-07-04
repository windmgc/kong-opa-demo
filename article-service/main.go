package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type article struct {
	ArticleID      string `json:"article_id"`
	ArticleTitle   string `json:"article_title"`
	ArticleContent string `json:"article_content"`
	UserID         string `json:"user_id"`
	PostedDate     string `json:"posted_date"`
}

// var users = make(map(string, User))
var articles = map[string]article{
	"1": {
		ArticleID:      "1",
		ArticleTitle:   "Article 1",
		ArticleContent: "Article 1 content",
		UserID:         "1",
		PostedDate:     "2020-01-01",
	},
	"2": {
		ArticleID:      "2",
		ArticleTitle:   "Article 2",
		ArticleContent: "Article 2 content",
		UserID:         "1",
		PostedDate:     "2020-01-02",
	},
	"3": {
		ArticleID:      "3",
		ArticleTitle:   "Article 3",
		ArticleContent: "Article 3 content",
		UserID:         "2",
		PostedDate:     "2020-01-03",
	},
}

func main() {
	app := fiber.New()

	app.Post("/articles", func(c *fiber.Ctx) error {
		// Dummy add article
		nextArticleID := fmt.Sprint(len(articles) + 1)
		articles[nextArticleID] = article{
			ArticleID:      nextArticleID,
			ArticleTitle:   c.FormValue("article_title"),
			ArticleContent: c.FormValue("article_content"),
			UserID:         c.FormValue("user_id"),
			PostedDate:     c.FormValue("posted_date"),
		}
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Get("/articles/:article_id", func(c *fiber.Ctx) error {
		articleID := c.Params("article_id")
		return c.JSON(articles[articleID])
	})

	app.Listen(":8082")
}
