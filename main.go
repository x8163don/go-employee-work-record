package main

import (
	"encoding/csv"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

type Member struct {
	no   string
	name string
}

func main() {

	var currentMember Member

	a := app.New()
	w := a.NewWindow("打卡系統")
	w.Resize(fyne.Size{Width: 640, Height: 320})

	members, _ := readMembers()

	message := widget.NewLabel("")
	vbox := container.NewVBox()
	hbox := container.NewHBox()
	lvbox := container.NewVBox()
	rvbox := container.NewVBox()

	hbox.Add(lvbox)
	hbox.Add(rvbox)

	vbox.Add(hbox)
	vbox.Add(message)

	for _, member := range members {
		tmpMember := member
		lvbox.Add(widget.NewButton(tmpMember.no+" "+tmpMember.name, func() {
			currentMember = tmpMember
			message.SetText(tmpMember.no + " " + tmpMember.name)
		}))
	}

	rvbox.Add(widget.NewButton("上班", func() { message.SetText(currentMember.no + " " + currentMember.name + "上班打卡記錄") }))
	rvbox.Add(widget.NewButton("下班", func() { message.SetText(currentMember.no + " " + currentMember.name + "下班打卡記錄") }))

	w.SetContent(vbox)
	w.ShowAndRun()
}

func readMembers() ([]Member, error) {
	currentDir, err := os.Getwd()
	file, err := os.Open(currentDir + "/config/members.csv")
	if err != nil {
		fmt.Printf(err.Error())
		exit(fmt.Sprintf("Failed to open the CSV file members"))
	}

	allMembers, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var data []Member
	for _, line := range allMembers {
		member := Member{}
		member.no = line[0]
		member.name = line[1]
		data = append(data, member)
	}
	return data, nil
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
