package project

type Model struct {
	ID          int    `sql:"id"`
	Company     int    `sql:"company"`
	Description string `sql:"description"`
	Schedule    string `sql:"schedule"`
}
