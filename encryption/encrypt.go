package encryption

import (
	utils "go-one-time-pad/utils"
)

func Encrypt(message string) (string, string) {
	message64 := utils.EncodeTo64(message)
	key := utils.GenerateRandomKey(len(message64))
	var result []rune

	for i := range message64 {
		char := message64[i]
		keyChar := key[i]
		var newChar rune

		if char < 48 {
			newChar = rune((int(char)+int(keyChar)-140)%5 + 43)
		} else if char >= 48 && char <= 57 {
			newChar = rune((int(char)+int(keyChar)-145)%10 + 48)
		} else if char >= 60 && char <= 63 {
			newChar = rune((int(char)+int(keyChar)-157)%4 + 60)
		} else if char >= 65 && char <= 90 {
			newChar = rune((int(char)+int(keyChar)-162)%26 + 65)
		} else {
			newChar = rune((int(char)+int(keyChar)-194)%26 + 97)
		}

		result = append(result, newChar)
	}

	return string(result), key
}
