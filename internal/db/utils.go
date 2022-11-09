package db

import "gorm.io/gorm"

func FilterTimelineByDate(db *gorm.DB, start_date string, end_date string) *gorm.DB {
	if start_date != "" {
		db = db.Where("timelines.timestamp >= ?", start_date)
	}
	if end_date != "" {
		db = db.Where("timelines.timestamp <= ?", end_date)
	}

	return db
}
