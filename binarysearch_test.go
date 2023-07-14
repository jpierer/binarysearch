package binarysearch

import "testing"

func TestSearch1(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	pos, err := Search(array, 1)
	if err != nil {
		t.Error("wrong")
	}
	if pos != 0 {
		t.Error("wrong")
	}
}

func TestSearch2(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	pos, err := Search(array, 2)
	if err != nil {
		t.Error(err)
	}
	if pos != 1 {
		t.Error("wrong")
	}
}

func TestSearch3(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	pos, err := Search(array, 4)
	if err != nil {
		t.Error("wrong")
	}
	if pos != 3 {
		t.Error("wrong")
	}
}

func TestSearch4(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	pos, err := Search(array, 6)
	if err == nil {
		t.Error("wrong")
	}
	if pos != -1 {
		t.Error("wrong")
	}
}
