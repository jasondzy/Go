package main

import (
    "fmt"
    "os/exec"
    "syscall"
    "os"

)

func main() {

    path, error := exec.LookPath("ls")
    if error != nil {
         panic("can not find cmd path")
    }

    fmt.Println("the cmd path is :", path)

    arg0 := []string{"ls", "-a", "-l", "-h"}
    env := os.Environ()

    err := syscall.Exec(path, arg0, env)

    if err != nil {
        panic("cmd wrong!")
    }


}
