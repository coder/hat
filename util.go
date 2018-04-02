package hat

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

// ReadAll reads all the data from a reader.
func ReadAll(t testing.TB, rd io.Reader) []byte {
	byt, err := ioutil.ReadAll(rd)
	require.NoError(t, err, "failed to read all")
	return byt
}
