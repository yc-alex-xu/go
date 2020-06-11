package main

import "fmt"

type skill []string

type human struct {
	name   string
	age    int
	weight int
}

type student struct {
	human      // 匿名字段，struct
	skill      // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func main() {
	//初始化一部分field
	jane := student{human: human{"Jane", 35, 100}, speciality: "Biology"}
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	jane.skill = []string{"anatomy"}
	jane.skill = append(jane.skill, "physics", "golang")
	fmt.Println("Her skill now are ", jane.skill)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}
