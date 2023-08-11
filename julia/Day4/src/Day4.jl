module Day4

using DelimitedFiles
g(x) = [parse(Int, x[1][1]):parse(Int, x[1][2]), parse(Int, x[2][1]):parse(Int, x[2][2])]
f(x) = g([split(x[1], "-"), split(x[2], "-")])
sections() = f.(split.(readdlm("input"), ","))

contain(x) = intersect(x[1], x[2]) == x[1] || intersect(x[1], x[2]) == x[2]
share(x) = !isdisjoint(x[1], x[2])

part1() = contain.(sections()) |> sum
part2() = share.(sections()) |> sum
end # module Day4
