package main

import (
	"fmt"
	"os"
	"testing"
)

func TestReverse(t *testing.T) {
	testcase := "hello world"

	rev, err := Reverse(testcase)
	if err != nil {
		t.Fatal(err)
	}

	doubleRev, err := Reverse(rev)
	if err != nil {
		t.Fatal(err)
	}

	if testcase != doubleRev {
		t.Fatal("not matching")
	}
}

func TestSmthg(t *testing.T) {
	for _, e := range os.Environ() {
		fmt.Println("--->", e)
	}
}
