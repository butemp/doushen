// Code generated by hertz generator.

package main

import (
	"dous/biz/model/query"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithMaxRequestBodySize(1000 << 20))
	Init()
	query.SetDefault(DB)
	register(h)
	h.Spin()
}