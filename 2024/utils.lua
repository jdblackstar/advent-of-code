local utils = {}

-- Time tracking
function utils.measure_time(func, ...)
    local start = os.clock()
    local results = {func(...)}  -- Capture all return values
    local end_time = os.clock()
    
    return {
        duration = end_time - start,
        results = results
    }
end

-- Memory tracking
function utils.get_memory_usage()
    collectgarbage("collect")  -- Force garbage collection
    return collectgarbage("count")
end

-- Statistics wrapper
function utils.with_stats(func)
    return function(...)
        collectgarbage("collect")  -- Clean up before measurement
        local initial_memory = utils.get_memory_usage()
        
        local result = utils.measure_time(func, ...)
        
        collectgarbage("collect")  -- Clean up after function execution
        local final_memory = utils.get_memory_usage()
        
        -- Print statistics
        print(string.format("\nPerformance Statistics:"))
        print(string.format("Time taken: %.6f seconds", result.duration))
        print(string.format("Memory usage: %.2f KB", final_memory))
        print(string.format("Peak memory delta: %.2f KB", final_memory - initial_memory))
        
        return table.unpack(result.results)
    end
end

return utils 