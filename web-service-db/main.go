package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
        AllowNativePasswords: true,
    }

    if db, err := sql.Open("mysql", cfg.FormatDSN()); err != nil {
        log.Fatal(db.Ping())
    } else {
        repo := Repo{DB: db}
        service := Service{Repo: &repo}
        controller := Controller{Service: &service}
        router := gin.Default()
        router.GET("/albums", controller.getAlbumns)
        router.GET("/artists", controller.getArtists)
        router.POST("/album", controller.postAlbum)
        router.DELETE("/album/:id", controller.deleteAlbum)
        router.Run("localhost:8080")
    }
}
