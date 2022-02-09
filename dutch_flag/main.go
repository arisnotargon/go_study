package main

import (
        "math/rand"

        "github.com/davecgh/go-spew/spew"
)

func main() {

        testArr := GetTestCase()

        spew.Dump(string(testArr))

        _ = DutchNationalFlagRes1(testArr)
        //
        //spew.Dump("res1:" + string(res))

        res := DutchNationalFlagRes2(testArr)
        spew.Dump("res2:" + string(res))
}

// 荷兰国旗问题1 冒泡法 由于已知只有三种元素,所以只要跑两次循环就行
func DutchNationalFlagRes1(inArr []int32) []int32 {
        n := len(inArr)
        p := 0
        for i := 0; i < n; i++ {
                if inArr[i] == 0 {
                        inArr[i], inArr[p] = inArr[p], inArr[i]
                        p++
                }
        }
        for i := p; i < n; i++ {
                if inArr[i] == 1 {
                        inArr[i], inArr[p] = inArr[p], inArr[i]
                        p++
                }
        }

        return inArr
}

// 只跑一次的解法
func DutchNationalFlagRes2(inArr []int32) []int32 {
        n := len(inArr)
        p1, i, p2 := -1, 0, n

        for i < p2 {
                if inArr[i] == 0 {
                        // 当i遍历到的元素为0,交换和p1的元素
                        p1++
                        inArr[i], inArr[p1] = inArr[p1], inArr[i]
                } else if inArr[i] == 2 {
                        // 为2的情况和p2的交换
                        p2--
                        inArr[i], inArr[p2] = inArr[p2], inArr[i]
                }

                i++
        }

        return inArr
}

// 生成随机地图
func GetTestCase() []int32 {
        n := 10
        ret := make([]int32, n)
        for i := 0; i < n; i++ {
                ret[i] = rand.Int31n(3)
        }

        return ret
}
