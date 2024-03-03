package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string //название операционной системы
}

// Добавьте в пакет electronic структуру applePhone. Сделайте так, чтобы она реализовывала интерфейсы Phone и Smartphone.
type applePhone struct {
	model string
	os    string
}

func (a *applePhone) Brand() string {
	return "apple"
}

func (a *applePhone) Model() string {
	return a.model
}

func (a *applePhone) Type() string {
	return "smartphone"
}

func (a *applePhone) OS() string {
	return a.os
}

func NewApplePhone(model, os string) applePhone {
	return applePhone{model: model, os: os}
}

// Добавьте в пакет electronic структуру androidPhone. Сделайте так, чтобы она также реализовывала интерфейсы Phone и Smartphone.
type androidPhone struct {
	brand string
	model string
	os    string
}

func (a *androidPhone) Brand() string {
	return a.brand
}

func (a *androidPhone) Model() string {
	return a.model
}

func (a *androidPhone) Type() string {
	return "smartphone"
}

func (a *androidPhone) OS() string {
	return a.os
}

func NewAndroidPhone(brand, model, os string) androidPhone {
	return androidPhone{brand: brand, model: model, os: os}
}

// Добавьте в пакет electronic структуру radioPhone. Сделайте так, чтобы она реализовывала интерфейсы Phone и StationPhone.
type radioPhone struct {
	brand        string
	model        string
	buttonsCount int
}

func (r *radioPhone) Brand() string {
	return r.brand
}

func (r *radioPhone) Model() string {
	return r.model
}

func (r *radioPhone) Type() string {
	return "station"
}

func (r *radioPhone) ButtonsCount() int {
	return r.buttonsCount
}

func NewRadioPhone(brand string, model string, buttons int) radioPhone {
	return radioPhone{brand: brand, model: model, buttonsCount: buttons}
}
