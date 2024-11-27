package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateEventCode(userId uint, eventName string) string {
	userIDStr := fmt.Sprintf("%d", userId)
	var first string
	if len(userIDStr) >= 2 {
		first = userIDStr[:2]
	} else {
		first = fmt.Sprintf("%02d", userId)
	}

	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	mid := string([]rune{letters[rand.Intn(26)], letters[rand.Intn(26)]})

	end := eventName[:2]

	return first + mid + strings.ToUpper(end)
}
