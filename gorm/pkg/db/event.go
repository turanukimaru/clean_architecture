package dummydb

import (
	"github.com/turanukimaru/ca/domain/pkg/cinema"
	"gorm.io/gorm"
	"time"
)

// TouchedSeat テーブル
type TouchedSeat struct {
	gorm.Model
	ScreenId uint
	No       uint
	StartAt  time.Time
	//	Name     string
	UserId uint
}

// ReservedSeat テーブル
type ReservedSeat struct {
	gorm.Model
	ScreenId uint
	No       uint
	StartAt  time.Time
	//	Name     string
	UserId uint
}

type ReserveEvent struct {
	ScreenId uint
	UserId   uint
	db       *gorm.DB
}

func (s *ReserveEvent) Touch(startTime time.Time, seats []cinema.Seat) error {
	for no, _ := range seats {
		// Create
		s.db.Create(&ReservedSeat{ScreenId: 1, No: uint(no), StartAt: startTime, UserId: s.UserId})
		if err := s.db.Error; err != nil {
			return err
		}
	}
	return nil
}
func (s *ReserveEvent) IsFirstTouch(startTime time.Time, seats []cinema.Seat) (bool, error) {
	return false, nil
}
func (s *ReserveEvent) LetGo(startTime time.Time, seats []cinema.Seat) error {
	return nil
}
func (s *ReserveEvent) Book(startTime time.Time, seats []cinema.Seat) error {
	return nil
}
func (s *ReserveEvent) Rollback(startTime time.Time, seats []cinema.Seat) error {
	s.db.Rollback()
	return s.db.Error
}

// どこで接続するかな…
//	dsn := "sqlserver://gorm:pass@127.0.0.1:1433?database=gorm"
//	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
