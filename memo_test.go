package main

import (
	"os"
	"testing"
)

func TestOptionCategoryByNewAction(t *testing.T) {
	makeDirectory(MEMO)
	defer os.RemoveAll(MEMO)
	filename := optionCategoryByNewAction("test.txt", "cate")
	if false == isExist(MEMO+"/cate") {
		t.Error("no create category directory")
	}
	if filename == "cate/text.txt" {
		t.Error("no much filename")
	}
}

func TestOptionTemplateByNewAction(t *testing.T) {
	makeDirectory(MEMO)
	makeDirectory(TEMPLATE)
	defer os.RemoveAll(MEMO)
	defer os.RemoveAll(TEMPLATE)
	dst, err := os.Create(MEMO + "test.txt")
	if err != nil {
		t.Error("create memo failed")
	}
	if _, err := os.Create(TEMPLATE + "template.txt"); err != nil {
		t.Error("create template failed")
	}
	optionTemplateByNewAction(dst, "template.txt")
}
