package stub_contacts

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPopulate(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	err = Populate(c)
	require.NoError(t, err)
	_, err = c.List()
	assert.NoError(t, err)
}

func TestPopulate2(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	err = Populate(c)
	require.NoError(t, err)
	err = Populate(c)
	assert.NoError(t, err)
}
