package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	Reset      = "\033[0m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Cyan       = "\033[36m"
	Bold       = "\033[1m"
	BoldRed    = "\033[1;31m"
	BoldGreen  = "\033[1;32m"
	BoldYellow = "\033[1;33m"
	BoldCyan   = "\033[1;36m"
)

func main() {
	clear_console()
	list_interfaces()
	fmt.Println("-----------------")
	start_server()
}

func list_interfaces() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, iface := range ifaces {
		fmt.Print("", iface.Name)
		fmt.Print(" : ", iface.HardwareAddr)
		fmt.Print(" ", iface.MTU)
		fmt.Print("  ", iface.Flags.String())
		fmt.Println()
	}
}

func start_server() {
	fmt.Println(cyan("Server started on port 8080"))

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func bold(input string) string {
	return Bold + input + Reset
}

func cyan(input string) string {
	return Cyan + input + Reset
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	date := time.Now().Format("15:04:05")

	formattedString := fmt.Sprintf("%s : %s", date, bold(string(body)))

	fmt.Println(formattedString)

	response := time.Now().Format("15:04:05") + "\n"

	switch string(body) {
	case "clear":
		clear_console()
	}

	_, err = w.Write([]byte(response))

	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}
