package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func exercise_3() {
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Desert Sand",
		},
		fourWheel: true,
	}

	fmt.Println(t)
	fmt.Println(t.doors)
	fmt.Println(t.color)
	fmt.Println(t.fourWheel)

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Silver",
		},
		luxury: false,
	}

	fmt.Println(s)
	fmt.Println(s.doors)
	fmt.Println(s.color)
	fmt.Println(s.luxury)
}
