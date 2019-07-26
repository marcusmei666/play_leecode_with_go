package sortHelper

type Sortable interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

// 定义IntArr类型
type IntArr []int
// 给IntArr提供Len方法
func (arr IntArr) Len() int {
	return len(arr)
}
// 给IntArr提供Less方法
func (arr IntArr) Less(i int, j int) bool {
	return arr[i] < arr[j]
}
// 给IntArr提供Swap方法
func (arr IntArr) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


type StrArr []string
func (arr StrArr) Len() int {
	return len(arr)
}
func (arr StrArr) Less(i int, j int) bool {
	return arr[i] < arr[j]
}
func (arr StrArr) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type FolArr []string
func (arr FolArr) Len() int {
	return len(arr)
}
func (arr FolArr) Less(i int, j int) bool {
	return arr[i] < arr[j]
}
func (arr FolArr) Swap(i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
