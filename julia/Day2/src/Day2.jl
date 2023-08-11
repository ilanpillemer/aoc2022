module Day2

using DelimitedFiles

turns() = readdlm("input", '\n', String)

function play(they, me)
    if they == me
        :draw
    elseif (they == :scissors && me == :paper) ||
           (they == :rock && me == :scissors) ||
           (they == :paper && me == :rock)
        :lost
    else
        :won
    end
end


function score(str)
    line = split(str)
    me = Dict("X" => 1, "Y" => 2, "Z" => 3)
    choice = Dict(
        "X" => :rock,
        "Y" => :paper,
        "Z" => :scissors,
        "A" => :rock,
        "B" => :paper,
        "C" => :scissors,
    )
    outcome = Dict(:lost => 0, :draw => 3, :won => 6)
    outcome[play(choice[line[1]], choice[line[2]])] + me[line[2]]
end

function should_play(they, result)
    w = Dict(:rock => :paper, :paper => :scissors, :scissors => :rock)
    l = Dict(:rock => :scissors, :paper => :rock, :scissors => :paper)
    if result == :win
        w[they]
    elseif result == :lose
        l[they]
    else
        they
    end
end

function score2(str)
    line = split(str)
    hint = Dict("X" => :lose, "Y" => :draw, "Z" => :win)
    value = Dict(:rock => 1, :paper => 2, :scissors => 3)
    choice = Dict(
        "X" => :rock,
        "Y" => :paper,
        "Z" => :scissors,
        "A" => :rock,
        "B" => :paper,
        "C" => :scissors,
    )
    my_move = should_play(choice[line[1]], hint[line[2]])
    outcome = Dict(:lost => 0, :draw => 3, :won => 6)
    outcome[play(choice[line[1]], my_move)] + value[my_move]
end

function part1()
    score.(turns()) |> sum
end

function part2()
    score2.(turns()) |> sum
end

end # module Day2
