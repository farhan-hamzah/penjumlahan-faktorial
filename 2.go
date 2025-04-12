package main
import "fmt"

func numSeq(n int)int{
	var i, j, hasilFak, total int
	hasilFak = n-1
	i = n
	for i > 0{
		j = i
		hasilFak = i-1
		for hasilFak > 0{
			j = j*hasilFak
			hasilFak--
		}
		total += sumNumSeq(j)
		i-=1
	}
	return total

}
func sumNumSeq(n int)int{
	var hasil int
	hasil += n
	return hasil
}
func main(){
	var num, hasil int
	fmt.Scan(&num)
	hasil = numSeq(num)
	fmt.Print(hasil)
}