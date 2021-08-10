package app

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/test", testBook)
	http.HandleFunc("/book", handleBook)
	//if err := http.ListenAndServe(":8080", pingHandler()); err != nil && err != http.ErrServerClosed {
	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		fmt.Errorf("something went wrong %s", err)
		fmt.Println("Server Not Started")
		return
	}
	fmt.Println("Server Started")
}

func handleBook(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusCreated)
	// writer.Write("Ticket created")
	// fmt.Println("Book Created!")
}

func testBook(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusInternalServerError)
	// writer.Write("Ticket created")
	// fmt.Println("Book Created!")
}

func pingHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}
