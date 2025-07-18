package golden

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// getSampleCompany creates deterministic sample data for testing
func getSampleCompany() Company {
	established, _ := time.Parse("2006-01-02", "1990-05-15")

	return Company{
		Name:        "TechCorp Inc.",
		Established: established,
		Revenue:     125000000.75,
		IsPublic:    true,
		Departments: []Department{
			{
				ID:     1,
				Name:   "Engineering",
				Budget: 5000000,
				Manager: Employee{
					ID:        101,
					FirstName: "Alice",
					LastName:  "Johnson",
					Position:  "CTO",
					Salary:    250000,
					Skills:    []string{"Leadership", "Go", "Architecture"},
					Contact: ContactInfo{
						Email: "alice.j@techcorp.com",
						Phone: "555-1001",
						Address: Address{
							Street:  "123 Tech Blvd",
							City:    "San Francisco",
							State:   "CA",
							ZipCode: "94105",
							Country: "USA",
						},
					},
					IsActive: true,
				},
				Employees: []Employee{
					{
						ID:        102,
						FirstName: "Bob",
						LastName:  "Smith",
						Position:  "Senior Engineer",
						Salary:    150000,
						Skills:    []string{"Go", "Docker", "Kubernetes"},
						Contact: ContactInfo{
							Email: "bob.s@techcorp.com",
							Phone: "555-1002",
							Address: Address{
								Street:  "456 Code Lane",
								City:    "Oakland",
								State:   "CA",
								ZipCode: "94612",
								Country: "USA",
							},
						},
						IsActive: true,
					},
				},
				Projects: []Project{
					{
						ID:           1001,
						Name:         "NextGen Platform",
						Budget:       2000000,
						Technologies: []string{"Go", "gRPC", "PostgreSQL"},
						StartDate:    time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
						EndDate:      time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
						IsCompleted:  false,
						Team:         []int{101, 102},
					},
				},
			},
			{
				ID:     2,
				Name:   "Marketing",
				Budget: 2000000,
				Manager: Employee{
					ID:        201,
					FirstName: "Carol",
					LastName:  "Williams",
					Position:  "CMO",
					Salary:    220000,
					Skills:    []string{"SEO", "Analytics", "Branding"},
					Contact: ContactInfo{
						Email: "carol.w@techcorp.com",
						Phone: "555-2001",
						Address: Address{
							Street:  "789 Market St",
							City:    "New York",
							State:   "NY",
							ZipCode: "10001",
							Country: "USA",
						},
					},
					IsActive: true,
				},
			},
		},
	}
}

func TestCompanyMarshalling(t *testing.T) {
	companyData := getSampleCompany()

	// Convert to JSON for golden file comparison
	got, err := json.MarshalIndent(companyData, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	goldenPath := filepath.Join("testdata", "TestCompanyStructure.golden")

	// Update golden files if env var is set
	if os.Getenv("UPDATE_GOLDEN") == "1" {
		if err := os.WriteFile(goldenPath, got, 0644); err != nil {
			t.Fatal(err)
		}
		return
	}

	// Read expected golden file
	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("Error reading golden file: %v", err)
	}

	require.Equal(t, string(want), string(got))
}
