package espnsportsapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetScoreboard(sport string, league string, date string) string {
	resp, err := http.Get(fmt.Sprintf("http://site.api.espn.com/apis/site/v2/sports/%s/%s/scoreboard?date=%s", sport, league, date))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb
}
