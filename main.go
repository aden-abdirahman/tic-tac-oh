package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 600
	screenHeight = 600
	cellSize     = 200
)

var (
	gameOver    = false
	draw        = false
	playersMove = true
	aiMove      = false
	x           rl.Texture2D
	o           rl.Texture2D
	winner      int
)

type Cell struct {
	t        rl.Texture2D
	rect     rl.Rectangle
	marked   bool
	charType string
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Tic Tac Toe")
	rl.SetMouseScale(1.0, 1.0)

	// x = rl.LoadTexture("assets/x.png")
	// o = rl.LoadTexture("assets/o.png")

	x = rl.LoadTexture("/Users/Abdirahman/Desktop/Go Projects/tic-tac-oh/assets/x.png")
	o = rl.LoadTexture("/Users/Abdirahman/Desktop/Go Projects/tic-tac-oh/assets/o.png")

	rl.SetTargetFPS(60)

	var positions int = 0
	board := [3][3]Cell{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j].rect.Width = cellSize
			board[i][j].rect.Height = cellSize

			board[i][j].rect.X = float32(j * cellSize)
			board[i][j].rect.Y = float32(i * cellSize)
		}

	}

	var mouseButtonPressed bool = false

	for !rl.WindowShouldClose() {
		if !gameOver && !draw {

			var mousePos rl.Vector2 = rl.GetMousePosition()

			// if mouse is pressed, check if cell is empty or not
			if mouseButtonPressed {
				for i := 0; i < len(board); i++ {
					for j := 0; j < len(board[i]); j++ {
						if rl.CheckCollisionPointRec(mousePos, board[i][j].rect) && !board[i][j].marked && playersMove {
							board[i][j].t = x
							board[i][j].marked = true
							board[i][j].charType = "x"
							playersMove = false
							aiMove = true
							positions++
						}

					}
					mouseButtonPressed = false
				}
			}

			// if mouse is not pressed, and cell is empty, ai will make a move
			if aiMove {
				for i := 0; i < len(board); i++ {
					for j := 0; j < len(board[i]); j++ {
						if !board[i][j].marked && aiMove {
							board[i][j].t = o
							board[i][j].marked = true
							board[i][j].charType = "o"
							aiMove = false
							playersMove = true
							positions++
						}
					}
				}
			}

			// super wordy logic to check if game is over or not and which player won
			if board[0][0].marked && board[0][1].marked && board[0][2].marked {
				if board[0][0].charType == board[0][1].charType && board[0][1].charType == board[0][2].charType {
					if board[0][0].charType == "x" {
						winner = 1
						gameOver = true
					} else {
						winner = 2
						gameOver = true
					}

				}
			}
			if board[1][0].marked && board[1][1].marked && board[1][2].marked {
				if board[1][0].charType == board[1][1].charType && board[1][1].charType == board[1][2].charType {
					if board[1][0].charType == "x" {
						winner = 1
						gameOver = true
					} else {
						winner = 2
						gameOver = true
					}
				}
			}
			if board[2][0].marked && board[2][1].marked && board[2][2].marked {
				if board[2][0].charType == board[2][1].charType && board[2][1].charType == board[2][2].charType {
					if board[2][0].charType == "x" {
						winner = 1
						gameOver = true
					} else {
						winner = 2
						gameOver = true
					}
				}
			}
			if board[0][0].marked && board[1][0].marked && board[2][0].marked {
				if board[0][0].charType == board[1][0].charType && board[1][0].charType == board[2][0].charType {
					if board[0][0].charType == "x" {
						winner = 1
						gameOver = true
					} else {
						winner = 2
						gameOver = true
					}
				}
			}
			if board[0][1].marked && board[1][1].marked && board[2][1].marked {
				if board[0][1].charType == board[1][1].charType && board[1][1].charType == board[2][1].charType {
					if board[0][1].charType == "x" {
						winner = 1
						gameOver = true
					} else {
						winner = 2
						gameOver = true
					}
				}
			}
			if board[0][2].marked && board[1][2].marked && board[2][2].marked {
				if board[0][2].charType == board[1][2].charType && board[1][2].charType == board[2][2].charType {
					if board[0][2].charType == "x" {
						winner = 1
						gameOver = true
					} else {
						winner = 2
						gameOver = true
					}
				}
			}
			if board[0][0].marked && board[1][1].marked && board[2][2].marked {
				if board[0][0].charType == board[1][1].charType && board[1][1].charType == board[2][2].charType {
					if board[0][0].charType == "x" {
						winner = 1
						gameOver = true
					} else {
						winner = 2
						gameOver = true
					}
				}
			}
			if board[0][2].marked && board[1][1].marked && board[2][0].marked {
				if board[0][2].charType == board[1][1].charType && board[1][1].charType == board[2][0].charType {
					if board[0][2].charType == "x" {
						winner = 1
						gameOver = true
					} else {
						winner = 2
						gameOver = true
					}
				}
			}

			if positions == 9 && !gameOver {
				draw = true
			}

		}

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) {
			mouseButtonPressed = true
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// draw grid on board
		rl.DrawLine(0, screenHeight/3, screenWidth, screenHeight/3, rl.White)
		rl.DrawLine(0, screenHeight*2/3, screenWidth, screenHeight*2/3, rl.White)
		rl.DrawLine(screenWidth/3, 0, screenWidth/3, screenHeight, rl.White)
		rl.DrawLine(screenWidth*2/3, 0, screenWidth*2/3, screenHeight, rl.White)

		// logic for drawing rectangles for each cell and adding textures

		// if !gameOver && !draw { commenting out condition because i want users to be able to see the last move, will update game over screen
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				// this function creates a color filled rectangle for each cell, it takes in the position of the cell, the size of the cell and the color as params
				rl.DrawRectangle(int32(board[i][j].rect.X), int32(board[i][j].rect.Y), int32(board[i][j].rect.Width-5), int32(board[i][j].rect.Height-5), rl.DarkBlue)
				// this function draws a texture on the rectangle, it takes in the position of the cell, the size of the cell and the texture as params. made sure it was centered
				rl.DrawTexture(board[i][j].t, int32(board[i][j].rect.X+board[i][j].rect.Width/2-float32(board[i][j].t.Width)/2), int32(board[i][j].rect.Y+board[i][j].rect.Height/2-float32(board[i][j].t.Height)/2), rl.White)

			}
		}
		// }

		if gameOver {
			if draw {
				rl.DrawText("It's a draw!", screenWidth/2-150, screenHeight/2, 50, rl.White)
			} else {
				rl.DrawText("Player "+strconv.Itoa(winner)+" wins!", screenWidth/2-150, screenHeight/2, 50, rl.White)
			}
		}

		rl.EndDrawing()

	}

	rl.CloseWindow()
}
