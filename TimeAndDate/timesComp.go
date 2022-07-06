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
	timee := time.Unix(8889498000, 10).Format("2006-January-02 3:4:5 pm")
	fmt.Println(timee)

}
