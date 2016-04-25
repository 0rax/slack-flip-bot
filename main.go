package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	answer, _ := json.Marshal(slackReturn{"in_channel", flipTable[rand.Intn(len(flipTable))] + thing})
	fmt.Fprintf(w, string(answer))
}

func flop(w http.ResponseWriter, r *http.Request) {

	thing := "┻━┻"
	if text, ok := r.URL.Query()["text"]; ok && text[0] != "" {
		thing = text[0]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	answer, _ := json.Marshal(slackReturn{"in_channel", thing + flopTable[rand.Intn(len(flopTable))]})
	fmt.Fprintf(w, string(answer))
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
	router.HandleFunc("/flop", flop).Methods("GET")

	n.UseHandler(router)
	n.Run(":4242")
}
