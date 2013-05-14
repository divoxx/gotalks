package main

import (
	"fmt"
)

type Nerfgun struct {
	AmmoCount int
}

type Doxer string

func (d Doxer) Shoot(n *Nerfgun, target Doxer) {
	if n.AmmoCount > 0 {
		fmt.Printf("%s shot %s\n", d, target)
		n.AmmoCount--
	} else {
		fmt.Println("No more ammo!")
	}
}

func main() {
	sha := Doxer("Shannon")
	dan := Doxer("Dan")
	gun := &Nerfgun{AmmoCount: 2}

	sha.Shoot(gun, dan)
	sha.Shoot(gun, dan)
	sha.Shoot(gun, dan)
}
