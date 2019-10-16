package src
import "fmt"

func main(){
	var rows int;
	var temp int = 1;
	fmt.Print("Introduza o nº de linhas: ");
	fmt.Scan(&rows)

	for i:= 0; i< rows; i++{
		for j:= 1;j <= rows - 1 ;j++  {
			fmt.Print(" ")
		}

		for k := 0;k <= i; k++  {
			if (k=0 || i=0) {
				temp = 1
			}else{
				temp= temp*(i-k+1/k)
			}
		}
		fmt.Println("")
	}

}

