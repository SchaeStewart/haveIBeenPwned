package main

import (
	"testing"
)

func TestCreateHash(t *testing.T) {
	if CreateHash("Password1") != "70CCD9007338D6D81DD3B6271621B9CF9A97EA00" {
		t.Error("Expected Password1 to equal 70CCD9007338D6D81DD3B6271621B9CF9A97EA00")
	}

	if CreateHash("") != "DA39A3EE5E6B4B0D3255BFEF95601890AFD80709" {
		t.Error("Expected \"\" to equal DA39A3EE5E6B4B0D3255BFEF95601890AFD80709")
	}
}

func TestFindPwnedPassword(t *testing.T) {
	// N.b. the hashes in hashlist are missing the first 5 characters because the haveibeenpwned api does not return them
	hashlist := "9007338D6D81DD3B6271621B9CF9A97EA00:10\r\n3EE5E6B4B0D3255BFEF95601890AFD80709:1\r\nB7CD6A8B88F5A13BDE26E12C7C2F348C920:1\r\n"

	if FindPwnedPassword(hashlist, "70CCD9007338D6D81DD3B6271621B9CF9A97EA00") != "10" {
		t.Error("Expected 10")
	}

	if FindPwnedPassword(hashlist, "F860A40C6221137A8E6DCE056CE9BBC06713019F") != "" {
		t.Error("Expected an empty string")
	}

}

// TODO: TestGetPwnedHashes
