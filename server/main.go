package main

import (
	"time"

	"vue-typescript-beego-admin/server/models"
	_ "vue-typescript-beego-admin/server/routers"

	"vue-typescript-beego-admin/server/utils"

	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
)

func main() {
	models.Init()

	utils.Che = cache.New(60*time.Minute, 120*time.Minute)

	beego.Run()
}
