package secret

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	length = 256
	chars  = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	outdateTime = 1 * time.Hour
	// outdateTime = 5 * time.Second
)

func NewKeyString() string {
	arr := make([]string, length)
	rand.Seed(time.Now().UnixNano())
	for i := range length {
		k := rand.Intn(len(chars))
		arr[i] = chars[k]
	}
	return strings.Join(arr, "")
}

func (key *Key) OutDated() bool {
	if key.ID == 0 {
		return true
	}

	if key.CreateAt.Add(outdateTime).Before(time.Now().UTC()) {
		log.Printf("key created at: %v", key.CreateAt.Local())
		return true
	}
	return false
}
