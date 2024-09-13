package main

import (
	"fmt"
	"imtixon5/alert"
	"imtixon5/cpu"
	"imtixon5/disk"
	"imtixon5/gpu"
	"imtixon5/logger"
	"imtixon5/memory"
	"imtixon5/network"
	"imtixon5/postgres"
	systemusage "imtixon5/processes"
	"imtixon5/restart"
	"imtixon5/signal"
	"log"
	"time"

	"github.com/google/uuid"
)

const (
	cpuThreshold     = 80.0
	diskThreshold    = 95.0
	networkThreshold = 100 * 1024 * 1024
	memoryThreshold  = 95.0
	gpuThreshold     = 90.0

	cpuprogressThreshold    = 5.0
	memoryprogressThreshold = 1 * 1024 * 1024 * 1024
	diskprogressThreshold   = 5 * 1024 * 1024
)

func main() {
	logs := logger.NewLogger()
	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal(err)
	}

	err = alert.InitBot("7257735336:AAEHYET1u4MGnPMD7R1Y9un-dqf_d1uQiLM")
	if err != nil {
		fmt.Println("Botni ishga tushirishda xatolik:", err)
		return
	}

	for {
		cpuUsage, err := cpu.GetCPUUsage()
		if err == nil && cpuUsage > cpuThreshold {
			alert.SendAlert(fmt.Sprintf("Alert: CPU yuklanishi %v dan oshdi\nHozirgi holati -> %v", cpuThreshold, cpuUsage))
		} else {
			query := `insert into cpu(id,current_situation) values($1,$2)`

			id := uuid.NewString()
			_, err := db.Exec(query, id, cpuUsage)
			if err != nil {
				logs.Error(fmt.Sprintf("cpu table ga xatolik: %s", err))
			}
		}

		diskUsage, err := disk.GetDiskUsage()
		if err == nil && diskUsage > diskThreshold {
			alert.SendAlert(fmt.Sprintf("Alert: Disk yuklanishi %v dan oshdi\nHozirgi holati -> %v", diskThreshold, diskUsage))
		} else {
			query := `insert into disk(id,current_situation) values($1,$2)`

			id := uuid.NewString()
			_, err := db.Exec(query, id, diskUsage)
			if err != nil {
				logs.Error(fmt.Sprintf("disk table ga xatolik: %s", err))
			}
		}

		networkUsage, err := network.GetNetworkStats()
		if err == nil && networkUsage > networkThreshold {
			alert.SendAlert(fmt.Sprintf("Alert: Tarmoq foydalanishi %v baytdan oshdi\nHozirgi holati -> %v", networkThreshold, networkUsage))
		} else {
			query := `insert into network(id,current_situation) values($1,$2)`

			id := uuid.NewString()
			networkUsageInt := int(networkUsage)
			if networkUsageInt > 0 {
				_, err := db.Exec(query, id, networkUsageInt)
				if err != nil {
					logs.Error(fmt.Sprintf("network table ga xatolik: %s", err))
				}
			}
		}

		memoryUsage, err := memory.GetMemoryUsage()
		if err == nil && memoryUsage > memoryThreshold {
			alert.SendAlert(fmt.Sprintf("Alert: RAM yuklanishi %v%% dan oshdi\nHozirgi holati -> %v%%", memoryThreshold, memoryUsage))
			signal.PlayErrorSound()
		} else {
			query := `insert into memory(id,current_situation) values($1,$2)`

			id := uuid.NewString()
			_, err := db.Exec(query, id, memoryUsage)
			if err != nil {
				logs.Error(fmt.Sprintf("memory table ga xatolik: %s", err))
			}
		}

		gpuUsage, err := gpu.GetGPUUsage()
		if err == nil && gpuUsage > gpuThreshold {
			alert.SendAlert(fmt.Sprintf("Alert: GPU yuklanishi %v%% dan oshdi\nHozirgi holati -> %v%%", gpuThreshold, gpuUsage))
		} else {
			query := `insert into gpu(id,current_situation) values($1,$2)`

			id := uuid.NewString()
			gpuint := int(gpuUsage)
			if gpuint > 0 {
				_, err := db.Exec(query, id, gpuint)
				if err != nil {
					logs.Error(fmt.Sprintf("gpu table ga xatolik: %s", err))
				}
			}
		}

		if cpuUsage > 95 && memoryUsage > 95 && diskUsage > 95 {
			alert.SendAlert("Alert: CPU, Memory va Disk uchchalasining yuklanishi 95%% dan oshdi. Dastur qayta ishga tushirilmoqda...")

			err := restart.RestartApplication()
			if err != nil {
				logs.Error(fmt.Sprintf("Qayta ishga tushirishda xatolik: %s", err))
			}
		}

		// progress statistikasi
		go func() {
			cpuprogressUsage, _ := systemusage.GetCPUUsage()
			for process, usage := range cpuprogressUsage {
				if usage > cpuprogressThreshold {
					alert.SendAlert(fmt.Sprintf("Alert: %s jarayoni CPU dan %v%% foydalanmoqda", process, usage))
				}
			}
		}()

		go func() {
			memoryprogressUsage, _ := systemusage.GetMemoryUsage()
			for _, process := range memoryprogressUsage {
				if process.WorkingSetSize > memoryprogressThreshold {
					alert.SendAlert(fmt.Sprintf("Alert: %s jarayoni %v GB xotira ishlatmoqda", process.Name, float64(process.WorkingSetSize)/1024/1024/1024))
				}
			}
		}()

		go func() {
			diskprogressUsage, _ := systemusage.GetDiskUsage()
			for process, usage := range diskprogressUsage {
				if usage > diskprogressThreshold {
					alert.SendAlert(fmt.Sprintf("Alert: %s jarayoni diskdan %v MB foydalanmoqda", process, usage/1024/1024))
				}
			}
		}()

		time.Sleep(50 * time.Millisecond)
	}
}