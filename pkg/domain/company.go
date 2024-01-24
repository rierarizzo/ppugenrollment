package domain

type Company struct {
	ID       int
	Name     string
	RUC      string
	ImageURL string
}

type Schedule struct {
	ID          int
	Code        string
	Description string
}
