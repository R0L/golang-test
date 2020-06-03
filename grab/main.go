package main

import (
	"time"

	"github.com/R0L/golang-test/grab/grab"
)

func main() {

	userIds := []string{
		"R022005050000300033",
	}

	for _, userId := range userIds {
		i := 5
		for i > 0 {
			for {
				go grab.NewGrab(
					grab.WithKey("1234567890123456"),
					grab.WithUrl("http://144.34.167.130:8080"),
					grab.WithUserId(userId),
				).GetSceneInfo("M012006040000660005")
				time.Sleep(400 * time.Millisecond)
				i--
			}
		}
	}
	select {}

}
