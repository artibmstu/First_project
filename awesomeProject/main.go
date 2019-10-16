package main

import "fmt"

func main () {
//////////////////////////////////////////
	var age,c int = 12,15
	var k int
	k = c - age
	fmt.Println(k)
/////////////////////////////////////////
	var lol float64 = 2.34764
	fmt.Println(k, lol)
////////////////////////////////////////
	var hel string = "Hello"
	fmt.Println(hel)
///////////////////////////////////////
	k = c + age	// плюс
	k = c - age	// минус
	k = c * age	// умножение
	k = c / age // деление
	k = c % age	// остаток при делении
	fmt.Println("Result is", k)
//////////////////////////////////////
	// const pi = 3.14
//////////////////////////////////////
	var web string = "Prog"
	fmt.Println(len(web)) // вывод количетсва букв в web
	fmt.Println(web + " \nis good") // объединение строк
//////////////////////////////////////
	var num float64 = 4.87346587
	fmt.Printf("%f boy\n", num)	// вывод в тексте
	fmt.Printf("%T \n", num)	// вывод формата числа
/////////////////////////////////////
	if age < k {
		fmt.Println("k > age")
	}	else if age == k {
		fmt.Println("age = k")
	}	else if (age > k) && (age <= 10000) {
		fmt.Println("age > k")
	}	else	{
		fmt.Println("age > 10000")
	}
////////////////////////////////////
	var proverka = 14
	switch proverka {
	case 0: fmt.Println("prov = 0")
	case 5: fmt.Println("prov = 5")
	case 10: fmt.Println("prov = 10")
	case 15: fmt.Println("prov = 15")
	default: fmt.Println("prov - ?")
	}
////////////////////////////////////
	var i1 = 0
	for i1 <= 10 {
		fmt.Println(i1)
		i1 ++
	}

	for i2 := 0; i2 <= 5; i2 ++ {
		fmt.Println(i2)
	}
////////////////////////////////////
	var arr[3] int
	arr[0] = 45
	arr[1] = 43
	arr[2] = 68
	fmt.Println(arr[1])
///////////////////////////////////
	nums := [3]float64 {4.23, 14.39, 3.78}
	for i3, value := range nums {
		fmt.Println(value, i3)
	}
//////////////////////////////////
	webSites := make(map[string]float64)

	webSites["itProger"] = 0.8
	webSites["yandex"] = 0.99
	fmt.Println(webSites["itProger"])

////////////////////////////////////////////////////////////////////////////////////////////////
	var kek_1, kek_2 = 48, 2
	var rum_1, rum_2 int
	rum_1, rum_2 = summ (kek_1, kek_2)
	fmt.Println(rum_1, rum_2)
///////////////////////////////////////////////////////////////////////////////////////////////
/////***
var x0 = 0
pointer (&x0)
fmt.Println (x0)
/////***
}
////////////////////////////////

//////////////////////////////////////////////////////////////////////////////////////////////
func summ (num_1 int, num_2 int) (int,int) {
	var res int
	res = num_1 + num_2
	return res, num_1 * num_2
}
///////////////////////////////

/////***
func pointer (x0 *int) {
	*x0 = 2
}
/////***
