package main

import (
	"fmt"
	foo_domain "repository_pattern/domain/foo"
	foo_repository "repository_pattern/infrastructure/pg/foo"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := parseConfigFromEnv()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repo := foo_repository.NewGateway(db)

	// 1. Generate a new Foo
	newFoo := foo_domain.GenerateFoo("Sample Foo Value")
	fmt.Printf("Generated Foo: ID=%s, Value=%s\n", newFoo.Id.String(), newFoo.Value.String())

	// 2. Save the Foo
	if err := repo.SaveFoo(newFoo); err != nil {
		fmt.Printf("Failed to save foo: %v\n", err)
		return
	}
	fmt.Println("Foo saved successfully")

	// 3. Find the saved Foo
	foundFoo, err := repo.FindFooById(newFoo.Id)
	if err != nil {
		fmt.Printf("Failed to find foo: %v\n", err)
		return
	}
	fmt.Printf("Found foo: ID=%s, Value=%s, Bars=%d, Bazs=%d\n",
		foundFoo.Id.String(), foundFoo.Value.String(), len(foundFoo.Bars), len(foundFoo.Bazs))

	// 4. Add some Bars and Bazs
	foundFoo.AddBarBaz(3)
	fmt.Printf("Added 3 Bars and 3 Bazs. Now Bars=%d, Bazs=%d\n", len(foundFoo.Bars), len(foundFoo.Bazs))

	// 5. Save the updated Foo
	if err := repo.SaveFoo(foundFoo); err != nil {
		fmt.Printf("Failed to save updated foo: %v\n", err)
		return
	}
	fmt.Println("Updated Foo saved successfully")

	// 6. Find the updated Foo to verify changes
	updatedFoo, err := repo.FindFooById(newFoo.Id)
	if err != nil {
		fmt.Printf("Failed to find updated foo: %v\n", err)
		return
	}
	fmt.Printf("Verified updated foo: ID=%s, Value=%s, Bars=%d, Bazs=%d\n",
		updatedFoo.Id.String(), updatedFoo.Value.String(), len(updatedFoo.Bars), len(updatedFoo.Bazs))

	// 7. Delete the Foo
	if err := repo.DeleteFoo(newFoo.Id); err != nil {
		fmt.Printf("Failed to delete foo: %v\n", err)
		return
	}
	fmt.Println("Foo deleted successfully")

	// 8. Verify deletion
	_, err = repo.FindFooById(newFoo.Id)
	if err != nil {
		fmt.Println("Foo deletion verified - record not found")
	} else {
		fmt.Println("Warning: Foo still exists after deletion")
	}
}
