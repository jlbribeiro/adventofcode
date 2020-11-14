package doorkey

import (
	"bufio"
	"io"
	"math"
	"sort"
	"unicode"
)

const (
	Space rune = '.'
	Wall  rune = '#'
	Me    rune = '@'
)

type Vec2 struct {
	Y int
	X int
}

func isKey(r rune) bool {
	return unicode.IsLower(r)
}

func isDoor(r rune) bool {
	return unicode.IsUpper(r)
}

func doorToKey(r rune) rune {
	return unicode.ToLower(r)
}

func keyToKeyID(r rune) int {
	return int(r - 'a')
}

func keyIDToBitMask(id int) int {
	return 1 << id
}

func keyToBitMask(r rune) int {
	return keyIDToBitMask(keyToKeyID(r))
}

type VisitPath struct {
	NSteps       int
	RequiredKeys KeySet
}

type VisitCell struct {
	Paths []VisitPath
}

func WalkTo(visitMap [][]*VisitCell, tunnels [][]rune, coords Vec2, requiredKeys KeySet, nSteps int) {
	switch {
	case coords.Y < 0,
		coords.Y >= len(tunnels),
		coords.X < 0,
		coords.X >= len(tunnels[coords.Y]):
		return
	}

	if tunnels[coords.Y][coords.X] == Wall {
		return
	}

	curCell := visitMap[coords.Y][coords.X]
	replacedWorsePath := false
	for i, path := range curCell.Paths {
		knowPathKeysAreSubset := path.RequiredKeys.IsSubsetOf(requiredKeys)
		if path.NSteps <= nSteps && knowPathKeysAreSubset {
			// We've reached this cell before requiring
			// fewer/same number of steps AND fewer/the same keys.
			return
		}

		// The current path is NOT worse than `path`; it might be better, or just
		// "different".
		// A path is better than another path in two situations:
		// 1. it requires    fewer keys AND costs fewer/same number of steps;
		// 2. it requires the same keys AND costs fewer steps.
		// Any other situation just means this is an alternative path:
		// it might be a longer path that requires fewer keys,
		// or a shorter path that requires more keys.

		isSubset, sameKeys := requiredKeys.IsSubsetOf(path.RequiredKeys), requiredKeys == path.RequiredKeys
		requiresFewerKeys := isSubset && !sameKeys
		if (requiresFewerKeys && nSteps <= path.NSteps) || (sameKeys && nSteps < path.NSteps) {
			curCell.Paths[i].NSteps = nSteps
			curCell.Paths[i].RequiredKeys = requiredKeys
			replacedWorsePath = true
			break
		}
	}

	// The current path is a different path.
	if !replacedWorsePath {
		curCell.Paths = append(curCell.Paths, VisitPath{
			NSteps:       nSteps,
			RequiredKeys: requiredKeys,
		})
	}

	// If a door is found, carry this "required key" on subsequent steps.
	if isDoor(tunnels[coords.Y][coords.X]) {
		keyForDoor := doorToKey(tunnels[coords.Y][coords.X])
		requiredKeys = requiredKeys.WithKey(keyForDoor)
	}

	//TODO: currently visiting all directions, because the criteria will exclude
	// extraneous visits; however, extra work could be avoided if we only visited
	// [direction, direction.Left() and direction.Right()], assuming we keep the
	// concept of direction.
	for _, direction := range []Direction{North, South, West, East} {
		dy, dx := direction.Offsets()
		ny, nx := coords.Y+dy, coords.X+dx
		WalkTo(visitMap, tunnels, Vec2{Y: ny, X: nx}, requiredKeys, nSteps+1)
	}
}

func PathsToKeysFrom(tunnels [][]rune, keysCoords map[rune]Vec2, coords Vec2) map[rune][]VisitPath {
	visitMap := make([][]*VisitCell, len(tunnels))
	for i := range visitMap {
		visitMap[i] = make([]*VisitCell, len(tunnels[i]))
		for j := range visitMap[i] {
			visitMap[i][j] = &VisitCell{}
		}
	}

	requiredKeys := KeySet(0)
	WalkTo(visitMap, tunnels, coords, requiredKeys, 0)

	bestPaths := make(map[rune][]VisitPath)
	for key, keyCoords := range keysCoords {
		// Skip path to itself
		if keyCoords == coords {
			continue
		}

		paths := visitMap[keyCoords.Y][keyCoords.X].Paths
		sort.Slice(paths, func(i, j int) bool {
			a := paths[i]
			b := paths[j]

			if a.NSteps != b.NSteps {
				return a.NSteps < b.NSteps
			}

			return a.RequiredKeys.IsSubsetOf(b.RequiredKeys)
		})
		bestPaths[key] = paths
	}

	return bestPaths
}

