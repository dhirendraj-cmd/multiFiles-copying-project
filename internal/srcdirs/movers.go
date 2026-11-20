package srcdirs

import (
	"fmt"
	"github.com/dhirendraj-cmd/multiFiles-copying-project/internal/destdirs"
)


func moveFromSrc(){
	fmt.Println("Movers code here.....")
}


func CallMoveFromSrc(){
	moveFromSrc()
	destdirs.ReceiveFromSrc()
}
