## passmanage: local password storage utility
store passwords in an encrypted file and retrieve from local machine.

### build 
build the binaries `make build` and run `sudo cp dist/<desired_os>/passmanage`

### usage
* create - `passmanage create <account_name>` and when prompted enter your password
* get - `passmanage get <account_name>` to copy to clip board add ` | pbcopy` at the end of the command
* delete - `passmanage delete <account_name> ` delete account + password pair from local storage
* list - `passmanage list` see all accounts you've stored passwords for
* generate - `passmanage generate` create a random password consisting of at least 1 upper case letter, 1 special charater, 1 number and length of 8 total characters 