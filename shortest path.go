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

const (
        WHITE = iota
        GREY  = iota
        BLACK = iota 
)

const (
    VALUE   = iota
    COLOR   = iota
    DIST    = iota
)

type Position struct{
    x int
    y int
}
type Node struct{
    value   int //1 if part of the network of nodes. 0 othervise
    color   int //white, grey or black. White if noot seen yet.
    dist    int //distance from starting node. 
    prevPos    Position    //the position of the cell which it came from. 
}
type NodeMatrix struct{
    width  int
    height int
    nodes[][]int
}

func createNodeMatrix(width,height int)[][]Node{
    matrix := make([][]Node, height) 
    for i := range matrix {
    	matrix[i] = make([]Node, width)
    }
    return matrix
}

func createAndInitNodeMatrix(width,height int, matrix[][]int)[][]Node{
    NodeMatrix := createNodeMatrix(width,height)
    for y:=0;y<len(NodeMatrix);y++{
        for x:=0;x<len(NodeMatrix[y]);x++{
            NodeMatrix[y][x].value = matrix[y][x]
            NodeMatrix[y][x].dist = -1
            NodeMatrix[y][x].prevPos=Position{-1,-1}
        }
    }
    
    return NodeMatrix
}

func createMatrixType(NodeMatrix [][]Node, matrixType int) [][]int{
    height := len(NodeMatrix)
    width  := len(NodeMatrix[0])
    matrix :=createMatrix(width,height)
    for y:=0;y<len(NodeMatrix);y++{
        for x:=0;x<len(NodeMatrix[y]);x++{
            switch matrixType{
                case VALUE:
                    matrix[y][x] = NodeMatrix[y][x].value
                case COLOR:
                    matrix[y][x] = NodeMatrix[y][x].color
                case DIST:
                    matrix[y][x] = NodeMatrix[y][x].dist
            }
        }
    }
    return matrix
}


func main() {
    matrix:=MatrixWithRandomIntegers(7,7,0.7)
    dispMatrix(matrix)
    
    m :=createAndInitNodeMatrix(7,7,matrix)
    valueMatrix :=createMatrixType(m,VALUE)
    dispMatrix(valueMatrix)
    
}

















