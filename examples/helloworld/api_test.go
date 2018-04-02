package helloworld

import (
	"testing"

	"github.com/codercom/hat"
	"github.com/stretchr/testify/assert"
	"go.coder.com/ctest/chttptest"
)

func TestAPI(tt *testing.T) {
	addr, close := chttptest.StartHTTPServer(tt, &API{})
	defer close()

	t := hat.New(tt, "http://"+addr)

	t.Run("Hello Echo", func(t hat.T) {
		t.Request(hat.GET, nil).Assert(func(r hat.Response) {
			byt := r.DuplicateBody()
			assert.Equal(t, "Hello /", string(byt))
		})

		t.RequestURL(hat.GET, "/dog_soup", nil).Assert(func(r hat.Response) {
			byt := r.DuplicateBody()
			assert.Equal(t, "Hello /dog_soup", string(byt))
		})
	})
}
