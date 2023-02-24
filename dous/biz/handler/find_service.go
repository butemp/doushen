package handler

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

// Find .
func Find(ctx context.Context, c *app.RequestContext) {
	//fileName := c.Param("path")
	fileName := c.Query("path")
	content, err := os.ReadFile("./file/upload/" + fileName)
	if err != nil {
		fmt.Println("read file error")
		return
	}
	words := strings.Split(fileName, ".")
	suffix := words[len(words)-1]
	if suffix == "mp4" {
		c.Data(200, "video/mp4", content)
	}
	if suffix == "jpg" {
		c.Data(200, "image/png", content)
	}
}
