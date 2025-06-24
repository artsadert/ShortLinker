package rest

import (
	"net/http"

	"github.com/artsadert/ShortLinker/internal/application/interfaces"
	"github.com/artsadert/ShortLinker/internal/interface/api/rest/dto"
	"github.com/gin-gonic/gin"
)

type LinkController struct {
	service interfaces.LinkService
}

func NewLinkController(e *gin.Engine, service interfaces.LinkService) *LinkController {
	controller := &LinkController{service: service}

	e.GET("/get_url", controller.GetLink)
	e.POST("/create_url", controller.CreateNewLink)

	return controller
}

func (l *LinkController) CreateNewLink(c *gin.Context) {
	//function to create short url from long
	var link dto.GetLinkRequest

	if err := c.ShouldBindBodyWithJSON(&link); err != nil {
		c.String(http.StatusBadRequest, `body must be as link request`)
	}
}

func (l *LinkController) GetLink(c *gin.Context) {
	// function to get long url from short

	var link dto.CreateLinkRequest

	if err := c.ShouldBindBodyWithJSON(&link); err != nil {
		c.String(http.StatusBadRequest, `body must be as link request`)
	}

}
