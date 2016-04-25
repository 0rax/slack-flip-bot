package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type slackReturn struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

var flipTable []string
var flopTable []string

func flip(w http.ResponseWriter, r *http.Request) {

	thing := "┻━┻"
	if text, ok := r.URL.Query()["text"]; ok && text[0] != "" {
		thing = text[0]
	}

	var flipText string
	switch r.URL.EscapedPath() {
	case "/flip":
		flipText = flipTable[rand.Intn(len(flipTable))] + thing
	case "/flop":
		flipText = thing + flopTable[rand.Intn(len(flopTable))]
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if strings.Contains(r.Header.Get("Accept"), "application/json") == true {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(slackReturn{"in_channel", flipText})
		fmt.Fprintf(w, string(response))
	} else {
		fmt.Fprintf(w, flipText)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	flipTable = []string{
		"(╯°□°）╯︵ ",
		"(┛◉Д◉)┛彡",
		"(ﾉ≧∇≦)ﾉ ﾐ ",
		"(ノಠ益ಠ)ノ彡",
		"(╯ರ ~ ರ）╯︵ ",
		"(┛ಸ_ಸ)┛彡",
		"(ﾉ´･ω･)ﾉ ﾐ ",
		"(ノಥ,_｣ಥ)ノ彡",
		"(┛✧Д✧))┛彡",
	}
	flopTable = []string{
		" ︵╰(°□°╰)",
		"ミ┗(◉Д◉┗)",
		" ﾐヽ(≧∇≦ヽ)",
		"ミヾ(ಠ益ಠヾ)",
		" ︵╰(ರ ~ ರ╰)",
		"ミ┗(ಸ_ಸ┗)",
		" ︵ ㇸ(･ω･´ㇸ)",
		"ミ∖(ಥ,_,ಥ∖)",
		"ミ┗((✧Д✧┗)",
	}

	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/flip", flip).Methods("GET")
	router.HandleFunc("/flop", flip).Methods("GET")

	n.UseHandler(router)
	n.Run(":4242")
}
