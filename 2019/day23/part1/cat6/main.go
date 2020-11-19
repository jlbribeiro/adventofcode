package cat6

import (
	"fmt"
	"io"

	"github.com/jlbribeiro/adventofcode/2019/day23/part1/intcode"
)

type Packet struct {
	destIP int
	x      int64
	y      int64
}

type Computer struct {
	cpu *intcode.CPU
	ip  int

	incoming []Packet
}

func (c *Computer) Boot(program []int64, ip int) error {
	c.cpu = intcode.NewCPU(program)
	c.ip = ip

	output, waitingInput := c.cpu.Exec([]int64{int64(ip)})
	if len(output) > 0 || !waitingInput {
		return fmt.Errorf("unexpected state: [ip:%2d] %v (waitingInput: %v)", ip, output, waitingInput)
	}

	return nil
}

func (c *Computer) Exchange() []Packet {
	var outgoing []Packet

	var output []int64
	if len(c.incoming) == 0 {
		output, _ = c.cpu.Exec([]int64{-1})
	} else {
		var input []int64
		for len(c.incoming) > 0 {
			var packet Packet
			packet, c.incoming = c.incoming[0], c.incoming[1:]
			input = append(input, []int64{packet.x, packet.y}...)
		}

		output, _ = c.cpu.Exec(input)
	}

	for len(output) > 0 {
		var wirePacket []int64
		wirePacket, output = output[:3], output[3:]
		if len(wirePacket) != 3 {
			panic(fmt.Sprintf("unexpected output from computer: %v", wirePacket))
		}

		outgoing = append(outgoing, Packet{
			destIP: int(wirePacket[0]),
			x:      wirePacket[1],
			y:      wirePacket[2],
		})
	}

	return outgoing
}

func (c *Computer) EnqueuePacket(packet Packet) {
	c.incoming = append(c.incoming, packet)
}

func FirstPacketTo(input io.Reader, nComputers int, targetIP int) int64 {
	debug := false

	program := intcode.ProgramFromInput(input)

	computers := make([]*Computer, nComputers)
	for i := range computers {
		computers[i] = &Computer{}

		ip := i
		if err := computers[i].Boot(program, ip); err != nil {
			panic(err)
		}
	}

	for {
		for _, computer := range computers {
			incomingPackets := len(computer.incoming)
			if debug && incomingPackets > 0 {
				fmt.Printf("[%2d] recv %d packets\n", computer.ip, incomingPackets)
			}
			packets := computer.Exchange()
			if debug && len(packets) > 0 {
				fmt.Printf("[%2d] send %d packets\n", computer.ip, len(packets))
			}

			for _, packet := range packets {
				if packet.destIP == targetIP {
					return packet.y
				}

				if packet.destIP < 0 || packet.destIP >= nComputers {
					panic(fmt.Sprintf("packet out of network range: %d", packet.destIP))
				}

				computers[packet.destIP].EnqueuePacket(packet)
			}
		}
	}
}
