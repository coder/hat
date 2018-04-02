package hat

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
