package main


import("fmt"
"math")

type Shape interface {
	Area() float64
	Perimeter() float64

}

type Rectangle struct {
	Width float64
	Height float64

}
type Circle struct {
	Radius float64

}

func(r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func(r Rectangle) Perimeter() float64{
	return 2*(r.Width + r.Height)
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius*c.Radius
}
func (c Circle) Perimeter() float64{
	return math.Pi*2*c.Radius
}

func main(){
	var s Shape
	s = Rectangle{Width: 5, Height: 3}
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
	s = Circle{Radius: 4}
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}


































// package main
// import "fmt"


// func main (){

// var dhiru [3]int = [3]int{1,2,3}
// d :=[...]string{"a","b","c"}


// fmt.Printf("%v \n",dhiru)
// fmt.Printf("%v \n",d)
// for i,v:= range d {
// 	fmt.Println(i,v)
// }

// var sliceee []int =[]int{1,2,3}
// fmt.Printf("%v \n",sliceee)
// var ss = make([]int,0,5)
// ss =append(ss,1)
// ss =append(ss,[]int{1,2,3,4,5}...)





// fmt.Printf("%v \n",ss);
// }