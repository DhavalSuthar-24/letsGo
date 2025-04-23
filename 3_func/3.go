package main
import "fmt"
func summ(numbers ...int)int{
	total :=0
	for _,num := range numbers{
		total +=num
	}
	return total
}
func main (){
	 dhiru :=summ(1,2,3)
	 fmt.Println(dhiru)
}