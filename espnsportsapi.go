package espnsportsapi

import (
	"io"
	"log"
	"net/http"
)

func GetScoreboard(sport string, league string, date string) string {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
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
