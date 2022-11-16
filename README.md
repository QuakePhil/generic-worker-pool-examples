Examples for https://github.com/QuakePhil/generic-worker-pool

## Test run
```
go test -v -race
```
```
=== RUN   ExampleSleepWorker
--- PASS: ExampleSleepWorker (1.03s)
=== RUN   ExampleSleepWorkers
--- PASS: ExampleSleepWorkers (0.10s)
=== RUN   ExamplePrimesNew
--- PASS: ExamplePrimesNew (1.32s)
=== RUN   ExamplePrimesNewConcurrent
--- PASS: ExamplePrimesNewConcurrent (0.33s)
=== RUN   ExamplePrimesNewCustomInputChannel
--- PASS: ExamplePrimesNewCustomInputChannel (0.32s)
PASS
ok  	github.com/quakephil/generic-worker-pool-examples	3.230s
```
