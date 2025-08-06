package main

import "fmt"

func kelvinParaCelsius(k float64) float64 {
	return k - 273.15

}
func main() {
	K := 373.15
	celsius := kelvinParaCelsius(K)
	fmt.Println("O ponto e ebulição em Celsius é:", celsius)

}