package summultiples

func main(lim int,divisores ...int)(sum int){
	for i :=1;i < lim;i++  {
		for _, d := range divisores {
			if d != 0 && i%d ==0{
				sum +=1
				break
			}
		}
	}
	return
}