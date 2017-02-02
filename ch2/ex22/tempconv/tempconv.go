package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meters float64
type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC Celsius = 0
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gF", f)
}

func (m Meters) String() string {
	return fmt.Sprintf("%gM", m)
}

func (p Pounds) String() string {
	return fmt.Sprintf("%gP", p)
}

func (k Kilograms) String() string {
	return fmt.Sprintf("%gK", k)
}
