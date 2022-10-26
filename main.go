package main

import (
	"gin-frame/cmd"
	_ "gin-frame/cmd/cmd_gorm_gen"
	_ "gin-frame/utility/log"
	_ "gin-frame/utility/pid"
	_ "gin-frame/utility/validator"
)

func main() {
	cmd.Execute()
}
