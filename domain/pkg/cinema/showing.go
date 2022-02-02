package cinema

import (
	"log"
	"time"
)

// Showing is a film and screen and start time.
type Showing struct {
	Film      Film
	Screen    Screen
	StartTime *time.Time
	logger    *log.Logger
}

// Seat is seat index in the Screen
type Seat uint

// Screen is a screen and room in the theater
type Screen struct {
	ID    uint
	Name  string
	Seats Reservable
}

// Film is Cinema
type Film struct {
	ID   uint
	Name string
}

// Reserve is 予約処理
func (s Showing) Reserve(seats []Seat) (err error) {
	defer func() { // エラーが発生したら全て取り消す
		if err != nil {
			s.Screen.Seats.Rollback()
		}
	}()
	// 全ての席に手を付ける
	if err = s.Screen.Seats.Touch(seats); err != nil {
		return
	}
	// 先着がいないことを確認する。
	doubleBooking := s.Screen.Seats.IsDouble(seats)
	// 先着がいたら諦めて手放す
	if doubleBooking {
		if err = s.Screen.Seats.LetGo(seats); err != nil {
			return
		}
		return
	}
	// 予約を確定して記帳する
	if err = s.Screen.Seats.Book(seats); err != nil {
		return
	}
	return
}

// Reservable is reserve
type Reservable interface {
	Touch(seats []Seat) error
	IsDouble(seats []Seat) bool
	LetGo(seats []Seat) error
	Book(seats []Seat) error
	Rollback() error
}
