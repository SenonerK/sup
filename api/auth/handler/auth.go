package handler

import (
	"github.com/senonerk/sup/api/auth/forms"
	"github.com/senonerk/sup/shared/aerr"

	"github.com/senonerk/sup/shared/http/util"

	"github.com/senonerk/sup/shared/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/senonerk/sup/srv/auth/proto"
)

// authApi struct
type authApi struct {
	Client auth.AuthService
}

const (
	FQDN      = "senonerk.sup.api.auth"
	userIDKey = "userID"
)

// New retuens new handeler for auth
func New() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.ErrorReporter())

	srv := authApi{
		Client: auth.NewAuthService("senonerk.sup.srv.auth", client.DefaultClient),
	}

	a := router.Group("/auth")
	a.POST("/login", srv.Login)
	a.POST("/register", srv.Register)
	a.PUT("/password", middlewares.AuthenticatedRoute(), srv.ChangePassword)
	a.GET("/token", middlewares.AuthenticatedRoute(), srv.NewToken)

	return router
}

func (api *authApi) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req auth.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
		return
	}

	if req.Username == "" || req.Password == "" {
		c.Error(util.ErrBadRequest())
		return
	}

	res, err := api.Client.Login(ctx, &req)
	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, gin.H{
		"token": res.Token,
	})
}

func (api *authApi) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var req auth.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
	}

	if req.Username == "" || req.Password == "" {
		c.Error(util.ErrBadRequest())
		return
	}

	if len(req.Password) < 8 {
		c.Error(&aerr.AppError{
			Code:    400,
			Message: "Password must be at least 8 characters long",
			Source:  FQDN,
		})
		return
	}

	_, err := api.Client.Register(ctx, &req)
	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, nil)
}

func (api *authApi) NewToken(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := api.Client.NewToken(ctx, &auth.NewTokenRequest{
		UserID:         c.GetString(userIDKey),
		ExpiresInHours: 48,
	})

	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, gin.H{
		"token": res.Token,
	})
}

func (api *authApi) ChangePassword(c *gin.Context) {
	ctx := c.Request.Context()

	var req forms.ChangePasswordForm

	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
		return
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		c.Error(util.ErrBadRequest())
		return
	}

	if req.OldPassword == req.NewPassword {
		c.Error(&aerr.AppError{
			Code:    409,
			Message: "Password cannot be the same",
			Source:  FQDN,
		})
		return
	}

	if len(req.NewPassword) < 8 {
		c.Error(&aerr.AppError{
			Code:    400,
			Message: "Password must be at least 8 characters long",
			Source:  FQDN,
		})
		return
	}

	_, err := api.Client.CheckPassword(ctx, &auth.CheckPasswordRequest{
		UserID:   c.GetString(userIDKey),
		Password: req.OldPassword,
	})

	if err != nil {
		c.Error(err)
		return
	}

	_, err = api.Client.ChangePassword(ctx, &auth.ChangePasswordRequest{
		UserID:      c.GetString(userIDKey),
		NewPassword: req.NewPassword,
	})

	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, nil)
}
