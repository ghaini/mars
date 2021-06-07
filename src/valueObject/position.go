package valueObject

type Position Point

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func (p *Position) AddX() {
	p.x++
}

func (p *Position) SubX() {
	p.x--
}

func (p *Position) AddY() {
	p.y++
}

func (p *Position) SubY() {
	p.y--
}

func (p *Position) GetX() int {
	return p.x
}

func (p *Position) GetY() int {
	return p.y
}

func (p *Position) IsInsideArea(maxPoint Point) bool {
	if p.x < 0 ||
		p.y < 0 ||
		p.x > maxPoint.x ||
		p.y > maxPoint.y {
		return false
	}

	return true
}
