package consts

type QuestionType int

const (
	MultipleChoice QuestionType = iota
	IdentifyingInformation
	Map
	NoteCompletion
	SpeakingTopic
	WritingTopic
)

func (t QuestionType) ToString() string {
	switch t {
	case MultipleChoice:
		return "Multiple Choice"
	case IdentifyingInformation:
		return "Identifying Information"
	case Map:
		return "Map"
	case NoteCompletion:
		return "Note Completion"
	case SpeakingTopic:
		return "Speaking Topic"
	case WritingTopic:
		return "Writing Topic"
	}
	return "-"
}
