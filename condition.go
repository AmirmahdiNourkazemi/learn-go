package main

import (
	"fmt"
	"time"
)

func ifCondition(x float64) string {
	if x > -5 && x < 0 {
		fmt.Println(x, "is negative")
	} else if x > 0 {
		fmt.Println(x, "is positive")
	}

	return "done"
}

func switchCaseCondition(x time.Weekday) string {
	switch time.Now().Weekday() {
	case x:
		return "It's the " + time.Now().Weekday().String()
	default:
		return "It's a weekday"
	}
}

func switchCasewithoutCondition() string {
	switch {
	case time.Hour < 12:
		return "It's before noon"
	case time.Hour > 12:
		return "It's after noon"
	default:
		return "It's a weekday"
	}
}
