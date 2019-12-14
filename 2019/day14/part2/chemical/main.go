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
	Coefficient int64
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
	coefficient, err := strconv.ParseInt(coefficientS, 10, 64)
	if err != nil {
		panic(err)
	}

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

func MinimumOreFromInput(in io.Reader) int64 {
	rbpn := ReactionsByProductNameFromInput(in)
	return MinimumOre(rbpn, 1)
}

func MaximumFuelPerTrillionOreFromInput(in io.Reader) int64 {
	return MaximumFuelPerAvailableOreFromInput(in, int64(1000000000000))
}

func MaximumFuelPerAvailableOreFromInput(in io.Reader, availableOre int64) int64 {
	rbpn := ReactionsByProductNameFromInput(in)
	return MaximumFuelPerAvailableOre(rbpn, availableOre)
}

func MinimumOre(rbpn ReactionsByProductName, targetFuel int64) int64 {
	requirements := make(map[string]int64)
	requirements["FUEL"] = targetFuel

	reacted := true
	for reacted {
		reacted = false

		for product, reqCoefficient := range requirements {
			if product == "ORE" || reqCoefficient <= 0 {
				continue
			}

			reacted = true
			reaction := rbpn[product]

			ratio := int64(math.Ceil(float64(reqCoefficient) / float64(reaction.Product.Coefficient)))

			productCoefficient := ratio * reaction.Product.Coefficient
			for _, reagent := range reaction.Reagents {
				requirements[reagent.Chemical] += ratio * reagent.Coefficient
			}

			requirements[product] -= productCoefficient
		}
	}

	return requirements["ORE"]
}

func MaximumFuelPerAvailableOre(rbpn ReactionsByProductName, availableOre int64) int64 {
	requiredOrePerFuelUnit := MinimumOre(rbpn, 1)

	// This is a lower bound, as accumulated waste along the reaction results in
	// a more efficient reaction overall.
	lowerBoundFuel := int64(math.Floor(float64(availableOre) / float64(requiredOrePerFuelUnit)))

	// Find an amount of fuel that requires more ore than availableOre;
	// adjust lowerBoundFuel along the way if possible.
	var upperBoundFuel int64
	for {
		fuel := lowerBoundFuel * 2
		requiredOre := MinimumOre(rbpn, fuel)
		if requiredOre > availableOre {
			upperBoundFuel = fuel
			break
		}
		lowerBoundFuel = fuel
	}

	// "Typical" binary search (with a somewhat modified stop condition).
	for upperBoundFuel-lowerBoundFuel > 1 {
		diff := upperBoundFuel - lowerBoundFuel
		fuel := lowerBoundFuel + diff/2
		requiredOre := MinimumOre(rbpn, fuel)

		if requiredOre > availableOre {
			upperBoundFuel = fuel
			continue

		} else if requiredOre < availableOre {
			lowerBoundFuel = fuel
			continue
		}

		// requiredOre is exactly availableOre
		return fuel
	}

	return lowerBoundFuel
}
