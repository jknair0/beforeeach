package examples

type Calculator struct {
	offset int
}

func NewCalculator() *Calculator {
	return &Calculator{
		offset: 0,
	}
}

func (c *Calculator) AddOffset(offset int)  {
	c.offset = c.offset + offset
}

func (c *Calculator) Add(a int, b int) int  {
	return a + b + c.offset
}
func (c *Calculator) Subtract(a int, b int) int  {
	return a - b + c.offset
}