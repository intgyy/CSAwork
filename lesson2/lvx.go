package main

import "fmt"

type Person struct {
	name  string
	level int
	exp   int
	hp    int
	power int
}

type Attacker interface {
	Attack(target *Person)
}

func (p *Person) Attack(target *Person) {
	fmt.Printf("%s攻击了%s\n", p.name, target.name)
	target.hp -= p.power
	fmt.Printf("%s受到了%d点伤害，血量剩余%d\n", target.name, p.power, target.hp)
}

func main() {
	p1 := Person{
		name:  "小明",
		level: 1,
		exp:   0,
		hp:    100,
		power: 10,
	}
	p2 := Person{
		name:  "小红",
		level: 1,
		exp:   0,
		hp:    100,
		power: 5,
	}

	p1.Attack(&p2)
}
