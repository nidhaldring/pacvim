package main

type Cursor struct {
	x, y int
}

func (c *Cursor) MoveUp() {
	c.y--
}

func (c *Cursor) MoveDown() {
	c.y++
}

func (c *Cursor) MoveLeft() {
	c.y++
}

func (c *Cursor) MoveRight() {
	c.y++
}

func (c *Cursor) GetPos() (int, int) {
	return c.x, c.y
}
