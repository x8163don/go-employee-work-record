package main

import (
	"encoding/csv"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"time"
)

type Member struct {
	no   string
	name string
}

func main() {
	fmt.Println(time.Now().Local())
	var currentMember Member

	a := app.New()
	w := a.NewWindow("打卡系統")
	w.Resize(fyne.NewSize(680, 480))

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

	workBtn := widget.NewButton("up", func() {
		if (currentMember == Member{}) {
			message.SetText("please choose person")
			return
		}
		message.SetText(fmt.Sprintf("%s - %s - %s - %s ", time.Now().Local(), currentMember.no, currentMember.name, "up"))
		writeRecord(currentMember, "up")
		currentMember = Member{}
	})

	offBtn := widget.NewButton("off", func() {
		if (currentMember == Member{}) {
			message.SetText("please choose person")
			return
		}
		message.SetText(fmt.Sprintf("%s - %s - %s - %s ", time.Now().Local(), currentMember.no, currentMember.name, "off"))
		writeRecord(currentMember, "off")
		currentMember = Member{}
	})

	rvbox.Add(offBtn)
	rvbox.Add(workBtn)

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

func writeRecord(member Member, recordType string) {
	currentDir, err := os.Getwd()
	file, err := os.OpenFile(currentDir+"/config/records.csv", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf(err.Error())
		exit(fmt.Sprintf("Failed to open the CSV file members"))
	}
	file.WriteString(fmt.Sprintf("%s,%s,%s,%s\n", time.Now().Local(), member.no, member.name, recordType))
	file.Close()
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
