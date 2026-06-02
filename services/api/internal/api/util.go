package api

import (
	"MockExamLab/internal/models"
	"github.com/gin-gonic/gin"
)

func getUserClaimsFromSession(c *gin.Context) *models.UserClaims {
	currentUser, isExist := c.Get("user")
	if !isExist {
		return nil
	}
	return currentUser.(*models.UserClaims)
}

func getQueryParameters(c *gin.Context) *models.QueryParams {
	var params models.QueryParams
	err := c.BindQuery(&params)
	if err != nil {
		return nil
	}
	return &params
}
