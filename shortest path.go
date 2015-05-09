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

const (
	SIDE_DISTANCE = 10
	DIAGONAL_DISTANCE =14
)

type Position struct{
    x int
    y int
}
type Node struct{
    value   int //1 if part of the network of nodes. 0 othervise
    color   int //white, grey or black. White if noot seen yet.
    dist    int //distance from starting node. 
	pos			Position
    prevPos    	Position    //the position of the cell which it came from. 
}

type Nodes struct{
    width  int
    height int
    nodeMatrix [][]Node
	
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
            NodeMatrix[y][x].dist = 1000
            NodeMatrix[y][x].prevPos= Position{-1,-1}
			NodeMatrix[y][x].pos = Position{x,y}
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

func AbsDiff(val1, val2 int) int{
	Diff := val1-val2
	if(Diff< 0){
		return -Diff
	}
	return Diff
}

func distToNeighboor(pos1, pos2 Position) int{
	xAbsDiff := AbsDiff(pos1.x,pos2.x)
	yAbsDiff := AbsDiff(pos1.y,pos2.y)
	if(xAbsDiff > 1 || yAbsDiff >1 ){
		fmt.Println("FEIL I distToNeighboor-funksjonen")
	}
	sumDiff := xAbsDiff + yAbsDiff
	switch sumDiff{
	case 0:
		return 0
	case 1:
		return SIDE_DISTANCE
	case 2: 
		return DIAGONAL_DISTANCE
	}
	fmt.Println("FEIL I distToNeighboor-funksjonen")
	return -1
}

func updateNeighboor(nodeFrom, nodeTo *Node){
	if (nodeFrom == nodeTo){
		return
	}
	if(nodeTo.value == 1){
		nodeToDistViaNodeFrom := nodeFrom.dist + distToNeighboor(nodeFrom.pos, nodeTo.pos)
		if( nodeTo.color == WHITE){
			nodeTo.dist = nodeToDistViaNodeFrom
			nodeTo.color = GREY
			nodeTo.prevPos = nodeFrom.pos
			//legg til nodeTo i liste og i rett posisjon.
		}else if(nodeTo.color == GREY){
			if(nodeToDistViaNodeFrom < nodeTo.dist){
				nodeTo.dist = nodeToDistViaNodeFrom
				nodeTo.prevPos = nodeFrom.pos
			}
		}
	}
}

func posWithinMatrix(width, height int, pos Position) bool {
	return (pos.x >= 0 && pos.y >=0 && pos.x<width && pos.y <height)
}


func (nodes * Nodes)updateNeighboors(pos Position){
	for y:=pos.y-1;y<=pos.y+1;y++{
		for x:=pos.x-1;x<=pos.x+1;x++{
			if(posWithinMatrix(nodes.width, nodes.height, Position{x,y})){
				updateNeighboor(&nodes.nodeMatrix[pos.y][pos.x],&nodes.nodeMatrix[y][x])
			}
		}
	}
}


func (nodes * Nodes)foo(pos Position){
	
	
	
	// sjekk at pos er på brett og har verdi 1
	
	// for hver naborute
		// sjekk at pos er på brett og har verdi 1
		// sjekk om naborute er ferdig
		// se om distansen til naboruten er korterene enn den dit fra før av
		// hvis det er tilfellet: oppdater distanse til naborute, oppdater farge, oppdater prevPos hos nabo.
		// legg naboen til i listen over uferdige noder. 
	
	// sjekk at naboruter er på brett og har verdi 1
	// hvis de er hvite, farg grå og legg til i liste på rett plass.
	// 
}



func main() {
    matrix:=MatrixWithRandomIntegers(7,7,0.7)
    //dispMatrix(matrix)
    
    nodeMatrix :=createAndInitNodeMatrix(7,7,matrix)
	nodes := Nodes{7,7,nodeMatrix}
	
    valueMatrix :=createMatrixType(nodeMatrix,VALUE)
    dispMatrix(valueMatrix)
    
	
	fmt.Println(nodes.nodeMatrix[1][1])
	nodes.updateNeighboors(Position{1,1})
	
	fmt.Println(nodes.nodeMatrix[1][1])
	fmt.Println(nodes.nodeMatrix[1][0])
	fmt.Println(nodes.nodeMatrix[1][2])
	fmt.Println(nodes.nodeMatrix[0][0])
	
	
	
	
	
	
	
	
	
	
}

















