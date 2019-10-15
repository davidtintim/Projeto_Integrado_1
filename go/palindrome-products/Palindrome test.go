package src

import (
	"fmt"
	"strconv"
	"strings"
)


func main()
{
var num, resto, temp int
var reverse int =0;

fmt.Scan(&num);

temp= num;

for
{
resto= num%10;
num= num/10;
reverse*10+resto;

if(num==0){
break;
}
}

if(temp==reverse)
{
fmt.Printf("%d is a Palindrome", temp)
}else
{
fmt.Printf("%d is not a Palindrome", temp)
}
}

