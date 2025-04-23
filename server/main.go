package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	controllers "github.com/Aritra640/ConnectSphere/server/internal/Controllers"
	ws "github.com/Aritra640/ConnectSphere/server/internal/WS/test_chat_room"
	"github.com/Aritra640/ConnectSphere/server/internal/auth"
	"github.com/Aritra640/ConnectSphere/server/internal/config"
	Internal_Validator "github.com/Aritra640/ConnectSphere/server/internal/validator"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error : no .env file found")
		panic(err)
	}

	Database_url := os.Getenv("DATABASE_URL")
	if Database_url == "" {
		log.Println("Error: database url empty!")
		panic("database empty!")
	}

	pg, err := sql.Open("postgres", Database_url)
	if err != nil {
		log.Println("Error: cannot open database connection")
		panic(err)
	}
	defer pg.Close()

	//Run migrations
	if err := goose.Up(pg, "./Database/migrations"); err != nil {
		log.Println("migrations failed")
		panic(err)
	} else {
		log.Println("migrations successfull!")
	}

	log.Printf("config.App is nil? %v\n", config.App == nil)

	jwt_key := os.Getenv("JWT")
	if jwt_key == "" {
		log.Println("JWT secret found empty!")
		panic(errors.New("jwt secret could not be fetched or is nil"))
	}
	config.App.JWT = []byte(jwt_key)
	config.App.DB = pg
	config.App.CTX = context.Background()
	config.App.QueryObj = db.New(pg)

	ws.Start_test_group()

  //Initialize auth service 
  auth.AuthSetup.Queries = config.App.QueryObj
  auth.AuthSetup.Rts = &auth.RefreshTokenService{Queries: config.App.QueryObj}
  auth.AuthSetup.Expiry = time.Hour * 24

	e := echo.New()

	//Register the custom validator
	e.Validator = &Internal_Validator.CustomValidatorService{
		Validator: validator.New(),
	}

	controllers.RoutesSetup(e)
	e.Logger.Fatal(e.Start(":8080"))
}
