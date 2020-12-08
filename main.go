package main

// func main() {
// 	str := []string{"I", "am", "Ethan"}
// 	for _, v := range str {
// 		go func() {
// 			fmt.Println(v)
// 		}()
// 	}
// 	time.Sleep(3 * time.Second)
// }

// 闭包里引用了不作为参数传递进去的值,都是引用传递...也就是说,println(v) 其实是引用了v的地址然后解引用,将值打印出来..等到这个goroutine执行println(v)的时候,v所指向的值已经是"Ethan"
// /这里说的不对，并不是地址。 根本原因是 for v := range , 这里的v可以认为是一个局部变量，for循环完成后，局部变量被设置为Ethan后， goroutine才被调度

// func main() {
// 	str := []string{"I", "am", "Ethan"}
// 	for _, v := range str {
// 		go func(v string) {
// 			fmt.Println(v)
// 		}(v)
// 	}
// 	time.Sleep(3 * time.Second)
// }

// func main() {
// 	slice := []int{0, 1, 2, 3}
// 	myMap := make(map[int]*int)

// 	for index, value := range slice {
// 		myMap[index] = &value
// 	}
// 	prtMap(myMap)
// }

// func prtMap(myMap map[int]*int) {
// 	for key, value := range myMap {
// 		fmt.Printf("map[%v]=%v\n", key, *value)
// 	}
// }

// 原因解释:但是由输出可以知道，映射的值都相同且都是3。其实可以猜测映射的值都是同一个地址，遍历到切片的最后一个元素3时，将3写入了该地址，所以导致映射所有值都相同。

//  其实真实原因也是如此，因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，就会导致错误，在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，所以值的地址总是相同的，导致结果不如预期。

// func main() {
// 	slice := []int{0, 1, 2, 3}
// 	myMap := make(map[int]*int)

// 	for index, value := range slice {
// 		v := value
// 		myMap[index] = &v
// 	}
// 	prtMap(myMap)
// }

// func prtMap(myMap map[int]*int) {
// 	for key, value := range myMap {
// 		fmt.Printf("map[%v]=%v\n", key, *value)
// 	}
// }
