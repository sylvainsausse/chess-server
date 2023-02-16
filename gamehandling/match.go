package gamehandling

import (
	"fmt"
	"net"
	"time"

	chess "github.com/sylvainsausse/chess-engine"
)

type Match struct {
	Board chess.Chessboard
	TimerW uint32
	TimerB uint32
}

func NewMatch() Match {
	return Match{Board : chess.NewChessboard(),TimerW: 100,TimerB: 100}
}

func (this Match) Start(j1 net.Conn,j2 net.Conn){
	//this.Board.Load("0x00000000000000000000000000000000000009000000A0000000000000001000")
	fmt.Print("Begin match...")
	conn := [2]net.Conn{j1,j2}
	buf := make([]byte,64)
	conn[0].Write([]byte{0x00})
	conn[1].Write([]byte{0x01})
	fmt.Println("OK!")
	for i := 0 ; !this.Board.CheckMate(i%2==1) && this.Board.Sum() > 10 && i < 10000 ; i++ {
		this.Board.Disp()
		fmt.Println(i,"-",this.Board.Sum())
		conn[i%2].Write(this.Board.Convert())
		time.Sleep(50*time.Millisecond)
		conn[i%2].Write([]byte{0xFF})
		
		conn[i%2].Read(buf)
		err := this.Board.Make_move(i%2==1,int(buf[0]),int(buf[1]),int(buf[2]),int(buf[3]))
		for err != nil {
			conn[i%2].Write([]byte{0x00})
			conn[i%2].Read(buf)
			err = this.Board.Make_move(i%2==1,int(buf[0]),int(buf[1]),int(buf[2]),int(buf[3]))
		}
		conn[i%2].Write([]byte{0xFF})
	}
	conn[0].Write([]byte{0x11})
	conn[1].Write([]byte{0x11})
	fmt.Println("-----FIN DE PARTIE-----")
	this.Board.Disp()
	if this.Board.CheckForChecks(chess.WHITE_TEAM) {
		fmt.Println("Victoire des noirs")
		fmt.Println(this.Board.GetAllPlaysDigest(chess.WHITE_TEAM))
	}else if  this.Board.CheckForChecks(chess.BLACK_TEAM) {
		fmt.Println("Victoire des Blancs")
		fmt.Println(this.Board.GetAllPlaysDigest(chess.BLACK_TEAM))

	} else {
		fmt.Println("Draw")
	}
	conn[0].Close()
	conn[1].Close()
}