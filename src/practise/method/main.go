package main

import "fmt"

const (
	//WHITE  enum 0
	WHITE = iota
	//BLACK enum 1
	BLACK
	//BLUE enum
	BLUE
	//RED  enum
	RED
	//YELLOW enum
	YELLOW
)

//Color  bye
type Color byte

//Box struct to describe a box with color
type Box struct {
	width, height, depth float64
	color                Color
}

//BoxList a slice of boxes
type BoxList []Box

func (b Box) volume() float64 {
	return b.width * b.height * b.depth
}

//指针作为receiver,修改其field的值
func (b *Box) setColor(c Color) {
	b.color = c //等价与*b.color
}

func (bl BoxList) biggestColor() Color {
	max := 0.00
	kMax := 0
	//k := Color(WHITE)
	for k, v := range bl {
		if t := v.volume(); t > max {
			max = t
			kMax = k
		}
	}
	return bl[kMax].color
}

func (bl BoxList) paint2Bllack() {
	for i := range bl {
		bl[i].setColor(BLACK)
	}
}

//通过实现Strnger interface 输出的就是字符如"YELLOW",而不是c/c++中的 0,1,2这种
func (c Color) String() string {
	//[...] instead of []: it ensures you get a (fixed size) array instead of a slice.
	strings := [...]string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].volume(), "cm³")

	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color)
	fmt.Println("The biggest one is", boxes.biggestColor())

	fmt.Println("Let's paint them all black")
	boxes.paint2Bllack()
	fmt.Println("The box list are:", boxes)
	fmt.Println("Obviously, now, the biggest one is", boxes.biggestColor())
}
