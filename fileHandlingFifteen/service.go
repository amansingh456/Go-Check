package filehandlingfifteen

import (
	"fmt"
	"io"
	"os"
)

func Smaple() {
	fmt.Println("file handling")

	// let's create a file
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// agr file create ho gya h to, humne jo resources use kiya hai "os" usko band krdo kyuki bad me issue create kr skta  hai isliye jab ye function complete ho jaye us se just ek step pehle close do

	//! yha se humne content add kr diya file me
	// content := "12345"
	// _, errs := io.WriteString(file, content)
	// if errs != nil {
	// 	fmt.Println(errs)
	// }

	//! now file read krna // this is long process
	file2, err := os.Open("example.txt")
	if err != nil {
		return
	}
	buffer := make([]byte, 1024) // buffer is temproray store to store data for reading and holding
	for {
		n, err := file2.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("error while reading the file")
		}

		fmt.Println(string(buffer[:n]), "file read") // buffer me jo bhi data hai usko poora read krna
	}
	defer file2.Close()
	//!

	//* this is a short process , if files size is very much then don't use it -  use the above one
	content, err := os.ReadFile("example.txt") // previously we uses this ioutil.Readfile()
	if err != nil {
		fmt.Println("error is :", err)
		return
	}
	// fmt.Println(content) // this will array of ascaivalue
	fmt.Println(string(content))
	//*

}
