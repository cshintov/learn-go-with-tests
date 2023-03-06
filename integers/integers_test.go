package integers

import "fmt"
/*import "testing"*/

/*func TestAdder(t *testing.T) {*/
	/*sum := Add(2, 2)*/
	/*want := 4*/
    /*assertCorrectMessage(t, sum, want)*/
/*}*/

/*func assertCorrectMessage(t testing.TB, got, want int) {*/
	/*t.Helper()*/
	/*if got != want {*/
		/*t.Errorf("got %q want %q", got, want)*/
	/*}*/
/*}*/

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func ExampleAddZeroes() {
	sum := Add(0, 0)
	fmt.Println(sum)
	// Output: 0
}

func ExampleAddNegatives() {
	sum := Add(-2, -1)
	fmt.Println(sum)
	// Output: -3
}

func ExampleAddMixed() {
	sum := Add(-2, 1)
	fmt.Println(sum)
	// Output: -1
}
