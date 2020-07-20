package functions

import (
	"flag"
)

// Read in command line variables
func ReadCredentials() (string, string, string, string) {
	usernamePtr := flag.String("jirauser", "", "Usename for JIRA connection")
	passwordPtr := flag.String("jirapassword", "", "Password for JIRA connection")
	jirateamPtr := flag.String("jirateam", "", "Jira Team Name")
	bugzillateamPtr := flag.String("bugzteam", "", "Buzilla Team Name")

	flag.Parse()
	return *usernamePtr, *passwordPtr, *bugzillateamPtr, *jirateamPtr
}
