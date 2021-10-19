package bind

import (
	"heroku-backend-a-cocreate/helper/header"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func SBJSON(c *gin.Context, model interface{}) (interface{}, error) {
	contentType := header.GetContentType(c)

	if contentType != appJSON {
		err := c.ShouldBindJSON(&model)
		return &model, err
	} else {
		err := c.ShouldBind(&model)
		return &model, err
	}
}
