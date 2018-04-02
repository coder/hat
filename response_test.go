package hat

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	req, err := http.NewRequest("GET", "google.com", nil)
	require.NoError(t, err)

	Body(strings.NewReader("test123"))(req)

	byt, err := ioutil.ReadAll(req.Body)
	require.NoError(t, err)

	assert.Equal(t, []byte("test123"), byt)
}

func TestResponse(t *testing.T) {
	resp := Response{
		Response: &http.Response{
			Body: ioutil.NopCloser(strings.NewReader("howdy")),
		},
		t: T{
			T: t,
		},
	}

	t.Run("DuplicateBody", func(t *testing.T) {
		for i := 0; i < 4; i++ {
			assert.Equal(t, "howdy", string(resp.DuplicateBody()))
		}
	})
}
