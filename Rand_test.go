package f64

import "math/rand"

func ExampleWithRandSign() {
	for i := 10; i > 0; i-- {
		for i := 5; i > 0; i-- {
			print(WithRandSign(rand.Float64()), "\t")
		}
		println()
	}
	//:-) Output:
}

func ExampleRand() {
	for i := 10; i > 0; i-- {
		for i := 5; i > 0; i-- {
			print(Rand(), "\t")
		}
		println()
	}
	//:-) Output:
}

func ExampleRandNormal() {
	for i := 10; i > 0; i-- {
		for i := 5; i > 0; i-- {
			print(RandNormal(), "\t")
		}
		println()
	}
	//:-) Output:
}
