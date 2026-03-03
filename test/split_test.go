package test

import (
	"strings"
	"testing"
)

func TestSplitJwt(t *testing.T) {
	auth := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	authSplit := strings.Split(auth, "Bearer")

	token := authSplit[1]

	t.Log(token)
}
