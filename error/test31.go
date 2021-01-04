package main

/**
自定义Error
*/
/**
errors包中源码内部实现
*/
// Package errors implements functions to manipulate errors.
/**
package errors

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
**/

/**
Example1:创建一个计算圆半径的简单程序，如果半径为负，他会返回一个错误
Explain:在下面的程序中,我们检查半径是否小于零.如果半径小于零，我们会返回等于 0 的面积,以及相应的错误信息.如果半径大于零,则会计算出面积,并返回值为 nil 的错误.
在 main 函数里，我们检查错误是否等于 nil.如果不是 nil,我们会打印出错误并返回,否则我们会打印出圆的面积.
*/
//func circle(radius float64) (float64, error) {
//	if radius < 0 {
//		return 0, errors.New("Area calculation failed, radius is less than zero")
//	}
//	return math.Pi * radius * radius, nil
//}
//func main() {
//	radius := -20.0
//	area, err := circle(radius)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Printf("Area of circle %0.2f", area)
//}

/**
Example2: 使用 Errorf 给错误添加更多信息
Explain: 上面的程序效果不错,但是如果我们能够打印出当前圆的半径,那就更好了.这就要用到 fmt 包中的 Errorf 函数了.Errorf 函数会根据格式说明符,规定错误的格式,并返回一个符合该错误的字符串.
*/
/**
func circle(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}
func main() {
	radius := -20.0
	area, err := circle(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
**/

/**
Example3:使用结构体类型和字段提供错误的更多信息
Explain：
*/
