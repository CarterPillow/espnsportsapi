package espnsportsapi

import (
	"io"
	"log"
	"net/http"
)

func get_scoreboard(sport string, league string, date string) string {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	return sb
}

func main() {
	sb := get_scoreboard("nfl", "nfl", "20190905")
	log.Println(sb)
}
