package main

import (
	"os"
	"path/filepath"
	pages "templ-echo-test/views/pages"

	"github.com/labstack/echo/v4"
	"github.com/russross/blackfriday/v2"
)

func ArticlePage(c echo.Context) error {
	articleName := c.Param("article")

	// Load the Markdown file
	markdownContent, err := os.ReadFile(filepath.Join("articles", articleName+".md"))
	if err != nil {
		return err
	}

	// Convert Markdown to HTML using blackfriday
	htmlContent := blackfriday.Run(markdownContent)

	return render(c, pages.Article(string(htmlContent)))
}
