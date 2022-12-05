package main

import (
	_ "gphoto/internal/packed"
	"gphoto/pkg/qiniu"

	"gphoto/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	qiniu.ListFile()
	cmd.Main.Run(gctx.New())
}
