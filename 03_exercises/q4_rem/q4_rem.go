package main

import "fmt"

func main(){
	var (
		num_little=""
		num_big=""
	)
	print("Enter a little number:")
	fmt.Scanln(&num_little)
	print("Enter a big number:")
	fmt.Scanln(&num_big)
	println("%i / %i = %i r %i", num_big, num_little, (int(num_big)/int(num_little)), num_big%num_little)

}
