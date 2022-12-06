package slices

// FindAll 在slice中搜索所有符合条件（condition函数返回true）的元素
func FindAll[T interface{}](slice []T, contdition func(T) bool) []T {
	result := []T{}
	for _, item := range slice {
		if contdition(item) {
			result = append(result, item)
		}
	}
	return result
}

// FindFirst 在slice中搜索第一个符合条件（condition函数返回true）的元素
func FindFirst[T interface{}](slice []T, contdition func(T) bool) T {
	for _, item := range slice {
		if contdition(item) {
			return item
		}
	}
	var t T
	return t
}

// FindLast 在slice中搜索最后一个符合条件（condition函数返回true）的元素
func FindLast[T interface{}](slice []T, contdition func(T) bool) T {
	for i := len(slice) - 1; i >= 0; i-- {
		if contdition(slice[i]) {
			return slice[i]
		}
	}
	var t T
	return t
}

// ForEach 对slice中的每个元素执行action函数，适用于读取slice场景。
func ForEach[T interface{}](slice []T, action func(e T)) {
	for i := 0; i < len(slice); i++ {
		action(slice[i])
	}
}

// ForEachPointer 对slice中的每个元素，通过其指针执行action函数，适用于改写slice场景。
func ForEachPt[T interface{}](slice []T, action func(e *T)) {
	for i := 0; i < len(slice); i++ {
		action(&slice[i])
	}
}

// ForEachParal 对slice中的每个元素多线程并行执行action函数，适用于读取slice场景。
// 参数n为并行协程数
func ForEachParal[T interface{}](slice []T, n int, action func(e T)) {
	if len(slice) <= 0 {
		return
	}
	if n <= 0 {
		n = 1 //若n设置错误，默认单线程执行
	} else {
		if n > len(slice) {
			n = len(slice) //若n小于slice的长度，则取n为slice的长度
		}
	}

	// 初始化token池
	token := make(chan bool, n)
	for i := 0; i < n; i++ {
		token <- true
	}
	// 执行action，每次执行时先取令牌，执行完后放回令牌
	for i := 0; i < len(slice); i++ {
		go func(index int) {
			<-token
			action(slice[index])
			token <- true
		}(i)
	}
	// 收回token，确保所有协程都已完成
	for i := 0; i < n; i++ {
		<-token
	}
}

// ForEachPtParal 对slice中的每个元素，通过其指针多线程并行执行action函数，适用于改写slice场景。
// 参数n为并行协程数
func ForEachPtParal[T interface{}](slice []T, n int, action func(e *T)) {
	if len(slice) <= 0 {
		return
	}
	if n <= 0 {
		n = 1 //若n设置错误，默认单线程执行
	} else {
		if n > len(slice) {
			n = len(slice) //若n小于slice的长度，则取n为slice的长度
		}
	}

	// 初始化token池
	token := make(chan bool, n)
	for i := 0; i < n; i++ {
		token <- true
	}
	// 执行action，每次执行时先取令牌，执行完后放回令牌
	for i := 0; i < len(slice); i++ {
		go func(index int) {
			<-token
			action(&slice[index])
			token <- true
		}(i)
	}
	// 收回token，确保所有协程都已完成
	for i := 0; i < n; i++ {
		<-token
	}
}

// Map 使用映射函数将 []T1 转换成 []T2。
// 映射函数 f 接受两个类型类型 T1 和 T2。
func Map[T1 interface{}, T2 interface{}](slice []T1, mapper func(T1) T2) []T2 {
	mapped := make([]T2, len(slice))
	for i := 0; i < len(slice); i++ {
		mapped[i] = mapper(slice[i])
	}
	return mapped
}

// MapParal 多线程并行使用映射函数将 []T1 转换成 []T2。
// 映射函数 f 接受两个类型类型 T1 和 T2。
func MapParal[T1 interface{}, T2 interface{}](slice []T1, n int, mapper func(T1) T2) []T2 {
	if n <= 0 {
		n = 1 //若n设置错误，默认单线程执行
	} else {
		if n > len(slice) {
			n = len(slice) //若n小于slice的长度，则取n为slice的长度
		}
	}
	mapped := make([]T2, len(slice))

	// 初始化token池
	token := make(chan bool, n)
	for i := 0; i < n; i++ {
		token <- true
	}
	// 执行action，每次执行时先取令牌，执行完后放回令牌
	for i := 0; i < len(slice); i++ {
		go func(index int) {
			<-token
			mapped[index] = mapper(slice[index])
			token <- true
		}(i)
	}
	// 收回token，确保所有协程都已完成
	for i := 0; i < n; i++ {
		<-token
	}

	return mapped
}

// Reduce 使用汇总函数将 []T1 切片汇总成一个结果。
func Reduce[T1, T2 interface{}](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter 使用过滤函数过滤切片中的数据。
// 该函数返回新的切片，只会保留调用 f 返回 true 的元素。
func Filter[T interface{}](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

// Reverse 反转Slice中元素的顺序。
func Reverse[T interface{}](s []T) []T {
	var r []T
	l := len(s)
	for i := range s {
		r = append(r, s[l-1-i])
	}
	return r
}

// 统计数组中各元素出现的频次，并将结果记入map
func GroupCount[T interface{}](s []T) map[interface{}]int {
	gmap := make(map[interface{}]int)
	for i := 0; i < len(s); i++ {
		gmap[s[i]]++
	}
	return gmap
}
