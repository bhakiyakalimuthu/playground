package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// json encoding
	// json decoding
	// json marshal
	//_, err := json.Marshal(_Val{in: "testing"})
	//if err != nil {
	//	fmt.Println("failed to marshal")
	//	return
	//}
	// json unmarshal
	var val *_Val
	if err := json.Unmarshal([]byte(`{"in":"testing"}`), &val); err != nil {
		fmt.Println("failed to unmarshal")
	}
	fmt.Println(*val)
	http.Client{}
}

type _Val struct {
	in string `json:"in"`
}
