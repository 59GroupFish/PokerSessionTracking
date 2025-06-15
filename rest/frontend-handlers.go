package rest

import (
	"log"
	"net/http"

	"github.com/59GroupFish/PokerSessionTracking/client/components"
	"github.com/59GroupFish/PokerSessionTracking/internal"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func GetHome(c echo.Context) error {
	return Render(c, http.StatusOK, components.Index("Home", components.HomeBody()))
}

func ViewGames(c echo.Context) error {
	games, err := internal.GetGames()
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error getting games.")
	}
	return Render(c, http.StatusOK, components.Index("Games", components.ViewGames(games)))
}

func ManagePlayers(c echo.Context) error {
	gameId := c.Param("gameId")

	// TODO: Verify game exists and user has access

	players, err := internal.GetPlayers(gameId)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error getting players.")
	}

	return Render(c, http.StatusOK, components.Index("Manage Players", components.ManagePlayers(players, gameId)))
}
