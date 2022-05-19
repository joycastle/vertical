package util

import (
	"fmt"
	"testing"
)

func TestCase_file(t *testing.T) {
	if names, err := ReadDirNames("/Users/mac2022/joycastle/code/vertical/util"); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		fmt.Println(names)
	}

	if names, err := ReadDirNamesWithRelativePath("/Users/mac2022/joycastle/code/vertical/util/"); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		fmt.Println(names)
	}

	if names, err := ReadDirNamesWithAbsoluePath("./"); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		fmt.Println(names)
	}

}
