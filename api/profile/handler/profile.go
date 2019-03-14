package handler

import (
	"github.com/senonerk/sup/api/profile/forms"
	"github.com/senonerk/sup/shared/aerr"
	"github.com/senonerk/sup/shared/http/util"
	"github.com/senonerk/sup/srv/profile/proto/profile"

	"github.com/senonerk/sup/shared/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

// profileApi struct
type profileApi struct {
	Client profile.ProfileService
}

const FQDN = "senonerk.sup.api.profile"

// New retuens new handler for auth
func New() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.ErrorReporter())

	srv := profileApi{
		Client: profile.NewProfileService("senonerk.sup.srv.profile", client.DefaultClient),
	}

	a := router.Group("/profile")
	a.PUT("/", middlewares.AuthenticatedRoute(), srv.UpdateInfo)
	a.POST("/status", middlewares.AuthenticatedRoute(), srv.UpdateStatus)

	return router
}

func (api *profileApi) UpdateInfo(c *gin.Context) {
	ctx := c.Request.Context()

	var req forms.UpdateInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
		return
	}

	if req.LastName == "" || req.FirstName == "" {
		c.Error(util.ErrBadRequest())
		return
	}

	_, err := api.Client.UpdateInfo(ctx, &profile.UpdateInfoRequest{
		UserID:    c.GetString(util.UserIDKey),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Birth:     req.BirthDate.Unix(),
	})

	if err != nil {
		c.Error(aerr.FromErr(err))
		return
	}

	util.Ok(c, nil)
}

func (api *profileApi) UpdateStatus(c *gin.Context) {
	ctx := c.Request.Context()

	var req forms.UpdateStatus
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
		return
	}

	if req.NewStatus == "" {
		c.Error(util.ErrBadRequest())
		return
	}

	_, err := api.Client.UpdateStatus(ctx, &profile.UpdateStatusRequest{
		UserID:    c.GetString(util.UserIDKey),
		NewStatus: req.NewStatus,
	})

	if err != nil {
		c.Error(aerr.FromErr(err))
		return
	}

	util.Ok(c, nil)
}
