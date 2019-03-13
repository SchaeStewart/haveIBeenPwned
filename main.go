// Takes a password as a command line arg and prints the number of times it has been pwned according to haveibeenpwned
//   Example: ./haveIBeenPwned Password1
//   Output: Your password has been pwned 111658 times
package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
 * TODO: Readme
 * TODO: Tests
 */

// CreateHash returns an sha1 hash of string.
func CreateHash(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	return strings.ToUpper(fmt.Sprintf("%x", hash.Sum(nil)))
}

// GetPwnedHashes gets a list of hashes from the pwnedpasswords api that are similair to the hastr.
// The pwnedpasswords api only requires the first 5 characters of the hash and returns all similair hashes.
func GetPwnedHashes(hashstr string) string {
	res, err := http.Get(fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", hashstr[0:5]))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

// FindPwnedPassword finds the given hash in a string of pwnedHashes.
// The pwnedHashes are expected to be formatted as HASH:TIMES_PWNED\r\n
func FindPwnedPassword(pwnedHashes string, hashstr string) string {
	for i, value := range strings.Split(pwnedHashes, "\r\n") {
		_ = i //TODO: idomatic handling of this?
		pwnedHash := strings.Split(value, ":")
		if pwnedHash[0] == hashstr[5:] {
			return pwnedHash[1]
		}
	}
	return ""
}

func main() {
	input := os.Args[1]
	hashstr := CreateHash(input)
	pwnedHashes := GetPwnedHashes(hashstr)
	passwordPwnedCount := FindPwnedPassword(pwnedHashes, hashstr)

	if passwordPwnedCount != "" {
		fmt.Printf("Your password has been pwned %s times", passwordPwnedCount)
	} else {
		fmt.Printf("Your password has not been pwned")
	}

}
