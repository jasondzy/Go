package main 

/*
#include <stdio.h>
#include <stdlib.h>
void hello()
{
	printf("hello everyone i will tell you some more\n");	
}
*/
import "C"

func main(){

	C.hello()
}