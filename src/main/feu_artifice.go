package main

import (
	"fmt"
	"math/rand"
	"time"
)

func color(ind int) string {
	switch ind {
	case 0:
		return "\x1B[31m"
	case 1:
		return "\x1B[32m"
	case 2:
		return "\x1B[33m"
	case 3:
		return "\x1B[34m"
	case 4:
		return "\x1B[35m"
	case 5:
		return "\x1B[36m"
	default:
		return "\x1B[37m"
	}
}

func power(ind int) string {
	switch ind {
	case 0:
		return "boom"
	case 1:
		return "Boom"
	case 2:
		return "BOOM"
	case 3:
		return "BOOOOOM"
	case 4:
		return "BOOOOOOOOM"
	case 5:
		return "."
	default:
		return "BOoOoOoOoOoOM"
	}
}

func duration(ind int) {
	switch ind {
	case 0:
		time.Sleep(1000 * time.Millisecond)
	case 1:
		time.Sleep(500 * time.Millisecond)
	case 2:
		time.Sleep(800 * time.Millisecond)
	case 3:
		time.Sleep(1000 * time.Millisecond)
	case 4:
		time.Sleep(15000 * time.Millisecond)
	case 5:
		time.Sleep(20000 * time.Millisecond)
	default:
		time.Sleep(1000 * time.Millisecond)
	}
}

func feuxArtifices(dur int, pow int, clr int) {
	duration(dur)
	fmt.Print(color(clr), power(pow))
}

func main() {
	for {
		d := rand.Intn(6)
		c := rand.Intn(6)
		p := rand.Intn(6)
		go feuxArtifices(d, p, c)
	}
}
