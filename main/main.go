package main

import (
    "runtime"
    "fmt"
    "bytes"
    "strconv"
)

func main(){
    fmt.Println(runtime.NumCPU())
    fmt.Println(runtime.NumGoroutine())
    go sayHello()
    go sayWorld()
    var str string
    fmt.Scan(&str)
}

func sayHello(){
    for i := 0; i < 10; i++{
        fmt.Print("hello ")
        runtime.Gosched()
    }
}

func sayWorld(){
    for i := 0; i < 10; i++ {
        fmt.Println("world ")
        runtime.Gosched()
    }
}


func getGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}