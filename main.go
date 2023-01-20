package main

import (
	"fmt"
	//"math"
)

func main() {

	//*Задание 1. Баллы ЕГЭ
	fmt.Println("Баллы ЕГЭ.")
	total := 0
	PassingScore := 275

	fmt.Println("Введитpwdе результат первого экзамена:")
	var ExamResult int
	fmt.Scan(&ExamResult)
	total += ExamResult

	fmt.Println("Введите результат второго экзамена:")
	fmt.Scan(&ExamResult)
	total += ExamResult

	fmt.Println("Введите результат третьего экзамена:")
	fmt.Scan(&ExamResult)
	total += ExamResult

	fmt.Println("Сумма проходных баллов:", PassingScore)
	fmt.Println("Количество набранных баллов:", total)

	if total >= PassingScore {
		fmt.Println("Поздвравляем, вы прошли вступительные экзамены")
	} else {
		fmt.Println("Вы не поступили.((((")
	}

}
