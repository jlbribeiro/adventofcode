package orbitmap

import (
	"bufio"
	"io"
	"strings"
)

type Object struct {
	Name           string
	InOrbitOf      *Object
	ObjectsInOrbit []*Object
}

type Objects map[string]*Object

func ObjectsFromInput(in io.Reader) Objects {
	var objects Objects = make(map[string]*Object)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		orbit := scanner.Text()
		orbitParts := strings.Split(orbit, ")")
		center, inOrbit := orbitParts[0], orbitParts[1]

		centerObj, ok := objects[center]
		if !ok {
			centerObj = &Object{
				Name: center,
			}
			objects[center] = centerObj
		}

		inOrbitObj, ok := objects[inOrbit]
		if !ok {
			inOrbitObj = &Object{
				Name: inOrbit,
			}
			objects[inOrbit] = inOrbitObj
		}

		inOrbitObj.InOrbitOf = centerObj
		centerObj.ObjectsInOrbit = append(centerObj.ObjectsInOrbit, inOrbitObj)
	}

	return objects
}

func (o Objects) CountOrbits() int {
	direct, indirect := o["COM"].CountOrbits(0)
	return direct + indirect
}

func (o *Object) CountOrbits(level int) (int, int) {
	direct := level
	indirect := 0
	for _, object := range o.ObjectsInOrbit {
		childDirect, childIndirect := object.CountOrbits(level + 1)
		indirect += childDirect + childIndirect
	}
	return direct, indirect
}

func (o Objects) OrbitalTransfersToSanta() int {
	youObj := o["YOU"]
	santaObj := o["SAN"]

	youPath := youObj.PathFromCOM()
	santaPath := santaObj.PathFromCOM()

	maxLen := len(youPath)
	if len(santaPath) < maxLen {
		maxLen = len(santaPath)
	}

	for i := 0; i < maxLen; i++ {
		if youPath[i] != santaPath[i] {
			return len(youPath) - i - 1 + len(santaPath) - i - 1
		}
	}

	// Edge case: one is orbiting (directly or indirectly) the other one
	// (i.e. one path is a substring of the other).
	a, b := len(youPath), len(santaPath)
	if a > b {
		a, b = b, a
	}

	return b - a
}

func (o *Object) PathFromCOM() []string {
	var path []string

	obj := o
	for obj != nil {
		path = append(path, obj.Name)
		obj = obj.InOrbitOf
	}

	// Reverse path (from o.Name, ..., COM to COM, ..., o.Name).
	for i := len(path)/2 - 1; i >= 0; i-- {
		opp := len(path) - 1 - i
		path[i], path[opp] = path[opp], path[i]
	}

	return path
}
