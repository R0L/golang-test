package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/R0L/golang-test/grab/grab"
)

func main() {

	userIds := []string{
		"R022005050000300034",
		"R022005170000440021",
		"R022005120000370013",
		"R022005060000310036",
		"R022005100000350017",
	}

	for _, userId := range userIds {
		ip := genIpaddr()
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

	// UserIds := []string{
	// 	"R022005050000300034",
	// 	"R022005170000440021",
	// 	"R022005120000370013",
	// 	"R022005060000310036",
	// 	"R022005100000350017",
	// }
	//
	// for _, userId := range UserIds {
	// 	xToken, _ := grab.NewGrab(
	// 		grab.WithKey("1234567890123456"),
	// 		grab.WithUrl("http://144.34.167.130:8080"),
	// 		grab.WithUserId(userId)).GetXToken()
	//
	// 	fmt.Printf("\"%s\",\n", xToken)
	// }

}

func genIpaddr() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}
