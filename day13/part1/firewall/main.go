package firewall

type Layer struct {
	Depth         int
	scanningRange int
}

func NewLayer(depth int, scanningRange int) *Layer {
	return &Layer{
		Depth:         depth,
		scanningRange: scanningRange,
	}
}

func (l *Layer) Detected(timeInstant int) bool {
	return timeInstant%(l.scanningRange*2-2) == 0
}

func (l *Layer) Severity() int {
	return l.Depth * l.scanningRange
}

type Firewall struct {
	layers []*Layer
}

func NewFirewall() *Firewall {
	return &Firewall{}
}

func (f *Firewall) AddLayer(l *Layer) {
	f.layers = append(f.layers, l)
}

func (f *Firewall) WalkThroughSeverity() int {
	severity := 0

	for _, layer := range f.layers {
		// It reaches layer with depth X on time instant X,
		// since it takes 1 picosecond to move one step (1 "depth unit").
		if layer.Detected(layer.Depth) {
			severity += layer.Severity()
		}
	}

	return severity
}
