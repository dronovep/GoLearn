package figures

// Важно!!! наследование методов произойдет именно в том случае, если компоненту не будет присвоено имя,
// а только указан его типа - который как правило с заглавной буквы и экспортируется
// Вот только в этом случае будем имитация наследования
type Circle struct {
	Point
	radius float64
}

func (o *Circle) SetRadius(radius float64) {
	o.radius = radius
}

func (o *Circle) GetRadius() float64 {
	return o.radius
}

func CreateCircle(x float64, y float64, radius float64) *Circle {
	return &Circle{Point{x, y}, radius}
}

// здесь мы переопределяем базовый метод GetLength поинта на специальный метод серкла
func (o *Circle) GetLength() float64 {
	return 2 * o.radius * 3.14
}

// вот так мы можем вызвать метод GetLength чисто из поинта
func (o *Circle) GetDistance() float64 {
	return o.Point.GetLength()
}
