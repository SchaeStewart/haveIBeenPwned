# HaveIBeenPwned

A simple command line program that tells you the number of times a password has been pwned according to haveibeenpwned.com  
Inspired by this [video](https://www.youtube.com/watch?v=hhUb5iknVJs)

## Usage

```
./haveIBeenPwned Password1
Your password has been pwned 111658 times
```

## Details

- This program does not send the password to the haveIBeenPwned API. Instead it sends the first 5 characters of a sha1 hash.
- The haveIBeenPwned API will return hashes that are similair to the given hash. This program then looks for any complete matches and returns the results.
