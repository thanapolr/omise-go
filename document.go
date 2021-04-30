package omise

// Document represents Omise's document object.
// See https://www.omise.co/documents-api for more information.
type Document struct {
	Base
	Filename    string `json:"filename"`
	Deleted     bool   `json:"deleted"`
	DownloadUri string `json:"download_uri"`
}
