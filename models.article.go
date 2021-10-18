// models.article.go
// Initializes application's article model

package main

import (
	"errors"
)

type article struct {
	ID      int    `json:"id"`     // ID number of the article
	Title   string `json:"title"`  // Title of the article
	Content string `json:content"` // Content of the article
}

// This section would normally taken up by initializing and connecting to Local Database
var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 Content"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 Content"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func createNewArticle(title, content string) (*article, error) {
	a := article{
		ID:      len(articleList) + 1,
		Title:   title,
		Content: content,
	}

	articleList = append(articleList, a)
	return &a, nil
}
