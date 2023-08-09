module Day1

greet() = print("Hello Ilan!")

function solve1()
    f = open("input")
    str = read(f, String)
    xs = split(str, "\n\n")
    process(el) = map(row -> parse(Int, row), split(el))
    println("Part 1: $(sum.(process.(xs)) |> maximum)")
    println("Part 2: $(sort(sum.(process.(xs)), rev = true)[1:3] |> sum)")
end

end # module Day1
