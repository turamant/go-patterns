package main

import (
	"fmt"
)

// Продукты
type Cake interface {
    GetName() string
}

type Frosting interface {
    GetName() string
}

// Конкретные продукты
type ChocolateCake struct{}

func (c *ChocolateCake) GetName() string {
    return "Chocolate Cake"
}

type VanillaFrosting struct{}

func (f *VanillaFrosting) GetName() string {
    return "Vanilla Frosting"
}

type StrawberryCake struct{}

func (c *StrawberryCake) GetName() string {
    return "Strawberry Cake"
}

type ChocolateFrosting struct{}

func (f *ChocolateFrosting) GetName() string {
    return "Chocolate Frosting"
}

// Абстрактная фабрика
type CakeFactory interface {
    CreateCake() Cake
    CreateFrosting() Frosting
}

// Конкретные фабрики
type ChocolateCakeFactory struct{}

func (f *ChocolateCakeFactory) CreateCake() Cake {
    return &ChocolateCake{}
}

func (f *ChocolateCakeFactory) CreateFrosting() Frosting {
    return &ChocolateFrosting{}
}

type StrawberryCakeFactory struct{}

func (f *StrawberryCakeFactory) CreateCake() Cake {
    return &StrawberryCake{}
}

func (f *StrawberryCakeFactory) CreateFrosting() Frosting {
    return &VanillaFrosting{}
}

func main() {
    var factory CakeFactory = &ChocolateCakeFactory{}
    cake := factory.CreateCake()
    frosting := factory.CreateFrosting()
    fmt.Printf("Created %s with %s\n", cake.GetName(), frosting.GetName())

    factory = &StrawberryCakeFactory{}
    cake = factory.CreateCake()
    frosting = factory.CreateFrosting()
    fmt.Printf("Created %s with %s\n", cake.GetName(), frosting.GetName())
}
