package main

import "fmt"

// 程式執行入口
func main() {
	type Celsius float64    // 攝氏温度
	type Fahrenheit float64 // 華氏温度

	const (
		AbsoluteZeroC Celsius = -273.15 // 絕對零度
		FreezingC     Celsius = 0       // 結冰溫度
		BoilingC      Celsius = 100     // 沸點
	)

	func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
	func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
	
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f == 0) // "true"
	// fmt.Println(c == f) // compile error: type mismatch

	//	func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
	//	c := FToC(212.0)
	//
	// fmt.Println(c.String()) // "100°C"
}
