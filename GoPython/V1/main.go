package main

import (
        "fmt"
        "os/exec"
)

func main() {

name:="User"

cmd := exec.Command("./dist/hello.exe",name)

out, err := cmd.CombinedOutput()

if err != nil {
    fmt.Println(err)
}

fmt.Println(string(out))

}