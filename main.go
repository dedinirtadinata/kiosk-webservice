package main

import (
	"fmt"
	_handler "github.com/dedinirtadinata/kiosk-webservice/delivery/http"
	"github.com/dedinirtadinata/kiosk-webservice/display/model"
	_displayRepo "github.com/dedinirtadinata/kiosk-webservice/display/repository"
	_displayUc "github.com/dedinirtadinata/kiosk-webservice/display/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	} else {
		//gin.SetMode(gin.ReleaseMode)
	}

	http.DefaultClient.Timeout = time.Second * time.Duration(viper.GetInt("context.timeout"))
}
func main() {
	r := gin.Default()
	db, err := gorm.Open(sqlite.Open("simontok.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.DisplayModel{})
	if err != nil {
		panic(err.Error())
	}
	//init
	displayRepo := _displayRepo.NewdisplayRepository(db)
	ucDisplay := _displayUc.NewdisplayUsecase(displayRepo)

	_handler.NewdisplayHandler(r, ucDisplay)
	//r.Use(cors.Middleware(cors.Config{
	//	Origins:         "*",
	//	Methods:         "GET, PUT, POST, DELETE",
	//	RequestHeaders:  "Origin, Authorization, Content-Type",
	//	ExposedHeaders:  "",
	//	MaxAge:          60 * time.Second,
	//	Credentials:     false,
	//	ValidateHeaders: false,
	//}))

	r.Run(viper.GetString("server.address"))
}
