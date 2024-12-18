package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	}
	return false
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if !archerIsAwake && prisonerIsAwake {
		return true
	}
	return false
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	//f Annalyn has her pet dog with her she can rescue the prisoner if the archer is asleep.
	//The knight is scared of the dog and the archer will not have time to get ready before Annalyn and the prisoner can escape.
	if (!archerIsAwake && petDogIsPresent) ||
		//Annalyn can free the prisoner if the prisoner is awake and the knight and archer are both sleeping
		(!petDogIsPresent && prisonerIsAwake && !knightIsAwake && !archerIsAwake) {
		return true
	}
	return false
}
