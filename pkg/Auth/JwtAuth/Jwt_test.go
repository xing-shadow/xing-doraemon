/*
 * @Time : 2020/11/16 10:36
 * @Author : wangyl
 * @File : Jwt_test.go
 * @Software: GoLand
 */
package JwtAuth

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("xing", "123456")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InhpbmciLCJwYXNzd29yZCI6IjEyMzQ1NiIsImV4cCI6MTYwNTQ5NDQ1NSwiaXNzIjoieGluZyJ9.BVXkLtTTrHnbU1axnYvJx2hjmhFAHiOuZqDR--KS-qw"
	claim, err := ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(claim)
}
