package delta

import "lb-replicator/lbaas/components"

func GetMissing(main []components.Hasher, secondary []components.Hasher) []components.Hasher {
	var missingComponents []components.Hasher
OUTER:
	for _, original := range main {
		for _, clone := range secondary {
			if original.Hash() == clone.Hash() {
				continue OUTER
			}
		}
		missingComponents = append(missingComponents, original)
	}
	return missingComponents
}

// Delta holds the information of what components.Hasher needs to be created and what components.Hasher needs to be deleted in
// order for both sides to be in sync.
type Delta struct {
	Create []components.Hasher
	Delete []components.Hasher
}

func NewDelta(main []components.Hasher, secondary []components.Hasher) Delta {
	return Delta{
		Create: GetMissing(main, secondary),
		Delete: GetMissing(secondary, main),
	}
}
