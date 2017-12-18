package duet

import (
	"strconv"
	"strings"
)

type MusicPlayer struct {
	registers map[rune]int
	lastSound int
}

func NewMusicPlayer() *MusicPlayer {
	return &MusicPlayer{
		registers: make(map[rune]int, 0),
		lastSound: 0,
	}
}

func (m *MusicPlayer) ValueOf(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		register := rune(value[0])
		val = m.registers[register]
	}

	return val
}

func (m *MusicPlayer) Play(instructions []string) int {
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]

		instructionParts := strings.Split(instruction, " ")
		instrCode := instructionParts[0]

		X := instructionParts[1]

		Y := ""
		if len(instructionParts) > 2 {
			Y = instructionParts[2]
		}

		switch instrCode {
		case "snd":
			m.lastSound = m.ValueOf(X)
		case "set":
			m.registers[rune(X[0])] = m.ValueOf(Y)
		case "add":
			m.registers[rune(X[0])] += m.ValueOf(Y)
		case "mul":
			m.registers[rune(X[0])] *= m.ValueOf(Y)
		case "mod":
			m.registers[rune(X[0])] %= m.ValueOf(Y)
		case "rcv":
			if m.registers[rune(X[0])] != 0 {
				return m.lastSound
			}
		case "jgz":
			if m.registers[rune(X[0])] > 0 {
				i += m.ValueOf(Y) - 1
			}
		}
	}

	return 0
}
