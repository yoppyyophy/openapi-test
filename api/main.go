package main

import (
	"log"
	"net/http"
	"openapi-test/generated"

	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetArtists(ctx echo.Context, _ generated.GetArtistsParams) error {
	albumsRecorded := 5
	artistGenre := "rock"
	artistName := "Yamada Taro"
	return ctx.JSON(http.StatusOK, &generated.Artist{
		AlbumsRecorded: &albumsRecorded,
		ArtistGenre:    &artistGenre,
		ArtistName:     &artistName,
		Username:       "hoge",
	})
}

func (Server) PostArtists(ctx echo.Context) error {
	return nil
}

func (Server) GetArtistsUsername(ctx echo.Context, username string) error {
	return nil
}

func main() {
	server := NewServer()

	e := echo.New()

	generated.RegisterHandlers(e, server)

	log.Fatal(e.Start("0.0.0.0:8080"))
}
