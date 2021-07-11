package main

import (
	"fmt"
	"math"
	"time"
)

const (
	WindowSize int64 = 10  	
	LatticeNum int64 = 10 	
	Limit      int64 = 20	
)


// 计算每个小格子内的数量
var CountLatticeNum [LatticeNum]int64

//开始时间
var StartTime = time.Now().Unix()

// 当前窗口索引
var index int64 = 0
func main() {
	go func() {
		for i := 1; i <= 200; i++ {
			time.Sleep(1 * time.Second)
		}
	}()

	// 模拟请求20次
	for i := 1; i <= 200; i++ {
		if IsPass() {
			fmt.Printf("第%v次请求放行\n", i)
		} else {
			fmt.Printf("第%v次请求拒绝\n", i)
		}
		fmt.Print("等待500 Millisecond\n")

		time.Sleep(400 * time.Millisecond)
	}

}

// IsPass 是否放行
func IsPass() bool {
	slideWindow()
	var c int64 = 0
	for _, v := range CountLatticeNum {
		c += v
	}

	if c >= Limit {
		return false
	} else {
		CountLatticeNum[index] ++
		return true
	}
}

//slideWindow 滑动窗口
func slideWindow() {
	requireTime := time.Now().Unix() + WindowSize
	//计算需要滑动的小格子数
	minWindowNum := int64(math.Max(float64(requireTime)-float64(StartTime)-float64(WindowSize), 0)) / (WindowSize / LatticeNum)
	if minWindowNum <= 0 {
		return
	}
	slideNum := int64(math.Min(float64(minWindowNum), float64(LatticeNum)))
	fmt.Printf("slideNum:%v\n", slideNum)
	for i := int64(0); i < slideNum; i++ {
		//计算当前窗口索引
		index = (index + 1) % LatticeNum
		CountLatticeNum[index] = 0
	}
	StartTime = StartTime + slideNum*(WindowSize/LatticeNum)
}