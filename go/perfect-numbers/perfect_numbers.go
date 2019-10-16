package perfect

import "fmt"

var i, num, sum int

sum = 0
func main()  {
	for i=1;i< num ;i++  {
		if (num % i = 0) {
			sum= sum + 1
		}
	}
	if(sum == num){
		fmt.Printf("%d is a Perfect Number", num)
	}
	else {
		fmt.Printf("%d not a Perfect Number", num)
	}

}

