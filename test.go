package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Resource struct {
	Name string
	IP   string
}

func main() {
	var (
		usernameTE, passwordTE  *walk.LineEdit
		connectPB, disconnectPB *walk.PushButton
		statusLB                *walk.Label
		resourceLB              *walk.ListBox
		inTE, outTE             *walk.TextEdit
	)

	resources := []Resource{
		{Name: "Server 1", IP: "192.168.0.1"},
		{Name: "Server 2", IP: "192.168.0.2"},
		{Name: "Server 3", IP: "192.168.0.3"},
	}

	sendPostRequest := func() {
		body := strings.NewReader(inTE.Text())

		resp, err := http.Post("http://localhost:8080", "text/plain", body)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(body)

		outTE.SetText("")                           // Clear the existing content
		outTE.AppendText(string(respBody) + "\r\n") // Append the current message
	}

	MainWindow{
		Title:  "Simple VPN App",
		Size:   Size{400, 500},
		Layout: VBox{},
		Children: []Widget{
			Label{Text: "Username:"},
			LineEdit{AssignTo: &usernameTE},
			Label{Text: "Password:"},
			LineEdit{AssignTo: &passwordTE, PasswordMode: true},
			PushButton{
				AssignTo: &connectPB,
				Text:     "Connect",
				OnClicked: func() {
					log.Println("Connecting...")
					statusLB.SetText("Status: Connected")
				},
			},
			PushButton{
				AssignTo: &disconnectPB,
				Text:     "Disconnect",
				OnClicked: func() {
					log.Println("Disconnecting...")
					statusLB.SetText("Status: Disconnected")
				},
			},
			Label{AssignTo: &statusLB, Text: "Status: Disconnected"},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{Text: "Available Resources:"},
					ListBox{
						AssignTo:      &resourceLB,
						Model:         resources,
						DisplayMember: "Name",
					},
				},
			},
			HSplitter{
				Children: []Widget{
					TextEdit{
						AssignTo: &inTE,
						OnKeyDown: func(key walk.Key) {
							if key == walk.KeyReturn {
								sendPostRequest()
							}
						},
					},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SEND",
				OnClicked: func() {
					sendPostRequest()
				},
			},
		},
	}.Run()
}

func getKeys(obj interface{}) []string {
	objType := reflect.TypeOf(obj)
	keys := make([]string, objType.NumField())

	for i := 0; i < objType.NumField(); i++ {
		keys[i] = objType.Field(i).Name
	}

	return keys
}
