package main

import "fmt"

var factor1 float64 = 12_76.12_23e-05
var factor2 float64 = 1.1223e-05
var factor3 float64 = 12_76.12_23e-08
var factor4 float64 = 1.1223e-03
var factor5 float64 = 1.3e+3
var factor6 float64 = 0x1.p-01
var factor7 float64 = 0x0.11p0

/** экспоненциальная запись с мантиссой и экспонентой: number = mantissa * base^exponenta
В случае десятичной системы счисления base = 10, вместо p пишется e
В случае системы счисления отличной от 10, base = 2, пишется p

mantissa - это все, что идет до "p или e" в соответствующей системе счисления
*/

func main() {
	fmt.Printf(
		"factor1 = %f\n"+
			"factor2 = %f\n"+
			"factor3 = %f\n"+
			"factor4 = %f\n"+
			"factor5 = %f\n"+
			"factor6 = %f\n"+
			"factor7 = %f\n",
		factor1,
		factor2,
		factor3,
		factor4,
		factor5,
		factor6,
		factor7,
	)
}
