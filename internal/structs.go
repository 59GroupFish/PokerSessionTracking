package internal

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	Id           uuid.UUID  `db:"id"`
	Name         string     `db:"name"`
	CreatedDate  time.Time  `db:"created_date"`
	InactiveDate *time.Time `db:"inactive_date"`
}

type Session struct {
	Id            uuid.UUID  `db:"id"`
	BBValue       float64    `db:"bb_value"`
	CashChipRatio float64    `db:"cash_chip_ratio"`
	CreatedDate   time.Time  `db:"created_date"`
	InactiveDate  *time.Time `db:"inactive_date"`
}

type PlayerSessionData struct {
	Id              uuid.UUID `db:"id"`
	BuyInCash       float64   `db:"buy_in_cash"`
	FinalStackChips float64   `db:"final_stack_chips"`
}

type Payment struct {
	Id           uuid.UUID  `db:"id"`
	Payer        string     `db:"payer"`
	Payee        string     `db:"payee"`
	Amount       float64    `db:"amount"`
	PaidDate     *time.Time `db:"paid_date"`
	CreatedDate  time.Time  `db:"created_date"`
	InactiveDate *time.Time `db:"inactive_date"`
}
