package xint

import (
	"math/rand"
	"time"
)

// 简单的概率算法用来生成红包
// 算法来自 https://blog.csdn.net/aaaadong/article/details/92839996
// 需求文档 https://shimo.im/sheets/XXWwHjt9XPdcG9y6/

// probabilitySample，每个档次一共几个概率事件，各自概率是多少，必须相加=10000
var probabilitySample = map[uint8][]int{
	1:  {8000, 1500, 200, 100, 100, 100},
	2:  {4900, 4400, 300, 200, 100, 100},
	10: {2500, 4000, 2500, 700, 200, 100},
}

// 标准红包概率分布，6个区间，每个区间的最大最小值
var redPacketSample = map[int]map[string]int{
	0: {"min": 1, "max": 10},
	1: {"min": 11, "max": 20},
	2: {"min": 21, "max": 40},
	3: {"min": 41, "max": 60},
	4: {"min": 61, "max": 80},
	5: {"min": 81, "max": 100},
}

// CalculateRedPacketProbability 传递概率分布的档次，获取概率值落在哪个区间
func CalculateRedPacketProbability(grade uint8) int {
	var start, end int

	probabilities := probabilitySample[grade]

	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(10000) //0-10000的随机数共10000个数

	for n, probability := range probabilities {
		end += probability
		if start <= rand && end > rand {
			return n
		}

		start = end
	}

	return 0
}

// CalculateRedPacketRatio 传递概率分布的档次，获得该区间的基础红包系数
func CalculateRedPacketRatio(grade uint8) float64 {
	probability := CalculateRedPacketProbability(grade)
	sample := redPacketSample[probability]

	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(sample["max"] - sample["min"])
	rand += sample["min"]

	return float64(rand) / 100
}
