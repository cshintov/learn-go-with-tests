# Pointers and Errors

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors

* At some point you may wish to use structs to manage state, exposing methods
  to let users change the state in a way that you can control.
* If a symbol (variables, types, functions et al) starts with a lowercase
  symbol then it is private outside the package it's defined in.
* When you call a function or a method the arguments are copied. So if you want
  to change the state of struct you will need to pass in a the pointer.
* Struct pointers are automatically dereferenced. So no need of
  `(****w).balance`.
* Use method reciever types consistent.
* If you want to indicate an error it is idiomatic for your function to return
  an err for the caller to check and act on.
* Errors can be `nil` because the return type of `Withdraw` will be `error`,
  which is an interface. If you see a function that takes arguments or returns
  values that are interfaces, they can be nillable.
* When you don't want a copy use a pointer: for ex: database connection pool or
  a large data structure
* Since pointers can be `nil` always check what a function returns if it
  returns a pointer
