package repositories

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"../../domains/repositories"
	"../../services"
	"../../utils/errors"
)

func CreateRepo(ctx *gin.Context) {
	var req repositories.CreateRepoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apiErr := errors.NewBadRequestApiError("invalid json body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(req)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func CreateRepos(ctx *gin.Context) {
	var req []repositories.CreateRepoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		apiErr := errors.NewBadRequestApiError("invalid json body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(req)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
