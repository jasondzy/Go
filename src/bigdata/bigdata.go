package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("usage:bigdata xxxxx\n")
    }

    for row := range bigdigits[0]{
        line := ""
        // fmt.Printf("%d:%s\n",row,str);
        for _,column := range os.Args[1]{
            // fmt.Printf("column:%d\n",column)
            digit := column - '0'
            if digit >=0 && digit <= 9{
                line += bigdigits[digit][row] + " "
            } else{
                fmt.Println("digit wrong!!!")
            }
            
        }
        fmt.Println(line)
    }

}


var bigdigits = [][]string{
    {"  000  ",
     " 0   0 ",
     "0     0",
     "0     0",
     "0     0",
     " 0   0 ",
     "  000  "},
    {" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
    {" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
    {" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
    {"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
        "   4  "},
    {"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
    {" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
    {"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
    {" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
    {" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}