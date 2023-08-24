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

func setMenu(app *tview.Application) *tview.List {

	return tview.NewList().
		AddItem("Select from the menu items below...", "", 'm', nil).
		AddItem("Memory Available", "Reads available memory in the system.", 'a', nil).
		AddItem("Available Disks", "Displays available disk space.", 'b', nil).
		AddItem("See Current User", "Displays current user.", 'c', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
}

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
		log.Fatalf(err.Error())
	}

	username := user.Username

	return username
}

func setResponse(index int, mainText string, secondaryText string, shortcut rune) {

	switch shortcut {
	case 'm':
		list.SetItemText(index, mainText, "select from below items...")
	case 'a':
		gb := memHandler()
		list.SetItemText(index, mainText, fmt.Sprintf("%v GB", gb))
	case 'b':
		disk := diskHandler()
		list.SetItemText(index, mainText, fmt.Sprintf("All: %v GB Used: %v GB Free: %v GB", int(disk.All)/int(GB), int(disk.Used)/int(GB), int(disk.Free)/int(GB)))
	case 'c':
		u := userHandler()
		list.SetItemText(index, mainText, fmt.Sprintf("Current User: %v", u))
	case 'q':
		list.SetItemText(index, mainText, "exiting...")
	default:
		fmt.Println("that isn't an available option")

	}
}

func bytesToGB(i uint64) uint64 {
	return i / 1024 / 1024 / 1024
}
func main() {
	app := tview.NewApplication()

	list = setMenu(app)

	list.SetChangedFunc(setResponse)

	if err := app.SetRoot(list, true).Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
