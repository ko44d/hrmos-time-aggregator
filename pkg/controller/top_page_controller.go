package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TopPageController interface {
	ShowForm(ctx *gin.Context)
	SetAPIKey(ctx *gin.Context)
}

type topPageController struct {
}

func NewTopPageController() TopPageController {
	return &topPageController{}
}

func (c *topPageController) ShowForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (c *topPageController) SetAPIKey(ctx *gin.Context) {
	apiKey := ctx.PostForm("api_key")
	ctx.SetCookie("api_key", apiKey, 3600, "/", "localhost", false, true)
	ctx.Redirect(http.StatusSeeOther, "/api/v1/aggregate")
}
