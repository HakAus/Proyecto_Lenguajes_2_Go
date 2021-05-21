package randgen

import "fmt"

// Function that returns a pseudo-random integer using the linear cnogruentional method outlined by Dromey in "How to solve it by computer" (1982)

// Closure
func randomIntGenerator(seed int) func() int {
	var a, b, m int = 109, 853, 4096
	return func() int {
		seed = ((a * seed + b) % m) 
		return seed
	}
}


func isValidSeed(seed int) bool {
	return isPrime(seed) && 11 <= seed && seed <= 101 
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

func isValidSize(size int) bool {
	return size >= 200 || size <= 1000
}

// Para que el resultado este en el rango 0 .. 199
func mapNumber(value, start1, stop1, start2, stop2 int) int {
	return int(float32((value - start1)) / float32((stop1 - start1)) * float32((stop2 - start2)) + float32(start2))
}

func GetRandomArray(seed int, size int) []int {
	// Check that seed is prime and between 11 and 101
	if isValidSeed(seed) && isValidSize(size){
		array := make([]int, size)
		// Generate n random numbers
		nextRandom := randomIntGenerator(seed)
		for index := 0; index < size; index++ {
			array[index] = mapNumber(nextRandom(),0,4095,0,199)
		}
		return array
	} else {
		fmt.Println("Verifique que la semilla es un valor primo entre 11 y 101. Además, el tamaño del arreglo debe estar entre 200 y 1000")
	}
	return nil
}