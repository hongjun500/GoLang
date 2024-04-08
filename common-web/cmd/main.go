package main

import (
	"fmt"

	_ "github.com/hongjun500/GoLang-master/common-web/docs"
	_ "github.com/hongjun500/GoLang-master/common-web/internal/conf"
	"github.com/hongjun500/GoLang-master/common-web/internal/web"
)

// @title					common-web API
// @description				common-web API
// @version					v1
// @BasePath				/
// @schemes					http https
// //@securityDefinitions.apikey	GinJWTMiddleware
// @in						header
// @name					Authorization
func main() {
	fmt.Println("hello common-web")
	web.Bootstrap()
}
