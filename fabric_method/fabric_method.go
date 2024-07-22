package main

import (
	"fmt"
)

// Продукт
type Cake interface {
    GetName() string
}

// Конкретные продукты
type ChocolateCake struct{}

func (c *ChocolateCake) GetName() string {
    return "Chocolate Cake"
}

type StrawberryCake struct{}

func (c *StrawberryCake) GetName() string {
    return "Strawberry Cake"
}

// Фабрика
type CakeFactory interface {
    CreateCake() Cake
}

type ChocolateCakeFactory struct{}

func (f *ChocolateCakeFactory) CreateCake() Cake {
    return &ChocolateCake{}
}

type StrawberryCakeFactory struct{}

func (f *StrawberryCakeFactory) CreateCake() Cake {
    return &StrawberryCake{}
}

func main() {
    var factory CakeFactory = &ChocolateCakeFactory{}
    cake := factory.CreateCake()
    fmt.Println(cake.GetName()) // Output: Chocolate Cake

    factory = &StrawberryCakeFactory{}
    cake = factory.CreateCake()
    fmt.Println(cake.GetName()) // Output: Strawberry Cake
}
