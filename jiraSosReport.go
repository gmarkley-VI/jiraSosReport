package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/openshift/gmarkley-VI/jiraSosRepot/functions"
	"log"
)

func main() {
	jiraURL := "https://issues.redhat.com"
	username, password, bugzilla_team, jira_team := functions.ReadCredentials()

	//Jira JQL are advanced search methods to pull the report information we need. These can be used in the WebUI how ever you will need to remove \ from before " when doing so.
	var jiraJQL [3][2]string
	jiraJQL[0][0] = "project = " + jira_team + " AND (resolved >= -7d OR (status in (Done, Pending) AND sprint in openSprints())) AND priority in (Urgent, Blocker, Critical, High) ORDER BY priority DESC"
	jiraJQL[0][1] = "--Completed\\Completing Last Week--"
	jiraJQL[1][0] = "project = " + jira_team + " AND (status in (\"In Progress\", \"Code Review\") AND sprint in openSprints()) AND priority in (Urgent, Blocker, Critical, High) ORDER BY priority DESC"
	jiraJQL[1][1] = "--Currently Active--"
	jiraJQL[2][0] = "project = " + jira_team + " AND (status in (\"To Do\") AND sprint in openSprints()) AND priority in (Urgent, Blocker, Critical, High) ORDER BY priority DESC"
	jiraJQL[2][1] = "--Remaining in Sprint--"

	//Create the client
	client, _ := functions.CreatTheClient(username, password, jiraURL)

	//Loop over the jiraJQL array and Request the issue objects
	for z := 0; z < len(jiraJQL); z++ {

		var issues []jira.Issue

		// append the jira issues to []jira.Issue
		appendFunc := func(i jira.Issue) (err error) {
			issues = append(issues, i)
			return err
		}

		// SearchPages will page through results and pass each issue to appendFunc taken from the Jira Example implementation
		// In this example, we'll search for all the issues with the provided JQL filter and Print the header that goes with it.
		err := client.Issue.SearchPages(fmt.Sprintf(`%s`, jiraJQL[z][0]), nil, appendFunc)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n%s\n", jiraJQL[z][1])

		for _, i := range issues {
			//fmt.Printf("%s: %s - https://issues.redhat.com/browse/%s\n", i.Fields.Type.Name, i.Fields.Summary, i.Key)
			fmt.Printf("%s\n", i.Fields.Summary)
		}
	}

	//Get Bugzilla bugs and Display them

	allbugs, versionbugs := functions.ReturnBugs(bugzilla_team)

	fmt.Printf("\n--Bugzilla Bugs for %s--\n", bugzilla_team)
	fmt.Printf("All Bugs: %s\n", allbugs)
	fmt.Printf("Current Version Bugs: %s\n", versionbugs)
}
