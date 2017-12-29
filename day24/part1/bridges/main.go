package bridges

import (
	"fmt"
	"strconv"
	"strings"
)

type Component struct {
	Ports []int
	inUse []bool
	score int
}

func NewComponent(ports []int) *Component {
	score := 0
	inUse := make([]bool, len(ports))

	for _, port := range ports {
		score += port
	}

	return &Component{
		Ports: ports,
		inUse: inUse,
		score: score,
	}
}

func NewComponentFromString(component string) *Component {
	portsS := strings.Split(component, "/")

	ports := make([]int, len(portsS))
	for i, portS := range portsS {
		port, err := strconv.Atoi(portS)
		if err != nil {
			panic(fmt.Errorf("unexpected port on component: %v", component))
		}

		ports[i] = port
	}

	return NewComponent(ports)
}

func (c *Component) String() string {
	portsS := make([]string, len(c.Ports))
	for i := range c.Ports {
		var state [2]string
		if c.inUse[i] {
			state = [2]string{"(", ")"}
		} else {
			state = [2]string{"", ""}
		}

		portsS[i] = fmt.Sprintf("%s%v%s", state[0], c.Ports[i], state[1])
	}

	return strings.Join(portsS, "/")
}

func (c *Component) HasFreePort(port int) bool {
	for i := range c.Ports {
		if c.Ports[i] == port && !c.inUse[i] {
			return true
		}
	}

	return false
}

func (c *Component) AttachPort(port int) error {
	skippedBecauseInUse := false
	for i, portPins := range c.Ports {
		if portPins == port {
			if c.inUse[i] {
				skippedBecauseInUse = true
				continue
			}

			c.inUse[i] = true
			return nil
		}
	}

	if skippedBecauseInUse {
		return fmt.Errorf("already in use: %v", port)
	}

	return fmt.Errorf("no such port: %v", port)
}

func (c *Component) DettachPort(port int) error {
	skippedBecauseNotInUse := false
	for i, portPins := range c.Ports {
		if portPins == port {
			if !c.inUse[i] {
				skippedBecauseNotInUse = true
				continue
			}

			c.inUse[i] = false
			return nil
		}
	}

	if skippedBecauseNotInUse {
		return fmt.Errorf("not in use: %v", port)
	}

	return fmt.Errorf("no such port: %v", port)
}

func (c *Component) FreeEnd() (int, error) {
	port := -1

	for i := range c.Ports {
		if !c.inUse[i] {
			if port != -1 {
				return -1, fmt.Errorf("multiple ports available")
			}

			port = c.Ports[i]
		}
	}

	return port, nil
}

func (c *Component) Score() int {
	return c.score
}

type BridgeBuilder struct {
	nComponents      int
	ComponentsByPort map[int][]*Component
}

func NewBridgeBuilder(componentsS []string) *BridgeBuilder {
	nComponents := 0
	byPort := make(map[int][]*Component, 0)

	for _, componentS := range componentsS {
		component := NewComponentFromString(componentS)
		nComponents++
		for _, port := range component.Ports {
			byPort[port] = append(byPort[port], component)
		}
	}

	return &BridgeBuilder{
		nComponents:      nComponents,
		ComponentsByPort: byPort,
	}
}

func (bb *BridgeBuilder) StrongestBridge() int {
	startComponents := bb.ComponentsByPort[0]

	strongestScore := 0
	currentScore := 0
	for _, startComponent := range startComponents {
		err := startComponent.AttachPort(0)
		if err != nil {
			panic(err)
		}

		currentScore = startComponent.Score()
		err = bb.buildBridge(startComponent, &currentScore, &strongestScore)
		if err != nil {
			panic(err)
		}

		err = startComponent.DettachPort(0)
		if err != nil {
			panic(err)
		}
	}

	return strongestScore
}

func (bb *BridgeBuilder) buildBridge(lastComponent *Component, currentScore *int, strongestScore *int) error {
	if *currentScore > *strongestScore {
		*strongestScore = *currentScore
	}

	endPins, err := lastComponent.FreeEnd()
	if err != nil {
		return err
	}

	lastComponent.AttachPort(endPins)

	nextComponents := bb.ComponentsByPort[endPins]
	for _, component := range nextComponents {
		if !component.HasFreePort(endPins) {
			continue
		}
		component.AttachPort(endPins)

		*currentScore += component.Score()
		err := bb.buildBridge(component, currentScore, strongestScore)
		if err != nil {
			return err
		}
		*currentScore -= component.Score()

		component.DettachPort(endPins)
	}

	lastComponent.DettachPort(endPins)

	return nil
}
