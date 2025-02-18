package models

type Choice struct {
	Selection string `json:"selection"`
}

type Item struct {
	Name  string
	Score int
}
