package mkatcproj

import (
	"os"
	"testing"
)

func TestCppObjMaker_File(t *testing.T) {
	fileName := "test"
	objectType := "file"

	cppCode := CppCode{
		FileName:   fileName,
		ObjectType: objectType,
	}

	cppCode.ObjMaker()

	_, err := os.Stat(fileName + ".cpp")
	if os.IsNotExist(err) {
		t.Errorf("ファイルが作成されませんでした")
	}

	os.Remove(fileName + ".cpp")
}

func TestGoObjMaker_File(t *testing.T) {
	fileName := "test"
	objectType := "file"

	goCode := GoCode{
		FileName:   fileName,
		ObjectType: objectType,
	}

	goCode.ObjMaker()

	_, err := os.Stat(fileName + ".go")
	if os.IsNotExist(err) {
		t.Errorf("ファイルが作成されませんでした")
	}

	os.Remove(fileName + ".go")
}
