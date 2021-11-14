package Interfaces

type IHaveCoordinates interface {
	GetX() float64
	GetY() float64
}

type IHaveLength interface {
	GetLength() float64
}

type IamFigure interface {
	IHaveCoordinates
	IHaveLength
}

type Figure = IamFigure
