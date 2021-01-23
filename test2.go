package main

import (
	"fmt"
	"math"
	"math/rand"
)
var c,python, java  bool


type Vertex struct {
	id           int
	userName     string
	passWord     string
}

/**
 * 变量定义和方法定义
 */
func add(x int ,y int) int{
	return  x+y;
}
/**
 * 主函数
 */
func main() {
	var i int
	//k: 用于定于确定类型数据
	 k := 3
	 fmt.Println("My favorite number is", rand.Intn(10))
	 fmt.Println(add(11,12))
	 fmt.Println(i,c,java,python,k)
	 initZero()
     conversion()
	 constant()
	 caculator()
	 deferFunction()
	testPointFunction()
}
/**
 *初始化数据  定义 func initZero() int
 */
func  initZero() int{
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	return i;
}

/***
  数据类型转化
 */
func conversion() {
	v := 42 // 修改这里！
	fmt.Printf("v is of type %T\n", v)


	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}

/**
  constant 转化   v:=6  常量
 */
const Pi = 3.14
func constant() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)
}

/**
 *caculator
 */
func caculator() {
	sum := 0;
	for i :=0; i<10; i++{
		sum += i
	}
	fmt.Println(sum)
	/**
	 * go 语言的前置和后置相关
	 */
	//a :=3;
	//	//for;a<1000; {
	//	//	a += a
	//	//}
	//	//fmt.Println(a)


	    a :=3
		for a<1000 {
			a += a
		}
		fmt.Println(a)
	    /***
	      if 语句
	     */
         if a>4{
         	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa")
		 }
	sqrt(2);
	sqrt(-4);
	deferMethod()
  }
  /**
    sqrt
   */
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/***
 *defer
 */
func deferMethod() {
	defer fmt.Println("world")
	fmt.Println("hello")
}


/***
   defer 相关配置
 */
func  deferFunction(){
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

/***
 * 数组指针
 */
func  testPointFunction(){
	i, j := 42, 2701
	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值

	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值
	p = &j         // 指向 j

	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

