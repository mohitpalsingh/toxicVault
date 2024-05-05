package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCopyEncryptDecrypt(t *testing.T) {
	payload := "Foo not bar!"

	src := bytes.NewReader([]byte(payload))
	dst := new(bytes.Buffer)

	key := newEncryptionKey()

	if _, err := copyEncrypt(key, src, dst); err != nil {
		t.Error(err)
	}

	out := new(bytes.Buffer)
	nw, err := copyDecrypt(key, dst, out)
	if err != nil {
		t.Error(err)
	}

	if nw != 16+len(payload) {
		t.Fail()
	}

	fmt.Printf("want %+v of size %+v have %+v of size %+v\n", payload, len(payload), out.String(), len(out.String()))

	if out.String() != payload {
		t.Error("decryption failed!!!")
	}
}
