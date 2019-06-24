package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func main() {

	//-- input --//
	maxCoupon := 50000
	randDistance := 8998
	minCodeLength := 3
	couponPrefix := "TDG"

	fmt.Println("--------------------------------------")
	fmt.Println("*** Start ***")

	var numCode int64
	var numCount int64
	var finishCount int
	var sb strings.Builder

	numCount = 0

	// var couponNum [maxCoupon]int64
	coupons := make([]string, maxCoupon)

	for i := 0; i < maxCoupon; i++ {

		distance := rand.Intn(randDistance)
		distance = distance + 1

		numCode = numCount + int64(distance)
		numCount = numCode

		hexCode := strconv.FormatInt(numCode, 16)
		hexCode = strings.ToUpper(hexCode)
		// hexCodeString := string(hexCode)
		hexCodeLen := len(hexCode)

		//Zero Padding
		if hexCodeLen < minCodeLength {

			for n := hexCodeLen; n < minCodeLength; n++ {
				sb.WriteString("0")
			}

			sb.WriteString(hexCode)
			hexCode = sb.String()
			sb.Reset()

		}

		//Concat Prefix
		sb.WriteString(couponPrefix)
		sb.WriteString(hexCode)
		coupon := sb.String()
		sb.Reset()

		coupons[i] = coupon

		finishCount++
		finishPercent := math.Round((float64(finishCount) * float64(100)) / float64(maxCoupon))

		// fmt.Printf("%d -> %d [%v](%d)\n", distance, numCode, string(hexCode), len(string(hexCode)))
		fmt.Printf("\rProgress -> %d/%d (%3.0f%%)", finishCount, maxCoupon, finishPercent)

	}

	fmt.Println("Summary")
	fmt.Printf("\nlen=%d cap=%d %v\n", len(coupons), cap(coupons), coupons)

	fmt.Println("*** Finish ***")
	fmt.Println("--------------------------------------")

}
