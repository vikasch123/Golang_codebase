package main

import "fmt"



// function without any return value but with parameters
func add(a int, b int){
	fmt.Println("Sum =",a+b)
}


// function with a return value

func multiply(a int , b int) int{
	return a*b
}

// function with multiple return values

func getNameAndAge() (string,int){
	return "Vikas",21
}

func isEven(a int) int{
	if a%2==0 {
		return 1;
	} else{
		return 0;
	}
}

func max(a,b int) int{
	if a>b {
		return a 
	} else{
		return b
	}
}

func main() {
	fmt.Println("Hi,Hello")
	// conditionals

	var age int = 21
	var gender string="male"
	var name string="Vikas"
	var day string="Sunday"


	if age < 18 {
		fmt.Println("You are a minor")
	} else if age >= 18 && age <= 64 {
		fmt.Println("You are an adult")

	} else {
		fmt.Println("You are too old!")

	}

	// switch case

	

	switch day {
	case "Monday":
		fmt.Println("Start of the week")

	case "Friday":
		fmt.Println("Weekend is near")

	case "Saturday":
		fmt.Println("It is a weekend")

	case "Sunday":
		fmt.Println("It is a weekend")

	default:
		fmt.Println("Just another day")

	}

	// for loop - Go's only loop

	for i := 5; i < 10; i++ {
		fmt.Println("i =", i)
	}

	nums := []int{10, 20, 30}

	nums = append(nums, 40)

	for i := 50; i < 100; i += 10 {
		nums = append(nums, i)
	}

	fmt.Println(nums)

//var gender string = "male"


fmt.Printf("Hi!, My name is %s, I am %d years old and i am a %s",name,age,gender)

fmt.Println()
// using sprintf to store formatted string in a variable


greetings := fmt.Sprintf("Hi, I am %s, How are you?",name)
fmt.Println(greetings)


add(4,5)
mult := multiply(4,3)
fmt.Println("the multiplication of 4 X 3 is :",mult)
name1,age1 := getNameAndAge()
fmt.Println(name1,age1)

// Input a number from the user and check for even or not

fmt.Println("Input a number :")
var number int
fmt.Scanln(&number)
result := isEven(number)

if result==1 {
	fmt.Println("the number %d  is even ",number)

} else{
	fmt.Printf("the number %d is not even i.e the number is odd",number)

}

fmt.Println("Enter two numbers ")
var x,y int
fmt.Scanln(&x,&y)
maximum := max(x,y)
fmt.Printf("The maximum number between %d and %d is %d",x,y,maximum)


}









