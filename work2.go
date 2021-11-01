package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

var CpuDataList []float64
var MemUseDataList []int64

var MaxCpu float64
var MinCpu  float64
var SumCpu float64
var MaxMemUse int64
var MinMemUse int64
var SumMemUse int64

func init(){
	MaxCpu = 0.0
	MinCpu = 100.0
	SumCpu = 0.0
	MaxMemUse = 0
	MinMemUse = int64(math.Pow(2,32)-1)
	SumMemUse = 0
}

func main(){
	data,err := os.Open("C:\\Users\\27761\\Desktop\\work2test10.txt")
	if err != nil{
		fmt.Println("open file err = ",err)
	}
	defer data.Close()
	CpuStrReg := regexp.MustCompile("\\d{0,3}.\\d.id")
	CpuDataReg := regexp.MustCompile("\\d{0,3}.\\d")
	MemStrlineReg := regexp.MustCompile("KiB Mem")
	MemUseDataStrReg := regexp.MustCompile("[0-9]+.used")
	MemUseDataReg := regexp.MustCompile("[0-9]+")
	if CpuStrReg == nil || CpuDataReg == nil || MemStrlineReg ==nil || MemUseDataStrReg == nil||MemUseDataReg == nil{
		fmt.Println("MustComplie err")
	}
	reader :=bufio.NewReader(data)
	for{
		Str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		CpuStr := CpuStrReg.FindString(Str)
		MemStrline := MemStrlineReg.FindString(Str)
		if CpuStr != ""{
			CpuData,_ := strconv.ParseFloat(CpuDataReg.FindString(CpuStr),64)
			CpuData = 100.0-CpuData
			if MaxCpu < CpuData{
				MaxCpu = CpuData
			}else if MinCpu > CpuData{
				MinCpu = CpuData
			}
			SumCpu += CpuData
			CpuDataList = append(CpuDataList,CpuData)
		}else if MemStrline != ""{
			MemUseDataStr := MemUseDataStrReg.FindString(Str)
			MemUseData,_ := strconv.ParseInt(MemUseDataReg.FindString(MemUseDataStr),10,64)
			MemUseData /= 1000
			if MaxMemUse < MemUseData{
				MaxMemUse = MemUseData
			}else if MinMemUse > MemUseData{
				MinMemUse = MemUseData
			}
			SumMemUse += MemUseData
			MemUseDataList = append(MemUseDataList,MemUseData)
		}
	}
	AvgCpu := SumCpu/float64(len(CpuDataList))
	fmt.Println("Cpu利用率")
	for _,v := range CpuDataList{
		fmt.Printf("%.1f\n",v)
	}
	fmt.Printf("Max:%.1f Min:%.1f ave:%.1f",MaxCpu,MinCpu,AvgCpu)

	AvgMemUse := SumMemUse/int64(len(MemUseDataList))
	fmt.Println("\n内存使用量")
	for _,v := range MemUseDataList{
		fmt.Printf("%d\n",v)
	}
	fmt.Printf("MaxMemUse:%d MinMemUse:%d AvgMemUse:%d",MaxMemUse,MinMemUse,AvgMemUse)
}