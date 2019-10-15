package src
import "fmt"

var total,prim,maxPeri int64

func tri(l1, l2, l3) {
	if p := l1 + l2 + l3; p <= maxPeri {
		prim++
		total += maxPeri / p
		tri(+1*s0-2*s1+2*s2, +2*s0-1*s1+2*s2, +2*s0-2*s1+3*s2)
		tri(+1*s0+2*s1+2*s2, +2*s0+1*s1+2*s2, +2*s0+2*s1+3*s2)
		tri(-1*s0+2*s1+2*s2, -2*s0+1*s1+2*s2, -2*s0+2*s1+3*s2)
	}
}

func main() {
	for maxPeri = 100; maxPeri <= 1e11; maxPeri *= 10{
		prim = 0;
		total =0;
		tri(3,4,5)
		fmt.Printf("%d %d, %d\n",maxPeri,total,prim)
	}

}



