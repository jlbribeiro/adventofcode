package shuffle

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type DeckRefImpl struct {
	cards []int64
}

func NewDeckRefImpl(nCards int64) *DeckRefImpl {
	cards := make([]int64, nCards)
	for i := int64(0); i < nCards; i++ {
		cards[i] = i
	}

	return &DeckRefImpl{
		cards: cards,
	}
}

func (d *DeckRefImpl) dealNewStack() {
	for i := len(d.cards)/2 - 1; i >= 0; i-- {
		d.cards[i], d.cards[len(d.cards)-i-1] = d.cards[len(d.cards)-i-1], d.cards[i]
	}
}

func (d *DeckRefImpl) cut(n int64) {
	for n < 0 {
		n += int64(len(d.cards))
	}
	n = n % int64(len(d.cards))

	newCards := make([]int64, len(d.cards))
	for i := range d.cards {
		newCards[i] = d.cards[(n+int64(i))%int64(len(d.cards))]
	}

	for i := range d.cards {
		d.cards[i] = newCards[i]
	}
}

func (d *DeckRefImpl) dealWithIncrement(increment int64) {
	newCards := make([]int64, len(d.cards))
	p := int64(0)
	for i := range d.cards {
		newCards[p] = d.cards[i]
		p = (p + increment) % int64(len(d.cards))
	}

	for i := range d.cards {
		d.cards[i] = newCards[i]
	}
}

func (d *DeckRefImpl) Shuffle(input io.Reader, nTimes int64) {
	sc := bufio.NewScanner(input)
	var instructions [][]string
	for sc.Scan() {
		technique := sc.Text()
		tokens := strings.Split(technique, " ")
		instructions = append(instructions, tokens)
	}

	for i := int64(0); i < nTimes; i++ {
		for _, tokens := range instructions {
			switch {
			case len(tokens) == 2 && tokens[0] == "cut":
				n, err := strconv.ParseInt(tokens[1], 10, 64)
				if err != nil {
					panic(err)
				}

				d.cut(n)

			case len(tokens) == 4 && tokens[0] == "deal" && tokens[3] == "stack":
				d.dealNewStack()

			case len(tokens) == 4 && tokens[0] == "deal" && tokens[2] == "increment":
				increment, err := strconv.ParseInt(tokens[3], 10, 64)
				if err != nil {
					panic(err)
				}

				d.dealWithIncrement(increment)

			default:
				panic(fmt.Sprintf("unexpected input: %v", tokens))
			}
		}
	}
}

func (d *DeckRefImpl) Cards() []int64 {
	return d.cards
}
