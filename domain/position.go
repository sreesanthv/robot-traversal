package domain

type Postion struct {
	X int
	Y int
}

func NewPostion(x, y int) *Postion {
	return &Postion{
		X: x,
		Y: y,
	}
}
