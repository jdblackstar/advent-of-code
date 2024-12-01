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
    -- In KB
    return collectgarbage("count")
end

-- Statistics wrapper
function utils.with_stats(func)
    return function(...)
        local initial_memory = utils.get_memory_usage()
        local result = utils.measure_time(func, ...)
        local final_memory = utils.get_memory_usage()
        
        -- Print statistics
        print(string.format("\nPerformance Statistics:"))
        print(string.format("Time taken: %.6f seconds", result.duration))
        print(string.format("Memory delta: %.2f KB", final_memory - initial_memory))
        
        return table.unpack(result.results)
    end
end

return utils 