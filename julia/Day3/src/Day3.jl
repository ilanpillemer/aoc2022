module Day3

using DelimitedFiles

backpacks() = readdlm("input")
compartments(a) = a[1:length(a)÷2], a[1+length(a)÷2:end]

priority(a) = intersect(a[1], a[2])
value(a) = findfirst(a[1], "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

threes() = [backpacks()[x:x+2] for x = 1:3:length(backpacks())]
badge(a) = intersect(a[1], a[2], a[3])

part1() = value.(priority.(compartments.(backpacks()))) |> sum
part2() = value.(badge.(threes())) |> sum

end # module Day3
