package virus

type CellStatus int

const (
	CLEAN CellStatus = iota + 1
	WEAKENED
	INFECTED
	FLAGGED
)

var cellStatuses = []CellStatus{
	CLEAN,
	WEAKENED,
	INFECTED,
	FLAGGED,
}

var cellNames = []string{
	".",
	"W",
	"#",
	"F",
}

func (c CellStatus) String() string {
	if c < CLEAN || c > FLAGGED {
		return "Unknown"
	}

	return cellNames[int(c)-1]
}

func (c CellStatus) Next() CellStatus {
	return cellStatuses[int(c)%len(cellStatuses)]
}
