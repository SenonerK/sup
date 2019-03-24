package handler

import (
	"github.com/senonerk/sup/api/chat/forms"
	"github.com/senonerk/sup/shared/http/util"
	"github.com/senonerk/sup/srv/chat/proto/chat"

	"github.com/senonerk/sup/shared/http/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
)

// profileAPI struct
type chatAPI struct {
	Client chat.ChatService
}

// FQDN this apis unique identifier
const FQDN = "senonerk.sup.api.chat"

// New retuens new handler for auth
func New() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.ErrorReporter())
	router.Use(middlewares.AuthenticatedRoute())

	srv := chatAPI{
		Client: chat.NewChatService("senonerk.sup.srv.chat", client.DefaultClient),
	}

	c := router.Group("/chat")
	c.POST("/", srv.Send)
	c.GET("/new", srv.ReceiveNew)

	return router
}

func (api *chatAPI) Send(c *gin.Context) {
	ctx := c.Request.Context()

	var req forms.Send
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
		return
	}

	if req.ToUserID == "" || req.Message == "" {
		c.Error(util.ErrBadRequest())
		return
	}

	_, err := api.Client.Send(ctx, &chat.SendRequest{
		FromUserID: c.GetString(util.UserIDKey),
		ToUserID:   req.ToUserID,
		Message:    req.Message,
	})

	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, nil)
}

func (api *chatAPI) ReceiveNew(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := api.Client.Receive(ctx, &chat.ReceiveRequest{
		UserID: c.GetString(util.UserIDKey),
	})

	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, res)
}
