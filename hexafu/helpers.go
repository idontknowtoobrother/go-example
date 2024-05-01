package hexafu

import "math/rand"

func GetRandomInt(n int) int { // เมื่อชื่อ func ขึ้นต้นด้วยตัว "พิมใหญ่" จะเป็น PUBLIC FUNCTION
	return getRandomIntPrivate(n)
}

func getRandomIntPrivate(n int) int { // เมื่อชื่อ func ขึ้นต้นด้วยตัว "พิมใหญ่" จะเป็น PRIVATE FUNCTION
	return rand.Intn(n)
}
