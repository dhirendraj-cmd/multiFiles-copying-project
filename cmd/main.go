package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/dhirendraj-cmd/multiFiles-copying-project/internal/srcdirs"
)


func main(){
	homeDir, _ := os.UserHomeDir()
	// homeDir = homeDir+"/"
	fmt.Println("home dir is: ", homeDir)

	// get source directry details
	fmt.Print("Enter Source Dir: ")
	srcDir := bufio.NewReader(os.Stdin)
	readSrc, _ := srcDir.ReadString('\n')
	readSrc = strings.TrimSpace(readSrc)
	readSrc = filepath.Join(homeDir, readSrc)
	fmt.Println("full src dir: ", readSrc)

	// calling movers func
	srcdirs.CallMoveFromSrc(readSrc)


	// get dest directry details
	// fmt.Print("Ennter Destination Dir: ")
	// destDir := bufio.NewReader(os.Stdin)
	// readDest, _ := destDir.ReadString('\n')
	// fmt.Println("Entered Dest Dir: ", readDest)



}