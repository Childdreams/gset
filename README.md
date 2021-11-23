# GSet

计算数组的差集,交集等

详情查看 [Test](simple_set_test.go)

```go
func TestSet_Add(t *testing.T) {
	s := set.New(1, 2, 3, 4, 5, 6)
	s.Add(7, 8, 9)
	s2 := set.New(1, 2, 3, 4, 5, 6, 7, 8, 9)

	if !reflect.DeepEqual(s, s2) {
		t.Fatalf("expect 'set, got %s", "1, 2, 3, 4, 5, 6, 7, 8, 9")
	}
}
```