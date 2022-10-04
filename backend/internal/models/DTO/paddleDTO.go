package DTO

type BasePaddleResponse struct {
	AlertId            string `json:"alert_id"`
	AlertName          string `json:"alert_name"`
	CancelUrl          string `json:"cancel_url"`
	CheckoutId         string `json:"checkout_id"`
	Currency           string `json:"currency"`
	CustomData         string `json:"custom_data"`
	Email              string `json:"email"`
	EventTime          string `json:"event_time"`
	Status             string `json:"status"`
	PSignature         string `json:"p_signature"`
	Passthrough        string `json:"passthrough"`
	SubscriptionId     string `json:"subscription_id"`
	SubscriptionPlanId string `json:"subscription_plan_id"`
	UserId             string `json:"user_id"`
}
type SubscriptionCreate struct {
	NextBillDate string `json:"next_bill_date"`
	Quantity     string `json:"quantity"`
	Source       string `json:"source"`
	UnitPrice    string `json:"unitPrice"`
	UpdateUrl    string `json:"updateUrl"`
}

type SubscriptionUpdate struct {
	MarketingConsent      string `json:"marketing_consent"`
	NewPrice              string `json:"new_price"`
	NewQuantity           string `json:"new_quantity"`
	NewUnitPrice          string `json:"new_unit_price"`
	NextBillDate          string `json:"next_bill_date"`
	OldNextBillDate       string `json:"old_next_bill_date"`
	OldPrice              string `json:"old_price"`
	OldQuantity           string `json:"old_quantity"`
	OldStatus             string `json:"old_status"`
	OldSubscriptionPlanId string `json:"old_subscription_plan_id"`
	OldUnitPrice          string `json:"old_unit_price"`
	UpdateUrl             string `json:"update_url"`
	PausedAt              string `json:"paused_at"`
	PausedFrom            string `json:"paused_from"`
	PausedReason          string `json:"paused_reason"`
}

type SubscriptionCanceled struct {
	MarketingConsent string `json:"marketing_consent"`
	Quantity         string `json:"quantity"`
	UnitPrice        string `json:"unit_price"`
}

type SubscriptionPaymentSucceeded struct {
	BalanceCurrency       string `json:"balance_currency"`
	BalanceEarnings       string `json:"balance_earnings"`
	BalanceFee            string `json:"balance_fee"`
	BalanceGross          string `json:"balance_gross"`
	BalanceTax            string `json:"balance_tax"`
	Country               string `json:"country"`
	Coupon                string `json:"coupon"`
	CustomerName          string `json:"customer_name"`
	Earnings              string `json:"earnings"`
	Fee                   string `json:"fee"`
	InitialPayment        string `json:"initial_payment"`
	Instalments           string `json:"instalments"`
	MarketingConsent      string `json:"marketing_consent"`
	NextBillDate          string `json:"next_bill_date"`
	NextPaymentAmount     string `json:"next_payment_amount"`
	OrderId               string `json:"order_id"`
	PaymentMethod         string `json:"payment_method"`
	PaymentTax            string `json:"payment_tax"`
	PlanName              string `json:"plan_name"`
	Quantity              string `json:"quantity"`
	ReceiptUrl            string `json:"receipt_url"`
	SaleGross             string `json:"sale_gross"`
	SubscriptionPaymentId string `json:"subscription_payment_id"`
	UnitPrice             string `json:"unit_price"`
}

type SubscriptionPaymentFailed struct {
	Amount                string `json:"amount"`
	AttemptNumber         string `json:"attempt_number"`
	Instalments           string `json:"instalments"`
	MarketingConsent      string `json:"marketing_consent"`
	NextRetryDate         string `json:"next_retry_date"`
	OrderId               string `json:"order_id"`
	Quantity              string `json:"quantity"`
	SubscriptionPaymentId string `json:"subscription_payment_id"`
	SubscriptionPlanId    string `json:"subscription_plan_id"`
	UnitPrice             string `json:"unit_price"`
	UpdateUrl             string `json:"update_url"`
}

type SubscriptionPaymentRefunded struct {
	Amount                  string `json:"amount"`
	BalanceCurrency         string `json:"balance_currency"`
	BalanceEarningsDecrease string `json:"balance_earnings_decrease"`
	BalanceGrossRefund      string `json:"balance_gross_refund"`
	BalanceTaxRefund        string `json:"balance_tax_refund"`
	EarningsDecrease        string `json:"earnings_decrease"`
	FeeRefund               string `json:"fee_refund"`
	GrossRefund             string `json:"gross_refund"`
	InitialPayment          int    `json:"initial_payment"`
	Instalments             string `json:"instalments"`
	MarketingConsent        string `json:"marketing_consent"`
	OrderId                 string `json:"order_id"`
}
