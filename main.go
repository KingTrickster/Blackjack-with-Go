package main

import {
	"fmt",
	"https://github.com/KingTrickster/Blackjack-with-Go/deck"
}

func PrintIntro() {
	fmt.Println("21 Blackjack! By Alex Trejo/KingTrickster, 2024")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("<http://www.apache.org/licenses/>")
}

func main() {

	var d deck.StandardDeck
	d.Initialize()
	for i := 5; i > 0; i-- {
		fmt.Printf("First Card: %+v\n", d.Cards[0])
		fmt.Printf("length of deck: %d\n", len(d.Cards))
		fmt.Printf("Shuffing...\n")
		e := d.Shuffle()
		if e != nil {
			panic(e)
		}
	}

	hand, e := d.Draw(5)
	if !(e == nil) {
		panic(e)
	}

	fmt.Printf("\n\n\nDrawn Hand:")
	for _, c := range hand {
		fmt.Printf(" %s ", c.ToStr())
	}
	fmt.Printf("\n\n")

	fmt.Printf("Length of deck: %d\n", len(d.Cards))

	pair := false
	for i, c := range hand {
		for j := i + 1; j <= len(hand)-1; j++ {
			if c.Equal(&hand[j]) {
				pair = true
			}
		}
	}

	if pair {
		fmt.Printf("You have a pair!\n")
	} else {
		fmt.Printf("You don't have a pair.\n")
	}

	return
}