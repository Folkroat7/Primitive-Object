package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Name string
	age  int
	ip   string
}

func (user *User) Adder(name string, age int, ip string) {
	user.Name = name
	user.age = age
	user.ip = ip
	fmt.Println("Создан юзер. Имя:", user.Name, "\n айпи:", user.ip)
}
func main() {
	fmt.Println("Введите значения:")
	name, ip := "", ""
	age := 0
	reader := bufio.NewReader(os.Stdin)
	chars, _ := reader.ReadString('\n')
	elements := strings.Fields(chars)
	for i, ch := range elements {
		switch i {
		case 0:
			name = ch
		case 1:
			age, _ = strconv.Atoi(ch)
		case 2:
			ip = ch
		}
	}
	user := User{}
	user.Adder(name, age, ip)
}
