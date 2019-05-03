package network

import (
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"gopkg.in/h2non/gentleman.v2"
	"time"
)

// Cache represent the process-global cache system, used for example to prevent repetitive HTTP requests from being sent to LibreNMS
var Cache = cache.New(5*time.Minute, 10*time.Minute)

// RunRequestIfNotCached will take a cache key and a request as parameter, it will only run it
// if the same request hasn't been executed already (checked using cache keys) and returns the
// cached response (or run it and cache the response before returning it)
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
