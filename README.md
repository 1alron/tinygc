# tinygc

> A tiny mark-and-sweep garbage collector written in Go.

`tinygc` is a compact, educational implementation of manual object management and tracing garbage collection. It provides a small virtual machine with a fixed-size stack, heap objects, object graphs, cycle handling, and an automatic mark-and-sweep collection phase.

The project is intentionally small enough to read in one sitting. It is designed for learning how a garbage collector discovers reachable objects, reclaims unreachable ones, and keeps its object list consistent.

## Highlights

- Mark-and-sweep collection with a stack as the root set
- Integer and pair objects
- Recursive traversal of nested object graphs
- Correct handling of cyclic references
- Explicit heap allocation through C's `malloc` and `free`
- Tests and a benchmark for the VM's allocation path

## How It Works

The VM owns a linked list of every allocated object. Values currently on the VM stack are the garbage collector's roots.

Collection has two phases:

1. **Mark**: walk from every stack value and recursively mark reachable pairs.
2. **Sweep**: walk the heap list, free unmarked objects, and unlink them from the list. Surviving objects are unmarked for the next collection.

The collector is tracing rather than reference-counted, so cycles are handled naturally. An object that is reachable through a cycle is retained, while an isolated cycle can eventually be reclaimed.

## Requirements

- Go `1.26.5` or compatible
- A C compiler and `cgo` enabled

The current implementation uses C allocation routines, so make sure `CGO_ENABLED=1` is available in your environment.

## Getting Started

Clone the repository and run the test suite:

```bash
git clone https://github.com/1alron/tinygc.git
cd tinygc
go test ./...
```

Run the benchmark:

```bash
go test ./pkg/gc -bench=.
```

## Example

```go
package main

import gc "github.com/1alron/tinygc/pkg/gc"

func main() {
	vm := gc.NewVM()
	defer gc.FreeVM(vm)

	vm.PushInt(10)
	vm.PushInt(20)
	vm.PushPair() // The pair becomes the top value on the VM stack.

	vm.Pop() // The pair is no longer rooted.
	gc.MarkAndSweep(vm)
}
```

`PushPair` consumes the two top stack values and creates a pair whose `Head` and `Tail` point to them. Values that are popped from the stack become collectible unless they remain reachable from another rooted object.

## Package Layout

```text
pkg/gc/
├── gc.go       # Mark-and-sweep implementation
├── types.go    # Object types and allocation
├── vm.go       # VM stack and object operations
└── gc_test.go  # Tests and benchmark
```

## Inspiration

This project was inspired by Bob Nystrom's article [*Baby's First Garbage Collector*](https://journal.stuffwithstuff.com/2013/12/08/babys-first-garbage-collector/). The article's clear, hands-on explanation of building a small mark-and-sweep collector was the starting point for this Go implementation.

## Scope

This is a learning project, not a drop-in replacement for Go's garbage collector. It currently uses a fixed-size stack, supports a small set of object types, and exposes VM internals to keep the implementation easy to inspect.

The most useful next steps for experimentation are adding more object types, improving root management, and separating the allocator from the collector.

## License

No license has been added yet.