package slices

// ForEach 对slice的每个元素执行action函数。
func ForEach[T any](slice []T, action func(T)) {
	for _, item := range slice {
		action(item)
	}
}

// Map 使用映射函数将 []T1 转换成 []T2。
// This function has two type parameters, T1 and T2.
// 映射函数 f 接受两个类型类型 T1 和 T2。
// 本函数可以处理所有类型的切片数据。
func Map[T1 any, T2 any](slice []T1, mapper func(T1) T2) []T2 {
	mapped := make([]T2, 0, len(slice))
	for _, item := range slice {
		mapped = append(mapped, mapper(item))
	}
	return mapped
}

// Reduce 使用汇总函数将 []T1 切片汇总成一个结果。
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter 使用过滤函数过滤切片中的数据 。
// 该函数返回新的切片，只会保留调用 f 返回 true 的元素。
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reverse 反转Slice中元素的顺序。
func Reverse[T any](s []T) []T {
	var r []T
	l := len(s)
	for i := range s {
		r = append(r, s[l-1-i])
	}
	return r
}

// 统计数组中各元素出现的频次，并将结果记入map
func GroupCount[T any](s []T) map[any]int {
	gmap := make(map[any]int)
	for i := 0; i < len(s); i++ {
		gmap[s[i]]++
	}
	return gmap
}
