package randgen

import "fmt"

// Function that returns a pseudo-random integer using the linear cnogruentional method outlined by Dromey in "How to solve it by computer" (1982)

// Closure
func RandomIntGenerator(seed int) func() int {
	var a, b, m int = 109, 853, 4096
	return func() int {
		seed = ((a * seed + b) % m) 
		return seed
	}
}

func isValidSeed(seed int) bool {
	if isPrime(seed) && 11 <= seed && seed <= 101 {
		return true
	} else {
		fmt.Println("Verifique que la semilla es un valor primo entre 11 y 101.")
		return false
	}
}

func isPrime(number int) bool {
	i := 2
	for i < number {
		if number%i == 0 {
			return false
		}
		i++
	}
	return true
}

/* 
Función que escala un número de un rango a otro.
En el proyecto se utiliza para mapear los valores del rango 0..4095
a 0..199
*/
func mapNumber(value, start1, stop1, start2, stop2 int) int {
	return int(float32((value - start1)) / float32((stop1 - start1)) * float32((stop2 - start2)) + float32(start2))
}

func GetRandomArray(seed, size int) []int {
	// Check that seed is prime and between 11 and 101
	if isValidSeed(seed) {
		array := make([]int, size)
		// Generate n random numbers
		nextRandom := RandomIntGenerator(seed)
		for index := 0; index < size; index++ {
			array[index] = mapNumber(nextRandom(), 0, 4096, 0, 200)
		}
		return array
	}
	return nil
}