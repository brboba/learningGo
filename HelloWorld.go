package main

import (
	"fmt"
	"math"
)

func main() {
	//1. Напишите программу для вычисления площади прямоугольника.
	//Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
	fmt.Println("task 1")
	var a, b int

	fmt.Print("Введите длинну стороны прямоугольника 1: ")
	fmt.Scanln(&a)

	fmt.Print("Введите длинну стороны прямоугольника 2: ")
	fmt.Scanln(&b)

	fmt.Println("Площадь прямоугольника будет равна : ", a * b)

	//2. Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
	//Площадь круга должна вводиться пользователем с клавиатуры.
	fmt.Println("task 2")

	var s, d float64
	fmt.Println("Введите площадь круга: ")
	fmt.Scanln(&s)

	d = math.Sqrt(float64(s/math.Pi)) * 2
	fmt.Println("Площадь прямоугольника будет равна : ", d)

	//3. С клавиатуры вводится трехзначное число.
	//Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
	fmt.Println("task 3")

	fmt.Println("Введите трехзначное число: ")
	fmt.Scanln(&a)

	fmt.Println("сотни : ", (a-a%100)%1000/100)
	fmt.Println("десятки : ", (a-a%10)%100/10)
	fmt.Println("еденицы : ", a%10)
}
