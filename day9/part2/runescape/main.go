package runescape

import (
	"bufio"
	"fmt"
	"io"
)

type Group struct {
	groups []*Group
	Score  int
}

func NewGroup() *Group {
	return &Group{
		groups: make([]*Group, 0),
		Score:  1,
	}
}

func (g *Group) AddGroup(nestedGroup *Group) {
	nestedGroup.Score = g.Score + 1
	g.groups = append(g.groups, nestedGroup)
}

func (g *Group) TotalScore() int {
	score := g.Score

	for _, childGroup := range g.groups {
		score += childGroup.TotalScore()
	}

	return score
}

type Stream struct {
	reader              *bufio.Reader
	stack               []*Group
	rootGroup           *Group
	insideGarbage       bool
	ignoreNextChar      bool
	removedGarbageCount int
}

func NewStream(r io.Reader) *Stream {
	reader := bufio.NewReader(r)

	return &Stream{
		reader:         reader,
		stack:          make([]*Group, 0),
		rootGroup:      nil,
		insideGarbage:  false,
		ignoreNextChar: false,
	}
}

func (s *Stream) Process() {
	for {
		char, _, err := s.reader.ReadRune()
		if err == io.EOF {
			break
		}

		if s.ignoreNextChar {
			s.ignoreNextChar = false
			continue
		}

		if !s.insideGarbage {
			switch char {
			case '{':
				s.pushGroup()
			case '}':
				s.popGroup()
			case '<':
				s.insideGarbage = true
			default:
				fmt.Println(fmt.Sprintf("Ignoring char %c\n", char))
			}

			continue
		}

		switch char {
		case '!':
			s.ignoreNextChar = true
		case '>':
			s.insideGarbage = false
		default:
			s.removedGarbageCount++
			fmt.Println(fmt.Sprintf("Ignoring char %c\n", char))
		}
	}
}

func (s *Stream) pushGroup() {
	g := NewGroup()

	if s.rootGroup == nil {
		s.rootGroup = g
	}

	if len(s.stack) > 0 {
		s.stack[len(s.stack)-1].AddGroup(g)
	}

	s.stack = append(s.stack, g)
}

func (s *Stream) popGroup() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stream) Score() int {
	return s.rootGroup.TotalScore()
}

func (s *Stream) RemovedGarbageCount() int {
	return s.removedGarbageCount
}
