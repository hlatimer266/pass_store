package parse

import (
	"fmt"
	"os/user"
	"syscall"

	"github.com/hlatimer266/passmanage/internal/generate"
	"github.com/hlatimer266/passmanage/internal/password"
	"github.com/pterm/pterm"
	"golang.org/x/term"
)

var (
	acctName string
)

func help() {
	pterm.FgGreen.Println("NAME:")
	pterm.FgGreen.Printf("\tpassmanage - manage local storage of passwords\n\n")
	pterm.FgGreen.Println("USAGE:")
	pterm.FgGreen.Printf("\tpassmanage [create|get|delete] {YOUR_ACCOUNT_NAME}\n\n")
	pterm.FgGreen.Println("COMMANDS:")
	pterm.FgGreen.Printf("\tcreate\t store password by account name and passphrase\n")
	pterm.FgGreen.Printf("\tget\t retrieve password by account name and passphrase\n")
	pterm.FgGreen.Printf("\tdelete\t remove password for specified account using passphrase\n")
}

func validArgs(args []string) (string, error) {
	var cmd string
	if len(args) == 1 {
		return "", fmt.Errorf("passmanage: try 'passmanage --help' for more information")
	}
	cmd = args[1]
	switch {
	case len(args) == 2 && (cmd == "list" || cmd == "generate"):
		return cmd, nil
	case len(args) == 2 && (cmd == "create" || cmd == "get" || cmd == "delete"):
		return "", fmt.Errorf("passmanage: try 'passmanage --help' for more information")
	case len(args) == 2 && (cmd != "create" && cmd != "get" && cmd != "delete" && cmd != "--help"):
		return "", fmt.Errorf("passmanage: try 'passmanage --help' for more information")
	default:
		return cmd, nil
	}
}

func CmdArgs(args []string) error {
	cmd, err := validArgs(args)
	if err != nil {
		return err
	}

	if cmd != "--help" && cmd != "list" && cmd != "generate" {
		acctName = args[2]
	}

	user, _ := user.Current()

	switch cmd {
	case "create":
		fmt.Print("PASSWORD:")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return err
		}
		fmt.Println()
		return password.Write(acctName, user.Username, string(bytePassword))
	case "get":
		return password.Get(acctName, user.Username)
	case "list":
		return password.List()
	case "generate":
		return generate.Password()
	case "delete":
		return password.Delete(acctName)
	case "--help":
		help()
		return nil
	default:
		return fmt.Errorf("passmanage: illegal option -- %s. try 'passmanage --help' for more information", args[1])
	}

}
