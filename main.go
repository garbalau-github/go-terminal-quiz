package main

// I/O lib
import "fmt"

func main() {
	fmt.Println("Welcome to Go Quiz!")
	fmt.Printf("Enter your name: ")

	var name string

	// Memory address reference
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game! \n", name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("You can't play, sorry")
		return
	} else {
		fmt.Println("You allowed to play, let's go!")
	}

	score := 0
	number_of_questions := 3;

	fmt.Printf("What is more powerfull, PS4 or PS5? ")
	var answer_one string
	fmt.Scan(&answer_one)

	if answer_one == "PS5" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
		score -= 1
	}

	fmt.Printf("Do you like chocolate ice-cream? 'Yes' or 'No'? ")
	var answer_two string
	fmt.Scan(&answer_two)

	if answer_two == "Yes" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
		score -= 1
	}

	fmt.Printf("Huimunple is a real word? 'Yes' or 'No'? ")
	var answer_three string
	fmt.Scan(&answer_three)

	if answer_three == "No" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Wrong!")
		score -= 1
	}

	fmt.Printf("You got %v questions out of %v \n", score, number_of_questions)

	// Get percentage
	percentage := (float64(score) / float64(number_of_questions)) * 100
	fmt.Printf("You scored: %v%%", percentage);
}