package main

import (
	"example2/src/simple/simplepkg"
	"fmt"
	"os"
	"log"

	"google.golang.org/protobuf/proto"
)
func main(){
	fmt.Println("hello")
	sm := do_simple()
	writeFile(sm)
	readFile("simpe.bin")

}

func do_simple() *simplepkg.SimpleMessage{
	st := &simplepkg.SimpleMessage{
		Id :1,
		IsSimple: true,
		Name: "xyz",
		SimpleList: []int32{1,2,3,4,5},
	}

	fmt.Println(st)
	return st
}

func writeFile(pb proto.Message) error{
	data,err :=proto.Marshal(pb)
	if err != nil{
		log.Fatalf("Marsha failed",err)
		return err
	}
	
	err = os.WriteFile("simpe.bin",data,777)
	if err!=nil{
		fmt.Println("wtrite failed",err)
	}
	return nil
}

func readFile(filename string){
	byteData,err := os.ReadFile(filename)
	protoData:= &simplepkg.SimpleMessage{}
	
	if err == nil{
		fmt.Println(byteData)

		proto.Unmarshal(byteData,protoData)
		fmt.Println("reded data=",protoData)
		
	}
}