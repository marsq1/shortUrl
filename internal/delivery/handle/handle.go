package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"net/http"
	"tinyUrlGRPC/pkg/proto"
)

const (
	port = "http://localhost:8000/"
)

type Handle struct {
	server *http.Server
	shorterService proto.ShorterServer
}

func NewHandle(shorterService proto.ShorterServer, port string) Handle {
	router := gin.New()
	rs := Handle{
		server: &http.Server{
			Addr: ":" + port,
			Handler: router,
		},
		shorterService: shorterService,
	}

	router.POST("/create", rs.createShortUrl)
	router.GET("/get", rs.getOriginalUrl)
	router.GET("/:redirect", rs.redirect)

	return rs
}

func (h *Handle) StartServer() error{
	return h.server.ListenAndServe()
}

func (h *Handle) createShortUrl(ctx *gin.Context) {
	var req proto.CreateRequest
	var save proto.SaveRequest
	if err := jsonpb.Unmarshal(ctx.Request.Body, &req); err != nil {
		ctx.String(http.StatusInternalServerError, "Error unmarshal request")
	}

	resp, err := h.shorterService.CreateTinyURL(ctx.Request.Context(), &req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error create response")
	}
	save.ShortUrl = resp.ShortUrl
	save.OriginalUrl = req.OriginalUrl
	_, err = h.shorterService.SaveUrl(ctx.Request.Context(), &save)

	ctx.JSON(http.StatusOK, gin.H{
		"shortUrl": port + resp.ShortUrl,
	})
}

func (h *Handle) getOriginalUrl(ctx *gin.Context) {
	var req proto.GetRequest
	if err := jsonpb.Unmarshal(ctx.Request.Body, &req); err != nil{
		ctx.String(http.StatusInternalServerError, "error unmarshal request")
	}

	original, err := h.shorterService.GetTinyURL(ctx.Request.Context(), &req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error get response")
	}

	m := jsonpb.Marshaler{}

	if err := m.Marshal(ctx.Writer, original); err != nil {
		ctx.String(http.StatusInternalServerError, "Error marshal response")
	}
}

func (h *Handle) redirect(ctx *gin.Context) {
	var req proto.GetRequest

	shortUrl := ctx.Param("redirect")
	req.ShortUrl = port + shortUrl
	log.Println(req.ShortUrl)
	originalUrl, err := h.shorterService.GetTinyURL(ctx.Request.Context(), &req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error get response")
	}

	ctx.Redirect(301, originalUrl.OriginalUrl)
}