type CandidatePath struct {
	key    rune
	nSteps int
}

type CacheKey struct {
	AtKey         rune
	RemainingKeys KeySet
}

func RemainingSteps(bestPaths map[rune]map[rune][]VisitPath, cache map[CacheKey]int, curKey rune, allKeys KeySet, curKeyset KeySet) int {
	if curKeyset == allKeys {
		return 0
	}

	remainingKeys := allKeys.WithoutKeySet(curKeyset)

	cacheKey := CacheKey{AtKey: curKey, RemainingKeys: remainingKeys}
	if minCollectionSteps, ok := cache[cacheKey]; ok {
		return minCollectionSteps
	}

	minCollectionSteps := math.MaxInt64
	for remainingKeyID := 0; keyIDToBitMask(remainingKeyID) <= int(allKeys); remainingKeyID++ {
		if !remainingKeys.HasKey(remainingKeyID) {
			continue
		}

		// remainingKeyID is one of the keys left to collect.

		remainingKey := rune('a' + remainingKeyID)

		minCostToRemainingKey := math.MaxInt64
		pathsToKey := bestPaths[curKey][remainingKey]
		for _, path := range pathsToKey {
			hasNecessaryKeys := path.RequiredKeys.IsSubsetOf(curKeyset)
			if !hasNecessaryKeys {
				continue
			}

			// This path is viable; the walker possesses the necessary keys.
			// As pathsToKey is sorted by nSteps, the first viable path is the
			// only path worth pursuing.
			minCostToRemainingKey = path.NSteps
			break
		}
		if minCostToRemainingKey == math.MaxInt64 {
			continue
		}

		collectionSteps := minCostToRemainingKey + RemainingSteps(bestPaths, cache, remainingKey, allKeys, curKeyset.WithKey(remainingKey))
		if collectionSteps < minCollectionSteps {
			minCollectionSteps = collectionSteps
		}
	}

	// By this point, `minCollectionSteps` should be the minimum number of steps
	// necessary to collect `remainingKeys`, starting from `curKey`.

	cache[cacheKey] = minCollectionSteps
	return cache[cacheKey]
}

func DiscoverAllPaths(tunnels [][]rune, keysWithCoords map[rune]Vec2, initialCoords Vec2) map[rune]map[rune][]VisitPath {
	bestPaths := make(map[rune]map[rune][]VisitPath)

	bestPaths[Me] = PathsToKeysFrom(tunnels, keysWithCoords, initialCoords)

	for key, keyCoords := range keysWithCoords {
		bestPaths[key] = PathsToKeysFrom(tunnels, keysWithCoords, keyCoords)
	}

	return bestPaths
}

func ShortestPathToKeys(tunnels [][]rune) int {
	var initialCoords Vec2
	foundInitialCoords := false

	keysWithCoords := make(map[rune]Vec2)
	for y, row := range tunnels {
		for x, cell := range row {
			if isKey(cell) {
				keysWithCoords[cell] = Vec2{Y: y, X: x}
				continue

			} else if cell == Me {
				initialCoords = Vec2{Y: y, X: x}
				tunnels[y][x] = Space
				foundInitialCoords = true
				continue
			}
		}
	}
	if !foundInitialCoords {
		return math.MaxInt64
	}

	allKeys := KeySet(0)
	for key := range keysWithCoords {
		allKeys = allKeys.WithKey(key)
	}

	bestPaths := DiscoverAllPaths(tunnels, keysWithCoords, initialCoords)

	// The starting point and the keys are nodes of a graph; bestPaths contains
	// the adjacency matrix between those nodes, where which path has a cost
	// (number of steps) and preconditions (the keys required to transverse it).
	// Two nodes might be linked via more than one edge; multiple edges
	// represent alternative paths that have different costs and required keys.
	// It should be noticed that a path that costs more steps but requires fewer
	// keys might be advantageous overall; same goes for the opposite situation
	// (more/different keys and fewer steps).
	// The adjacency matrix discovery algorithm ensures that the set of paths
	// for each pair of nodes is the simplest/smallest possible:
	// no path in the set is "better" than another
	// (e.g. same keys and fewer steps, fewer keys and same steps, ...).

	cache := make(map[CacheKey]int)
	return RemainingSteps(bestPaths, cache, Me, allKeys, KeySet(0))
}

func ShortestPathFromInput(in io.Reader) int {
	var tunnels [][]rune

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		tunnels = append(tunnels, row)
	}

	return ShortestPathToKeys(tunnels)
}
