package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	p1 := s.start.x + int(s.a)
	p2 := s.start.y + int(s.a)

	return Point{p1, p2}
}

func (s *Square) Area() uint {
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	return s.a * 4
}
