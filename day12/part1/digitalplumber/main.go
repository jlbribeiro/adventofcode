package digitalplumber

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Program struct {
	ID                  int
	ConnectedProgramIDs []int
}

func NewProgramFromString(input string) *Program {
	regex, err := regexp.Compile(`(\d+) <-> ((\d+)(, \d+)*)`)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	matches := regex.FindStringSubmatch(input)
	programID64, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		panic(err)
	}
	programID := int(programID64)

	connectedProgramIDsStr := strings.Split(matches[2], ", ")
	connectedProgramIDs := make([]int, len(connectedProgramIDsStr))

	for i, progStr := range connectedProgramIDsStr {
		connectedProgramID, err := strconv.ParseInt(progStr, 10, 64)
		if err != nil {
			panic(err)
		}

		connectedProgramIDs[i] = int(connectedProgramID)
	}

	return &Program{
		ID:                  programID,
		ConnectedProgramIDs: connectedProgramIDs,
	}
}

type ProgramNetwork struct {
	programs    []*Program
	adjacencies [][]bool
	flooded     bool
}

func NewProgramNetworkFromInput(input string) *ProgramNetwork {
	r := strings.NewReader(input)
	return NewProgramNetworkFromReader(r)
}

func NewProgramNetworkFromReader(r io.Reader) *ProgramNetwork {
	scanner := bufio.NewScanner(r)

	programs := []*Program(nil)

	for scanner.Scan() {
		programStr := scanner.Text()

		program := NewProgramFromString(programStr)
		programs = append(programs, program)
	}

	adjacencies := make([][]bool, len(programs))
	for i := range adjacencies {
		adjacencies[i] = make([]bool, len(adjacencies))
	}

	return &ProgramNetwork{
		programs:    programs,
		adjacencies: adjacencies,
		flooded:     false,
	}
}

func (pn *ProgramNetwork) flood() {
	for _, program := range pn.programs {
		a := program.ID
		pn.adjacencies[a][a] = true
		for _, connectedProgramID := range program.ConnectedProgramIDs {
			b := connectedProgramID
			pn.adjacencies[a][b] = true
			pn.adjacencies[b][a] = true
		}
	}

	// Floyd-Warshall
	for k := 0; k < len(pn.adjacencies); k++ {
		for i := 0; i < len(pn.adjacencies); i++ {
			for j := 0; j < len(pn.adjacencies); j++ {
				pn.adjacencies[i][j] = pn.adjacencies[i][j] || pn.adjacencies[i][k] && pn.adjacencies[k][j]
			}
		}
	}
}

func (pn *ProgramNetwork) NConnectionsOf(node int) int {
	if !pn.flooded {
		pn.flood()
		pn.flooded = true
	}

	count := 0

	for i := 0; i < len(pn.adjacencies); i++ {
		if pn.adjacencies[node][i] {
			count++
		}
	}

	return count
}
