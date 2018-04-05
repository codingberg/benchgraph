# benchgraph
Visualization of Golang benchmark output using Google charts

## Introduction
In Golang we can analyze algorithm efficiency by writing benchmark functions and looking at execution time in ns/op. This task might become significantly hindered by increasing number of benchmark tests. One way to handle this is to visualize multiple benchmark results and track the function curve on a graph. The `benchgraph` reads benchmark output lines, prepare data for the graph, and upload data to remote server, which enables online view and html embedding. Graph turns out to be very handy in case of many algorithms that are tested against many arguments, especially if you are studing internal algorithm design.

## Installation

```bash
git clone https://github.com/CodingBerg/benchgraph.git
cd ./benchmark
go get ./
go install
```

## Naming convention
In order for `benchgraph` to work a coder is required to follow the **naming convention** when coding benchmark functions:
```go
// Naming convention
func Benchmark[Function_name]_[Function_argument](b *testing.B){
...
}
```
For example, if we take one line from the benchmark output,
```bash
BenchmarkF1_F-4       	30000000	        53.7 ns/op
```
it will be parsed and plotted on graph as function `F1(F)=53.7`, taking `F` as an argument and `53.7` as function result. 
In short, X-axis shows function arguments, while Y-axis shows function execution time in ns/op.

## Usage
The output of benchmark is piped through `benchgraph`:

```bash
go test -bench .|benchgraph -title="Graph: F(x) in ns/op"
testing: warning: no tests to run
? PASS
√ BenchmarkF1_F-4       	30000000	        53.7 ns/op
√ BenchmarkF1_FF-4      	20000000	        62.9 ns/op
√ BenchmarkF1_FFF-4     	20000000	        70.0 ns/op
√ BenchmarkF1_FFFF-4    	20000000	        80.3 ns/op
√ BenchmarkF1_FFFFF-4   	20000000	        90.8 ns/op
√ BenchmarkF1_FFFFFF-4  	20000000	        99.5 ns/op
...
Waiting for server response ...
=========================================

http://benchgraph.codingberg.com/1

=========================================
```

In front of every line `benchgraph` places indicator whether line is parsed correctly, or not.
When you see red marks `-` or `?`, it means, either you do not follow the **naming convention** from above, or the line doesn't contain benchmark test at all. At the end, `benchgraph` returns URL to the graph. From there, follow instructions how to embed graph into custom HTML page. Also, you can just share the graph link.

## Help

```bash
benchgraph -help
Usage of benchgraph:
  -apiurl string
    	url to server api (default "http://benchgraph.codingberg.com")
  -local
      generates the response locally
  -oba value
    	comma-separated list of benchmark arguments (default [])
  -obn value
    	comma-separated list of benchmark names (default [])
  -title string
    	title of a graph (default "Graph: Benchmark results in ns/op")
```

You can filter out which functions and against which arguments you want to display on graph by passing `-obn` and `-oba` arguments. This can be very handy in case when performing many benchmark tests.

```bash
go test -bench .|benchgraph -title="Graph1: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="F,FF,FFF,FFFF,FFFFF,FFFFFF,FFFFFFF,FFFFFFFF"
```

## Hints on productivity

You can first save benchmark output and then use it later for drawing graphs. This is very handy if your benchmark tests take some time to complete.

```bash
go test -bench . > out

cat out|benchgraph -title="Graph1: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="F,FF,FFF,FFFF,FFFFF,FFFFFF,FFFFFFF,FFFFFFFF"
cat out|benchgraph -title="Graph2: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="0F,F0,F00,F000,F0000,F00000,F000000,F0000000"
```

To have all in local, you can also use the **-local** option :

```bash
go test -bench . > out

cat out|benchgraph -title="Graph1: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="F,FF,FFF,FFFF,FFFFF,FFFFFF,FFFFFFF,FFFFFFFF" -local
cat out|benchgraph -title="Graph2: Benchmark F(x) in ns/op" -obn="F2,F3,F4" -oba="0F,F0,F00,F000,F0000,F00000,F000000,F0000000" -local
```

It will generates on your temp folder, a local html file.

## Online Demo

Here we analyze efficiency of different algorithms for computing parity of uint64 numbers:

http://codingberg.com/golang/interview/compute_parity_of_64_bit_unsigned_integer

There are two graphs embedded into page behind above link:

http://benchgraph.codingberg.com/1

http://benchgraph.codingberg.com/2

*Both above links can be also shared without emebeding into HTML page.*

