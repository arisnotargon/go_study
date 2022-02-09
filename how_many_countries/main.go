package main

import (
        "math/rand"
        "strconv"

        "github.com/davecgh/go-spew/spew"
)

func main() {
        testMap := getRandomMap()

        for _, line := range testMap {
                spew.Dump(string(line))
        }

        res := sumCountries(testMap)

        spew.Dump("结果为", res)
}

func sumCountries(inMap [][]int32) int {
        countriesCount := 0
        cache := make(map[string]int)

        n := len(inMap)

        // 从左上开始遍历
        for i := 0; i < n; i++ {
                for j := 0; j < n; j++ {
                        tempKey := strconv.Itoa(i) + "_" + strconv.Itoa(j)
                        tempCountry, ok := cache[tempKey]
                        if !ok {
                                // 不存在的话认为是一个新国家
                                countriesCount++
                                tempCountry = countriesCount
                                cache[tempKey] = tempCountry
                        }

                        // 如果不是最后一行,检查它下 面的是否相同颜色
                        if i < n-1 {
                                if inMap[i+1][j] == inMap[i][j] {
                                        cache[strconv.Itoa(i+1)+"_"+strconv.Itoa(j)] = tempCountry
                                        spew.Dump("相同", cache, i, j)
                                }
                        }

                        // 如果不是最后一列,检查它右 边的是否相同颜色
                        if j < n-1 {
                                if inMap[i][j+1] == inMap[i][j] {
                                        cache[strconv.Itoa(i)+"_"+strconv.Itoa(j+1)] = tempCountry
                                        spew.Dump("相同", cache, i, j)
                                }
                        }
                }
        }

        return countriesCount
}

// 生成随机地图
func getRandomMap() [][]int32 {
        n := 30
        ret := make([][]int32, n)
        for i := 0; i < n; i++ {
                ret[i] = make([]int32, n)
                for j := 0; j < n; j++ {
                        ret[i][j] = rand.Int31n(4)
                }
        }

        return ret
}
