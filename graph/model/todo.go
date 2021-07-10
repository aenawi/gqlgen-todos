package model

type Todo struct {
	Text		string `json:"text"`
	ID			string `json:"id"`
	Done		bool `json:"done"`
	UserID	string `json:"user"`
}
