package utilities

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

//f, _ := os.Create("cpu.prof")
//pprof.StartCPUProfile(f)
//defer pprof.StopCPUProfile()
//

//f1, _ := os.Create("trace.out")
//defer f.Close()
//
//trace.Start(f1)
//defer trace.Stop()

//import _ "net/http/pprof"
//
//http://localhost:6060/debug/pprof/heap - информация о расходе памяти.
//http://localhost:6060/debug/pprof/profile - CPU-профилирование.
//http://localhost:6060/debug/pprof/goroutine - информация о горутинах и их состоянии.
