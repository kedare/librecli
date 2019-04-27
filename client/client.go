package client

import (
	"os"

	"gopkg.in/h2non/gentleman.v2"
)

func BuildAPIClient() gentleman.Client {
	base := gentleman.New()
	base.BaseURL(os.Getenv("LIBRECLI_URL"))
	base.SetHeader("X-Auth-Token", os.Getenv("LIBRECLI_TOKEN"))
	return *base
}
