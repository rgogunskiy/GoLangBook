package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FToM(f Feet) Meters {
	return Meters(f * 0.3084)
}

func MToF(m Meters) Feet {
	return Feet(m / 0.3048)
}

func PToK(p Pounds) Kilograms {
	return Kilograms(p * 0.4536)
}

func KToP(k Kilograms) Pounds {
	return Pounds(k / 0.4536)
}
