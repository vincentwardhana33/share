package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ExampleServer_customHTTPStatusCode() {
	type CustomResponse struct {
		Status int `json:"status"`
	}
	exp := CustomResponse{
		Status: http.StatusBadRequest,
	}

	opts := &ServerOptions{
		Address: "127.0.0.1:8123",
	}

	testServer, err := NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err = testServer.Start()
		if err != nil {
			log.Println(err)
		}
	}()

	defer testServer.Stop(5 * time.Second)

	epCustom := &Endpoint{
		Path:         "/error/custom",
		RequestType:  RequestTypeJSON,
		ResponseType: ResponseTypeJSON,
		Call: func(res http.ResponseWriter, req *http.Request, reqBody []byte) (
			resbody []byte, err error,
		) {
			res.WriteHeader(exp.Status)
			return json.Marshal(exp)
		},
	}

	err = testServer.registerPost(epCustom)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the server fully started.
	time.Sleep(1 * time.Second)

	client := NewClient("http://127.0.0.1:8123", nil, false)

	httpRes, resBody, err := client.PostJSON(nil, epCustom.Path, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", httpRes.StatusCode)
	fmt.Printf("%s\n", resBody)
	// Output:
	// 400
	// {"status":400}
}
