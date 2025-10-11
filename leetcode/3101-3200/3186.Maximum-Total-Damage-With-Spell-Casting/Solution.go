package Solution

import "sort"

func Solution(power []int) int64 {
	if len(power) == 0 {
		return 0
	}

	// 1. Calculate Total Damage per Value and collect unique damage values.
	damageSums := make(map[int]int64)
	for _, p := range power {
		damageSums[p] += int64(p)
	}

	// Get all unique damage values and sort them.
	var uniqueDamages []int
	for d := range damageSums {
		uniqueDamages = append(uniqueDamages, d)
	}
	sort.Ints(uniqueDamages)

	M := len(uniqueDamages)
	if M == 0 {
		return 0
	}

	// 2. Dynamic Programming Array F.
	// F[i] stores the max total damage considering spells with damage
	// up to and including uniqueDamages[i].
	F := make([]int64, M)

	// F[0] Base Case:
	F[0] = damageSums[uniqueDamages[0]]

	// 3. DP Iteration:
	for i := 1; i < M; i++ {
		currentDamage := uniqueDamages[i]
		D_i := damageSums[currentDamage]

		// Option 1: Skip spells of currentDamage.
		// Max damage is F[i-1].
		skipDamage := F[i-1]

		// Option 2: Cast all spells of currentDamage.
		// Gain D_i. Must add to the max damage achieved up to damage currentDamage - 3.

		// Find the index 'j' of the *largest* unique damage value <= currentDamage - 3.
		// We use sort.Search to find the first element > currentDamage - 3, then step back.

		target := currentDamage - 3

		// Search for the first index 'k' where uniqueDamages[k] > target.
		// Since sort.Search is not available for a custom search over a slice
		// and we need to find the index for uniqueDamages, we use sort.SearchInts
		// which is equivalent to finding the insertion point.

		// 'k' is the index of the first element STRICTLY GREATER THAN target.
		k := sort.SearchInts(uniqueDamages, target+1)

		// The index of the largest element <= target is k-1.
		j := k - 1

		var prevMaxDamage int64 // Will be 0 if no such spell exists (j < 0).
		if j >= 0 {
			// prevMaxDamage is the max damage considering spells up to uniqueDamages[j].
			prevMaxDamage = F[j]
		}

		castDamage := D_i + prevMaxDamage

		// F[i] = max(Option 1, Option 2)
		F[i] = max(skipDamage, castDamage)
	}

	// The result is the maximum damage considering all unique damage values.
	return F[M-1]
}
