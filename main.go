package main

import (
	"fmt"
	"time"
)

func main() {
	//var bar progressBar.Bar
	//bar.NewOption(0,100)
	//for i := 0; i <=100; i++ {
	//	time.Sleep(100*time.Millisecond)
	//	bar.Play(int64(i))
	//}
	//bar.Finish()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("hhh\r%d",i)
	}
}
