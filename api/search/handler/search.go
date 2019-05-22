package handler

import (
	"github.com/senonerk/sup/shared/http/util"
	"github.com/senonerk/sup/srv/profile/proto/profile"

	"github.com/senonerk/sup/shared/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

// searchAPI struct
type searchAPI struct {
	Client profile.ProfileService
}

// New retuens new handler for auth
func New() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.ErrorReporter())
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.AuthenticatedRoute())

	srv := searchAPI{
		Client: profile.NewProfileService("senonerk.sup.srv.profile", client.DefaultClient),
	}

	a := router.Group("/search")
	a.GET("/", srv.Search)

	return router
}

func (api *searchAPI) Search(c *gin.Context) {
	ctx := c.Request.Context()

	query := c.Query("query")
	if query == "" {
		c.Error(util.ErrBadRequest())
		return
	}

	p, err := api.Client.Search(ctx, &profile.SearchRequest{
		Query: query,
	})

	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, p.Users)
}
