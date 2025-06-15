package internal

import (
	"log"
	"time"

	"github.com/59GroupFish/PokerSessionTracking/db"
	"github.com/google/uuid"
)

type Game struct {
	Id           uuid.UUID  `json:"id" db:"id"`
	Name         string     `json:"name" db:"name" binding:"required"`
	CreatedDate  time.Time  `json:"created_date" db:"created_date"`
	InactiveDate *time.Time `json:"inactive_date" db:"inactive_date"`
}

func CreateGame(newGame *Game) (Game, error) {
	var response Game

	insertTx, err := db.DB.Beginx()
	if err != nil {
		return response, err
	}
	defer insertTx.Rollback()

	sql := `
		insert into games (name)
		VALUES ($1)
		returning id, name, created_date;
	`

	err = insertTx.Get(&response, sql, newGame.Name)
	if err != nil {
		return response, err
	}

	if err := insertTx.Commit(); err != nil {
		log.Printf("error committing transaction: %s", err.Error())
		return response, err
	}

	return response, nil
}

func GetGames() ([]Game, error) {
	var response []Game

	sql := `
		select id, name, created_date, inactive_date from games
		where inactive_date is null
	`

	err := db.DB.Select(&response, sql)
	if err != nil {
		return response, err
	}

	return response, nil
}
