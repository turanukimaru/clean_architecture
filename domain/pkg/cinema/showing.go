package cinema

import (
	"github.com/pkg/errors"
	"log"
	"time"
)

// Showing is a film and screen and start time.
type Showing struct {
	Film      Film
	Screen    Screen
	StartTime time.Time
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
			err = errors.Wrap(err, "Rollbackエラー："+s.Screen.Seats.Rollback(s.StartTime, seats).Error())
		}
	}()
	// 全ての席に手を付ける
	if err = s.Screen.Seats.Touch(s.StartTime, seats); err != nil {
		return
	}
	// 先着がいないことを確認する。
	notDoubleBooking, err := s.Screen.Seats.IsFirstTouch(s.StartTime, seats)
	// 先着がいたら諦めて手放す
	if !notDoubleBooking {
		if err = s.Screen.Seats.LetGo(s.StartTime, seats); err != nil {
			return
		}
		return
	}
	// 予約を確定して記帳する
	if err = s.Screen.Seats.Book(s.StartTime, seats); err != nil {
		return
	}
	return
}

// Reservable is reserve
type Reservable interface {
	Touch(startTime time.Time, seats []Seat) error
	IsFirstTouch(startTime time.Time, seats []Seat) (bool, error)
	LetGo(startTime time.Time, seats []Seat) error
	Book(startTime time.Time, seats []Seat) error
	Rollback(startTime time.Time, seats []Seat) error
}
