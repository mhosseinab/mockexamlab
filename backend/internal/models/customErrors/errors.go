package customErrors

import "MockExamLab/internal/models/consts"

type SubscriptionError struct {
	Status consts.SubscriptionState
}

func (m *SubscriptionError) Error() string {
	switch m.Status {
	case consts.Canceled:
		return "canceled"
	case consts.Expired:
		return "expired"
	case consts.Paused:
		return "Paused"
	case consts.Refund:
		return "Refund"
	default:
		return ""
	}
}
