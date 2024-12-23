package Mosdef


func dfsHelper(graph map[string][]string, current string, end string, path []string, allPaths *[][]string) {
    // If we've reached the end node, add the current path to allPaths
    if current == end {
        // Create a copy of the path to avoid reference issues
        pathCopy := make([]string, len(path))
        copy(pathCopy, path)
        *allPaths = append(*allPaths, pathCopy)
        return
    }

    // Explore neighbors
    for _, neighbor := range graph[current] {
        // Check if neighbor is not already in the path to prevent cycles
        if !contains(path, neighbor) {
            // Add neighbor to the path and continue DFS
            dfsHelper(graph, neighbor, end, append(path, neighbor), allPaths)
        }
    }
}

// Helper to check if a slice contains a specific element
func contains(slice []string, element string) bool {
    for _, item := range slice {
        if item == element {
            return true
        }
    }
    return false
}

// Example Usage
func DepthFirstSearch(graph map[string][]string, start, end string) [][]string {
    allPaths := [][]string{}
    dfsHelper(graph, start, end, []string{start}, &allPaths)
    //func dfsHelper(graph map[string][]string, current string, end string, path []string, allPaths *[][]string)
    return allPaths
}
