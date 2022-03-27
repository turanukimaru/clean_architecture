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
type FasterTouchedSeat struct {
	gorm.Model
	ScreenId uint
	No       uint
	StartAt  time.Time
	//	Name     string
	UserId uint
	Faster uint
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

func (s *ReserveEvent) Touch(startAt time.Time, seats []cinema.Seat) error {
	for no := range seats {
		s.db.Create(&ReservedSeat{ScreenId: 1, No: uint(no), StartAt: startAt, UserId: s.UserId})
		if err := s.db.Error; err != nil {
			return err
		}
	}
	return nil
}

// IsFirstTouch は seats が全て先に確保できたことを示す。
func (s *ReserveEvent) IsFirstTouch(startAt time.Time, seats []cinema.Seat) (bool, error) {
	var touchedSeats []FasterTouchedSeat
	// 同じ席同じ時間で自分より id の若いものを検出する。
	// 想定SQL:SELECT id, (SELECT ifnull(min(id),0) FROM touched_seat other WHERE other.id < base.id and other.screen_id = base.screen_id and other.no = base.no) as faster FROM (select * from `touched_seat` WHERE screen_id = 1 and start_at = '2022-03-25 09:00:00' and user_id = 2 and no in (1,2)) as base;
	query := "id, (SELECT IFNULL(min(id),0) FROM touched_seat other WHERE other.id < base.id and other.screen_id = base.screen_id and other.no = base.no and other.start_at = base.start_at) as faster"
	table := s.db.Model(&TouchedSeat{}).Select("*").Where("user_id = ? and screen_id = ? and start_at = ?", s.UserId, s.ScreenId, startAt).Where("no IN ?", seats)
	s.db.Select(query).Table("(?) as base", table).Find(&touchedSeats)
	for i := range touchedSeats {
		if touchedSeats[i].Faster > 0 {
			// バグ検出のために重複があったらLog出力しておくとよい
			return false, nil
		}
	}
	return true, nil
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
