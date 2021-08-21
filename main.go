package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
)

var y int

func main() {
	//x := 10
	//y := "Hello World!"
	//result, _ := fmt.Printf("x: %v, %T\n", x, x)
	//fmt.Printf("y: %v, %T\n", y, y)

	//x = 20
	//fmt.Printf("x: %v, %T\n", x, x)

	//test_func(result)

	http_example()
}

func test_func(result int) {
	fmt.Printf("Result: %v\n", result)

	fmt.Println(y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func http_example() {
	http.HandleFunc("/hello", func(response http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(response, "Hello, playground")
	})

	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(http.Serve(l, nil))
	}()

	log.Println("Sending request...")
	res, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Reading response...")
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}
}
