package main

import (
	//"cerberus/app/person"
	"cerberus/services/ipfs"
	"fmt"
)

func main() {

	//institution.TestInst()
	//person.TestPers()

	newDir, err := ipfs.CreateGroupAccountsIpfsDirectory("institutionAccounts")
	fmt.Println(err)
	fmt.Println(newDir)
}
