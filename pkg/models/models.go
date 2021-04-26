package models

type Annotation struct {
	Pos   string   `json:"pos"`
	Word  []string `json:"words"`
	Gloss string   `json:"glossary"`
}
