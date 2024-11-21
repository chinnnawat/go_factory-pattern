package usecases

import (
	"fmt"

	"example.com/m/v2/interfaces"
)

func NewBook(bookType string) (interfaces.IBook, error) {
	switch bookType {
	case "comics":
		return NewCommic(), nil
	}
	return nil, fmt.Errorf("unknown book type")
}
