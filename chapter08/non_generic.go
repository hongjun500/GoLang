package main

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats32(m map[string]float32) float32 {
	var s float32
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats64(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
