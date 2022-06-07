package storage

import (
	"homework/pkg/student"
)

type Storage struct { //Создали собсвенный тип хранилища для студентов
	Database map[string]*student.Student
}

//Для этого типа хранилища реализуем метод put

func (s *Storage) Put(value *student.Student) {
	s.Database[value.Name] = value
}

func (s *Storage) Get(name string) *student.Student {
	return s.Database[name]
}
