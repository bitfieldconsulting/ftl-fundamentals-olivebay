# 01 - Testing Times

In this exercise, all we need to do is make sure our Go environment is set up and everything is working right.

We have a Go package in the [`hello.go`](hello.go) file, and some accompanying tests in [`hello_test.go`](hello_test.go). We're going to run these tests and see if they pass.

## 01.01 - Good to Go

In this directory, run the command:

**`go test`**

If everything is good, we will see this output:

```
PASS
ok      hello   0.186s
```

This tells us that all the tests passed in the package called `hello`, and that running them took 0.186 seconds. (The exact time can vary.)

## Done

You're done! Hopefully that was pretty easy. Go on to the [next exercise](../02/README.md).