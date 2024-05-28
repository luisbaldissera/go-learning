package main

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type Controller struct {
    Service *Service
}

func (c *Controller) getAlbumns(ctx *gin.Context) {
    albums, err := c.Service.ReadAlbums()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Some error occured."})
        return
    }
    ctx.JSON(http.StatusOK, albums)
}

func (c *Controller) getArtists(ctx *gin.Context) {
    artists, err := c.Service.Artists()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Some error occured."})
        return
    }
    ctx.JSON(http.StatusOK, artists)
}

func (c *Controller) postAlbum(ctx *gin.Context) {
    var album Album
    if err := ctx.BindJSON(&album); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Some error occured."})
    }
    createdAlbum, err := c.Service.AddAlbum(album)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Some error occured."})
    }
    ctx.JSON(http.StatusCreated, createdAlbum)
}

func (c *Controller) deleteAlbum(ctx *gin.Context) {
    id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Some error occured."})
    }
    album, err := c.Service.DeleteAlbum(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Some error occured."})
    }
    ctx.JSON(http.StatusOK, album)
}

// vim: fdl=0
