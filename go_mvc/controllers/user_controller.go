package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"../services"
	"../utils"
)

func  GetUser(c *gin.Context) {
	userIdParam := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.ResponseError(c, appErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.ResponseError(c, apiErr)
		return
	}
	// return user
	utils.Response(c, http.StatusOK, user)
}
