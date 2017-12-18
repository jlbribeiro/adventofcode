package duet

import (
	"fmt"
	"strconv"
	"strings"
)

type Duet struct {
	MusicPlayer0 *MusicPlayer
	MusicPlayer1 *MusicPlayer
}

func NewDuet() *Duet {
	musicPlayer0 := NewMusicPlayer(0)
	musicPlayer1 := NewMusicPlayer(1)

	musicPlayer0.SendTo(musicPlayer1)
	musicPlayer1.SendTo(musicPlayer0)

	return &Duet{
		MusicPlayer0: musicPlayer0,
		MusicPlayer1: musicPlayer1,
	}
}

func (d *Duet) Play(instructions []string) int {
	lastIPointer0, lastIPointer1 := -1, -1

	for {
		curIPointer0 := d.MusicPlayer0.Play(instructions)
		curIPointer1 := d.MusicPlayer1.Play(instructions)

		if curIPointer0 == lastIPointer0 && curIPointer1 == lastIPointer1 {
			fmt.Println("Deadlock detected!")
			break
		}

		lastIPointer0, lastIPointer1 = curIPointer0, curIPointer1
	}

	return d.MusicPlayer1.SentCount
}

type MusicPlayer struct {
	id        int
	registers map[rune]int
	SentCount int
	queue     []int
	receiver  *MusicPlayer
	iPointer  int
}

func NewMusicPlayer(id int) *MusicPlayer {
	registers := make(map[rune]int, 0)
	registers['p'] = id

	queue := make([]int, 0)

	return &MusicPlayer{
		id:        id,
		registers: registers,
		SentCount: 0,
		queue:     queue,
		iPointer:  0,
	}
}

func (m *MusicPlayer) SendTo(o *MusicPlayer) {
	m.receiver = o
}

func (m *MusicPlayer) Send(val int) {
	fmt.Println(fmt.Sprintf("[%v](%v) About to send a value to receiver's queue.", m.id, m.iPointer))
	m.receiver.OnReceive(val)
	m.SentCount++
}

func (m *MusicPlayer) OnReceive(val int) {
	m.queue = append(m.queue, val)
	fmt.Println(fmt.Sprintf("[%v](%v) Received a value into the queue(len=%v)", m.id, m.iPointer, len(m.queue)))
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
	if m.iPointer < 0 || m.iPointer >= len(instructions) {
		return m.iPointer
	}

	instruction := instructions[m.iPointer]

	instructionParts := strings.Split(instruction, " ")
	instrCode := instructionParts[0]

	X := rune(instructionParts[1][0])

	Y := ""
	if len(instructionParts) > 2 {
		Y = instructionParts[2]
	}

	switch instrCode {
	case "snd":
		m.Send(m.ValueOf(string(X)))
	case "set":
		m.registers[X] = m.ValueOf(Y)
	case "add":
		m.registers[X] += m.ValueOf(Y)
	case "mul":
		m.registers[X] *= m.ValueOf(Y)
	case "mod":
		m.registers[X] %= m.ValueOf(Y)
	case "rcv":
		if len(m.queue) == 0 {
			fmt.Println(fmt.Sprintf("[%v](%v) Waiting for a value.", m.id, m.iPointer))
			return m.iPointer
		}

		fmt.Println(fmt.Sprintf("[%v](%v) About to read a value from the queue(len=%v).", m.id, m.iPointer, len(m.queue)))

		m.registers[X] = m.queue[0]
		m.queue = m.queue[1:]
	case "jgz":
		if m.ValueOf(string(X)) > 0 {
			m.iPointer += m.ValueOf(Y)
			return m.iPointer
		}
	}

	m.iPointer++
	return m.iPointer
}
