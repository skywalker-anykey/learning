package hw8_8_2

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type UnitType string

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		switch t {
		case CM:
			value *= 2.54
		case Inch:
			value /= 2.54
		}
	}
	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type MDim struct {
	length Unit
	width  Unit
	height Unit
}

func (m MDim) Length() Unit {
	return m.length
}

func (m MDim) Width() Unit {
	return m.width
}

func (m MDim) Height() Unit {
	return m.height
}

func NewMDimensions(length, width, height float64) MDim {
	return MDim{Unit{length, CM}, Unit{width, CM}, Unit{height, CM}}
}

type DDim struct {
	length Unit
	width  Unit
	height Unit
}

func (d DDim) Length() Unit {
	return d.length
}

func (d DDim) Width() Unit {
	return d.width
}

func (d DDim) Height() Unit {
	return d.height
}

func NewDDimensions(length, width, height float64) DDim {
	return DDim{length: Unit{length, Inch}, width: Unit{width, Inch}, height: Unit{height, Inch}}
}
