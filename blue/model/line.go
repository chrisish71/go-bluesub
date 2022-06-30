package model

type Line struct {
	Text          string `json:"text"`
	Color         string `json:"color"`
	Justification string `json:"justification"`
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Underline     bool   `json:"underline"`
	BoxingColor   string `json:"boxingColor"`
}
