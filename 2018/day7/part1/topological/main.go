package topological

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"unicode/utf8"
)

var InstructionRegex = regexp.MustCompile(`Step (.*) must be finished before step (.*) can begin.`)

func Order(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)

	requisites := make(map[int]map[int]struct{})
	nodes := make(map[int]struct{})
	for scanner.Scan() {
		instruction := scanner.Text()
		parts := InstructionRegex.FindStringSubmatch(instruction)

		a, _ := utf8.DecodeLastRuneInString(parts[1])
		b, _ := utf8.DecodeLastRuneInString(parts[2])

		requisite := int(a - 'A')
		dependent := int(b - 'A')

		if _, ok := requisites[dependent]; !ok {
			requisites[dependent] = make(map[int]struct{})
		}

		requisites[dependent][requisite] = struct{}{}
		nodes[requisite] = struct{}{}
		nodes[dependent] = struct{}{}
	}

	for node := range nodes {
		if _, ok := requisites[node]; !ok {
			requisites[node] = make(map[int]struct{})
		}
	}

	nNodes := len(nodes)

	var sequence bytes.Buffer
	for sequence.Len() < nNodes {
		minRequisites := nNodes
		nextNode := nNodes

		for node := range nodes {
			nRequisites := len(requisites[node])

			if nRequisites > minRequisites {
				continue
			}

			if nRequisites == minRequisites && node > nextNode {
				continue
			}

			minRequisites = nRequisites
			nextNode = node
		}

		for node := range nodes {
			delete(requisites[node], nextNode)
		}
		delete(nodes, nextNode)

		sequence.WriteRune(rune('A' + nextNode))
	}

	return sequence.String()
}
