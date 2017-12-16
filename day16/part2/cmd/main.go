package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jlbribeiro/adventofcode/day16/part2/promenade"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	steps := strings.Split(strings.TrimSpace(string(input)), ",")

	dancers := promenade.NewDancers(16)
	initialAlignment := dancers.Alignment()

	// We want to dance 1 billion dances.
	nDances := 1000000000

	loopFound := false
	for i := 1; i <= nDances; i++ {
		dancers.Dance(steps)
		if !loopFound && dancers.AlignmentEqualTo(initialAlignment) {
			loopFound = true

			// We get to the initialAlignment after i dances;
			// let's make this promenade quick then.
			nDances %= i

			// Reset i so that it becomes 1 (the initial value) in the next
			// iteration, since we've just changed the value used in the
			// condition.
			i = 0
		}
	}

	fmt.Println(dancers)
}
