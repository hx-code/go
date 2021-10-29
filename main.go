package main

import (
	// "debug/dwarf"
// "debug/dwa	rf"
	"fmt"
	// "image/draw"
	// "math/rand"
	// "time"
)

// import "fmt"

// func main() {
// 	const ligthSpeed = 299792
// 	var distance = 5600000
// 	fmt.Println(distance/ligthSpeed, "seconds")
// 	distance = 401000000
// 	fmt.Println(distance/ligthSpeed, "seconds")
// }
// //可以声明多个变量 var distance,speed =56000000,1000000
// func main() {
// 	var weigth = 149.0
// 	weigth *= 0.3783
// }
// //在go里边只有age++

// "Atoi"

// "strconv"
// "strconv"
//  "unicode/utf8"
// "math/big"

// "math"

// func main() {go
// 	var num = rand.Intn(10) + 1
// 	fmt.Println(num)
// 	num = rand.Intn(10) + 1
// 	fmt.Println(num)

// }

// func main() {
// 	fmt.Println("There is a sign near the entrance taht reads 'no minors")
// 	var age = 41
// 	var minor = age < 18
// 	fmt.Printf("At age %v, am I a minor %v/n", age, minor)
// }
// func main() {
// 	var command = "go east"
// 	if command == "go esst" {
// 		fmt.Println("you head further up the mountain")
// 	} else if command == "go inside" {
// 		fmt.Println("you enter the cave where you live out the rest of your life ")
// 	} else if command == "go inside1" {
// 		fmt.Println("you enter the cave where you live out the rest of your life ")
// 	} else if command == "go inside2" {
// 		fmt.Println("you enter the cave where you live out the rest of your life")
// 	} else {
// 		fmt.Println("Didn't quite get that")
// 	}
// }
// func main() {
// 	fmt.Println("This years is 2100,shoud you leave ?")
// 	var year = 2100
// 	var leap = year%400 == 0 || (year%4 == 0 && year%100 == 0)
// 	if leap {
// 		fmt.Println("look before you leap")
// 	} else {
// 		fmt.Println("keep your feet on your groud")
// 	}

// }
// func mian() {
// 	var haveTorch = true
// 	var litTorch = false
// 	if !haveTorch || !litTorch {
// 		fmt.Printf("Nothing to see here")
// 	}

// }
//fallthrough
// func main() {
// 	var room = "lack"
// 	switch {
// 	case room == "cave":
// 		fmt.Println("你好牛")
// 	case room == "lack":
// 		fmt.Println("你好棒")
// 		fallthrough
// 	case room == "underwater":
// 		fmt.Println("可以的")
// 	}
// }
// func main() {
// 	var count = 10
// 	for count <0 {
// 	fmt.Println(count)
// 		time.Sleep(time.Second)//停留一秒
// 	count--
//  	}
//  	fmt.Println("Litoff")
//  }
//变量作用域有短声明count：=10
// func main(){
// 	var count =0
// 	for count <10{
// 		var num = rand.Intn(10)+1
// 		fmt.Println(num)
// 		count++
// 	}

// }
//era 变量是在main函数外声明的，他拥有package有多个函数作用域 短声明不可以声明package 作用域的变量
// var era ="AD"
// func main(){
// 	year :=2018
// 	switch  month := rand.Intn(12)+ 1; month {
// 	case 2:
// 		day :=rand.Intn(28) + 1
// 		fmt.Println(era,year,month,day)
// 	case 4,6,9,11:
// 		day := rand.Intn(30)+1
// 		fmt.Println(era,year,month,day)
// 	default:
// 		day :=rand.Intn(31) +1
// 		fmt.Println(era ,year ,month,day)

// 	}
// }
//整数来初始化某个变量两种浮点类型float64双精度 float 32 当处理3d游戏数千个顶点 math包里都使用64
// func main(){
// 	// var pi64 = math.Pi
// 	// var pi32 float32 = math.Pi
// 	// fmt.Println(pi64)
// 	// fmt.Println(pi32)
// 	// var price float64
// 	// fmt.Println("price")
// }
// func main(){
// 	piggyBank := 0.1
// 	piggyBank +=0.2
// 	fmt.Println(piggyBank == 0.3)
// 	fmt.Println(math.Abs(piggyBank-0.3)<0.0001)

