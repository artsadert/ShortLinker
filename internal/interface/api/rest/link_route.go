package rest

import (
	"fmt"
	"net/http"

	"github.com/artsadert/ShortLinker/internal/application/interfaces"
	"github.com/artsadert/ShortLinker/internal/interface/api/rest/dto/mapper"
	link_request "github.com/artsadert/ShortLinker/internal/interface/api/rest/dto/request"
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
	var link link_request.CreateLinkRequest

	if err := c.ShouldBindBodyWithJSON(&link); err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, `body must be as link request`)
		return
	}

	res, err := l.service.CreateNewLink(link_request.ToCreateLinkCommand(&link))
	if err != nil {
		fmt.Println(err)
		c.String(500, "Server error occupied")
		return
	}

	c.JSON(200, mapper.ToCreateLinkResponse(res))
}

func (l *LinkController) GetLink(c *gin.Context) {
	// function to get long url from short

	var link link_request.GetLinkRequest

	if err := c.ShouldBindBodyWithJSON(&link); err != nil {
		fmt.Println(err)
		c.String(500, `body must be as link request`)
		return
	}

	res, err := l.service.GetLink(link_request.ToGetLinkCommand(&link))
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, "Server error occupied")
		return
	}

	c.JSON(200, mapper.ToGetLinkResponse(res))
}
