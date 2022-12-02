package main

import "fmt"

type Points struct {
	 lnt int
	 lag int
}

func main() {
	Points start
	start.lnt = 0
	start.lag = 0
	fmt.Println(start.lnt, start.lag)
}
