package handler

import (
	"regexp"
	"time"

	"github.com/senonerk/sup/api/profile/forms"
	"github.com/senonerk/sup/shared/aerr"
	"github.com/senonerk/sup/shared/http/util"
	"github.com/senonerk/sup/srv/profile/proto/profile"

	"github.com/senonerk/sup/shared/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

// profileAPI struct
type profileAPI struct {
	Client profile.ProfileService
}

// FQDN this apis unique identifier
const FQDN = "senonerk.sup.api.profile"

var emailRegex *regexp.Regexp

func init() {
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
}

// New retuens new handler for auth
func New() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.ErrorReporter())
	router.Use(middlewares.AuthenticatedRoute())

	srv := profileAPI{
		Client: profile.NewProfileService("senonerk.sup.srv.profile", client.DefaultClient),
	}

	a := router.Group("/profile")
	a.GET("/", srv.GetInfo)
	a.PUT("/", srv.UpdateInfo)
	a.POST("/status", srv.UpdateStatus)
	a.PUT("/email", srv.UpdateEmail)
	a.POST("/email", srv.ConfirmEmail)

	return router
}

func (api *profileAPI) UpdateInfo(c *gin.Context) {
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

func (api *profileAPI) UpdateStatus(c *gin.Context) {
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

func (api *profileAPI) UpdateEmail(c *gin.Context) {
	ctx := c.Request.Context()

	var req forms.UpdateEmail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
		return
	}

	if !emailRegex.MatchString(req.NewEmail) {
		c.Error(&aerr.AppError{
			Code:    400,
			Message: "Invalid email",
			Source:  FQDN,
		})
		return
	}

	_, err := api.Client.UpdateEmail(ctx, &profile.UpdateEmailRequest{
		UserID:   c.GetString(util.UserIDKey),
		NewEmail: req.NewEmail,
	})

	if err != nil {
		c.Error(aerr.FromErr(err))
		return
	}

	util.Ok(c, nil)
}

func (api *profileAPI) ConfirmEmail(c *gin.Context) {
	ctx := c.Request.Context()

	var req forms.ConfirmEmail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
		return
	}

	if req.Token == "" {
		c.Error(util.ErrBadRequest())
		return
	}

	_, err := api.Client.ConfirmEmail(ctx, &profile.ConfirmEmailRequest{
		UserID:     c.GetString(util.UserIDKey),
		EmailToken: req.Token,
	})

	if err != nil {
		c.Error(aerr.FromErr(err))
		return
	}

	util.Ok(c, nil)
}

func (api *profileAPI) GetInfo(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := api.Client.GetInfo(ctx, &profile.GetInfoRequest{
		UserID: c.GetString(util.UserIDKey),
	})

	if err != nil {
		c.Error(aerr.FromErr(err))
		return
	}

	util.Ok(c, struct {
		FirstName string    `json:"firstname"`
		LastName  string    `json:"lastname"`
		BirthDate time.Time `json:"birthdate"`
		Status    string    `json:"status"`
		Email     string    `json:"email"`
	}{
		FirstName: res.FirstName,
		LastName:  res.LastName,
		BirthDate: time.Unix(res.Birth, 0),
		Status:    res.Status,
		Email:     res.Email,
	})
}
