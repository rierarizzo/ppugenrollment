package user

import "time"

type Model struct {
	ID           int       `db:"id"`
	IDCardNumber string    `db:"id_card_number"`
	Name         string    `db:"name"`
	Surname      string    `db:"surname"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	Role         string    `db:"role"`
	DateOfBirth  time.Time `db:"date_of_birth"`
	IsAGraduate  bool      `db:"is_a_graduate"`
	Level        int       `db:"level"`
}
