package consts

type PaddleAlert string

const (
	SubscriptionCreated          PaddleAlert = "subscription_created"
	SubscriptionUpdated                      = "subscription_updated"
	SubscriptionCancelled                    = "subscription_cancelled"
	SubscriptionPaymentSucceeded             = "subscription_payment_succeeded"
	SubscriptionPaymentFailed                = "subscription_payment_failed"
	SubscriptionPaymentRefunded              = "subscription_payment_refunded"
)
