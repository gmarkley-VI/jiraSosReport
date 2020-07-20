# jiraSosReport
A Scrum of Scrum generator for Jira written in GO. This is my first GO application. Feel free to help me improve it.

## Build
```# make build```

## Example Flags
```
#./jiraSosReport -h                                                                                      
   Usage of ./jiraSosReport:
     -bugzteam string
           Bugzilla Team Name
     -jirapassword string
           Password for JIRA connection
     -jirateam string
           Jira Team Name
     -jirauser string
           Usename for JIRA connection
```

## Example output
```
#./jiraSosReport -jirauser=<Username> -jirapassword=<password> -jirateam=WINC -bugzteam=WindowsContainers

--Completed\Completing Last Week--
Add "prepare-upgrade" command to WMCB
Bump and react to WNI SSH changes
Handle Windows Node draining
Watch Windows Machines 
Windows webServer deployment fails to come up in CI

--Currently Active--
Add a version sub-command
Investigate node upgrades
Add option to run the hybrid overlay as a Windows service
Remove WNI dependency and Windows Machine Config CR
Remove tracker configmap and Windows Secrets

--Remaining in Sprint--

--Bugzilla Bugs for WindowsContainers--
All Bugs: 2
Current Version Bugs: 2

```
