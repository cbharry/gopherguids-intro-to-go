//For each collection type, array, slice, and map, write a test function that does the following:
//
//Setup an exp variable filled with a minimum of 4 pieces of data. You may use any data type you wish. Example: exp := [4]string{"John", "Paul", "George", "Ringo"}
//Setup an empty act variable.
//Iterate through the exp variable and copy its values into the act variable.
//Iterate through the act variable and assert that its contents match that of the exp variable.
//For slice and map types assert that the length of act and exp are the same.
//For map when iterating through act to assert it has the same contents of exp,
//assert that the key being requested from exp exists. Hint: Use the “magic ok”.

//
package module02

import "testing"

func TestArray(t *testing.T) {
	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	for i, beatle := range exp {
		act[i] = beatle
	}

	for i, _ := range act {
		if act[i] != exp[i] {
			t.Fatalf("Elements differ at index: %d, %s != %s", i, act[i], exp[i])
		}
	}
}

func TestSlice(t *testing.T) {
	exp := []string{"John", "Paul", "George", "Ringo"}
	act := []string{}

	for _, beatle := range exp {
		act = append(act, beatle)
	}

	for i, _ := range act {
		if act[i] != exp[i] {
			t.Fatalf("Elements differ at index: %d, %s != %s", i, act[i], exp[i])
		}
	}

	if len(act) != len(exp) {
		t.Fatalf("Length of exp and act differ")
	}
}

func TestMap(t *testing.T) {
	exp := map[int]string{
		0: "John",
		1: "Paul",
		2: "George",
		3: "Ringo",
	}
	act := make(map[int]string)

	for k, beatle := range exp {
		act[k] = beatle
	}

	for k, _ := range act {
		if _, ok := exp[k]; !ok {
			t.Fatalf("Key %d doesn't exist in exp", k)
		}
		if act[k] != exp[k] {
			t.Fatalf("Elements differ at key: %d, %s != %s", k, act[k], exp[k])
		}
	}

	if len(act) != len(exp) {
		t.Fatalf("Length of exp and act differ")
	}
}
