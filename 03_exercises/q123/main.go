package main

import "fmt"

//var myname = "chris"

/*func main()
	fmt.Println("hello,", name)

}

*/
func main (){
	fmt.Println("what is your name?")
	//var askname
	var askname = ""  //, _ = fmt.Scanln()
	fmt.Scanln(&askname)
	fmt.Println("hi,",*&askname)

}
