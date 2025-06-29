package tasks

import (
	"fmt"
	"time"
)

func Task() {
	time.Sleep(time.Second * 10)
	fmt.Println("task")
}
