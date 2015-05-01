package main

import "fmt"
import . "math"


const (
    BOARD_SIZE int = 8   //boardSize    
)

type Square struct{
    x int
    y int
}

func OnBoard(s Square) bool{
    return (s.x>= 0 && s.x< BOARD_SIZE && s.y>= 0 && s.y<BOARD_SIZE)     
}

func NewSquare(x int ,y int) Square{
    m := new(Square)
    m.x = x
    m.y = y
    if(!OnBoard(*m)){
        fmt.Println("Not on board")
    }
    return *m
}

func (sq Square)print(){
    fmt.Println("x: ",sq.x,",  y: ",sq.y)    
}



func KnightMove(oldSquare Square, newSquare Square) bool{
    if(OnBoard(newSquare)){
        var xDiff int = int(Abs(float64(newSquare.x-oldSquare.x)))
        var yDiff int = int(Abs(float64(newSquare.y-oldSquare.y)))
        if(xDiff == 2 && yDiff == 1 ||  xDiff == 1 && yDiff == 2 ){
            return true
        }  
    }
    return false
}


func LineMove(oldSquare Square, newSquare Square) bool {
    if(OnBoard(newSquare)){
        var xDiff int = int(Abs(float64(newSquare.x-oldSquare.x)))
        var yDiff int = int(Abs(float64(newSquare.y-oldSquare.y)))
        if(xDiff == 0 && yDiff == 0){
            return false    
        }else if(xDiff == 0 || yDiff == 0){
            return true    
        }  
    }
     return false
}

func DiagonalMove(oldSquare Square, newSquare Square) bool {
     if(OnBoard(newSquare)){
        var xDiff int = int(Abs(float64(newSquare.x-oldSquare.x)))
        var yDiff int = int(Abs(float64(newSquare.y-oldSquare.y)))
        if(xDiff == yDiff && xDiff != 0){
            return true
        }  
    }
    return false
}

func QueenMove(oldSquare Square, newSquare Square) bool {
    if(LineMove(oldSquare,newSquare) || DiagonalMove(oldSquare,newSquare)){
        return true
    }
    return false
}


func main(){
    
    
    oldSquare := Square{1,1}
    newSquare := Square{3,1}
    
    fmt.Println(LineMove(oldSquare,newSquare))
    fmt.Println(DiagonalMove(oldSquare,newSquare))
    fmt.Println(QueenMove(oldSquare,newSquare))
    
}
