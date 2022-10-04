package consts

type UserRole int

const (
	Basic UserRole = iota
	ExamModifier
	Marker
	Admin
)
