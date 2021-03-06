package controller

import (
	"heroku-backend-a-cocreate/dto"
	"heroku-backend-a-cocreate/helper/response"
	"heroku-backend-a-cocreate/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateArticle(c *gin.Context) {
	var articleRequest *dto.Article
	if err := c.ShouldBindJSON(&articleRequest); err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}

	article, err := service.CreateArticle(articleRequest)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusCreated, "Success", article)
	return
}

func DeleteArticle(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}

	err = service.DeleteArticle(id)
	result := map[string]int{
		"id_artikel_deleted": id,
	}

	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Success", result)
	return
}

func GetArticleById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", err.Error())
		return
	}

	article, err := service.GetArticleById(id)
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Success", article)
	return
}

func GetAllArticle(c *gin.Context) {
	articles, err := service.GetAllArticle()
	if err != nil {
		log.Error(err)
		response.BuildErrResponse(c, http.StatusInternalServerError, "Failed to process request", err.Error())
		return
	}

	response.BuildResponse(c, http.StatusOK, "Success", articles)
	return
}
