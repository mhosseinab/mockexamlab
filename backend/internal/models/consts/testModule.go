package consts

type TestModule int

const (
	IeltsAcademic TestModule = 0
	IeltsGeneral             = 1
)

func (t TestModule) ToString() string {
	switch t {
	case IeltsAcademic:
		return "Ielts Academic"
	case IeltsGeneral:
		return "IeltsGeneral"
	}
	return "-"
}
