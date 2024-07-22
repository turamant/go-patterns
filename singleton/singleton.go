package main

import (
    "fmt"
    "sync"
)

type singleton struct {
    value int
}

var (
    instance *singleton
    once     sync.Once
)

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{value: 42}
    })
    return instance
}

func main() {
    s1 := GetInstance()
    s2 := GetInstance()

    fmt.Println("s1 address:", s1)
    fmt.Println("s2 address:", s2)
    fmt.Println("s1.value:", s1.value)
    fmt.Println("s2.value:", s2.value)
}
