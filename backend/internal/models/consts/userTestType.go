package consts

type UserTestType int

const (
	FullTest       UserTestType = 0
	IeltsSpeaking               = 1
	IeltsListening              = 2
	IeltsReading                = 3
	IeltsWriting                = 4
)

func (t UserTestType) ToString() string {
	switch t {
	case FullTest:
		return "Full Test"
	case IeltsSpeaking:
		return "Ielts Speaking"
	case IeltsListening:
		return "Ielts Listening"
	case IeltsReading:
		return "Ielts Reading"
	case IeltsWriting:
		return "Ielts Writing"
	}
	return "-"
}
