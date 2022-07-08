package password

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hlatimer266/passmanage/internal/encryption"
	"github.com/pterm/pterm"
)

const (
	ConfigFilePath = "/usr/local/etc/passmanage_config.json"
)

var (
	Passwords map[string][]byte
)

func init() {
	Passwords = make(map[string][]byte)
}

func fileToStruct() (map[string][]byte, error) {
	// check if file exists
	if _, err := os.Stat(ConfigFilePath); err != nil {
		fmt.Println(err.Error())

		f, err := os.Create(ConfigFilePath)
		if err != nil {
			fmt.Println(err)
			return Passwords, err
		}
		f.Write([]byte("{}"))
	}

	data, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		return Passwords, err
	}
	json.Unmarshal([]byte(data), &Passwords)
	return Passwords, nil
}

func List() error {
	passwords, err := fileToStruct()
	if err != nil {
		fmt.Println("error reading config file")
		return err
	}
	var accts []string
	for k := range passwords {
		accts = append(accts, k)
	}
	pterm.FgGreen.Printf("%v\n", accts)
	return nil
}

func Write(acct, masterPassword, password string) error {
	pterm.EnableDebugMessages()
	passBytes := encryption.Encrypt([]byte(password), masterPassword)

	passwordsMap, err := fileToStruct()
	if err != nil {
		return fmt.Errorf("passmange: unable to access stored passwords file [%s]", err)
	}
	passwordsMap[acct] = passBytes

	file, _ := json.MarshalIndent(passwordsMap, "", " ")
	err = ioutil.WriteFile(ConfigFilePath, file, 0777)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	pterm.FgGreen.Printf("stored password for [%s] account\n", acct)
	return nil
}

func Get(acct, passphrase string) error {
	passWordsMap, err := fileToStruct()
	if err != nil {
		return fmt.Errorf("passmange: unable to access stored passwords file [%s]", ConfigFilePath)
	}
	hashedPassword := passWordsMap[acct]
	if hashedPassword == nil {
		return fmt.Errorf("passmange: no record for account [%s]", acct)
	}
	plainTextPass, err := encryption.Decrypt(hashedPassword, passphrase)
	if err != nil {
		return fmt.Errorf("passmange: no record for account:[%s] passphrase:[%s] pair", acct, passphrase)
	}
	fmt.Printf("%s", plainTextPass)
	return nil
}

func Delete(acct string) error {
	passwords, err := fileToStruct()
	if err != nil {
		fmt.Println("error reading config file")
		return err
	}
	if _, ok := passwords[acct]; !ok {
		return fmt.Errorf("[%s] account does not exist", acct)
	}

	delete(passwords, acct)
	file, _ := json.MarshalIndent(passwords, "", " ")
	err = ioutil.WriteFile(ConfigFilePath, file, 0777)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}
