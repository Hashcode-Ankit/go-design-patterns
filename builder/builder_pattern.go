package builder

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportsWheels Wheels = "sports"
	SteelWheels         = "steel"
)

type Car struct {
	color  Color
	wheels Wheels
	speed  Speed
}

type CarBuilder struct {
	car Car
}
type Builder interface {
	Drive() error
	Stop() error
}

func NewBuilder(speed Speed, wheels Wheels) CarBuilder {
	return CarBuilder{
		car: Car{
			speed:  speed,
			wheels: wheels,
		},
	}
}
