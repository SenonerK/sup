package middlewares

import (
	"strings"

	"github.com/senonerk/sup/shared/aerr"

	"github.com/micro/go-micro/client"
	"github.com/senonerk/sup/srv/auth/proto"

	"github.com/gin-gonic/gin"
	"github.com/senonerk/sup/shared/http/util"
)

// AuthenticatedRoute is a middleware for authentication
func AuthenticatedRoute() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		header := c.GetHeader("authorization")
		if header != "" {
			bearer := strings.Split(header, " ")
			if len(bearer) == 2 {

				api := auth.NewAuthService("senonerk.sup.srv.auth", client.DefaultClient)
				res, err := api.VerifyToken(ctx, &auth.VerifyTokenRequest{
					Token: bearer[1],
				})

				if err != nil {
					c.Error(aerr.FromErr(err))
					c.Abort()
					return
				}

				c.Set("userID", res.UserID)
				c.Next()
				return
			}
		}
		c.Error(util.ErrUnauthorized())
		c.Abort()
	}
}
