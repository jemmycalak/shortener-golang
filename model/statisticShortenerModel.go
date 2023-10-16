package model

import "time"

type StatisticShortener struct {
	Id          int64 `gorm:"primary_key"`
	ShortenerId string
	Shortener   Shortener `gorm:"references:Id"`
	IpAddress   string
	Date        time.Time `gorm:"type:DATE;not null;default:CURRENT_TIMESTAMP"`
}
