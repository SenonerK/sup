package handler

import (
	"strconv"

	"github.com/senonerk/sup/shared/aerr"

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
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.AuthenticatedRoute())

	srv := chatAPI{
		Client: chat.NewChatService("senonerk.sup.srv.chat", client.DefaultClient),
	}

	c := router.Group("/chat")
	c.POST("/", srv.Send)
	c.GET("/", srv.Receive)
	c.GET("/new", srv.ReceiveNew)
	c.PUT("/:what", srv.Update)

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

	res, err := api.Client.ReceiveNew(ctx, &chat.ReceiveNewRequest{
		UserID: c.GetString(util.UserIDKey),
	})

	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, res)
}

func (api *chatAPI) Receive(c *gin.Context) {
	ctx := c.Request.Context()

	amount, err := strconv.Atoi(c.Query("amount"))
	if err != nil {
		c.Error(&aerr.AppError{
			Code:    400,
			Message: "Invalid amount specified",
		})
		return
	}

	skip, err := strconv.Atoi(c.Query("skip"))
	if err != nil {
		c.Error(&aerr.AppError{
			Code:    400,
			Message: "Invalid skip specified",
		})
		return
	}

	from, err := strconv.Atoi(c.Query("from"))
	if err != nil {
		c.Error(&aerr.AppError{
			Code:    400,
			Message: "Invalid from specified",
		})
		return
	}

	if amount < 10 {
		c.Error(util.ErrBadRequest())
		return
	}

	res, err := api.Client.Receive(ctx, &chat.ReceiveRequest{
		UserID: c.GetString(util.UserIDKey),
		Amount: int32(amount),
		Skip:   int32(skip),
		From:   int64(from),
	})

	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, res)
}

func (api *chatAPI) Update(c *gin.Context) {
	ctx := c.Request.Context()

	var req forms.Update
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(util.ErrInavlidForm())
		return
	}

	if req.ChatID == "" || req.Timestamp.Unix() <= 0 {
		c.Error(util.ErrBadRequest())
		return
	}

	rpcreq := &chat.UpdateRequest{
		UserID: c.GetString(util.UserIDKey),
		ChatID: req.ChatID,
		When:   req.Timestamp.Unix(),
	}

	var err error
	if what := c.Param("what"); what == "read" {
		_, err = api.Client.Read(ctx, rpcreq)
	} else if what == "receive" {
		_, err = api.Client.Received(ctx, rpcreq)
	} else {
		c.Error(&aerr.AppError{
			Code:    400,
			Message: "Incorrect parameter specified",
		})
		return
	}

	if err != nil {
		c.Error(err)
		return
	}

	util.Ok(c, nil)
}
