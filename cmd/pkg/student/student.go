package student

import "strconv"

type Student struct { //Тип данных student
	Name  string
	Age   int
	Grade int
}

func (s Student) NewStudent() string { //Создали структуру Student при помощи функции newStudent
	return s.Name + " " + strconv.Itoa(s.Age) + " " + strconv.Itoa(s.Grade)
}