// }
//go语言里边有十种 var year int= 2000    无符号数 var month uint =2 uint8可以表示八种颜色 0x表示十六进制
// func main(){
// 	year:=2018
// 	fmt.Println("Type %T for %v\n",year,year)
// 	a:="text"
// 	fmt.Println("Type %T for %[1]v\n",a)
// }
// func main(){
// 	var green uint8 = 3
// 	fmt.Println("%08b\n",green)
//     fmt.Println("%08b\n",green)
// }
//更大数可以用big包
// func main() {
// 	ligthSpeed := big.NewInt(299792)
// 	secondPerDay := big.NewInt(86400)
// 	fmt.Println(ligthSpeed, secondPerDay)
// }

//多语言文本
//声明字符串： peace :="peace" var peace ="peace"  var peace string = "peace" 字符串零值 var bank string
//字符串字面值/原始字符面值go
// func main(){
// 	fmt.Println("peace be opon\n upun")
// 	fmt.Println(`string upon \n123`)
// }
//go语言提供了rune这个类型 它是int32的一个别名 而byte是uinut类型的别名 用于二进制
// func main(){
// var pi rune = 960
// var alpha rune = 940
// var omega rune = 969
// var bang byte = 33
// fmt.Printf("%v %v %v %v\n",pi,alpha,omega,bang)
// fmt.Printf("%c %c %c %c\n",pi,alpha,omega,bang)
// }
//凯撒加密法：对于加密有效的办法就是把每个字母移动固定位置
// func main(){
// 	c:='a'
// 	c = c+3
// 	fmt.Printf("%c",c )
// 	if c>'2'{
// 		c=c-26
// 	}
// }
//rot13 他会把字母替换+13对应的字符
//go的内置函数len
// func main(){
// 	message := "ua vagreangvbany fcnpr fgngvba"
// 	for i:=0;i<len(message);i++{
// 		c:=message[i]
// 		if c>='a'&&c<='z' {
// 			c=c+13
// 			if c>'z'{
// 				c=c-26
// 			}
// 			fmt.Printf("%v" , c)

// 		}
// 	}

// }
// func main(){
// 	question :="hello"
// 	fmt.Println(len(question),"byte")
// }
// func main(){
// 	question :="hello"
// 	fmt.Println(utf8.RuneCountInString(question),"runes")
// 	c, size := utf8.DecodeRuneInString(question)
// 	fmt.Printf("fist rune :%v %c byte",c,size)
// }
// func main(){
// 	earthDays:=365.2425
// 	fmt.Printf(int(earthDays))
// // }
// func main(){
// 	if v >=0 && v<=Math.MaxUint8{
// 		v8 :=uint8(v)
// 		fmt.Println("converted",v8)
// 	}
// }
// func main(){
// 	var pi rune =960
// 	var alpha rune =940
// 	var omega rune = 969
// 	var bang byte = 33
// 	fmt.Printf(string(pi),string(alpha),string(omega),string(bang))
// }
//想把数值转化成有意义的数值strconv包的ltoa就可以做这件事
// func main(){
// countdown :=10;
// str := "Launch in T minus"+strconv.Itoa(countdown)+"seconds"
// fmt.Print(str)
// }
// func main(){
// 	countdown:=9
// 	str:=fmt.Sprintf("Launch in T minus %v seconds.",countdown)
// 	fmt.Println(str)
// }
//strconv包里有个Atoi函数由于字符串里面包含任意字符或者是转化的数太大，所以Atoi函数会发生保错

// func main(){
// 	countdown,err :=strconv.Atoi("10")
// 	if err nil{

// 	}
// 	fmt.Println("countdown")
// }
// func main() {
// Launch:=false
// LaunchText := fmt.Sprintf("v%",Launch)
// fmt.Println("ready for launch:",LaunchText)
// var yesNo string
// if Launch{
// 	yesNo ="yes"
// }else{
// 	yesNo="no"
// }
// fmt.Println("Ready for Launch:",yesNo)
// }
//函数：每一件事通有很多步骤，每个步骤可以分为独立的函数，函数可以复用 func intn(n int) int
//函数可以有多个func Unix(sec int64 , nses int64)Time 调用 future:=time.Unix(12622780800,0)简化func Unix(sec,nsec int64)Time countdown,err :=strconv.Atoi("10") func Atoi(s string)(int error)
//printfln是一个特殊的参数，它可以接受一个 ，两个甚至多个参数
//...表示函数的参数数量是可变的   参数a的类型为interface{}，是一个空接口
// func kelvinToCelsius(k float64)float64{
// 	k-=273.15
// 	return k
// }
// func main(){
// 	kelvin := 294.0
// 	celsius := kelvinToCelsius(kelvin)
// 	fmt.Print(kelvin," k is ",celsius,"c")
// }
//方法 关键字type来声明 type celsius float64  //var temperature celsius = 207
// func main(){
// 	type celsius float64
// 	const degress = 20
// 	const temperature celsius = degress
// 	 +=10
// 	fmt.Println(temperature)
// }
// type kelvin float64
// func fakeSensor()kelvin{
// 	return kelvin(rand.Intn(150)+150)
// }
// func realSensor()kelvin{
// 	return 0

