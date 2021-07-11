package main

import "fmt"

func main() {
	// 声明map
	var aMap map[string]int
	if aMap == nil {
		fmt.Println("aMap is nil, need make")
		// 创建map
		aMap = make(map[string]int)
		fmt.Println("aMap:", aMap)
		// aMap: map[]

		// 插入数据
		aMap["a"] = 1
		aMap["b"] = 2
		aMap["c"] = 3

		fmt.Println("aMap:", aMap)
		// aMap: map[a:1 b:2 c:3]

		// aMap 大小
		fmt.Println("aMap length is", len(aMap))
		// aMap length is 3

		// 访问map
		fmt.Println("aMap c value is:", aMap["c"])
		// aMap c value is: 3

		// 访问不存在的key
		fmt.Println("aMap w value is:", aMap["w"])
		// aMap w value is: 0

		// 检查key是否存在
		value, ok := aMap["w"]
		if ok == true {
			fmt.Println("w of aMap", "is", value)
		} else {
			fmt.Println("w of aMap not found")
		}
		// w of aMap not found

		// 遍历aMap
		for key, value := range aMap {
			fmt.Printf("aMap[%s] = %d \n", key, value)
		}
		/**
		aMap[a] = 1
		aMap[b] = 2
		aMap[c] = 3
		*/

		// 删除key
		delete(aMap, "c")
		fmt.Println("aMap:", aMap)
		// aMap: map[a:1 b:2]

	}
}
