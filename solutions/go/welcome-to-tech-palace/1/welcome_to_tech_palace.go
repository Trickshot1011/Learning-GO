package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	solution := ""
    solution = strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
    return solution
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    cleanMsg := ""
    var st string
	for i := 0; i < len(oldMsg); i++ {
        st = string(oldMsg[i])
    	if st != "*" && st != "\n"  {
            cleanMsg += st
        }
    }
    return strings.TrimSpace(cleanMsg)
}
