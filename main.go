package main


import (
	"fmt"
	"awesomeProject/slicewithminimalaverage"
	"math/rand"
)


func main()  {
	A := []int{4, 2, 2 ,5 ,1 ,5 , 8}
	fmt.Println(slicewithminimalaverage.SliceWithMinimalAverage(A))

	//test for big array
	B := []int{}
	for i:=0; i<=10000000; i++ {
		num := rand.Intn(100000)
		B = append(B, num)
	}
	fmt.Println(slicewithminimalaverage.SliceWithMinimalAverage(B[:]))

}


