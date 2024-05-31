package hw8_8_2

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type Bmw struct {
	brand       string
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

func (b Bmw) Brand() string {
	return b.brand
}

func (b Bmw) Model() string {
	return b.model
}

func (b Bmw) Dimensions() Dimensions {
	return b.dimensions
}

func (b Bmw) MaxSpeed() int {
	return b.maxSpeed
}

func (b Bmw) EnginePower() int {
	return b.enginePower
}

func NewBMW(model string, maxSpeed, enginePower int, dimensions Dimensions) Bmw {
	return Bmw{brand: "BMW", model: model, maxSpeed: maxSpeed, enginePower: enginePower, dimensions: dimensions}
}

type Mercedes struct {
	brand       string
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

func (m Mercedes) Brand() string {
	return m.brand
}

func (m Mercedes) Model() string {
	return m.model
}

func (m Mercedes) Dimensions() Dimensions {
	return m.dimensions
}

func (m Mercedes) MaxSpeed() int {
	return m.maxSpeed
}

func (m Mercedes) EnginePower() int {
	return m.enginePower
}

func NewMercedes(model string, maxSpeed, enginePower int, dimensions Dimensions) Mercedes {
	return Mercedes{brand: "Mercedes", model: model, maxSpeed: maxSpeed, enginePower: enginePower, dimensions: dimensions}
}

type Dodge struct {
	brand       string
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

func (d Dodge) Brand() string {
	return d.brand
}

func (d Dodge) Model() string {
	return d.model
}

func (d Dodge) Dimensions() Dimensions {
	return d.dimensions
}

func (d Dodge) MaxSpeed() int {
	return d.maxSpeed
}

func (d Dodge) EnginePower() int {
	return d.enginePower
}

func NewDodge(model string, maxSpeed, enginePower int, dimensions Dimensions) Dodge {
	return Dodge{brand: "Dodge", model: model, maxSpeed: maxSpeed, enginePower: enginePower, dimensions: dimensions}
}
