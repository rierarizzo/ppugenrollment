package project

type Model struct {
	ID          int    `db:"id"`
	Company     int    `db:"company"`
	Description string `db:"description"`
	Schedule    string `db:"schedule"`
}
