package omise

type Barcode struct {
	Object string    `json:"object"`
	Type   string    `json:"type"`
	Image  *Document `json:"image"`
}
