package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	a := app.New()
	w := a.NewWindow("Get IP Address Information")
	title := canvas.NewText("Your IP Address", color.White)
	input := widget.NewEntry()
	ip := getIp()
	input.SetText(ip)
	fmt.Println(ip)
	input.SetPlaceHolder("Enter password length")
	copybtn := widget.NewButtonWithIcon("Copy IP Address", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(input.Text)
	})
	w.SetContent(container.NewVBox(
		title,
		input,
		copybtn,
	))
	w.Resize(fyne.NewSize(300, 100))
	w.ShowAndRun()
}

func getIp() string {
	response, err := http.Get("https://api.ipify.org")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(responseData)
}
