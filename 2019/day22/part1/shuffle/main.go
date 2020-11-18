package shuffle

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Deck struct {
	cards []int
}

func NewDeck(nCards int) *Deck {
	cards := make([]int, nCards)
	for i := 0; i < nCards; i++ {
		cards[i] = i
	}
	return &Deck{
		cards: cards,
	}
}

func (d *Deck) dealNewStack() {
	for i := len(d.cards)/2 - 1; i >= 0; i-- {
		d.cards[i], d.cards[len(d.cards)-i-1] = d.cards[len(d.cards)-i-1], d.cards[i]
	}
}

func (d *Deck) cut(n int) {
	for n < 0 {
		n += len(d.cards)
	}
	n = n % len(d.cards)

	newCards := make([]int, len(d.cards))
	for i := range d.cards {
		newCards[i] = d.cards[(n+i)%len(d.cards)]
	}

	for i := range d.cards {
		d.cards[i] = newCards[i]
	}
}

func (d *Deck) dealWithIncrement(increment int) {
	newCards := make([]int, len(d.cards))
	p := 0
	for i := range d.cards {
		newCards[p] = d.cards[i]
		p = (p + increment) % len(d.cards)
	}

	for i := range d.cards {
		d.cards[i] = newCards[i]
	}
}

func (d *Deck) Shuffle(input io.Reader) {
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		technique := sc.Text()
		tokens := strings.Split(technique, " ")

		switch {
		case len(tokens) == 2 && tokens[0] == "cut":
			n, err := strconv.ParseInt(tokens[1], 10, 64)
			if err != nil {
				panic(err)
			}

			d.cut(int(n))

		case len(tokens) == 4 && tokens[0] == "deal" && tokens[3] == "stack":
			d.dealNewStack()

		case len(tokens) == 4 && tokens[0] == "deal" && tokens[2] == "increment":
			increment, err := strconv.ParseInt(tokens[3], 10, 64)
			if err != nil {
				panic(err)
			}

			d.dealWithIncrement(int(increment))

		default:
			panic(fmt.Sprintf("unexpected input: %s", technique))
		}
	}
}

func (d *Deck) FindCard(card int) int {
	for i := range d.cards {
		if d.cards[i] == card {
			return i
		}
	}

	return -1
}

func (d *Deck) Cards() []int {
	cards := make([]int, len(d.cards))
	for i := range d.cards {
		cards[i] = d.cards[i]
	}
	return cards
}

func FindCardAfterShuffle(nCards int, input io.Reader, card int) int {
	deck := NewDeck(nCards)
	deck.Shuffle(input)

	return deck.FindCard(card)
}
