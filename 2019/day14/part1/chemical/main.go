package chemical

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type Molecule struct {
	Coefficient int
	Chemical    string
}

type Reaction struct {
	Reagents []Molecule
	Product  Molecule
}

type ReactionsByProductName map[string]*Reaction

func NewMoleculeFromString(s string) (Molecule, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return Molecule{}, fmt.Errorf("invalid chemical: %s", s)
	}

	coefficientS, name := parts[0], parts[1]
	coefficient64, err := strconv.ParseInt(coefficientS, 10, 32)
	if err != nil {
		panic(err)
	}

	coefficient := int(coefficient64)

	return Molecule{
		Coefficient: coefficient,
		Chemical:    name,
	}, nil
}

func ReactionsByProductNameFromInput(in io.Reader) ReactionsByProductName {
	rbpn := ReactionsByProductName{}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		reactionS := scanner.Text()
		reactionParts := strings.Split(reactionS, " => ")
		if len(reactionParts) != 2 {
			panic(fmt.Errorf("invalid input: %s", reactionS))
		}

		reagents := strings.Split(reactionParts[0], ", ")
		productS := reactionParts[1]

		reaction := &Reaction{}
		for _, reagentS := range reagents {
			reagent, err := NewMoleculeFromString(reagentS)
			if err != nil {
				panic(err)
			}

			reaction.Reagents = append(reaction.Reagents, reagent)
		}

		product, err := NewMoleculeFromString(productS)
		if err != nil {
			panic(err)
		}

		reaction.Product = product
		rbpn[product.Chemical] = reaction
	}

	return rbpn
}

func MinimumOreFromInput(in io.Reader) int {
	rbpn := ReactionsByProductNameFromInput(in)
	return MinimumOre(rbpn)
}

func MinimumOre(rbpn ReactionsByProductName) int {
	requirements := make(map[string]int)
	requirements["FUEL"] = 1

	reacted := true
	for reacted {
		reacted = false

		for product, reqCoefficient := range requirements {
			if product == "ORE" || reqCoefficient <= 0 {
				continue
			}

			reacted = true
			reaction := rbpn[product]

			ratio := int(math.Ceil(float64(reqCoefficient) / float64(reaction.Product.Coefficient)))

			productCoefficient := ratio * reaction.Product.Coefficient
			for _, reagent := range reaction.Reagents {
				requirements[reagent.Chemical] += ratio * reagent.Coefficient
			}

			requirements[product] -= productCoefficient
		}
	}

	return requirements["ORE"]
}
