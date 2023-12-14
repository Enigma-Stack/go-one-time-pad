package encryption

import (
	utils "go-one-time-pad/utils"
)

func Decrypt(message string, key string) string {
    var result64 []rune

    for i := range message {
        char := message[i]
        keyChar := key[i]
        var newChar rune

        if char < 48 {
            newChar = rune((int(char) - int(keyChar) + 54) % 5 + 43)
        } else if char >= 48 && char <= 57 {
            newChar = rune((int(char) - int(keyChar) + 49) % 10 + 48)
        } else if char >= 60 && char <= 63 {
            newChar = rune((int(char) - int(keyChar) + 37) % 4 + 60)
        } else if char >= 65 && char <= 90 {
            newChar = rune((int(char) - int(keyChar) + 32) % 26 + 65)
        } else {
            newChar = rune((int(char) - int(keyChar) + 26) % 26 + 97)
        }

        result64 = append(result64, newChar)
    }

    decrypted := utils.DecodeFrom64(string(result64))
    return decrypted
}
