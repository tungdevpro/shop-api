package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tungdevpro/shop-api/common"
	"github.com/tungdevpro/shop-api/config"
	"github.com/tungdevpro/shop-api/constant"
	"github.com/tungdevpro/shop-api/database"
	authBiz "github.com/tungdevpro/shop-api/modules/auth/business"
	authRepo "github.com/tungdevpro/shop-api/modules/auth/repository/inmem"
	authRest "github.com/tungdevpro/shop-api/modules/auth/transport/rest"
)

func main() {
	cfg := config.NewConfig()
	db, err := database.CreateMysql(cfg)
	if err != nil {
		log.Fatalln("Database cannot connect")
	}
	fmt.Println("Database was connected")

	baseCtx := common.NewBaseContext(cfg, db)

	app := fiber.New()

	authApi := authRest.NewApi(authBiz.NewBusiness(authRepo.NewAuthRepository(baseCtx)))

	api := app.Group(constant.Api)
	{
		api.Post(fmt.Sprintf("%s%s", constant.Auth, constant.SignUp), authApi.SignUpHandler())
	}

	app.Listen(fmt.Sprintf(":%v", cfg.Port))
}
