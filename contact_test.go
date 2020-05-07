package stub_contacts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSameContact_Yes_Empty(t *testing.T) {
	c1 := EmptyContact
	c2 := EmptyContact
	assert.True(t, isSameContact(c1, c2))
}

func Test_isSameContact_Yes_Full(t *testing.T) {
	c1 := Contact{
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
	c2 := Contact{
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
	assert.True(t, isSameContact(c1, c2))
}

func Test_isSameContact_Yes_DiffIDS(t *testing.T) {
	c1 := Contact{
		ID:        10,
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
	c2 := Contact{
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
	assert.True(t, isSameContact(c1, c2))
}

func Test_isSameContact_No_First(t *testing.T) {
	c1 := Contact{
		FirstName: "Joey",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
	c2 := Contact{
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
	assert.False(t, isSameContact(c1, c2))
}

func Test_isSameContact_No_Last(t *testing.T) {
	c1 := Contact{
		FirstName: "Eric",
		LastName:  "DeMaio",
		Phone:     "call-me-123",
	}
	c2 := Contact{
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
	assert.False(t, isSameContact(c1, c2))
}

func Test_isSameContact_No_Phone(t *testing.T) {
	c1 := Contact{
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
	c2 := Contact{
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-234",
	}
	assert.False(t, isSameContact(c1, c2))
}

func Test_isSameContact_No_Empty(t *testing.T) {
	c1 := Contact{
		FirstName: "Eric",
		LastName:  "",
		Phone:     "call-me-123",
	}
	c2 := Contact{
		FirstName: "",
		LastName:  "Adams",
		Phone:     "call-me-234",
	}
	assert.False(t, isSameContact(c1, c2))
}
