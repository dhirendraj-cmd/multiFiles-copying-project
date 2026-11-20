package main

import (
	"bufio"
	"fmt"
	"os"
	// "github.com/dhirendraj-cmd/multiFiles-copying-project/internal/srcdirs"
)


func main(){
	homeDir, _ := os.UserHomeDir()
	fmt.Println("home dir is: ", homeDir)

	// get source directry details
	fmt.Print("Enter Source Dir: ")
	srcDir := bufio.NewReader(os.Stdin)
	readSrc, _ := srcDir.ReadString('\n')
	fmt.Println("Entered Source Dir: ", readSrc)

	// get dest directry details
	fmt.Print("Ennter Destination Dir: ")
	destDir := bufio.NewReader(os.Stdin)
	readDest, _ := destDir.ReadString('\n')
	fmt.Println("Entered Dest Dir: ", readDest)

}