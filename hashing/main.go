package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

// hashString function hashes a given string using SHA-256 and returns the hexadecimal representation
func hashString(input string) string {
    hash := sha256.New()             // Create a new SHA-256 hash
    hash.Write([]byte(input))        // Write the input data to the hash
    hashBytes := hash.Sum(nil)       // Calculate the hash
    return hex.EncodeToString(hashBytes) // Encode the hash to a hexadecimal string
}

func main() {
    // Test strings
    testStrings := []string{
        "hello",
        "world",
        "GoLang",
        "hashing",
    }

    // Hash each test string and print the result
    for _, str := range testStrings {
        hashedStr := hashString(str)
        fmt.Printf("Original: %s, SHA-256 Hash: %s\n", str, hashedStr)
    }
}
