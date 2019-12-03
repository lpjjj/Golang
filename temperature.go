package temperature

import "fmt"

func tempCtoF(f){
	int c
	c = print("input the degree of Celsius?")
	var f 
	f=c*9/5+32
	fmt.Printf("%d degree of Celsius= %d degree of Fahrenheit \n",c,f )
}

func tempFtoC(f){
	int f
	f = print("input the degree of Fahrenheit?")
	var c 
	c=(f-32)*5/9
	fmt.Printf("%d degree of Fahrenheit= %d degree of Celsius \n",c,f )
}
