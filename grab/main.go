package main

import (
	"math/rand"
	"time"

	"github.com/R0L/golang-test/grab/grab"
)

func main() {

	userIdIps := map[string]string{
		"R022005050000300034": "192.168.31.105",
		"R022005170000440021": "192.168.31.106",
		"R022005120000370013": "10.0.1.107",
		"R022005060000310036": "192.168.31.108",
		"R022005100000350017": "192.168.31.109",
	}

	for userId, ip := range userIdIps {
		go func(userId, ip string) {
			i := 5
			for i > 0 {
				for {
					go grab.NewGrab(
						grab.WithKey("1234567890123456"),
						grab.WithUrl("http://144.34.167.130:8080"),
						grab.WithUserId(userId),
						grab.WithIp(ip),
					).GetSceneInfo("M012006040000660007")
					time.Sleep(time.Duration(300+rand.Intn(500)) * time.Millisecond)
					i--
				}
			}
		}(userId, ip)
	}
	select {}

}
