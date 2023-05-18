package main

import "fmt"

func test1() {
	m1 := make(map[string]string)
	m1["name"] = "yukai"
	m1["age"] = "18"
	fmt.Printf("m1: %v\n", m1)
}

func test2() {
	m2 := map[string]string{
		"name": "yukai",
		"age":  "18",
	}
	fmt.Printf("m2: %v\n", m2)
}

func test3() {
	m3 := map[string]string{
		"name": "yukai",
		"age":  "18",
	}

	v, ok := m3["name"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("------------------")
	v, ok = m3["age1"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
}

func test4() {
	m4 := map[string]string{
		"name": "yukai",
		"age":  "18",
	}

	for k, v := range m4 {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

}

func main() {
	test1()
	test2()
	test3()
	test4()
}
