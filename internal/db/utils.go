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

// generic filter function that takes an array of strings and a field,
// if the field is in the array, it will return true
func filterByField(field string, array []string) bool {
	if len(array) == 0 {
		return true
	}

	for _, v := range array {
		if v == field {
			return true
		}
	}
	return false
}
