package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

	labelIpAddress := canvas.NewText("Your IP Address", color.White)
	inputIpAddress := widget.NewEntry()
	inputIpAddress.SetPlaceHolder("0.0.0.0")

	labelIPType := canvas.NewText("IP Type", color.White)
	inputIPType := widget.NewEntry()
	inputIPType.SetPlaceHolder("IP Type")

	labelOrg := canvas.NewText("Provider", color.White)
	inputOrg := widget.NewEntry()
	inputOrg.SetPlaceHolder("Provider")

	labelCountry := canvas.NewText("Country", color.White)
	inputCountry := widget.NewEntry()
	inputCountry.SetPlaceHolder("Country")

	labelRegion := canvas.NewText("Region", color.White)
	inputRegion := widget.NewEntry()
	inputRegion.SetPlaceHolder("Region")

	labelCity := canvas.NewText("City", color.White)
	inputCity := widget.NewEntry()
	inputCity.SetPlaceHolder("Your City")

	progress := widget.NewProgressBar()

	locationGrid := container.New(layout.NewGridLayout(2), labelCountry, inputCountry, labelRegion, inputRegion, labelCity, inputCity)

	getInformationButton := widget.NewButtonWithIcon("Get IP Information", theme.SearchIcon(), func() {
		ip := getIp()
		inputIpAddress.SetText(ip)
		progress.SetValue(0.5)
		if ip != "" {
			ipInfo := getIpInfo(ip)
			inputIPType.SetText(ipInfo.Type)
			inputOrg.SetText(ipInfo.Connection.Org)
			inputCity.SetText(ipInfo.City)
			inputCountry.SetText(ipInfo.Country)
			inputRegion.SetText(ipInfo.Region)
		}
		progress.SetValue(1.0)
	})

	cleanFieldsButton := widget.NewButtonWithIcon("Clean Fields", theme.ContentClearIcon(), func() {
		inputIpAddress.SetText("")
		inputIPType.SetText("")
		inputOrg.SetText("")
		inputCity.SetText("")
		inputCountry.SetText("")
		inputRegion.SetText("")
		progress.SetValue(0)
	})

	copyIpAddressButton := widget.NewButtonWithIcon("Copy IP Address", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(inputIpAddress.Text)
	})

	buttonsGrid := container.New(layout.NewGridLayout(3), getInformationButton, cleanFieldsButton, copyIpAddressButton)

	w.SetContent(container.NewVBox(
		labelIpAddress,
		inputIpAddress,
		labelIPType,
		inputIPType,
		labelOrg,
		inputOrg,
		locationGrid,
		buttonsGrid,
		progress,
	))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}

type Response struct {
	Status string
	Code   string
	Total  int
}

type IpInfo struct {
	Ip         string
	Type       string
	Continent  string
	Country    string
	Region     string
	City       string
	Connection Connection
}

type Connection struct {
	Org    string
	Domain string
}

func red_button() *fyne.Container { // return type
	btn := widget.NewButton("Visit", nil) // button widget
	// button color
	btn_color := canvas.NewRectangle(
		color.NRGBA{R: 255, G: 0, B: 0, A: 255})
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
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

func getIpInfo(ip string) IpInfo {
	res, err := http.Get("https://ipwho.is/" + ip)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var myStoredVariable IpInfo
	json.Unmarshal([]byte(responseData), &myStoredVariable)
	return myStoredVariable
}
