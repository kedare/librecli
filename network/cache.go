package network

import (
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gentleman.v2"
	"time"
)

var Cache = cache.New(5*time.Minute, 10*time.Minute)

func RunRequestIfNotCached(key string, req *gentleman.Request) (*gentleman.Response, error) {
	res, found := Cache.Get(key)
	if found {
		log.Debugf("CACHE HIT %v", key)
		return res.(*gentleman.Response), nil
	} else {
		log.Debugf("CACHE MISS %v", key)
		res, err := req.Send()
		if err != nil {
			return nil, err
		}
		Cache.Set(key, res, cache.NoExpiration)
		return res, nil
	}
}
