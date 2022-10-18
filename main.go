package main

import (
	"gin-frame/cmd"
	_ "gin-frame/cmd/cmd_gorm_gen"
)

func main() {
	cmd.Execute()
}
