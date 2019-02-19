package main

import "fmt"

func reversedstring (inputstr string)(result string){

	for _, v := range inputstr {
		result = string(v) + result
	}
	return result

}

func main(){
	fmt.Println(reversedstring("fghjk"))
}