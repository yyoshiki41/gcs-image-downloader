package entity

type GcsResponse struct {
	Items []Link
}

type Link struct {
	Link string
}

func NewGcsResponse() *GcsResponse {
	return &GcsResponse{}
}
