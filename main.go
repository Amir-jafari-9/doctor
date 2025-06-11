package main
#
import (
	"bufio"
	"fmt"
	"math/rand"
	"myapp/doctor"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	intro := doctor.Introductions[rand.Intn(len(doctor.Introductions))]
	fmt.Println(intro)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "quit" {
			break
		}

		foundMatch := false
		for pattern, responses := range doctor.Psychobabble {
			re := regexp.MustCompile("(?i)" + pattern) // (?i) makes it case-insensitive
			matches := re.FindStringSubmatch(userInput)
			if len(matches) > 1 {
				fragment := matches[1]
				responseTemplate := responses[rand.Intn(len(responses))]
				response := fmt.Sprintf(responseTemplate, fragment)
				fmt.Println(response)
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			// Fallback to a default response
			fallback := doctor.DefaultResponses[rand.Intn(len(doctor.DefaultResponses))]
			response := fmt.Sprintf("%s  \"%s\"", fallback, userInput)

			fmt.Println(response)
		}
	}
}
