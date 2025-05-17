package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	controllers "github.com/Aritra640/ConnectSphere/server/internal/Controllers"
	mail "github.com/Aritra640/ConnectSphere/server/internal/Mail-Server"
	gcs "github.com/Aritra640/ConnectSphere/server/internal/WS/Group_Messages"
	pcs "github.com/Aritra640/ConnectSphere/server/internal/WS/Personal_Messages"
	tcr "github.com/Aritra640/ConnectSphere/server/internal/WS/test_chat_room"
	"github.com/Aritra640/ConnectSphere/server/internal/auth"
	"github.com/Aritra640/ConnectSphere/server/internal/cachestore"
	"github.com/Aritra640/ConnectSphere/server/internal/config"
	Internal_Validator "github.com/Aritra640/ConnectSphere/server/internal/validator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	config.App.QueryObj = db.New(config.App.DB)

	//personal chat service setup
	pcs.PersonalMessageSetup.Queries = config.App.QueryObj
	pcs.PersonalMessageSetup.WS_store = pcs.NewPersonalChatStore()
	pcs.PersonalMessageSetup.CUID = make(map[uuid.UUID]pcs.PersonalChatID_UIDmap)

	tcr.Start_test_group()

	//Initialize auth service
	auth.AuthSetup.Queries = config.App.QueryObj
	auth.AuthSetup.Rts = &auth.RefreshTokenService{Queries: config.App.QueryObj}
	auth.AuthSetup.Expiry = time.Hour * 24

	//Group Service Setup
	gcs.GroupChatMessageSetup.JWT = config.App.JWT

	config.App.PCS = pcs.PersonalMessageSetup
	config.App.GCS = gcs.GroupChatMessageSetup

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	config.App.PCS.WS_store.RunWS(ctx)
	config.App.GCS.RunAll(ctx)

  //Setting up mail server 
  mail.MailSetup.Email = os.Getenv("GOOGLE_EMAIL")
  mail.MailSetup.Password = os.Getenv("GOOGLE_APP_PASSWORD")
  mail.MailSetup.TestEmail = os.Getenv("TEST_EMAIL")
  mail.MailSetup.Smtp = os.Getenv("SMTP")

  go TimersRefresh(ctx)

	e := echo.New()

	//Test web socket connection with authentication
	e.Any("/test_ws", tcr.TestChatRoom)

	//Register the custom validator
	e.Validator = &Internal_Validator.CustomValidatorService{
		Validator: validator.New(),
	}

	//Use CORS
	e.Use(middleware.CORS())

	controllers.RoutesSetupV1(e)
	e.Logger.Fatal(e.Start(":8080"))
}

//Refresh cache storage every 10 minutes
func TimersRefresh(ctx context.Context){
  timer := time.NewTicker(10 * time.Minute)
  defer timer.Stop()

  var wg = &sync.WaitGroup{}
  log.Println("Refresh cache service started, task will eun every 10 minutes")

  for t := range timer.C {
    log.Printf("Running refresh cache task at %v" , t)
    wg.Add(2)
    go cachestore.CacheService.RefreshOtpStorage(ctx , wg)
    go cachestore.CacheService.RefreshUnverifiedUserData(ctx,wg)
  }

  wg.Wait()
  log.Println("Gracefully shutdown refresh-timer")
}
