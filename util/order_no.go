package util

import "strconv"

func GenerateOrderNo(projectId int64, userId int64, count int64) string {
	return "01" + spell(projectId, 5) + spell(userId, 9) + spell(count, 4)
}

func spell(id int64, len int) (s string) {
	for i := len - 1; i > 0; i-- {
		if id/int64(exponent(10, uint64(i))) > 0 {
			break
		}
		s += "0"
	}
	s += strconv.FormatInt(id, 10)
	return
}

func exponent(a, n uint64) uint64 {
	result := uint64(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}
