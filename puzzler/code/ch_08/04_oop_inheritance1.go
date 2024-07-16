package main

import "fmt"

type Player struct {
	Name  string
	Age   int
	Score int
}

func (p *Player) ShowInfo() {
	fmt.Println(p.Name, p.Age, p.Score)
}

func (p *Player) SetScore(score int) {
	p.Score = score
}

type Striker struct {
	Player
}

func (s *Striker) goal() {
	fmt.Println("Goaling!")
}

func main() {
	p := Striker{}
	p.Player.Name = "Jack"
	p.Player.Age = 18
	p.goal()
	p.Player.SetScore(1)
	p.Player.ShowInfo()
}
