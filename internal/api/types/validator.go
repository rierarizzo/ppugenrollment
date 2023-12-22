package types

import (
	"fmt"
	"ppugenrollment/internal/domain"
	"regexp"
)

type Validator struct {
	appErr *domain.AppError
}

func (v *Validator) MustBeGreaterThan(high, value int) bool {
	if v.appErr != nil {
		return false
	}

	if value <= high {
		v.appErr.Type = domain.BadRequestError
		v.appErr.Err = fmt.Errorf("must be greater than %d", high)

		return false
	}

	return true
}

func (v *Validator) MustNotBeEmpty(value string) bool {
	if v.appErr != nil {
		return false
	}

	if value == "" {
		v.appErr.Type = domain.BadRequestError
		v.appErr.Err = fmt.Errorf("must not be empty")

		return false
	}

	return true
}

func (v *Validator) MustNotBeZero(value int) bool {
	if v.appErr != nil {
		return false
	}

	if value == 0 {
		v.appErr.Type = domain.BadRequestError
		v.appErr.Err = fmt.Errorf("must not be zero")

		return false
	}

	return true
}

func (v *Validator) MustBeEmail(value string) bool {
	if v.appErr != nil {
		return false
	}

	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)

	if !re.MatchString(value) {
		v.appErr.Type = domain.BadRequestError
		v.appErr.Err = fmt.Errorf("must be a valid email")

		return false
	}

	return true
}
