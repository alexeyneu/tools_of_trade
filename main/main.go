package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/alexeyneu/rino2/on_green"
	)

type bomb struct {
	PrivateKey string `json:"privateKey"`
	Chain      string `json:"chain"`     
	Testnet    bool   `json:"testnet"`
}

type merv map[string]bomb
type sierra map[string]string

func main() {

	var tetr merv
	fast := make(sierra)
	data, _:= os.ReadFile("western.json")
	json.Unmarshal(data, &tetr)
	for id,bomb_m := range tetr {
		tyd := on_green.Make_from(bomb_m.PrivateKey)
		fast[tyd] = id
	}
	f,_ := json.Marshal(fast)
	fmt.Println(string(f))
	os.WriteFile("western.json", f, 0777)
}
