package omise

// Source represents Omise's source object.
// See https://www.omise.co/source-api for more information.
type Source struct {
	Object        string   `json:"object"`
	ID            string   `json:"id"`
	Live          bool     `json:"livemode"`
	Location      *string  `json:"location"`
	Type          string   `json:"type"`
	Flow          string   `json:"flow"`
	Amount        int64    `json:"amount"`
	Currency      string   `json:"currency"`
	ScannableCode *Barcode `json:"scannable_code"`
	Barcode       *string  `json:"barcode"`
	Email         *string  `json:"email"`
	Name          *string  `json:"name"`
	MobileNumber  *string  `json:"mobile_number"`
	PhoneNumber   *string  `json:"phone_number"`
}
