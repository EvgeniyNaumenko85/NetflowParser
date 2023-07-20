package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("приложени в Fyne приложение")
	label := widget.NewLabel("Привет, Fyne!")
	label2 := widget.NewLabel("Lable2")

	w.SetContent(container.NewVBox(
		label,
		label2,
	))

	w.ShowAndRun()
}
