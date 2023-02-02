// 宣告程式屬於哪個 package
package main

// 引入套件
import (
	"fmt"
)

// 程式執行入口
func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)

	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 34
	delete(ages, "alice")
	fmt.Println(ages)
	fmt.Println(ages["charlie"])
}
