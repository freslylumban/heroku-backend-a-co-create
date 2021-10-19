package controller

import (
	"net/http"

	"heroku-backend-a-cocreate/helper/header"
	"heroku-backend-a-cocreate/helper/response"
	"heroku-backend-a-cocreate/model"
	"heroku-backend-a-cocreate/service"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func TestRouter(c *gin.Context) {
	response.BuildResponse(c, http.StatusOK, "Welcome to API Banking Innovation Co-Create", "API Test OK!")
}

func InjectFirstRole(c *gin.Context) {
	var roles = []model.Role{{Role: "admin"}, {Role: "user"}}

	contentType := header.GetContentType(c)
	var errBind error
	if contentType == appJSON {
		errBind = c.ShouldBindJSON(&roles)
	} else {
		errBind = c.ShouldBind(&roles)
	}

	if errBind != nil {
		response.BuildErrResponse(c, http.StatusBadRequest, "Failed to process request", errBind.Error())
		return
	}

	anyRole := service.FindRoleForInject()

	for _, all := range roles {
		if all.Role == anyRole.Role {
			response.BuildErrResponse(c, http.StatusUnprocessableEntity, "Failed to process request", "Role has registered. This endpoint has done.")
			return
		}
	}

	createdFirstRoles := service.CreateRoles(roles)
	response.BuildResponse(c, http.StatusCreated, "Create First Roles OK!", createdFirstRoles)
}
