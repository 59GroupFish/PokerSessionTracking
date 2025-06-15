package rest

import (
	"log"
	"net/http"

	"github.com/59GroupFish/PokerSessionTracking/internal"
	"github.com/labstack/echo/v4"
)

func CreateGameHandler(c echo.Context) error {
	reqBody := new(internal.Game)
	if err := c.Bind(reqBody); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Invalid body.")
	}

	response, err := internal.CreateGame(reqBody)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusNotAcceptable, "Error creating game.")
	}

	return c.JSON(http.StatusOK, &response)
}

func GetGamesHandler(c echo.Context) error {
	response, err := internal.GetGames()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Errror getting games.")
	}

	return c.JSON(http.StatusOK, &response)
}

func GetSimpleDebt(c echo.Context) error {
	// call simple debt here
	return nil
}
