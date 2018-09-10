package main

import (
	"fmt"

	delay109 "github.com/nukosuke/go-delay109"
)

func main() {
	client := delay109.New(nil)
	status, _ := client.Search(&delay109.Query{
		Line:       delay109.DT,
		Year:       2018,
		Month:      9,
		Day:        4,
		Direction:  delay109.DirectionUp,
		TimePeriod: delay109.Before10am,
	})

	fmt.Println(status.LineNameJa())
	fmt.Println(status.DirectionText(), "方面")
	fmt.Println(status.DelayDuration().Minutes(), "分遅延")
}
