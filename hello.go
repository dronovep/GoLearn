package main

type mp1 = int
type mp2 = int

type mfl1 = float64
type mfl2 = float64

var varfl1 mfl1 = 12.2
var varfl2 mfl2 = varfl1

var a mp1 = 3
var b mp2 = 2
var c mp1 = b
var d mp2 = mp2(a)

type Point1 struct {
	x float64
	y float64
}

type Point2 struct {
	x float64
	y float64
}

type Point3 struct {
	a float64
	b float64
}

var p1 Point1 = Point1{3, 2}
var p2 Point2 = Point2(p1) //так можно
//var p3 Point2 = Point3{15, 25} // а так нельзя
var p3 Point3 = Point3{15, 25}

var m float32 = 5.0

//var n float64 = m  // неявно преобразовать не получится
var n float64 = float64(m)
var k float32 = float32(n) // явным преобразованием снизить точность можно

var pp1 *Point1 = &p1

//var pp2 *Point2 = &p1	// неявно преобразовать не получится
var pp2 *Point2 = (*Point2)(&p1)

//var pp3 *Point3 = (*Point3)(&p1)  // так преобразовать указатели не получится.
// по факту  можно преобразовать указатель на тип А в указатель на тип в Б только если можно преобразовать значение
// типа А в значение типа Б

func main() {
	//var avar bool = bool(5) // преобразовать int в bool не получится
	//var bvar int = 3
	//
	//if bvar {		// ну и соответственно использовать в условиях не получится
	//
	//}

	//var mbytes  = [...]rune{'a', 'l', 'p', 'h', 'a'}    //можно и так
	//var mbytes  = [...]byte{'a', 'l', 'p', 'h', 'a'}  // и так

	//var mbytes = [5]rune{'a', 'l', 'p', 'h', 'a'} //можно и так
	//var msl = mbytes[:3]
	//var mstr = string(msl)
	////var mstr = string(mbytes)					// массив сконвертировать в строку не получится, только слайс можно
	//fmt.Printf("mstr = %s\n", mstr)
	//
	//var mstr2 string = "father"
	//var syms []rune = []rune(mstr2)
	//syms[0] = 'm'
	//syms[1] = 'o'
	//fmt.Printf("father : f->m, a->o : %s\n", string(syms))
	//fmt.Printf("sysms: len = %d, cap = %d, value = %v\n", len(syms), cap(syms), syms)
	//
	//var mbytes2 = [...]byte{116, 104, 101, 114}
	//var msl2 = mbytes2[:]
	//fmt.Printf("numbers equal to string: %s\n", string(msl2))
	//
	//var marr1 = [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	////var par  = (*[3]int)(&marr1[5]) // делать указатель на кусок массива не получится
	//var mslar = marr1[:]
	//var par = (*[3]int)(mslar)
	//fmt.Printf("par points to %v\n", *par)
	//par = (*[3]int)(marr1[5:])
	//fmt.Printf("par points to %v\n", *par)
	//par[0] = 125
	//fmt.Printf("marr1 = %v\n", marr1)
	//
	//var par2 = &par[1] // здесь тип *int
	//fmt.Printf("par2: %v\n", *par2)

	// Резюме! Из масива можно получить слайс, из слайса можно взять указатель на массив
	// Из массива получить указатель на подмассив не получится
	// Из слайса можно получить указатель на подслайс

	var ip = []byte{255, 255, 255, 255}
	//var ipn = int32(ip) // так не получится
	var smth interface{} = ip
	//var ipn int32 = int32(smth) // так тоже не получится
	var ipn = (*int32)(&ip[0]) // и даже так не получится
}
