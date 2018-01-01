package main

import(
	"fmt"
	"bufio"
	"os"
)

type console map[string]func()

var (
	cons console
)

func init(){
	cons =  console(map[string]func(){
		"accueils": accueils,
	})
	
}

func affe(){
	fmt.Println('\320')
}

func accueils(){
	reader := bufio.NewReader(os.Stdin)
	var text string
	affe()
	for text != "quit\n"{
		text, _ = reader.ReadString('\n')
		
		
	}
}

func consoleme(a string){
	if aa,bb := cons[a];bb{
		aa()
	}
	
}

func main(){
	ar := makeFields(4,3)
}
