package iteration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepeat(t *testing.T) {
	require.Equal(t, Repeat('a', 5), "aaaaa")
	require.Equal(t, Repeat('c', 4), "cccc")
}
