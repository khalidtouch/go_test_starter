package starter

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)




func TestSayHello(t *testing.T) {
	greeting := SayHello("William")
	assert.Equal(t, "Hello William. Welcome!", greeting)

	greeting2 := SayHello("Khalid Isah")
	assert.Equal(t, "Hello Khalid Isah. Welcome!", greeting2)
}

func TestOddOrEven(t *testing.T) {
	t.Run("Check Non-negative numbers", func(t *testing.T) {
		//non-negative tests 
		assert.Equal(t, "45 is an odd number", OddOrEven(45))
		assert.Equal(t, "40 is an even number", OddOrEven(40))
		assert.Equal(t, "16 is an even number", OddOrEven(16))
	})

	t.Run("Check for negative numbers", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", OddOrEven(-45))
		assert.Equal(t, "-40 is an even number", OddOrEven(-40))	
	})
}


func TestCheckHealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T){
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder() 
		CheckHealth(writer, req)
		response := writer.Result() 
		body, err := io.ReadAll(response.Body)


		assert.Equal(t, "health check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, 
		"text/plain; charset=utf-8", response.Header.Get("Content-Type"))
		assert.Equal(t, nil, err)
	})
}