package hash

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	password := "myPassword123"
	for cost := 10; cost <= 14; cost++ {
		start := time.Now()
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
		elapsed := time.Since(start)
		fmt.Printf("cost=%d, time=%v\n", cost, elapsed)
		fmt.Printf("hash=%s\n", hash)
	}
}
