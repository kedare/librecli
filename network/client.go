package network

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
)

// BuildHTTPClient build the default HTTP client used in the application and inject a debug logging
func BuildHTTPClient() *gentleman.Client {
	base := gentleman.New()

	base.UseResponse(func(ctx *context.Context, h context.Handler) {
		log.Debugf("HTTP %v %v", ctx.Response.Status, ctx.Request.URL)
		h.Next(ctx)
	})

	return base
}

// BuildAPIClient build the HTTP client used for LibreNMS, it automatically inject the base URL and authentication token.
func BuildAPIClient() *gentleman.Client {
	base := BuildHTTPClient()
	baseURL := os.Getenv("LIBRECLI_URL")
	if baseURL == "" {
		panic("LIBRECLI_URL is not defined")
	}
	base.BaseURL(os.Getenv("LIBRECLI_URL"))

	token := os.Getenv("LIBRECLI_TOKEN")
	if token == "" {
		panic("LIBRECLI_TOKEN is not defined")
	}

	base.SetHeader("X-Auth-Token", os.Getenv("LIBRECLI_TOKEN"))
	return base
}
