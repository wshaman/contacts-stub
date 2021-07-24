package stub_contacts

import (
	"errors"
	"fmt"
)

func Populate(yp YellowPages) error {
	persons := []Contact{
		{FirstName: "Eric", LastName: "Adams", Phone: "call-me-123"},
		{FirstName: "Joey", LastName: "DeMaio", Phone: "call-me-234"},
		{FirstName: "Attila", LastName: "Dorn", Phone: "call-me-345"},
		{FirstName: "Falk Maria", LastName: "Schlegel", Phone: "call-me-345"},
		{FirstName: "James", LastName: "Hetfield", Phone: "call-me-456"},
		{FirstName: "Kirk", LastName: "Hammett", Phone: "call-me-567"},
		{FirstName: "Luca", LastName: "Turilli", Phone: "call-me-678"},
		{FirstName: "Joakim", LastName: "Brodén", Phone: "call-me-789"},
		{FirstName: "Pär", LastName: "Sundström", Phone: "call-me-789"},
	}
	for _, v := range persons {
		if _, err := yp.Save(v); err != nil {
			if errors.Is(err, ErrDuplicate) {
				fmt.Printf("Skipping existing contact: %s\n", v.String())
				continue
			}
			return err
		}
	}
	return nil
}
