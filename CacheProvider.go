package main

import "apodExample/cacheService"

type CacheServiceForApodClient struct {}

var cacheServer *cacheService.Cacher

func NewCacheServiceForApodClient(IPAddress string, Port string) *CacheServiceForApodClient {
	self := CacheServiceForApodClient{}
	cacheServer = cacheService.NewCacher(IPAddress, Port)
	return &self
}

func (c *CacheServiceForApodClient) Get(key string, callback func() string ) string {
	//c := cacheService.NewCacher(config.CacheServer.IPAddress, config.CacheServer.Port)
	return cacheServer.FetchFromCache(key, callback)
}
