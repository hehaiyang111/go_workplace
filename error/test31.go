package main

import "fmt"

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
Explain：错误还可以用实现error的接口的结构体来表示.这种方式可以更加灵活的处理错误.在上面例子中,如果我们希望访问引发错误的半径，现在唯一的办法就是解析错误的描述信息Area calculation failed, radius -20.00 is less than zero。这样做不太好，因为一旦描述信息发生变化，程序就会出错。
method:创建一个实现error接口的结构体类型，并使用它的字段来提供关于错误的更多信息。
*/
/**
method1: 创建一个表示错误的结构体类型
*/
/**
type areaError struct {
	err string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{err: "radius is negative",radius: radius}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			// 如果是半径错误
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		//如果是别的错误
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f", area)
}
*/

/*
	method2
*/
type areaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}
func main() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
			return
		}
	}
	fmt.Println(area)
}
