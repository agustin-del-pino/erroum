# erroum
Light error handling for golang.

# Install

````shell
go get -u github.com/agustin-del-pino/erroum
````

# What's Erroum?

This packages provides light way for error handling in golang. 

Golang errors are just an implementation of an interface *( `Error() string` )* so they're just strings, because of that, it's not necessary to use a structure or some other complex thing than a simple **String Pointer**.

# Quick start

Erroum provides a way to create and handle go-errors.

````go
var (
    ErrNotFound = erroum.New("the resource was not found")
)

func do() error {
    return ErrNotFound
}

func main() {

    err := do()

    if erroum.Is(err, ErrNotFound) {
        println(err.Error())
    }
}
````

# Reference

## New | *`New(d string) error`*
This function creates a new error implementation.

Where *`d`* is the error description.

````go
var err = erroum.New("i'm an error")
````

## From | *`From(c string, d string, e ...error) error`*
This function creates an error from a description and another ones.

Where *`c`* is the concatenation, *`d`* is the error description and *`e`* are the errors.

````go
var (
    ErrServiceFailure = erroum.New("the service got a failure")
    ErrNotFound = erroum.New("the resource was not found") 
    ErrNotFoundAtService = erroum.From("->", "service foo fail", ErrServiceFailure, ErrNotFound)
    // service foo fail -> the service got a failure -> the resource was not found
)
````

## Merge | *`Merge(c string, e ...error) error`*
This function creates an error from merging many errors.

Where *`c`* is the concatenation and *`e`* are the errors.

````go
var (
    ErrServiceFailure = erroum.New("the service got a failure")
    ErrNotFound = erroum.New("the resource was not found") 
    ErrNotFoundAtService = erroum.Merge("->", ErrServiceFailure, ErrNotFound)
    // the service got a failure -> the resource was not found
)
````

## Is | *`Is(e error, t error) bool`*
This function compares if two errors are equal.

Where *`e`* is the error to evaluate and *`t`* the target error.

````go
if erroum.Is(e, err) {}
````

## IsAny | *`IsAny(e error, t error) bool`*
This function evaluates if the error isn't `nil`.

Where *`e`* is the error to evaluate and *`t`* the target error.

````go
if erroum.IsAny(e) {}
````

## IsSome | *`IsSome(e error, t ...error) bool`*
This function evaluate if the error is one of the target errors.

Where *`e`* is the error to evaluate and *`t`* the target errors.

````go
if erroum.IsSome(e, err1, err2) {}
````

## InCaseOfAny | *`InCaseOfAny[T any](e error, r T, d T) T`*
This functions evaluate if the error is not `nil`. Will return a given value or a default value in case of true or false, respectively.

Where *`e`* is the error to evaluate, *`r`* the value to return in case of true and *`d`* the default value.

````go
status := erroum.InCaseOfAny[int](err, 500, 200)
````

## InCaseOf | *`InCaseOf[T any](e error, t error, r T, d T) T`*
This function combines: `Is` and `InCaseOfAny` functions.

Where *`e`* is the error to evaluate, *`t`* the target error, *`r`* the value to return in case of true and *`d`* the default value.

````go
status := erroum.InCaseOf(err, ErrNotFound, 404, 200)
````

## InCaseOfSome | *`InCaseOfSome[T any](e error, r T, d T, t ...error) T`*
This function combines: `IsSome` and `InCaseOfAny` functions.

Where *`e`* is the error to evaluate, *`r`* the value to return in case of true, *`d`* the default value and *`t`* the target errors.

````go
status := InCaseOfSome[int](err, 400, 200, ErrInvalidParam, ErrInvalidPayload, ErrInvalidQuery)
````



