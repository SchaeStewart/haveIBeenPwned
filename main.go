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
 * TODO: Version control
 * TODO: Comments
 * TODO: Tests
 */

func createHash(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	hashstr := strings.ToUpper(fmt.Sprintf("%x", hash.Sum(nil)))
	return hashstr
}

func getPwnedHashes(hashstr string) string {
	res, err := http.Get(fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", hashstr[0:5]))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	bs := string(body)
	return bs
}

func findPwnedPassword(pwnedHashes string, hashstr string) string {
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
	hashstr := createHash(input)
	pwnedHashes := getPwnedHashes(hashstr)
	passwordPwnedCount := findPwnedPassword(pwnedHashes, hashstr)

	if passwordPwnedCount != "" {
		fmt.Printf("Your password has been pwned %s times", passwordPwnedCount)
	} else {
		fmt.Printf("Your password has not been pwned")
	}

}
