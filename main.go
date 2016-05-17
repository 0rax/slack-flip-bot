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

var (
	flipTable     []string
	flopTable     []string
	flipFlopTable []string
)

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
	case "/flipflop":
		flipText = thing + flipFlopTable[rand.Intn(len(flipFlopTable))] + thing
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
	flipFlopTable = []string{
		" ︵╰(°□°）╯︵ ",
		"ミ┗(◉Д◉)┛彡",
		" ︵╰(ರ ~ ರ）╯︵ ",
		"ミ┗(ಸ_ಸ)┛彡",
		"ミ⎝(´･ω･`)⎠彡",
		"ミ┗((✧Д✧))┛彡",
	}
	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/flip", flip).Methods("GET")
	router.HandleFunc("/flop", flip).Methods("GET")
	router.HandleFunc("/flipflop", flip).Methods("GET")

	n.UseHandler(router)
	n.Run(":4242")
}
