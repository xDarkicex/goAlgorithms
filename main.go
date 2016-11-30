// Euclids Algorithm Greatest common demonmiator
// Solves GCD, LCM

//By Gentry
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	c "github.com/xDarkicex/goAlgorithms/Calc"
)

func init() {
	fmt.Println("====================================")
	fmt.Println("           New Calculator           ")
	fmt.Println("====================================")
	fmt.Println("Options: ")
	fmt.Println("GCD: greatest common demonmiator")
	fmt.Println("Enter option: ")
}

var reader = bufio.NewReader(os.Stdin)
var input string

func main() {
	input, _ = reader.ReadString('\n')
	i := strings.TrimRight(input, "\r\n")
	i = strings.ToLower(i)
	switcher(i)

}
func switcher(input string) {
	switch input {
	case "gcd":
		c.SolveGCD()
	case "lcm":
		c.SolveLCM()
	case "factorial":
		c.SolveFactorial()
	case "prime":
		c.SolveSieve()
	case "guessprime":
		c.GuessPrime()
	default:
		fmt.Println("GoodBye...")
	}
}
