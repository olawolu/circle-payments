package payments

type CardData struct {
	// IdempotencyKey string `json:"idempotencyKey"`
	// KeyID          string `json:"keyId"`
	ExpMonth       int32  `json:"expMonth"`
	ExpYear        int32  `json:"expYear"`
	CardDetails
	BillingDetails
	MetaData
}

type CardDetails struct {
	CardNumber string `json:"number"`
	CVV        string `json:"cvv"`
}

type BillingDetails struct {
	CardHolderName string `json:"name"`
	City           string `json:"city"`
	Country        string `json:"country"`
	AddressLine1   string `json:"line1"`
	AddressLine2   string `json:"line2"`
	District       string `json:"district"`
	PostalCode     string `json:"postalCode"`
}

type MetaData struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	SessionID   string `json:"sessionId"`
	IPAddress   string `json:"ipAddress"`
}

type PaymentData struct {
	IdempotencyKey string `json:"idempotencyKey"`
	KeyID          string `json:"keyId"`
	Amount         bill   `json:"amount"`
	Verification   string `json:"verification"`
	Source         source `json:"source"`
	Description    string `json:"description"`
	CardDetails
	MetaData
}

type bill struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type source struct {
	ID   string `json:"id"` //card-id returned from create card call
	Type string `json:"type"`
}