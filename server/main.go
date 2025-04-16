package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Aritra640/ConnectSphere/server/Database/db"
	"github.com/Aritra640/ConnectSphere/server/internal/config"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {

	err := godotenv.Load(); if err != nil {
    log.Println("Error : no .env file found")
    panic(err)
  }

  Database_url := os.Getenv("DATABASE_URL"); if Database_url == "" {
    log.Println("Error: database url empty!")
    panic("database empty!")
  }
  
  pg,err := sql.Open("postgres" , Database_url); if err != nil {
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

  config.App.DB = pg
  config.App.CTX = context.Background()
  config.App.QueryObj = db.New(pg)

  e := echo.New() 
  e.GET("/hello" , func(c echo.Context) error {
    return c.JSON(200 , "hi hello there")
  })

  e.Logger.Fatal(e.Start(":8080"))
}
