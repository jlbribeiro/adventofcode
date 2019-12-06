package orbitmap

import (
	"bufio"
	"io"
	"strings"
)

type Object struct {
	Name           string
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
