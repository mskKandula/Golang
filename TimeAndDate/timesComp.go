package main

import (
	"fmt"
	"time"
)

func main() {
	// var timestamp int64
	// timestamp = 1596621423
	// after 6 hrs
	// 1 6 0 2 1 5 4 6 2 0
	// timee := time.Unix(8889498000, 10).Format("2006-January-02 3:4:5 pm")
	// fmt.Println(timee)

	// tim := time.Duration(1661945400000).Format("2006-January-02 3:4:5 pm")
	// fmt.Println(tim)

	// i, err := strconv.ParseInt("1661945400000", 10, 64)
	// if err != nil {
	// 	panic(err)
	// }
	// tm, _ := time.Parse("01-02-2006", "30-08-2022")

	tm := time.Unix(1661945400000, 1000)
	fmt.Println(tm)
	fmt.Println(tm.Before(time.Now()))

	// fmt.Println(tim.Sub(time.Now()))

}
