package main


type par struct{
	Size int
	Bomb int
}

var (
	Pars = []par{
		par{5,10},
		par{10,20},
		par{20,30},
		par{30,40},
	}
)

type Par int
const(
	easy = Par(iota)
	medium
	difficult
	hard
)
