package array

import "sort"

// DiffInt int整形切片比较差异
// 参数为两个int切片
// 返回 第一个切片中存在,第二个切片不存在的一个新切片
//
func DiffInt(first, second []int) (result []int) {
	sort.Ints(second)
	secondMap := make(map[int]int)
	for _, v := range second {
		secondMap[v] = v
	}
	for _, v := range first {
		if _ ,ok := secondMap[v]; !ok {
			result = append(result,v)
		}
	}
	return
}
// DiffInt64 int64整形切片比较差异
// 参数为两个int64切片
// 返回 第一个切片中存在,第二个切片不存在的一个新切片
//
func DiffInt64(first, second []int64) (result []int64) {
	secondMap := make(map[int64]int64)
	for _, v := range second {
		secondMap[v] = v
	}
	for _, v := range first {
		if _ ,ok := secondMap[v]; !ok {
			result = append(result,v)
		}
	}
	return
}
// DiffString 字符切片比较差异
// 参数为两个string切片
// 返回 第一个切片中存在,第二个切片不存在的一个新切片
func DiffString(first, second []string) (result []string) {
	secondMap := make(map[string]string)
	for _, v := range second {
		secondMap[v] = v
	}
	for _, v := range first {
		if _ ,ok := secondMap[v]; !ok {
			result = append(result,v)
		}
	}
	return
}