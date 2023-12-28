package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(":::: Example related to time in go ::::")

	timeNow := time.Now()

	fmt.Println(timeNow)

	formattedDateStr := timeNow.Format("02-01-2006")
	fmt.Println("Formatted date is :::: ", formattedDateStr)

	formattedDateTime := timeNow.Format("02-01-2006 15:04:05")
	fmt.Println("Formatted date time is ::: ", formattedDateTime)

	formattedDateTimeWithDay := timeNow.Format("02-01-2006 15:04:05 Monday")
	fmt.Println("Formatted date time with day is ::: ", formattedDateTimeWithDay)

	formattedDateTimeWith12hrFormat := timeNow.Format("02-01-2006 03:04:05 Monday")
	fmt.Println("Formatted date time with 12 hr format is ::: ", formattedDateTimeWith12hrFormat)

	formattedDateTimeWithRFC822 := timeNow.Format(time.RFC822)
	fmt.Println("formattedDateTimeWithRFC822 is ::: ", formattedDateTimeWithRFC822)

	formattedDateTimeWithRFC3339 := timeNow.Format(time.RFC3339)
	fmt.Println("formattedDateTimeWithRFC3339 is ::: ", formattedDateTimeWithRFC3339)

	formattedDateTimeWithDateTimeConst := timeNow.Format(time.DateTime)
	fmt.Println("formattedDateTimeWithDateTimeConst is ::: ", formattedDateTimeWithDateTimeConst)

	/* Create date */

	birthday := time.Date(2024, 02, 13, 01, 15, 00, 0, time.UTC)
	fmt.Println("My birthday is at :::: ", birthday)

	/* Milliseconds */

	millis := time.Now().UnixMilli()
	fmt.Println("Milliseconds ::: ", millis)
}
