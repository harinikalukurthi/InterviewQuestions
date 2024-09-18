Is golang is platform independent? what is platform independent? how can we achieve in go?
Do we need a compiler to run Go program
what is go build, what go run, how can you run the binaries created from build. How can you run the binary file created on windows can be run on macos
what is statatically typed?
why is goroutines have only 2 kb while threads have 2mb?
what is stack and heap?
->command to get stack size: ulimit -s
-> vm_stat will give us the entire memory usage of the computer.
what is segmented stack?
what is defer and how they are executed. what will happen if we have more than one defer in a single function. what will happen to defer statements when an exception arises? will those defers gets executed
what is param and Query? where do we use that. what is the main difference for that?
sort the array without any inbuilt method.
sol: func sorting(x [10]int){
	for i:=0;i<len(x);i++{
		for j:=0;j<len(x);j++{
			if x[i]<x[j]{
				x[i],x[j]=x[j],x[i]
			}
		}
	}
	fmt.Println(x)
}
 how can we improve the above code since it is O(n)2

 how can we implement polymorphism in go , what is polymorphism
 lets say we have a struct as 
 type student struct{
	add string
	rollno int
	name string
	age int

} I want to sort the data first by name then by age, how can we do that?
what is comparator in golang?
what is array, slice and difference.
reference type and static type

go concurrency patterns

https://medium.com/@gopinathr143/go-concurrency-patterns-a-deep-dive-a2750f98a102