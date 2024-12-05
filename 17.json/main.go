package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"张三", "李四"}}
	jsonStr, err := json.Marshal(movie)
	// jsonStr是字节码，需要string(jsonStr)转成字符串
	if err != nil {
		fmt.Println("json marshal error:", err)
	} else {
		fmt.Printf("jsonStr=%v\n", string(jsonStr))
	}

	m := Movie{}
	err = json.Unmarshal(jsonStr, &m)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
	} else {
		fmt.Println(m)
		return
	}
}
