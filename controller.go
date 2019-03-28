package web_tracking

import (
	"github.com/joaosoft/web"
)

type Controller struct {
	service *WebTracking
}

func NewController(service *WebTracking) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) DummyHandler(ctx *web.Context) error {
	return ctx.Response.NoContent(web.StatusNoContent)
}
