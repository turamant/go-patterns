package main

import (
    "fmt"
    "sync"
)

type Singleton struct {
    value int
}

var (
    instance *Singleton
    once     sync.Once
)

func GetInstance(value int) *Singleton {
    once.Do(func() {
        instance = &Singleton{value: value}
    })
    return instance
}

func main() {
    s1 := GetInstance(42)
    s2 := GetInstance(24)

    fmt.Println("s1 address:", s1)
    fmt.Println("s2 address:", s2)
    fmt.Println("s1.value:", s1.value)
    fmt.Println("s2.value:", s2.value)
    fmt.Println("s1 == s2:", s1 == s2)
}
