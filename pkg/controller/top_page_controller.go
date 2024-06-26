package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ko44d/hrmos-time-aggregator/pkg/usecase"
	"net/http"
)

type TopPageController interface {
	ShowForm(ctx *gin.Context)
	SetAPIKey(ctx *gin.Context)
}

type topPageController struct {
	atu usecase.AuthenticationTokenUsecase
}

func NewTopPageController(atu usecase.AuthenticationTokenUsecase) TopPageController {
	return &topPageController{atu: atu}
}

func (c *topPageController) ShowForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func (c *topPageController) SetAPIKey(ctx *gin.Context) {
	apiKey := ctx.PostForm("api_key")
	companyURL := ctx.PostForm("company_url")
	res, err := c.atu.GetToken(apiKey, companyURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("token", res.Token, 3600, "/", "localhost", false, true)
	ctx.SetCookie("company_url", companyURL, 3600, "/", "localhost", false, true)
	ctx.Redirect(http.StatusSeeOther, "/aggregate")
}
