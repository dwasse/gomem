# GoMem

GoMem is a minimalistic utility for determining the size of a generalized struct in Go.

After doing some memory profiling, I was frustrated with how little tooling there is for determining the total memory consumption of a variable in Go, so I made a utility to do just that. `gomem.GetMemoryConsumption()` takes an arbitrary struct as an `interface` and recursively records the memory consumption of all fields, dereferencing all pointers.

To run tests, do `go test`.

Feel free to let me know if you find any errors in the code- it's certainly been useful to me!
