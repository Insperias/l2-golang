Что выведет программа? Объяснить вывод программы.

package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
Ответ:

В результате выполнения будет выведено error.
В го утиная типизация (если объект говорит как утка, летает как утка и плавает как утка, то это утка).
Таким образом, тип customError удовлетворяет интерфейсу error, поэтому может использоваться. Однако при присвоении
переменная err будет не с типом nil, а с типом *customError, и, соответственно, перестанет быть равна nil.