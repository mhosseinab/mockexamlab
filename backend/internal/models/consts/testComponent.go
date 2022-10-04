package consts

type TestComponent int

const (
	Speaking  TestComponent = 0
	Listening               = 1
	Reading                 = 2
	Writing                 = 3
)

func (t TestComponent) ToString() string {
	switch t {
	case Speaking:
		return "Speaking"
	case Listening:
		return "Listening"
	case Reading:
		return "Reading"
	case Writing:
		return "Writing"
	}
	return "-"
}
