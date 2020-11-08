# Rough Go
A rough local Go (baduk, weiqi) game using [Locra](https://github.com/zserge/lorca)  and a [baduk](http://godoc.org/github.com/acityinohio/baduk) library

## How do you run it?
build and then execute. A few examples:
roughgo.exe        --> will list out the acceptable command line arguments
roughgo.exe 13     --> will start a new game on a 13x13 board
roughgo.exe 15     --> will start a new game on a 15x15 board
roughgo.exe b DWJAA0zoAgyMDNgAI5parKqYsJuIaRYTLmsQABAAAP__    --> will resume the game encoded by DWJAA0zoAgyMDNgAI5parKqYsJuIaRYTLmsQABAAAP__ starting with black

## Dependencies
Locra using the chrome dev toolkit behind the scenes. This will require having chrome installed and granting the appropriate permissions during runtime.

## There is no website? no way to play against each other? 
Yeap. This is a pure visual represntation of a go board. You can share this over a video chat and play the game as you would in person. The baduk library includes features such as taking pieces, scoring, encoding, and the SVG rendering. RoughGo was meant as a way to turn that library into a working game with minimal effort.