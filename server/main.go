package main

import (
	"context"
	"fmt"
	"gallery/server/database"
	"gallery/server/domain"
	"gallery/server/handler"
	"gallery/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {

	var (
		databaseUser     = os.Getenv("DATABASE_USER")
		databaseName     = os.Getenv("DATABASE_NAME")
		databaseHost     = os.Getenv("DATABASE_HOST")
		databasePort     = os.Getenv("DATABASE_PORT")
		databasePassword = os.Getenv("DATABASE_PASSWORD")
	)

	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, databasePort, databaseUser, databaseName, databasePassword)

	conn := database.OpenDB(dbConn)

	appAddr := ":" + os.Getenv("APP_PORT")

	r := gin.Default()
	r.Use(utils.CORSMiddleware()) //For CORS

	imgDomain := domain.NewImageService(conn)

	h := handler.NewHandlerService(imgDomain)

	r.GET("/images", h.ListImages)
	r.POST("/image_info", h.SaveImageInfo)
	r.GET("/image_info/:imageId", h.GetImageInfo)
	r.DELETE("/image_info/:imageId", h.DeleteImageInfo)

	//Starting and Shutting down Server
	srv := &http.Server{
		Addr:    appAddr,
		Handler: r,
	}

	go func() {
		//service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	//Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
