package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/pbnjay/memory"
	"github.com/rivo/tview"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var list *tview.List

func memHandler() uint64 {
	m := memory.TotalMemory()
	gb := bytesToGB(m)
	return gb
}

func diskHandler() *DiskStatus {
	disk := &DiskStatus{}
	err := disk.diskUsage("/")
	if err != nil {
		log.Fatalf("%v", err)
	}

	return disk
}

func userHandler() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("%v", err)
	}

	username := user.Username

	return username
}

func bytesToGB(i uint64) uint64 {
	return i / 1024 / 1024 / 1024
}

func menu(app *tview.Application) *tview.List {

	list := tview.NewList().
		AddItem("Select an item from below...", "", 'm', nil).
		AddItem("Get available memory.", "", 'a', nil).
		AddItem("Get availabe disk space.", "", 'b', nil).
		AddItem("Get the current logged in user.", "", 'c', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	return list

}
func setResponse(index int, mainText string, secondaryText string, shortcut rune) {
	switch shortcut {
	case 'm':
		list.SetItemText(index, mainText, "select item from below...")
	case 'a':
		mem := memHandler()
		list.SetItemText(index, mainText, fmt.Sprintf("%v GB", mem))
	case 'b':
		disk := diskHandler()
		list.SetItemText(index, mainText, fmt.Sprintf("All: %v GB Used: %v GB Free: %v GB", int(disk.All)/int(GB), int(disk.Used)/int(GB), int(disk.Free)/int(GB)))
	case 'c':
		u := userHandler()
		list.SetItemText(index, mainText, fmt.Sprintf("Current User: %v", u))
	default:
		fmt.Println("that is not an available option")

	}
}

func main() {
	app := tview.NewApplication()

	list = menu(app)

	list.SetChangedFunc(setResponse)

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
