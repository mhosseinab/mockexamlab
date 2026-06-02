package consts

type PaddleStatus string

const (
	Active       PaddleStatus = "active "
	Trialing                  = "trialing"
	PastDue                   = "past_due"
	PaddlePaused              = "paused"
	Deleted                   = "deleted"
)
