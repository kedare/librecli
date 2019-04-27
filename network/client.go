package network

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"os"
)

func BuildHTTPClient() *gentleman.Client {
	base := gentleman.New()

	base.UseResponse(func(ctx *context.Context, h context.Handler) {
		log.Debugf("HTTP %v %v", ctx.Response.Status, ctx.Request.URL)
		h.Next(ctx)
	})

	return base
}

func BuildAPIClient() *gentleman.Client {
	base := BuildHTTPClient()
	base.BaseURL(os.Getenv("LIBRECLI_URL"))
	base.SetHeader("X-Auth-Token", os.Getenv("LIBRECLI_TOKEN"))
	return base
}
