package Compound

type Ruble struct {
	name string
}

func (o *Ruble) GetValute() string {
	return "This is Ruble valute"
}

type Dollar struct {
	value int
}

func (o *Dollar) GetValute() string {
	return "This is Dollar valute"
}

type Compound struct {
	Ruble
	Dollar
}

func CreateCompound() *Compound {
	return &Compound{Ruble{"RU"}, Dollar{100}}
}

//Не получится вызвать неоднозначный метод, даже если определить собственный метод
//func GetValute(o *Compound) string {
//	return "This is compound valute"
//}
