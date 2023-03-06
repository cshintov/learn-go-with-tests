package iteration

import (
	"fmt"
)

func Example() {
	fmt.Println(ContainsAny("team", "i"))
	fmt.Println(ContainsAny("fail", "ui"))
	fmt.Println(ContainsAny("ure", "ui"))
	fmt.Println(ContainsAny("failure", "ui"))
	fmt.Println(ContainsAny("foo", ""))
	fmt.Println(ContainsAny("", ""))

    /*Output:*/
    /*false*/
    /*true*/
    /*true*/
    /*true*/
    /*false*/
    /*false*/
}