// }
// func main(){
// 	sensor := fakeSensor
// 	fmt.Println(sensor())
// 	sensor = realSensor
// 	fmt.Println(sensor())
// }
//将函数传递给其他函数
// type kelvin float64
// func measureTemprature(samples int,sensor func() kelvin){
// 	for i:=0;i<samples;i++{
// 		k := sensor
// 		fmt.Printf("%v k\n",k)
// 		time.Sleep(time.Second)
// 	}
// }
// func fakeSensor()kelvin{
// 	return kelvin(rand.Intn(151)+150)
// }
// func main(){
// 	measureTemprature(3,fakeSensor)
// }
//声明函数类型
// type  sensor func() kelvin 所以：func measureTemperture(simples int,s func() kelvin)
//可以精简为：func measureTemperature(simple int, ssensor)
// func drawTable (row int,getRow string)
//匿名函数
// var f =func ()  {
// 	fmt.Println("dress up the aaa")

// }
// func main(){
// 	f()
// }
// func main(){
// 	f:=func(message string){
// 	fmt.Println(message)
// }
// f("go to the pa")
// }
// func main(){
// 	func(){
// 		fmt.Println("hello")
// 	}()
// }
//闭包就是由于匿名函数封闭并包围作用域中的变量而得名
//  type kelvin float32
//  type sensor func() kelvin
//  func realSensor() kelvin{
// 	 return 0
//  }
//  func calibrate(s sensor,offset kelvin)sensor{
// 	 return func() kelvin {
// 		 return s()+offset

// 	 }
//  }
//  func main(){
// 	 sensor := calibrate(realSensor,5)
// 	 fmt.Println(sensor())

//  }
// func main(){
// 	var planets [8]string
// 	planets[0] ="mercury"
// 	planets[1] ="Venus"
// 	planets[2] ="Earth"
// 	earth :=planets[2]
// 	fmt.Println(earth)
// 	fmt.Println(len(planets))
// 	fmt.Println(planets[3]=="")
// }
//复合字面值是一种符合类型初始化紧凑语法 go的符合字面语法允许我们只用一步就完成数组声明和初始化
//便利数组
// func main(){
// 	dwarf := [5]string{"Ceres","Pluto","Haumes","makemake","erid"}
// 	for i := 0; i < len(dwarf); i++ {
// 		dwarf:=dwarf[i]
// 		fmt.Println(i,dwarf)
// 	}
// }
//数组的拷贝
// func main(){
// 	planets :=[...]string{
// 		"Mercuy",
// 		"Venus",
// 		"Earth",
// 		"Mars",

// 	}
// 	planetsMarkII := planets
// 	planets[3] ="whhoa"
// 	fmt.Println(planets)
// 	fmt.Println("planetsMarkII")
// }
// func terraform(planets[8]string){
// 	for i:=range planets{
// 		planets[i]="new"+planets[i]
// 	}
// }
// func main(){
// 	planets:=[...]string{
// 		"mercurry",
// 		"vens",
// 		"earth",
// 		"mars",
// 		"jupy",
// 		"spa",
// 		"mm",
// 		"aa",
// 	}
// 	terraform(planets)
// 	fmt.Println(planets)
//terrestrial :=planets[0:4]
// }
//数组的长度有也是数组的一部分函数一般使用slice而不是用数组作为参数
//切片slice指向数组的窗口
//忽略掉slice的起始索引，表示数组从开始切分 忽略掉slice的结束索引，使用数组长度作为索引
//slice索引不能为复数
// func main(){
// 	dwarfArray :=[...]string{"co","promise","proxy","vue"}
// 	drawSlice:=dwarArray[:]
//  	draws := []string{"hheh","heheh","www","jjjj"}
//  	fmt.Println(dwarfSlice,dwarfs)
//  } 
