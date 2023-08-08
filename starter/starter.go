package starter

import (
	"fmt"
	"math"
	"net/http"
	//"net/http/httptest"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name) 
}


func OddOrEven(number int) string {
	criteria := math.Mod(float64(number), 2)
	if criteria == 1 || criteria == -1 {
		return fmt.Sprintf("%v is an odd number", number)
	}

	return fmt.Sprintf("%v is an even number", number)
}


func CheckHealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "health check passed")
}