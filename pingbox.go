package main
import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 8080
	log.Println("Starting up pingbox on port", port, "...")
	http.HandleFunc("/", handlr)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}

func handlr(writer http.ResponseWriter, r *http.Request) {
	remote_addr := r.Header.Get("X-Forwarded-For")
	if remote_addr == "" {
		remote_addr = r.RemoteAddr
	}
	log.Println(r.Method, "request from", remote_addr)
	fmt.Fprintf(writer, "OK - pingbox")
}
