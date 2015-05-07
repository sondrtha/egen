package main

import "fmt"
import "math/rand"

func createMatrix(width, height int) [][]int{
    matrix := make([][]int, height) 
    for i := range matrix {
    	matrix[i] = make([]int, width)
    }
    return matrix
}

//if p = 0.6 the function will return 1 60% of the times. 
func RandomNumber(p float64) int{
    probInInt :=int(p*100)
    if(rand.Intn(100) < probInInt){
        return 1
    }else{
        return 0    
    }  
}

func MatrixWithRandomIntegers(width, height int, probability float64) [][]int{
    matrix :=createMatrix(width,height)
    for y:=0;y<height;y++{
        for x:=0;x<width;x++{
            matrix[y][x]=RandomNumber(probability)
        }
    }
    return matrix
}

func dispMatrix(matrix[][]int){
    fmt.Println()
    for y:=0;y<len(matrix);y++{
        for x:=0;x<len(matrix[y]);x++{
            fmt.Print(matrix[y][x], " ")
        }
        fmt.Println()
    }
    fmt.Println()
}

func main() {
    matrix:=MatrixWithRandomIntegers(7,7,0.7)
    dispMatrix(matrix)
}
