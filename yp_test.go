package stub_contacts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func _getEric() Contact {
	return Contact{
		FirstName: "Eric",
		LastName:  "Adams",
		Phone:     "call-me-123",
	}
}

func TestYellowPages_Save_New(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	contact := _getEric()
	_, err = c.Save(contact)
	assert.NoError(t, err)
	l, err := c.List()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(l))
}

func TestYellowPages_Save_New_Duplicate(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)

	contact := _getEric()
	_, _ = c.Save(contact)
	_, err = c.Save(contact)
	assert.Error(t, err)
}

func TestYellowPages_Save_Update(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)

	contact := _getEric()
	id, _ := c.Save(contact)
	assert.Equal(t, uint(1), id)
	contact = Contact{
		FirstName: "Louis",
		LastName:  "Marullo",
		Phone:     "call-me-123",
	}
	contact.ID = id
	_, err = c.Save(contact)
	assert.NoError(t, err)
	joey, err := c.Load(id)
	assert.NoError(t, err)
	assert.True(t, isSameContact(joey, contact))
}

func TestYellowPages_Save_Duplicate(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	contact := _getEric()
	id, _ := c.Save(contact)
	contact.ID = id
	_, err = c.Save(contact)
	assert.Error(t, err)
}

func TestYellowPages_Delete(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	contact := _getEric()
	id, _ := c.Save(contact)
	l, err := c.List()
	require.NoError(t, err)
	assert.Equal(t, 1, len(l))
	err = c.Delete(id)
	assert.NoError(t, err)
	l, err = c.List()
	require.Error(t, err)
	assert.Equal(t, ErrNotFound, err)
}

func TestYellowPages_Delete_Duplicate(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	contact := _getEric()
	id, _ := c.Save(contact)
	l, err := c.List()
	require.NoError(t, err)
	assert.Equal(t, 1, len(l))
	err = c.Delete(id)
	assert.NoError(t, err)
	err = c.Delete(id)
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)
}

func TestYellowPages_FindByPhone_YES(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	contacts := []Contact{
		{FirstName: "Eric", LastName: "Adams", Phone: "call-me-123"},
		{FirstName: "Joey", LastName: "DeMaio", Phone: "call-me-234"},
		{FirstName: "Attila", LastName: "Dorn", Phone: "call-me-234"},
		{FirstName: "Falk Maria", LastName: "Schlegel", Phone: "call-me-345"},
	}
	for _, v := range contacts {
		_, err = c.Save(v)
		require.NoError(t, err)
	}
	list, err := c.FindByPhone("123")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(list))
	assert.Equal(t, "Eric", list[0].FirstName)
}

func TestYellowPages_FindByPhone_YES_Begin(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	contacts := []Contact{
		{FirstName: "Eric", LastName: "Adams", Phone: "call-me-123"},
		{FirstName: "Joey", LastName: "DeMaio", Phone: "call-me-234"},
		{FirstName: "Attila", LastName: "Dorn", Phone: "call-me-234"},
		{FirstName: "Falk Maria", LastName: "Schlegel", Phone: "call-me-345"},
	}
	for _, v := range contacts {
		_, err = c.Save(v)
		require.NoError(t, err)
	}
	list, err := c.FindByPhone("call")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(list))

}

func TestYellowPages_FindByPhone_NO(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	contacts := []Contact{
		{FirstName: "Eric", LastName: "Adams", Phone: "call-me-123"},
		{FirstName: "Joey", LastName: "DeMaio", Phone: "call-me-234"},
		{FirstName: "Attila", LastName: "Dorn", Phone: "call-me-234"},
		{FirstName: "Falk Maria", LastName: "Schlegel", Phone: "call-me-345"},
	}
	for _, v := range contacts {
		_, err = c.Save(v)
		require.NoError(t, err)
	}
	_, err = c.FindByPhone("777777")
	assert.Error(t, err)
	assert.Equal(t, ErrNotFound, err)
}

func TestYellowPages_FindByPhone_YES_MULTY(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	contacts := []Contact{
		{FirstName: "Eric", LastName: "Adams", Phone: "call-me-123"},
		{FirstName: "Joey", LastName: "DeMaio", Phone: "call-me-234"},
		{FirstName: "Attila", LastName: "Dorn", Phone: "call-me-234"},
		{FirstName: "Falk Maria", LastName: "Schlegel", Phone: "call-me-345"},
	}
	for _, v := range contacts {
		_, err = c.Save(v)
		require.NoError(t, err)
	}
	list, err := c.FindByPhone("234")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(list))
}
