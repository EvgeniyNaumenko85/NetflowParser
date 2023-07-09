//package utilites

//import (
//	"fmt"
//	"github.com/shirou/gopsutil/cpu"
//)
//
//func CpuInfo() {
//	info, err := cpu.Info()
//	if err != nil {
//		fmt.Printf("Ошибка при получении информации о процессоре: %s\n", err)
//		return
//	}
//
//	//physicalCores := info[0].PhysicalCores
//	logicalCores := info[0].Cores
//
//	//fmt.Printf("Количество физических ядер: %d\n", physicalCores)
//	fmt.Printf("Количество логических ядер: %d\n", logicalCores)
//}

package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Моё первое в Fyne приложение")
	label := widget.NewLabel("Привет, Fyne!")
	label2 := widget.NewLabel("Lable2")

	w.SetContent(container.NewVBox(
		label,
		label2,
	))

	w.ShowAndRun()
}
