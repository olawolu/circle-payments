package payments

// PaymentRequest sent to cicle's api
type PaymentRequest struct {
	PaymentData
	CardData
}

// PaymentResponse as returned from cicrle's api
type PaymentResponse struct {
	PaymentID  string   `json:"id"`
	Source     Source   `json:"source"`
	Amount     Bill     `json:"amount"`
	Status     string   `json:"status"`
	CreateDate string   `json:"createDate"`
	UpdateDate string   `json:"updateDate"`
	Metadata   MetaData `json:"metadata"`
}

// 
// Custom helper types
type CardData struct {
	ExpiryMonth    int32 `json:"expMonth"`
	ExpiryYear     int32 `json:"expYear"`
	CardDetails
	BillingDetails
	MetaData
}

type PaymentData struct {
	Amount      string `json:"amount"`
	Description string `json:"description"`
}


type CardDetails struct {
	CardNumber string `json:"number"`
	CVV        string `json:"cvv"`
}

type BillingDetails struct {
	Name string `json:"name"`
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

type Source struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Bill struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
