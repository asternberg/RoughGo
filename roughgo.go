package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/acityinohio/baduk"
	"github.com/zserge/lorca"
)

const usageMessage = "Usage: \n\tfor a new game use roughgo BOARDSIZE ie roughgo 13\n\tto load a game with white starting use roughgo w ENCODING \n\tto load a game with black starting use roughgo b ENCODING"

func main() {

	// init logging
	file, err := os.OpenFile("RoughGoInfo.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Started logging...")

	b, blackToMove, err := initBoard(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Board Initialized")
	log.Print("Board size is: " + strconv.Itoa(b.Size))

	// create ui screen
	ui, err := lorca.New("", "", 600, 800)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// bind js actions
	ui.Bind("doClickEmptyVertex", func(x, y int) {

		log.Println("Clicked empty " + strconv.Itoa(x) + " " + strconv.Itoa(y))
		var err error
		if blackToMove {
			err = b.SetB(x, y)
		}
		if !blackToMove {
			err = b.SetW(x, y)
		}
		if err != nil {
			log.Print(err)
		} else {
			blackToMove = !blackToMove
		}
		ui.Load(creatPageContent(b, blackToMove, false, false))
		return
	})
	ui.Bind("doOnKeyUp", func(key string) {
		if key == "s" {
			ui.Eval(`document.getElementById("boardgraphic").innerHTML="Scoring...";`)
			ui.Load(creatPageContent(b, blackToMove, true, false))
		}
		if key == "e" {
			ui.Eval(`document.getElementById("boardgraphic").innerHTML="Encoding...";`)
			ui.Load(creatPageContent(b, blackToMove, false, true))

		}
		if key == " " {
			blackToMove = !blackToMove
			ui.Load(creatPageContent(b, blackToMove, false, false))

		}
		return
	})

	ui.Load(creatPageContent(b, blackToMove, false, false))

	// Create UI with basic HTML passed via data URI
	// Wait until UI window is closed
	<-ui.Done()
}

// initializes board from argument array
func initBoard(args []string) (baduk.Board, bool, error) {
	argLength := len(args)
	var b baduk.Board
	var boardSize int
	var blackToMove bool
	blackToMove = true
	if argLength == 2 {
		boardSize, _ = strconv.Atoi(args[1])
		b.Init(boardSize)
	} else if argLength == 3 {
		firstArg := (strings.ToLower(args[1]))
		if firstArg == "w" {
			blackToMove = false
		} else if firstArg == "b" {
			blackToMove = true
		} else {
			fmt.Println(usageMessage)
			return b, false, errors.New("invalid parameter - run roughgo without any parameters for usage instructions")
		}
		secondArg := (args[2])
		b.Decode(secondArg)
	} else {
		fmt.Println(usageMessage)
		return b, false, errors.New("invalid parameter - run roughgo without any parameters for usage instructions")

	}
	return b, blackToMove, nil
}

// creates html/js content for board state
func creatPageContent(b baduk.Board, blackToMove bool, score bool, encode bool) string {
	var svgString string
	svgString = b.PrettySVG()

	var title string
	if blackToMove {
		title = "Black to move."
	}
	if !blackToMove {
		title = "White to move."
	}
	title += "\npress space to pass."
	if score {
		whitescore, blackscore := b.Score()
		title += "\nBlack score: " + strconv.Itoa(blackscore) + "\nWhite score: " + strconv.Itoa(whitescore) + "\n"
	} else {
		title += "\npress 's' to score."
	}

	if encode {
		encoding, _ := b.Encode()
		title += "\board encoding:<br>" + encoding + ""
	} else {
		title += "\npress 'e' to encode."
	}

	log.Print(title)
	pageBody := "data:text/html," + url.PathEscape(`
	<html>
		<head><title>Rough Go</title></head>
		<body><div id ="topdiv">`+title+`</div><div id="boardgraphic">`+svgString+`</div></body>
		<script>
		function onKeyUp(event) { doOnKeyUp(event.key);}
		function clickEmptyVertex() { doClickEmptyVertex(parseInt(this.dataset.x),parseInt(this.dataset.y));}
		var emptyvertices = document.getElementsByClassName("empty-vertex")
		for(var i = 0; i < emptyvertices.length; i++) {
            var emptyvertex = emptyvertices[i];
			emptyvertex.onclick = clickEmptyVertex;
		}
		document.addEventListener("keyup", onKeyUp);
		</script>
		<style>rect.empty-vertex:hover {opacity: 0.5}</style>
	</html>
	`)
	return pageBody
}
