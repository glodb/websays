package validators

import (
	"errors"
	"websays/app/models"
)

type ArticleValidator struct {
}

func (art *ArticleValidator) Validate(apiName string, data interface{}) error {
	articleData := data.(models.Article)
	switch apiName {
	case "/api/addArticle":
		if articleData.Title == "" {
			return errors.New("Title can't be empty")
		}

		if articleData.Body == "" {
			return errors.New("Body can't be empty")
		}
	case "/api/upateArticle":
		if articleData.ID <= 0 {
			return errors.New("ID is not valid")
		}

		if articleData.Title == "" {
			return errors.New("Title can't be empty")
		}

		if articleData.Body == "" {
			return errors.New("Body can't be empty")
		}
	}

	return nil
}
