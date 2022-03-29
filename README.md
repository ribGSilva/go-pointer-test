# GO POINTER TESTS

I made a few tests about whether to use pointer of the struct and pass it through the apps or passing a copy of the struct

## The test

My test was:
- Init a struct, assing values, pass it to another function and assing new values (I know in the case of passing a copy, I won’t change the original struct).
- Run that code 1M times
- And if enabled, run the gc after it


I made 4 scenarios:

- Passing copy of the original struct to the function
- Passing the mem addr of a struct to the function
- Create a new() struct and pass it to the function
- Create a var *struct, assign and pass to the function

I made all these scenarios with 3 different structs:

- With 2 strings pointers internally
- With 2 strings internally
- With 1 string and 1 string pointer internally

## Results

Here are the results:

- With 2 strings pointers internally:
```sh
BenchmarkRun-8              3634            340269 ns/op               0 B/op          0 allocs/op
BenchmarkRunRef-8             19          58365423 ns/op        32000197 B/op    2000001 allocs/op
BenchmarkRunP-8               19          69821741 ns/op        32000084 B/op    2000001 allocs/op
BenchmarkRunNew-8             10         108008314 ns/op        64000140 B/op    4000002 allocs/op
```
- With 2 strings internally:
```sh
BenchmarkRunN-8             3687            331109 ns/op               0 B/op          0 allocs/op
BenchmarkRunNRef-8          3680            340118 ns/op               0 B/op          0 allocs/op
BenchmarkRunNP-8            3459            351436 ns/op               0 B/op          0 allocs/op
BenchmarkRunNNew-8          3375            367857 ns/op               0 B/op          0 allocs/op
```
- With 1 string and 1 string pointer internally:
```sh
BenchmarkRunM-8             3272            365833 ns/op               0 B/op          0 allocs/op
BenchmarkRunMRef-8            33          32482126 ns/op        16000029 B/op    1000000 allocs/op
BenchmarkRunMP-8              36          33233946 ns/op        16000032 B/op    1000000 allocs/op
BenchmarkRunMNew-8            18          70569504 ns/op        32000049 B/op    2000000 allocs/op
```

I made the same tests running the gc after the executions, here are the results:

- With 2 strings pointers internally:
```sh
BenchmarkRun-8              2294            711065 ns/op               8 B/op          0 allocs/op
BenchmarkRunRef-8             19          60193381 ns/op        32000085 B/op    2000001 allocs/op
BenchmarkRunP-8               18          67637015 ns/op        32000087 B/op    2000001 allocs/op
BenchmarkRunNew-8              8         142124176 ns/op        64000186 B/op    4000002 allocs/op
```
- With 2 strings internally:
```sh
BenchmarkRunN-8             1708            678103 ns/op              14 B/op          0 allocs/op
BenchmarkRunNRef-8          1718            683389 ns/op              15 B/op          0 allocs/op
BenchmarkRunNP-8            1786            820259 ns/op              12 B/op          0 allocs/op
BenchmarkRunNNew-8          1510            820217 ns/op              11 B/op          0 allocs/op
```
- With 1 string and 1 string pointer internally:
```sh
BenchmarkRunM-8             1516            894396 ns/op              11 B/op          0 allocs/op
BenchmarkRunMRef-8            31          36674718 ns/op        16000029 B/op    1000000 allocs/op
BenchmarkRunMP-8              31          36132979 ns/op        16000032 B/op    1000000 allocs/op
BenchmarkRunMNew-8            16          71783320 ns/op        32000020 B/op    2000000 allocs/op
```

*I hope I did any mistake in the tests*

## Conclusion

Well, after a lot of tests, I concluded:

- If we have no pointers inside our struct, it doen’t make much difference to pass a pointer or not, but passing a copy is slightly faster than pass a pointer.
- If we have at least one pointer inside our struct, we should pass a copy of the struct (of course, if we don’t have to change the original value)

So as the difference is not really great, is easier to always pass the copy of the struct, and if we need to change the original value, we can ether receive the pointer to the object, or return a copy of the struct with the change in the value.
