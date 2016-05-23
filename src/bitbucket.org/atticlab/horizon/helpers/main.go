//Package helpers contains miscellaneous helper functions
package helpers

import (
    "time"
)

// SameWeek resturns true if both of the timestamps are within the same week 
func SameWeek (t1 time.Time, t2 time.Time) bool {
    year1, week1 := t1.ISOWeek()
	year2, week2 := t2.ISOWeek()
			
	return week1 == week2 && year1 == year2 
}

// MaxDate returns latest of two timestamps
func MaxDate(t1 time.Time, t2 time.Time) time.Time{
    if t1.Unix() > t2.Unix() {
        return t1
    }
    
    return t2
}