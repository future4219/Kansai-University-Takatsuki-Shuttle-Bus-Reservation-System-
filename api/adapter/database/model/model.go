package model

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey;size:26;not null"` // PK
	StudentID string `gorm:"unique;not null"`             // 学籍番号
	IdmUniv   string `gorm:"size:255;not null"`           // 学生証のIDm
	IdmBus    string `gorm:"size:255:not null"`           // バスカードのIDm
	UserType  string `gorm:"size:255;not null"`           // ユーザータイプ
	IsDeleted bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Bus struct {
	ID            string    `gorm:"primaryKey;size:26;not null"` // PK
	BusID         string    `gorm:"size:255;not null"`           // バスID
	DepartureTime time.Time `gorm:"not null"`                    // 出発時刻
	Seats         []Seat    `gorm:"constraint:OnDelete:CASCADE;"`
	IsUp          bool      `gorm:"not null"`
}

type Seat struct {
	ID           string `gorm:"primaryKey;size:26;not null"` // PK
	Number       int32  `gorm:"not null"`                    // 座席番号
	BusID        string `gorm:"size:26;not null"`
	ReservedUser *User  `gorm:"constraint:OnDelete:CASCADE;"` // 予約ユーザー（nilであれば予約なし）
}

type File struct {
	FileID        string `gorm:"primaryKey;size:26;not null"`
	FileKind      string `gorm:"size:255;not null"`
	FileName      string `gorm:"size:1024;not null"`
	ContentType   string `gorm:"size:255;not null"`
	FileSize      int32  `gorm:"not null"`
	IsDeleted     bool   `gorm:"default:false"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedUserID string // 追加したユーザー
	CreatedUser   User   `gorm:"constraint:OnDelete:CASCADE;"`
	UpdatedUserID string // 更新したユーザー
	UpdatedUser   User   `gorm:"constraint:OnDelete:CASCADE;"`
}
