package srcdirs

import (
	"fmt"
	"os"
	"path/filepath"
	// "github.com/dhirendraj-cmd/multiFiles-copying-project/internal/destdirs"
)


func moveFromSrc(srcDir string){
	fmt.Println(srcDir)
	// files, err:=os.ReadDir(srcDir)
	abs, _:=filepath.Abs(srcDir)
	fmt.Println("Resolved path:", abs)

	files, err := os.ReadDir(srcDir)
	if err!=nil{
		fmt.Println("Can't Read directory!!!!", err)
		return
	}

	for _, file := range files{
		fmt.Println("file name: ", file.Name(), file.IsDir())
	}
}


func CallMoveFromSrc(srcDir string){
	moveFromSrc(srcDir)

}
