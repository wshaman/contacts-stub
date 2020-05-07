package stub_contacts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPopulate(t *testing.T) {
	c, err := NewYellowPages()
	require.NoError(t, err)
	err = Populate(c)
	require.NoError(t, err)
	_, err = c.List()
	require.NoError(t, err)
}
