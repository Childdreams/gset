package set_test

import (
	"fmt"
	"github.com/Childdreams/gset"
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := set.New(1, 2, 3, 4, 5, 6)
	s.Add(7, 8, 9)
	s2 := set.New(1, 2, 3, 4, 5, 6, 7, 8, 9)

	if !reflect.DeepEqual(s, s2) {
		t.Fatalf("expect 'set, got %s", "1, 2, 3, 4, 5, 6, 7, 8, 9")
	}
}

func TestSet_Remove(t *testing.T) {
	s := set.New(1, 2, 3, 4, 5, 6)
	s.Remove(1, 2, 3)
	s2 := set.New(4, 5, 6)

	if !reflect.DeepEqual(s, s2) {
		t.Fatalf("expect 'set, got %s", "4, 5, 6")
	}
}

func TestSet_Has(t *testing.T) {
	s := set.New(1, 2, 3, 4, 5, 6)
	if !s.Has(6) {
		t.Fatalf("expect 'set, got %s", "6")
	}
}

func TestSet_Count(t *testing.T) {
	s := set.New(1, 2, 3, 4, 5, 6)
	count := s.Count()
	if count != 6 {
		t.Fatalf("expect 'set, got %s", "4, 5, 6")
	}
}

func TestSet_Clear(t *testing.T) {
	s := set.New(1, 2, 3, 4, 5, 6)
	s.Clear()
	if s.Count() != 0 {
		t.Fatalf("expect 'set, got %s", "map[interface{}]bool")
	}
}

func TestSet_List(t *testing.T) {
	s := set.New(1, 2, 3, 4, 5, 6)
	s1 := s.List()
	if len(s1) != s.Count() {
		t.Fatalf("expect 'set, got %d", s.Count())
	}
}

func TestSet_Union(t *testing.T) {
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(10, 20, 3, 4)
	s3 := s1.Union(s2)
	if !reflect.DeepEqual(s3, set.New(1, 2, 3, 4, 10, 20)) {
		t.Fatalf("expect 'set, got %s", "1, 2, 3, 4, 10, 20")
	}
}

func TestSet_Minus(t *testing.T) {
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(10, 20, 3, 4)
	s3 := s1.Minus(s2)
	if !reflect.DeepEqual(s3, set.New(1, 2)) {
		t.Fatalf("expect 'set, got %s", "1, 2")
	}
}

func TestSet_Intersect(t *testing.T) {
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(10, 20, 3, 4)
	s3 := s1.Intersect(s2)
	if !reflect.DeepEqual(s3, set.New(3, 4)) {
		t.Fatalf("expect 'set, got %s", "3, 4")
	}
}

func TestSet_Complement(t *testing.T) {
	s1 := set.New(1, 2, 3, 4)
	s2 := set.New(10, 20, 3, 4)
	s3 := s1.Complement(s2)
	fmt.Println(s3)
	if !reflect.DeepEqual(s3, set.New(10, 20)) {
		t.Fatalf("expect 'set, got %s", "10, 20")
	}
}
