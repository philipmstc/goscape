package model

type Board struct {
	Tiles map[*Tile]Position
}

type Position struct {
	x int
	y int
}
func (pos Position) equals(other Position) bool {
	return pos.x == other.x && pos.y == other.y
}

type MoveAction struct  {
	deltaX int
	deltaY int
	name string
}

func (move MoveAction) Do(player *Player) {
	player.move(move.deltaX, move.deltaY)
}

func (move MoveAction) GetName() string {
	return move.name
}

var MoveLeft = MoveAction{-1, 0, "Go left"}
var MoveRight = MoveAction{1, 0, "Go right"}
var MoveUp = MoveAction{0, -1, "Go up"}
var MoveDown = MoveAction{0, 1, "Go down"}

func NewPosition(x int, y int) Position {
	return Position{x, y}
}

func NewGameBoard() Board {
	return Board{make(map[*Tile]Position)}
}
