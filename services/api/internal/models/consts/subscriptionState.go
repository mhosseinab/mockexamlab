package consts

type SubscriptionState int

const (
	Canceled SubscriptionState = iota
	Expired
	NonSubscriber
	Subscriber
	Refund
	Paused
)

func (t SubscriptionState) ToString() string {
	switch t {
	case Canceled:
		return "Canceled"
	case Expired:
		return "Expired"
	case NonSubscriber:
		return "Non Subscriber"
	case Subscriber:
		return "Subscriber"
	case Refund:
		return "Refund"
	case Paused:
		return "Paused"
	}
	return "-"
}
