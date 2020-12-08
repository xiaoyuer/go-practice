package main

import (
	"fmt"
	"time"
)

func main() {
	//var p *int
	p := 233
	foo1(&p)
	foo2(p)
	foo1(&p)
	foo2(p)
	foo1(&p)()
	foo1(&p)()
	foo1(&p)()
	foo2(p)()
	foo2(p)()
	foo2(p)()
	fmt.Println("foo3")
	foo3()
	fmt.Println("foo4")
	foo4()
	fmt.Println("foo5")
	foo5()
	time.Sleep(3 * time.Second)

	go foo6()
	foo6Chan <- 1
	foo6Chan <- 2
	foo6Chan <- 3
	foo6Chan <- 5

	foo6Chan <- 11
	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 12
	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 13
	time.Sleep(time.Duration(1) * time.Nanosecond)
	foo6Chan <- 15
	// 微秒
	foo6Chan <- 21
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 22
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 23
	time.Sleep(time.Duration(1) * time.Microsecond)
	foo6Chan <- 25
	time.Sleep(time.Duration(10) * time.Second)
	// 毫秒
	foo6Chan <- 31
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 32
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 33
	time.Sleep(time.Duration(1) * time.Millisecond)
	foo6Chan <- 35
	time.Sleep(time.Duration(10) * time.Second)
	// 秒
	foo6Chan <- 41
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 42
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 43
	time.Sleep(time.Duration(1) * time.Second)
	foo6Chan <- 45
	time.Sleep(time.Duration(10) * time.Second)
	// 实验完毕，最后记得关闭channel
	close(foo6Chan)
	//毫秒和秒的两组非常确定，顺序输出。但是微妙就不一定了，有时候是顺序输出，大部分时候是随机输出如“22，22，23，25”或者“21，22，25，25”之类的
	f7s := foo7(11)
	for _, f7 := range f7s {
		f7()
	}

	foo0()()
	foo8()
}

//闭包的延迟绑定
//在执行的时候去外部环境寻找最新的数值，那x的最新数值就是11
func foo0() func() {
	x := 1
	f := func() {
		fmt.Printf("foo0 val = %d\n", x)
	}
	x = 11
	return f
}

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Printf("foo3 val = %d\n", val)
	}
}

func show(v interface{}) {
	fmt.Printf("foo4 val = %v\n", v)
}

func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go show(val)
	}
}

func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %v\n", val)
		}()
	}
}

// 其实这个问题的本质同闭包的延迟绑定，或者说，这段匿名函数的对象就是闭包。在我们调用go func() { xxx }()的时候，只要没有真正开始执行这段代码，那它还只是一段函数声明。而在这段匿名函数被执行的时候，才是内部变量寻找真正赋值的时候。

// 在case5中，for-loop的遍历几乎是“瞬时”完成的，4个Go Routine真正被执行在其后。矛盾是不是产生了？这个时候for-loop结束了呀，val生命周期早已结束了，程序应该报错才对呀？

//既然说Go Routine执行的时候比for-loop慢，那如果我在遍历的时候增加sleep机制呢
var foo6Chan = make(chan int, 10)

func foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", val)
		}()
	}
}

func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})
	}
	return fs
}

func foo8() {
	for i := 1; i < 10; i++ {
		curTime := time.Now().UnixNano()
		go func(t1 int64) {
			t2 := time.Now().UnixNano()
			fmt.Printf("foo8 ts = %d us \n", t2-t1)
		}(curTime)
	}
}

//5微秒~60微秒
