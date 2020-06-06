package main

import "fmt"

type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human  //匿名字段
	school string
	loan   float32
}

type employee struct {
	human   //匿名字段
	company string
	money   float32
}

//human实现SayHi方法
func (h human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//human实现Sing方法
func (h human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//employee重载human的SayHi方法
func (e employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface men被human,student和employee实现
// 因为这三个类型都实现了这两个方法
type men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := student{human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := student{human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := employee{human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := employee{human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义了slice men
	fmt.Println("Let's use a slice of men and see what happens")
	x := make([]men, 4)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2], x[3] = paul, sam, mike, tom

	songs := []string{"November rain", "Born to be wild"}
	for k, v := range x {
		v.SayHi()
		v.Sing(songs[k%len(songs)]) //顺序选歌
	}
}
