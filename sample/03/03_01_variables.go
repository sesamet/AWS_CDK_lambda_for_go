// 宣告程式屬於哪個 package
package main

// 引入套件
import (
	"fmt"
)

// 程式執行入口
func main() {
	// var apples int32 = 1
	// var oranges int16 = 2
	// var sum int = apples + oranges // compile error
	// var sum = int(apples) + int(oranges)
	// var sum = int(apples) + int(oranges)
	// fmt.Println(sum)
	// 不同位元的整數一起運算，會出現編譯錯誤

	var apples int = 1
	var oranges int = 2
	var sum int = apples + oranges // compile error
	fmt.Println(sum)
}
