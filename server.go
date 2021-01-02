package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func CreateServer(config *Config) *fasthttp.Server {
	handler := &Handler{Config: config}
	if config.HttpCacheTTL == 0 {
		handler.CacheControlHeader = []byte("private, no-cache, no-store, must-revalidate")
	} else {
		handler.CacheControlHeader = []byte(fmt.Sprintf("max-age=%d", config.HttpCacheTTL))
	}
	handler.TaskManager = NewTaskManager(config.ConvertConcurrency)
	return &fasthttp.Server{
		Handler:               handler.handleRequests,
		ErrorHandler:          handler.handleErrors,
		NoDefaultServerHeader: true,
		MaxRequestBodySize:    config.MaxUploadedImageSize * 1024 * 1024,
		ReadTimeout:           time.Duration(5 * time.Second),
	}
}
