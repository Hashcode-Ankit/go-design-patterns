package main

import (
	"fmt"
	"main/state"
)

// Harsh
func main() {
	// create object of meetingStore
	stateStr := `[{"type":"STREAM","stream":{"stream_descriptor":{"name":"Golang - HS Shortlisted"},"stream_state":{"__ab_no_cursor_state_message":true}}}]`
	if err := state.ParseState(stateStr, "Golang - HS Shortlisted"); err != nil {
		fmt.Println("here we received err while parsing state: ", err)
	}
}

// func onMain(w http.ResponseWriter, r *http.Request) {
// 	data := map[string]string{
// 		"Name": "Ankit",
// 	}

// 	response, err := json.Marshal(data)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 	}
// 	w.Write(response)
// }

// func onHello(w http.ResponseWriter, r *http.Request) {
// 	data := map[string]string{
// 		"Hi": "Hello",
// 	}

// 	response, err := json.Marshal(data)
// 	if err != nil {
// 		w.Write([]byte(err.Error()))
// 	}
// 	w.Write(response)
// }

// // Builder pattern
// newCarBuilder := builder.NewBuilder(1, "Sports Car Wheels")
// fmt.Println(newCarBuilder.Stop())

// // Factory Desing pattern

// // simple server in golang
// mux := http.NewServeMux()
// mux.HandleFunc("/", onMain)
// mux.HandleFunc("/hello", onHello)

// err := http.ListenAndServe(":3333", mux)
// if err != nil {
// 	fmt.Println("failed with error: ", err)
// }
