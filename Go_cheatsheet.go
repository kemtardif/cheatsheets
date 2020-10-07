package main

import (
	"fmt" //Import user output
	"bufio" //Allow user input
	"os"
	"strconv"
	"math"

//go run file.cs compile and run the program
//go build file.cs compile and create an .exe file

func main() {

	//USER INPUT AND OUTPUT////////
	fmt.Println("Hello World") //Print 
	var a string = fmt.Sprintf("%T %v", 5.45, 3) //Return data type of 5.45 and the value 3. (%q get string value)
	scanner := bufio.NewScanner(os.Stdin) //Creater input object
	scanner.Scan() //Let user input something and save it in scanner object
	input := scanner.Text() //New variable storing the input
	input, _ = strconv.ParseInt(input) //Convert string to base 10 integer, stting to 0 if can't convert

//VARIABLES AND DATA TYPES//////

var x string = "Homies"
var y int //different types of integers : uin8, uint16, int8, int 16, int32
y = 5
var z = 3.45 //implicit type, GO will guess the type
x1 := "Yep yep" //Shortcut to declare implicit variable
var bl bool = true
var fl float32 = 3.1416

//MUTABLE AND IMMUTABLE DATA TYPES//////////////////

var x int := 5
y := x
y = 5   //Doesn't change value of x, since integers (and strings, booleans, arrays) are immutable

var x []int = []int{1, 2, 3}
y := x
 y[0] = 2 //Change x[0] to 2, since slices (and maps) are mutable

//MATHEMATICS//////////////

//Arithmetics is  made on same data type variables
//Give answer of same data type so 3/2 gives 1

var a float32 = 3.1416
var b int = 5

value := a/float32(b)

math.Abs(n)
math.Max(a,b)
math.min(a,b)

//IF STATEMENTS/////////////////

if condition1 && condition2 {
	Do something
} else if condition3 || condition3 {
	Do something
} else {
		Do something
	}

//SWITCH STATEMENTS////////////

ans := 10
switch ans {
case 1 :
	Do something
case 2:
	Do something
default :
Do something if other cases are not satisfied

}

// FOR LOOPS/////////////

for x := 0; x < 5; x++{
	Do something
	break //end loop
	continue // Go back to top of loop without executing anything below
}

//ARRAYS/////////////

var arr [n] int //array containing 0's. Default for strings is empty space.
arr1 := [3]string {"one", "two", "three"}
arr1[1] = "two"
arr2D := [2][3]int {{1, 2, 3}, {4, 5, 6} //2-dimentsional array}
arr2D[0][1] = 2
len(array) //# of elements

//SLICES//////////////

var x [5]int = [5]{1, 2, 3, 4, 5}
var s[]int = x[:] //Make copy of x
var t[]int = x[a:b] //Make an sub-array of elements with indices a, a+1, ...b-1
var q[int] = x[1:3] //Create array [2,3]
q[:cap(q)] //Extend q to [2, 3, 4, 5]

//Slices are used to make arrays with undefined # of elements:
var a[]int = []int {2,3,4,5,6}
b := append (a, 10) //Return new slice containing elements of a and 10 at the end
c := make([]int, 5) //Create new slice of 5 0's

//RANGE/////////////////

var a []int = []int{1, 2, 3, 4, 5}

for i, item := range a{  //Loop over indices of a[i] 
	Do something
}

for _, item := range a{ //Loop over the item values 
	Do something 
}

//MAPS (HASHES)//////////////////////

var mp map[string]int = map[string]int {
	"one":1,
	"two":2,
	"three":3,
}
mp["one"] = 1 //Change value
mp["four"] = 900 //Add new key/value pair
delete(mp, "four") //Delete key/pair if it exists
len(mp) //Return # of pairs in map

mp1 := make(map[string]int) //Create empty map

val, ok := mp["five"] //Check if "five key/pair exists and if not, create one woth default value"

//FUNCTIONS///////////////////////////

func newFunction(x string, y int) (int, int) { //Declared outside main
	Do something
	return total, reminder
}
ans1, ans2 := newFunction(5, 7) //Declare inside main, return integer total and reminder

func secondFunction (x int, y int) (z1 int, z2 int){
	defer fmt.println("Hello") ///Execute this line at end of function
	z1 = something
	z2 = something else   ///Return z1 and z2
	return

}

thirdFunction := func(x int){ //Different ways to show functions are variables
	fmt.Println(x)
}(4)
x := third Function
x(1)

func fourthFunction( thirdFunction(x int) int) { //Functions as variables
	return thirdFunction(x)
}

//POINTERS, &, AND DEREFERENCE, *, OPERATORS//////////////////

x := 7 //x referrence where the value 7 is stored in memory
y := &x //y point the location of 7 in the memory, 0xc0000100a0
*y = 8 // Dereference the location of in memory to 8, effectively changing x to 8

func changeValue (str *string){ //Change value of immutable data. Passing value of variable in functions
	*str = "changed"            //Doesn't give info about the variable. We need pointers.
}
//STRUCTS AND CUSTOM TYPES///////////////////

type NewStruct struct {
	attr1 int
	attr2 string
}

var p1 NewStruct = NewStruct{attr2 : "two"} //Only specify attr2, setting attr1 to default value 0
p1.attr1 = 3

//To change value of something inside a function, we need to pass the pointer of that something

type SecondStruct struct {  //Embedding a struc in a struc
	attr3 float64
	*NewStruct

}

p1 := &NewStruct{3, "Hi"}
c1 := SecondStruct{3.1416, p1} // &NewStruct{3, "Hi"}
c1.attr1 //Acces attr1 of struc attribute

//STRUCT METHODS/////////////////

func (s NewStruct) firstFunction() int { //Getter
	return s.attr1
}

func (s *NewStruct) secondFunction (n int) int { //Setter
	s.attr1 = n
}
}
