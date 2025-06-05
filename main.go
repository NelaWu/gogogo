package main

import (
	"fmt"
	"os"
)

type student struct {
	id    int
	name  string
	class string
}

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type allStudent struct {
	students []*student
}

func (s *allStudent) addStudent(new *student) {
	s.students = append(s.students, new)
}

func (s *allStudent) modStudents(mod *student) {
	for i, student := range s.students {
		if student.id == mod.id {
			s.students[i] = mod
			return
		}
	}
}

func (s *allStudent) showStudents() {
	for _, student := range s.students {
		fmt.Println("id:", student.id, "name:", student.name, "class:", student.class)
	}
}

func showMenu() {
	fmt.Println("歡迎來到Life校務系統")
	fmt.Println("請問你要幹嘛？")
	fmt.Println("1. 新增life學生")
	fmt.Println("2. 修改life學生")
	fmt.Println("3. 看學生列表")
	fmt.Println("4. Quit")
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("學生學號：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("學生姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("學生班級：")
	fmt.Scanf("%s\n", &class)
	stu := newStudent(id, name, class) //调用student的构造函数构造一个学生
	return stu
}

func main() {
	allStu := new(allStudent)
	allStu.students = make([]*student, 0)
	for {
		showMenu()
		var input int
		fmt.Print("請輸入你要做什麼：")
		fmt.Scanf("%d\n", &input)
		switch input {
		case 1:
			fmt.Println("<新增life學生>")
			stu := getInput()
			allStu.addStudent(stu)
		case 2:
			fmt.Println("<get 2>")
			stu := getInput()
			allStu.modStudents(stu)
		case 3:
			fmt.Println("<看學生列表 3>")
			allStu.showStudents()
		case 4:
			os.Exit(0)
		}
	}
}
