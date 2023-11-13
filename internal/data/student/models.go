package student

import "time"

type Model struct {
	ID           int       `sql:"id"`
	IDCardNumber string    `sql:"id_card_number"`
	Name         string    `sql:"name"`
	Surname      string    `sql:"surname"`
	Email        string    `sql:"email"`
	Password     string    `sql:"password"`
	Role         string    `sql:"role"`
	DateOfBirth  time.Time `sql:"date_of_birth"`
	IsAGraduate  bool      `sql:"is_a_graduate"`
	Level        int       `sql:"level"`
}
