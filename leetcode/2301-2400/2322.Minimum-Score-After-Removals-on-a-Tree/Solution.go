package Solution

import (
	"math"
	"sort"
)

// minimumScore calculates the minimum score achievable by removing two distinct edges from a tree.
// The score is defined as the difference between the largest and smallest XOR values of the three resulting components.
func Solution(nums []int, edges [][]int) int {
	n := len(nums) // Number of nodes

	// 1. Build adjacency list for graph traversal
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u) // Undirected graph
	}

	// 2. Calculate the total XOR sum of all node values
	totalXorSum := 0
	for _, num := range nums {
		totalXorSum ^= num
	}

	// 3. Prepare slices for DFS pre-calculation
	// subtreeXor[i]: XOR sum of nodes in the subtree rooted at i (including i)
	subtreeXor := make([]int, n)
	// tin[i]: Entry time of node i in DFS traversal
	tin := make([]int, n)
	// tout[i]: Exit time of node i in DFS traversal (after visiting all its descendants)
	tout := make([]int, n)
	timer := 0 // Global timer for assigning tin/tout values

	// DFS function to compute subtreeXor, tin, and tout
	var dfs func(u, p int) // u: current node, p: parent node
	dfs = func(u, p int) {
		timer++
		tin[u] = timer // Mark entry time for node u

		subtreeXor[u] = nums[u] // Initialize subtree XOR with current node's value

		// Traverse all neighbors of u
		for _, v := range adj[u] {
			if v == p { // Skip if v is the parent of u
				continue
			}
			dfs(v, u)                      // Recursively call DFS for child v
			subtreeXor[u] ^= subtreeXor[v] // Aggregate child's subtree XOR into current node's
		}

		timer++
		tout[u] = timer // Mark exit time for node u
	}

	// Start DFS from node 0 (arbitrary root), with -1 as its virtual parent
	dfs(0, -1)

	// Initialize minimum score to a very large value
	minScore := math.MaxInt32

	// 4. Enumerate all possible pairs of distinct nodes (i, j)
	// These nodes represent the "roots" of the two subtrees that will be cut off.
	// The edges being cut are implicitly (parent[i], i) and (parent[j], j).
	// We start from node 1 because node 0 is the root and does not have a "parent edge" to cut.
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ { // j must be different from i

			// Check if i is an ancestor of j using DFS timestamps
			// (tin[i] < tin[j] and tout[i] > tout[j])
			isIAncestorOfJ := (tin[i] < tin[j] && tout[i] > tout[j])
			// Check if j is an ancestor of i using DFS timestamps
			isJAncestorOfI := (tin[j] < tin[i] && tout[j] > tout[i])

			var xorVal1, xorVal2, xorVal3 int // The three component XOR sums

			if isIAncestorOfJ {
				// Case 1: Node i is an ancestor of node j.
				// This means j's subtree is entirely contained within i's subtree.
				// The three components are:
				//   - Component A: The subtree rooted at j. (xorVal1)
				//   - Component B: The part of i's subtree *outside* j's subtree. (xorVal2)
				//   - Component C: The rest of the tree (outside i's subtree). (xorVal3)
				xorVal1 = subtreeXor[j]
				xorVal2 = subtreeXor[i] ^ subtreeXor[j] // XOR sum of i's subtree excluding j's subtree
				xorVal3 = totalXorSum ^ subtreeXor[i]   // Total XOR sum excluding i's subtree
			} else if isJAncestorOfI {
				// Case 2: Node j is an ancestor of node i.
				// This is symmetric to Case 1. Swap roles of i and j for calculation.
				xorVal1 = subtreeXor[i]
				xorVal2 = subtreeXor[j] ^ subtreeXor[i] // XOR sum of j's subtree excluding i's subtree
				xorVal3 = totalXorSum ^ subtreeXor[j]   // Total XOR sum excluding j's subtree
			} else {
				// Case 3: Nodes i and j are not ancestors of each other.
				// Their respective subtrees are disjoint.
				// The three components are:
				//   - Component A: The subtree rooted at i. (xorVal1)
				//   - Component B: The subtree rooted at j. (xorVal2)
				//   - Component C: The rest of the tree (outside both i's and j's subtrees). (xorVal3)
				xorVal1 = subtreeXor[i]
				xorVal2 = subtreeXor[j]
				xorVal3 = totalXorSum ^ subtreeXor[i] ^ subtreeXor[j] // Total XOR sum excluding both i's and j's subtrees
			}

			// Store the three XOR values in a slice, sort them, and calculate the difference.
			currentXors := []int{xorVal1, xorVal2, xorVal3}
			sort.Ints(currentXors) // Sort to easily find min and max

			// Calculate the score for this pair of edge removals
			diff := currentXors[2] - currentXors[0] // Largest XOR - Smallest XOR

			// Update the minimum score found so far
			if diff < minScore {
				minScore = diff
			}
		}
	}

	return minScore
}
