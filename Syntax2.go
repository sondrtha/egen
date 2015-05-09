SLICE - eksempel
package main
import "fmt"
// S :=make([]int,10)
// D = append(D,4)

func sortSlice(S []int) []int{
	val :=S[len(S)-1]
	rS :=make([]int,0)
	for i:=0;i<len(S)-1;i++{
		if(S[i] > val){
			lowPart := S[:i]
			heighPart:= S[i:len(S)-1]
			rS=append(rS,lowPart...)
			rS=append(rS,val)
			rS=append(rS,heighPart...)
			return rS
		}
	}
	fmt.Println("ERROR sortSlice")
	return rS
}

func main() {
	D :=[]int{1,2,3,5,6}
	D = append(D,4)
	S:=sortSlice(D)
	fmt.Println(S)
}
