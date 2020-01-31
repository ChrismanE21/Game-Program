package main
import "fmt"

type piece int

const ( //sets up goable varables
    blank piece = iota
    x
    o
    cat
)

type TicTacToeBoard struct { //sets up array type

	board []piece
}

func NewTicTacToeBoard() *TicTacToeBoard { //New game
    return &TicTacToeBoard{[]piece{blank,blank,blank,blank,blank,blank,blank,blank,blank}}
}

func (b TicTacToeBoard) String() string {
    s:=""
    for i,v :=  range b.board {
		
        if v == blank {//illastrates moves
            s+="   "
        } else if v == x {
            s+=" x "
        } else if v == o {
            s+=" o "
        }
        
        if i%3 < 2 {//illastrates board
            s+="|"
        }
        
        if i%3 == 2 && i/3 < 2 {//illastrates board
            s+="\n---+---+---\n"
        }
    }
    return s //fmt.Sprintf("%", b)
}

func (b *TicTacToeBoard) place(location int) {
    
}

func (b *TicTacToeBoard) getPlayer() piece {//keeps track of player turn and when the board is filled out 
    c := 9 
    for _ , v :=  range b.board {
        if v == blank {
            c--
        }
    }
    if c%2 == 0 {
        return x
    }
    return o
}

func (b *TicTacToeBoard) getWinner() int {//Chenks when win
	win := 0 //0 no one won //1 x won//2 o won//
	
		if(b.board[0] == x && b.board[1] == x && b.board[2] == x){ //X Wins
			win = 1
		}
		if(b.board[3] == x && b.board[4] == x && b.board[5] == x){ //X Wins
			win = 1
		}
		if(b.board[6] == x && b.board[7] == x && b.board[8] == x){ //X Wins
			win = 1
		}
		if(b.board[0] == x && b.board[3] == x && b.board[6] == x){ //X Wins
			win = 1
		}
		if(b.board[1] == x && b.board[4] == x && b.board[7] == x){ //X Wins
			win = 1
		}
		if(b.board[2] == x && b.board[5] == x && b.board[8] == x){ //X Wins
			win = 1
		}
		if(b.board[0] == x && b.board[4] == x && b.board[8] == x){ //X Wins
			win = 1
		}
		if(b.board[6] == x && b.board[4] == x && b.board[2] == x){ //X Wins
			win = 1
		}
		////////////////////////////////////////////////////////////////////
		
		if(b.board[0] == o && b.board[1] == o && b.board[2] == o){ //O Wins
			win = 2
		}
		if(b.board[3] == o && b.board[4] == o && b.board[5] == o){ //O Wins
			win = 2
		}
		if(b.board[6] == o && b.board[7] == o && b.board[8] == o){ //O Wins
			win = 2
		}
		if(b.board[0] == o && b.board[3] == o && b.board[6] == o){ //O Wins
			win = 2
		}
		if(b.board[1] == o && b.board[4] == o && b.board[7] == o){ //O Wins
			win = 2
		}
		if(b.board[2] == o && b.board[5] == o && b.board[8] == o){ //O Wins
			win = 2
		}
		if(b.board[0] == o && b.board[4] == o && b.board[8] == o){ //O Wins
			win = 2
		}
		if(b.board[8] == o && b.board[4] == o && b.board[2] == o){ //O Wins
			win = 2
		}
		
    return win
}

func main() {
	var move int
	var win int = 0
	
	for true{
		board:=NewTicTacToeBoard()
		for i :=0; i < 10; i++ {
			fmt.Print("Enter your move.  Move(1-9): ")
			fmt.Scanln(&move)
			
			board.board[move - 1] = board.getPlayer()
			
			fmt.Println(board)
			
			win = board.getWinner()
			
			if win == 1 {
				i = 10
				fmt.Println("")
				fmt.Println("X Wins")
				fmt.Println("")
				fmt.Println("///////////////////////////////////////////////////")
				fmt.Println("New Game Started")
				fmt.Println("")
			}
			if win == 2 {
				i = 10
				fmt.Println("")
				fmt.Println("O Wins")
				fmt.Println("")
				fmt.Println("///////////////////////////////////////////////////")
				fmt.Println("New Game Started")
				fmt.Println("")
			}
			if win == 0 && i == 8 {
				fmt.Println("")
				fmt.Println("Draw!")
				fmt.Println("")
				fmt.Println("///////////////////////////////////////////////////")
				fmt.Println("New Game Started")
				fmt.Println("")
			}
		}
	}
}
		/*
		fmt.Println(board.getPlayer())
		board.board[3]=board.getPlayer()
		fmt.Println(board.getPlayer())    Keeping this for refence
		board.board[6]=board.getPlayer()
		board.board[1]=board.getPlayer()
		fmt.Println(board)
		*/
	

