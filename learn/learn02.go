package learn

import "time"

func GoroutineTest(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person
}
