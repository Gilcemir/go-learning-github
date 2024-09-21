package validator

import (
	"unicode"
	"unicode/utf8"
)

const (
	MinLenErr = "A senha deve ter oito ou mais caracteres."
	MaiusErr  = "Sua senha deve conter pelo menos um caracter maiúsculo."
	MinsErr   = "Sua senha deve conter pelo menos um caracter minúsculo."
	NumErr    = "Sua senha deve conter pelo menos um número."
	SpecErr   = "Sua senha deve conter pelo menos um caracter especial."
)

func IsValidPassword(password *string) (bool, *ValidationError) {
	result := make([]bool, 5)
	if utf8.RuneCountInString(*password) >= 8 {
		result[0] = true
	}

	for _, r := range *password {
		if r >= 'A' && r <= 'Z' {
			result[1] = true
			continue
		}
		if r >= 'a' && r <= 'z' {
			result[2] = true
			continue
		}
		if r >= '0' && r <= '9' {
			result[3] = true
			continue
		}
		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			result[4] = true
		}
	}

	countErrors := 0
	validationError := newValidationError()

	for idx, r := range result {
		switch {
		case idx == 0 && !r:
			countErrors++
			validationError.Messages[0] = MinLenErr
		case idx == 1 && !r:
			countErrors++
			validationError.Messages[1] = MaiusErr
		case idx == 2 && !r:
			countErrors++
			validationError.Messages[2] = MinsErr
		case idx == 3 && !r:
			countErrors++
			validationError.Messages[3] = NumErr
		case idx == 4 && !r:
			countErrors++
			validationError.Messages[4] = SpecErr
		}
	}

	if countErrors == 0 {
		return true, nil
	}
	return false, validationError
}
