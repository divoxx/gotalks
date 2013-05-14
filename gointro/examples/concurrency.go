package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	menu   = []string{"Earthquake", "Corpse Reviver #2", "Irish Car Bomb", "Orgasm", "Caipirinha", "Gin Sour"}
	doxer  = []string{"Body", "Gobhi", "Chris", "Dan", "Nil", "Shannon", "Jey", "Ellis"}
	drinks = make(chan string)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Bartender() {
	size := len(menu)
	for {
		drinks <- menu[rand.Intn(size)]
		time.Sleep(500 * time.Millisecond)
	}
}

func Doxer(name string) {
	for drink := range drinks {
		fmt.Printf("%s drank %s\n", name, drink)
	}
}

func main() {
	go Bartender()
	go Bartender()

	for _, d := range doxer {
		go Doxer(d)
	}

	time.Sleep(10 * time.Second)
}
