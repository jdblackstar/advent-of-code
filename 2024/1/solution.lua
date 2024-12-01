local utils = require("2024.utils")

local function read_input(file_path)
    local file = assert(io.open(file_path, "r"))
    local content = file:read("*all")
    file:close()
    return content
end

-- Part 1 solution
local function solve_part1(input)
    local list1 = {}
    local list2 = {}
    
    for line in input:gmatch("[^\n]+") do
        local num1, num2 = line:match("(%d+)%s+(%d+)")
        table.insert(list1, tonumber(num1))
        table.insert(list2, tonumber(num2))
    end
    
    table.sort(list1)
    table.sort(list2)
    
    local total = 0
    for i = 1, #list1 do -- #list1 is length of list1
        local diff = math.abs(list1[i] - list2[i])
        total = total + diff
    end
    
    return total, #list1
end

-- Part 2 solution
local function solve_part2(input)
    local list1 = {}
    local list2 = {}
    
    for line in input:gmatch("[^\n]+") do
        local num1, num2 = line:match("(%d+)%s+(%d+)")
        table.insert(list1, tonumber(num1))
        table.insert(list2, tonumber(num2))
    end
    
    table.sort(list1)
    table.sort(list2)
    
    local similarity_score = 0
    
    for i = 1, #list1 do -- #list1 is length of list1
        local current = list1[i]
        local occurrences = 0
        
        for k = 1, #list2 do
            if list2[k] == current then
                occurrences = occurrences + 1
            end
        end
        
        -- Add to similarity score (number × times it appears)
        similarity_score = similarity_score + (current * occurrences)
    end
    
    return similarity_score, #list1
end

-- Main execution
local input = read_input("2024/1/input.txt")

-- Wrap solutions with statistics
local part1_with_stats = utils.with_stats(solve_part1)
local part2_with_stats = utils.with_stats(solve_part2)

print("Part 1:")
local result1, pairs = part1_with_stats(input)
print(string.format("Result: %d (processed %d pairs)", result1, pairs))

print("\nPart 2:")
local result2 = part2_with_stats(input)
print("Result:", result2) 