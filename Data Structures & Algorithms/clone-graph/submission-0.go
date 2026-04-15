/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
		return nil
	}
	cloneMap := make(map[*Node]*Node)
	return dfs(node, cloneMap)
}
func dfs(node *Node, cloneMap map[*Node]*Node) *Node {
	if clonedNode, exists := cloneMap[node]; exists {
		return clonedNode
	}

	clonedNode := &Node{Val: node.Val}
	cloneMap[node] = clonedNode

	for _, neighbor := range node.Neighbors {
		clonedNeighbor := dfs(neighbor, cloneMap)
		clonedNode.Neighbors = append(clonedNode.Neighbors, clonedNeighbor)
	}
	return clonedNode
}
