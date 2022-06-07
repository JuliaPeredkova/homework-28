package main

import (
	"bufio"
	"fmt"
	"homework/pkg/storage"
	"homework/pkg/student"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {

	studentMap := storage.Storage{Database: make(map[string]*student.Student)}

	fmt.Println("Введите имя студента, его возраст и группу")

	for { //Получаем именя студентов, возраст и номер группы в бесконечном цикле
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			studentDataScan := sc.Text()
			studentDataScanSlice := strings.Fields(studentDataScan)
			//fmt.Println(studentDataScanSlice)

			if len(studentDataScanSlice) < 3 {
				fmt.Println("Ошибка при входных данных: недостаточное количество аргументов")
				return
			}

			studentName := studentDataScanSlice[0]
			studentAge, errStudentAge := strconv.Atoi(studentDataScanSlice[1])
			studentGrade, errstudentGrade := strconv.Atoi(studentDataScanSlice[2])

			if errStudentAge != nil || errstudentGrade != nil {
				fmt.Println("Ошибка при обработке возраста и группы студента")
			}

			student := student.Student{
				Name:  studentName,
				Age:   studentAge,
				Grade: studentGrade,
			}

			//С помощью метода put записываем в хранилище структуру student
			//put(studentMap, &student)
			studentMap.Put(&student)

			// if EOF - выходим из цикла и реализуем метод get

			if err := sc.Err(); err != nil {
				if err == io.EOF {
					studentMap.Get(student.Name)
					//get(studentMap, student.name)
				}
				return
			}
		}

		for _, v := range studentMap.Database {
			fmt.Println(v.NewStudent())
		}
	}
}
