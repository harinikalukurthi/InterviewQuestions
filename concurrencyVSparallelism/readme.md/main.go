Goroutine: Goroutines are designed for concurrency. They run concurrently, meaning they are executed in an interleaved manner on a single OS thread. This concurrency allows efficient multitasking on a single CPU core.
Traditional Thread: Threads are typically used for parallelism. They are executed in parallel on multiple CPU cores, allowing true parallel execution of tasks. Threads often have higher overhead in terms of creation and management.


what makes goroutines to have very less in size?
goroutines uses a segemented stack which can grow or shrink as needed by the goroutines.This is actually making goroutines to have a 2KB size.