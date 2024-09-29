package main

func main() {
	game := SetUpGame()

	game.LoadFromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	game.PrintState()
}
