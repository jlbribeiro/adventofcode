package shuffle

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
)

// sumOfPowersMod calculates the sum of powers modulo mod.
// Reference: https://stackoverflow.com/a/31221590
func sumOfPowersMod(base int64, power int64, mod int64) int64 {
	if power&1 == 1 {
		// (base * sumOfPowersMod(base, power-1, mod) + 1) % mod

		sum := sumOfPowersMod(base, power-1, mod) // partial_sum (see below)

		result := big.NewInt(sum)
		result.Mul(result, big.NewInt(base))     //  base * partial_sum
		result.Add(result, big.NewInt(int64(1))) //  base * partial_sum + 1
		result.Mod(result, big.NewInt(mod))      // (base * partial_sum + 1) % mod

		if !result.IsInt64() {
			panic("result should be int64!")
		}

		return result.Int64()

	} else if power > 0 {
		// ((base + 1) * sumOfPowersMod(base * base % mod, power // 2, mod)) % mod
		pb := big.NewInt(mod)

		result := big.NewInt(base)
		result.Mul(result, big.NewInt(base)) // base * base
		result.Mod(result, pb)               // base * base % mod

		if !result.IsInt64() {
			panic("result should be int64!")
		}

		halfN := power / 2
		sum := sumOfPowersMod(result.Int64(), halfN, mod)

		result.SetInt64(sum)                   // partial_sum (see below)
		result.Mul(result, big.NewInt(base+1)) // (base + 1) * partial_sum
		result.Mod(result, pb)                 // ((base + 1) * partial_sum) % mod

		if !result.IsInt64() {
			panic("result should be int64!")
		}

		return result.Int64()
	}

	return 0
}

// offsetN calculates the offset after nIterations.
// The iterative/recursive version of the function can be transformed into a
// closed form (which involves a partial sum of a geometric series).
//
// E.g. deriving it by induction (factor[0]=1134, offset[0]=4518):
//     offset[3] = 5815077138 * 1134 + 4518
//               = 5815077138 * 1134 + 4518
//               = ((1134 * 4518 + 4518) * 1134 + 4518) * 1134 + 4518
//               = ((a * b + b) * a + b) * a + b
//               = a * a * a * b + a * a * b + a * b + b
//               = b (a * a * a + a * a + a + 1)
//               = b (a^3 + a^2 + a^1 + a^0)
//
// Reference (for the closed form of a partial sum of a geometric series):
// https://math.stackexchange.com/a/971770
func offsetN(factor0 int64, offset0 int64, nIterations int64, nCards int64) int64 {
	// offset[n] = (offset[n-1] * factor[0] + offset[0]) % d.nCards
	// or
	// offset[n] = offset[0] * (factor[0]^n + factor[0]^(n - 1) + factor[0]^(n - 2) + ... + factor[0]^0)
	//           = offset[0] * ((factor[0]^(n + 1) - 1) / (factor[0] - 1))
	sum := sumOfPowersMod(factor0, nIterations, nCards)
	result := big.NewInt(offset0)
	result.Mul(result, big.NewInt(sum))
	result.Mod(result, big.NewInt(nCards))

	if !result.IsInt64() {
		panic("result should be int64!")
	}

	return result.Int64()
}

type DeckShuffler interface {
	dealNewStack()
	cut(n int64)
	dealWithIncrement(increment int64)
	Shuffle(input io.Reader, nTimes int64)
}

var _ DeckShuffler = &DeckRefImpl{}
var _ DeckShuffler = &Deck{}

type Deck struct {
	nCards int64
	factor int64
	offset int64
	signal int64
}

func NewDeck(nCards int64) *Deck {
	return &Deck{
		nCards: nCards,
		factor: 1,
		offset: 0,
		signal: 1,
	}
}

func (d *Deck) dealNewStack() {
	d.signal = -d.signal
	d.offset = -d.offset
	for d.offset < 0 {
		d.offset += d.nCards
	}
	d.offset %= d.nCards
}

func (d *Deck) cut(n int64) {
	d.offset += n
	for d.offset < 0 {
		d.offset += d.nCards
	}
	d.offset %= d.nCards
}

