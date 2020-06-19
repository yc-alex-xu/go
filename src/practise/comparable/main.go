package main

import (
	"fmt"
)

/*
c++ code
  string s1 = "abc";
  string s2 = "abc";
  cout << (s1 == s2) << endl; //"1"
  string *s3 = new string("abc");
  string *s4 = new string("abc");
  cout << (s3 == s4) << endl;//"0"
*/

func stringCompare() {
	s1, s2 := "abc", "abc"
	fmt.Println(s1 == s2) //"true"
	s3 := new(string)
	*s3 = "abc"
	s4 := new(string)
	*s4 = "abc"
	fmt.Println(s3 == s4) // "false"

}

func arrayCompare() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b != c) // "true false true"
}

func main() {
	stringCompare()
	arrayCompare()
}
