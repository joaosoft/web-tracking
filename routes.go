package web_tracking

import (
	"encoding/json"

	"github.com/joaosoft/manager"
	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

func (c *Controller) RegisterRoutes(w manager.IWeb) error {
	if err := w.AddRoutes(
		manager.NewRoute(string(web.MethodPost), "/api/v1/dummy", c.DummyHandler),
	); err != nil {
		return err
	}

	w.AddFilter("*", string(web.PositionBefore), c.MiddlewareTracking(), string(web.MethodAny))

	return nil
}

func (c *Controller) MiddlewareTracking() web.MiddlewareFunc {
	return func(next web.HandlerFunc) web.HandlerFunc {
		return func(ctx *web.Context) error {

			if ctx.Request.GetParam(trackingParam) == trackingEnabled {
				go c.handleTracking(ctx)
			}

			return next(ctx)
		}
	}
}

func (c *Controller) handleTracking(ctx *web.Context) {
	trackingRequest := TrackingRequest{}

	// parameters on url
	err := ctx.Request.BindParams(&trackingRequest.Tracking)
	if err != nil {
		c.service.logger.Error(ErrorHandlingTracking.Format(err.Error()))
		return
	}

	// override parameters on body
	err = ctx.Request.Bind(&trackingRequest)
	if err != nil {
		c.service.logger.Error(ErrorHandlingTracking.Format(err.Error()))
		return
	}

	if errs := validator.Validate(trackingRequest); len(errs) > 0 {
		c.service.logger.Error(ErrorHandlingTracking.Format(err.Error()))
		return
	}

	// handle tracking
	request, err := c.service.webClient.NewRequest(web.MethodPost, c.service.config.TrackingHost, web.ContentTypeApplicationJSON, nil)
	if err != nil {
		c.service.logger.Error(ErrorHandlingTracking.Format(err.Error()))
		return
	}

	bytes, err := json.Marshal(trackingRequest.Tracking)
	if err != nil {
		c.service.logger.Error(ErrorHandlingTracking.Format(err.Error()))
		return
	}

	response, err := request.WithBody(bytes).Send()
	if err != nil {
		c.service.logger.Error(ErrorHandlingTracking.Format(err.Error()))
		return
	}

	if response.Status >= web.StatusBadRequest {
		c.service.logger.Error(ErrorTrackingStatus.Format(response.Status, response.StatusText).Error())
		return
	}
}
