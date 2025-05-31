package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %0.2f кг.\nРост: %0.2f м.\n\n", p.Name, p.Weight, p.Height)
}
