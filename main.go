package main

import (
	"fmt"
	"os"
	"bufio"
	
)



func drawTable( tic [][]string) {
		fmt.Println(tic)
}

func modifyTable(tic [][]string, position string, Type string) [][]string {

		
		switch position {
		case "a1": tic[0][0]=Type
		case "a2": tic[0][1]=Type
		case "a3": tic[0][2]=Type
		case "b1": tic[1][0]=Type
		case "b2": tic[1][1]=Type
		case "b3": tic[1][2]=Type
		case "c1": tic[2][0]=Type
		case "c2": tic[2][1]=Type
		case "c3": tic[2][2]=Type

		}
		return tic
}

func makeMove(tic [][]string, computerType string) [][]string  {
	turn :=true //keep the track of turn
	for i:=range tic{
		for j:=range tic[i] {
			if tic[i][j] ==" " {
				if turn==true{ //is it your turn?
			tic[i][j]=computerType
			turn=false
			break
			}
		}
		
		
	}
	
	
}
return tic
}

func winner(tic [][]string){

}

func display(tic [][]string){
	fmt.Println(tic[0][0], "|", tic[0][1], "|", tic[0][2])
	fmt.Println(tic[1][0], "|", tic[1][1], "|", tic[1][2])
	fmt.Println(tic[2][0], "|", tic[2][1], "|", tic[2][2])
}

func main() {
	p:=fmt.Println
	var computerType string
	computerType="O"
	tic := [][]string{
		{" "," "," "},
		{" "," "," "},
		{" "," "," "},
	}

	///////////////////////////////////////////////////////////
	//The player determines its sign, X or O
	p("Please enter your preferred sign ( X or O ) : ")
	reader := bufio.NewReader(os.Stdin)
	humanType,_ := reader.ReadString('\n')
		fmt.Println("Human Type: ",humanType)
	///////////////////////////////////////////////////////////

	///////////////////////////////////////////////////////////
	//The player determines the starter of the game
	p("Do you want to be the starter of the game ? Please enter Y or N : ")
	reader = bufio.NewReader(os.Stdin)
	starter, _, _ := reader.ReadLine()
	starter2 := string(starter)
	//p("starter = ", starter)
	if starter2 =="Y" {
		p("You start the game") 
	} else if starter2 =="N"{ 
		p("Computer starts the game")
	}

	/////////////////////////////////////////////////////////////

	tic = modifyTable(tic, "b1", "X" )
	
	tic = makeMove(tic, computerType)
	/*drawTable(tic)	
	tic = makeMove(tic, computerType)
	drawTable(tic)
	tic = makeMove(tic, computerType)
	drawTable(tic)
	tic = makeMove(tic, computerType)
	drawTable(tic) */
	display(tic)
}