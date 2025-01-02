# install rover
~$`curl -sSL https://rover.apollo.dev/nix/latest | sh`

# If using MongoDB and air
`$ air admin adminpass`
- os.Args[0] is the path to the binary
- os.Args[1] is db username
- os.Args[2] is db pass
- Optionally, if you want to run the binary with `go run main.go`, use the credentials() function to capture the user and pass from the command-line
