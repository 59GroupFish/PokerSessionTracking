package internal

import (
	"log"
	"time"

	"github.com/59GroupFish/PokerSessionTracking/db"
	"github.com/google/uuid"
)

type Player struct {
	Id           uuid.UUID  `json:"id" db:"id"`
	Name         string     `json:"name" db:"name" binding:"required"`
	CreatedDate  time.Time  `json:"created_date" db:"created_date"`
	InactiveDate *time.Time `json:"inactive_date" db:"inactive_date"`
	GameId       uuid.UUID  `json:"game_id" db:"game_id" binding:"required"`
}

func AddPlayer(newPlayer *Player) (Player, error) {
	var response Player

	insertTx, err := db.DB.Beginx()
	if err != nil {
		return response, err
	}
	defer insertTx.Rollback()

	sql := `
		insert into players (name, game_id)
		VALUES ($1, $2)
		returning id, name, created_date, inactive_date, game_id;
	`
	log.Printf("gameId: %s", newPlayer.GameId)
	err = insertTx.Get(&response, sql, newPlayer.Name, newPlayer.GameId)
	if err != nil {
		return response, err
	}

	if err := insertTx.Commit(); err != nil {
		log.Printf("error committing transaction: %s", err.Error())
		return response, err
	}

	return response, nil
}

func GetPlayers(gameId string) ([]Player, error) {
	var response []Player

	sql := `
		select id, name, created_date, inactive_date from players
		where inactive_date is null
		and game_id = $1
		order by name
	`

	err := db.DB.Select(&response, sql, gameId)
	if err != nil {
		return response, err
	}

	return response, nil
}
