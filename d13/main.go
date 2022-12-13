package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("d12")

	file, err := os.ReadFile("INPUT")
	if err != nil {
		log.Fatal(err)
	}
	str := string(file) // convert content to a 'string'
	//  part 1
	//  fmt.Println(str)
	//  pairs := strings.Split(str, "\n\n")
	//  fmt.Println(pairs)
	//	fmt.Println("total=0")
	//	for i, x := range pairs {
	//		ys := strings.Fields(x)
	//		fmt.Printf("if (%s < %s); total = total + %d; end \n", ys[0], ys[1], i+1)
	//	}
	//	fmt.Println("total")

	// part2
	fmt.Println("xs = [")
	items := strings.Fields(str)
	for _, item := range items {
		fmt.Println(item, ",")
	}
	fmt.Println("[[2]],")
	fmt.Println("[[6]]")
	fmt.Println("]")
	fmt.Println("sort!(xs)")
	fmt.Println("a = findall(x->x==[[2]],xs)")
	fmt.Println("b = findall(x->x==[[6]],xs)")
	fmt.Println("a[1]*b[1]")

}

func trim(str string) string {
	return strings.TrimSpace(str)
}

// override base functions in Julia and then run the generated Julia Code

//function Base.isless(x::Int64, y::Vector{Int64})
// return [x] < y
// end
//
//function Base.isless(x::Vector{Int64}, y::Int64)
// return x < [y]
//end
//
//function Base.isless(x::Int64, y::Vector{Any})
// return [x] < y
//end
//
//function Base.isless(x::Vector{Any}, y::Int64)
// return x < [y]
//end
//
//function Base.isless(x::Vector{Vector{Int64}}, y::Int64)
// return x < [y]
//end
//
//function Base.isless(x::Int64, y::Vector{Vector{Any}})
// return [x] < y
//end
//
//function Base.isless(x::Vector{Vector{Any}}, y::Int64)
//  return x < [y]
//end

//function Base.isless(x::Int64, y::Vector{Vector{Int64}})
// return [x] < y
//end

//function Base.isless(x::Int64, y::Vector{Vector{Vector{Int64}}})
//  return [x] < y
//end
