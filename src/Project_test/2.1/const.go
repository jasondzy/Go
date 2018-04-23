package main 

import (
	"fmt"
	"strings"
)

type Bitflag int //Bitflag and int are dirrerent type  int != Bitflag

const (

	Active Bitflag = 1<<iota
	Send 
	Receive
)

func (flags Bitflag) String() string{    //you can add function for any type

	var flag []string

	if flags & Send == Send{
		flag = append(flag,"Send")
	}
	if flags & Active == Active{
		flag = append(flag,"Active")
	}

	if flags & Receive == Receive{
		flag = append(flag,"Receive")
	}

	if len(flag)>0{
		return fmt.Sprintf("%d(%s)",flags,strings.Join(flag,"|"))
	}

	return "0()"

}

func main(){

	flag := Active | Send

	fmt.Println(flag)
}

