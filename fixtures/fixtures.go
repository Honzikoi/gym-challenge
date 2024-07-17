package fixtures

import (
	"log"
)

// FixtureLauncher launches all the fixtures
func FixtureLauncher() {
	if err := CreateUserFixture(); err != nil {
		log.Fatalf("Could not load user fixtures: %v", err)
	}

	if err := LoadWorkoutFixtures(); err != nil {
		log.Fatalf("Could not load workout fixtures: %v", err)
	}

	if err := LoadSessionFixtures(); err != nil {
		log.Fatalf("Could not load session fixtures: %v", err)
	}

	// Add other fixture calls here
}
