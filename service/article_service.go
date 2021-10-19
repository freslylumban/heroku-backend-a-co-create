package service

import (
	"heroku-backend-a-cocreate/dto"
	"heroku-backend-a-cocreate/model"
	"heroku-backend-a-cocreate/repository"

	log "github.com/sirupsen/logrus"
)

type IArticleService interface {
	CreateArticle(article *dto.Article) (*model.Article, error)
	DeleteArticle(idArticle int) error
	GetArticleById(idArticle int) (*model.Article, error)
	GetAllArticle() ([]*model.Article, error)
}

func CreateArticle(article *dto.Article) (*model.Article, error) {
	articleToCreate, err := repository.CreateArticle(article)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return articleToCreate, nil
}

func DeleteArticle(idArticle int) error {
	err := repository.DeleteArticle(idArticle)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GetArticleById(idArticle int) (*model.Article, error) {
	article, err := repository.FindArticleById(idArticle)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return article, nil
}

func GetAllArticle() ([]*model.Article, error) {
	articles, err := repository.FindAllArticle()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return articles, nil
}