func (d *Deck) dealWithIncrement(increment int64) {
	d.factor = (d.factor * increment) % d.nCards

	d.offset *= increment
	if d.signal < 0 {
		d.offset += increment - 1
	}

	for d.offset < 0 {
		d.offset += d.nCards
	}
	d.offset %= d.nCards
}

func (d *Deck) Shuffle(input io.Reader, nTimes int64) {
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
			panic(fmt.Sprintf("unexpected input: %s", technique))
		}
	}

	factor0 := d.factor
	offset0 := d.offset

	// factor can be calculated using a closed formula (pseudo-code):
	//   pow(factor0, nTimes) % nCards
	//
	// That is the equivalent to Golang's
	//   (big.Int).Exp(factor0, nTimes, nCards).
	//
	//
	// offsetN() calculates the offset using a sum of powers formula
	// (and contains the equation in the comments).

	// This is the iterative version.
	// for i := int64(1); i < nTimes; i++ {
	// 	d.factor = (d.factor * factor0) % d.nCards
	// 	d.offset = (d.offset*factor0 + offset0) % d.nCards
	// }

	if nTimes > 1 {
		if d.signal < 0 {
			panic("not able to fast-forward inverted stacks")
		}

		res := big.NewInt(factor0)
		res.Exp(res, big.NewInt(nTimes), big.NewInt(d.nCards))
		if !res.IsInt64() {
			panic("result should fit in an int64")
		}

		d.factor = res.Int64()
		d.offset = offsetN(factor0, offset0, nTimes, d.nCards)
	}
}

func (d *Deck) FindCard(card int64) int64 {
	if card < 0 || card >= d.nCards {
		return -1
	}

	ind := (d.signal*d.factor*card - d.offset)
	if d.signal < 0 {
		ind--
	}
	for ind < 0 {
		ind += d.nCards
	}
	ind %= d.nCards
	return ind
}

func (d *Deck) Get(index int64) int64 {
	// The formula used in FindCard for calculating the index is
	//   ind = (factor * card - offset) % nCards
	//
	// This is equivalent to (assume `x` is the card number from here on)
	//   factor * x = index + offset (mod nCards)
	//
	// e.g. (factor=21, offset=19, nCards=53; with index=3)
	//   21x - 19 = 3 (mod 53)
	//   21x      = 22 (mod 53)
	//
	// Solving this linear congruence is a matter of calculating
	// the modular inverse between 21 and 53 (using the example above).
	// Remember that nCards is known to be prime.
	// e.g. ModInverse(21, 53) = 48
	// So multiplying 48 on both sides yields
	//   x = 48 * 22 (mod 53)
	//     = 1056 (mod 53)
	//     = 49
	// 49 is the card at index 3, for factor 21, offset 19.

	factorb := big.NewInt(d.factor)
	indexb := big.NewInt(index)
	offsetb := big.NewInt(d.offset)
	modb := big.NewInt(d.nCards)

	res := new(big.Int)
	res.Add(indexb, offsetb)

	k := new(big.Int)
	k.ModInverse(factorb, modb)

	res.Mul(res, k)
	res.Mod(res, modb)

	if !res.IsInt64() {
		panic("expected to be int64")
	}

	return res.Int64()
}

func (d *Deck) Cards() []int64 {
	cards := make([]int64, d.nCards)
	for i := range cards {
		cards[d.FindCard(int64(i))] = int64(i)
	}
	return cards
}

func FindCardAfterShuffleUsingReference(nCards int64, input io.Reader, nTimes int64, card int64) int64 {
	deck := NewDeckRefImpl(nCards)
	deck.Shuffle(input, nTimes)

	for i := range deck.cards {
		if deck.cards[i] == card {
			return int64(i)
		}
	}

	return -1
}

func FindCardAfterShuffleFast(nCards int64, input io.Reader, nTimes int64, card int64) int64 {
	deck := NewDeck(nCards)
	deck.Shuffle(input, nTimes)
	return deck.FindCard(card)
}

func GetAfterShuffle(nCards int64, input io.Reader, nTimes int64, index int64) int64 {
	deck := NewDeck(nCards)
	deck.Shuffle(input, nTimes)
	return deck.Get(index)
}
