package main

import "tt/src/demo/client"

const (
	IP   = "192.168.1.209:53"
	PORT = 1010
)

func main() {
	client.Ma()
	//var inTE, outTE *walk.TextEdit
	//
	//MainWindow{
	//	Title:   "SCREAMO",
	//	MinSize: Size{600, 400},
	//	Layout:  VBox{},
	//	Children: []Widget{
	//		HSplitter{
	//			Children: []Widget{
	//				TextEdit{AssignTo: &inTE},
	//go get -u golang.org/x/image/				TextEdit{AssignTo: &outTE, ReadOnly: true},
	//			},
	//		},
	//		PushButton{
	//			Text: "SCREAM",
	//			OnClicked: func() {
	//				outTE.SetText(strings.ToUpper(inTE.Text()))
	//			},
	//		},
	//	},
	//}.Run()
}
