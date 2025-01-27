package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v5"
)

// GenerateEmployeeData generates employee data and writes it to a CSV file
func GenerateEmployeeData(numRecords int, fileName string) error {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Open or create the CSV file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	// Initialize CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{"Name", "Age", "Address", "Phone", "Email", "Salary"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("could not write header: %v", err)
	}

	// Generate and write employee records
	for i := 0; i < numRecords; i++ {
		name := gofakeit.Name()
		age := strconv.Itoa(rand.Intn(48) + 18) // Age between 18 and 65
		address := gofakeit.Address().Address
		phone := gofakeit.Phone()
		email := gofakeit.Email()
		salary := strconv.Itoa(rand.Intn(100001) + 50000) // Salary between 50,000 and 150,000

		record := []string{name, age, address, phone, email, salary}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("could not write record: %v", err)
		}
	}

	fmt.Printf("Successfully generated %d records in %s\n", numRecords, fileName)
	return nil
}

func InsertEmployeeData(fileName, dbURL string) error {
	// Open the CSV file
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Read CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}

	// Skip the header
	records = records[1:]

	// Connect to PostgreSQL
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return fmt.Errorf("could not connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Prepare bulk insert
	batch := &pgx.Batch{}

	for _, record := range records {
		age, _ := strconv.Atoi(record[1])
		salary, _ := strconv.ParseFloat(record[5], 64)

		// Add insert statement to the batch
		batch.Queue(
			"INSERT INTO employees (name, age, address, phone, email, salary) VALUES ($1, $2, $3, $4, $5, $6)",
			record[0], // Name
			age,       // Age
			record[2], // Address
			record[3], // Phone
			record[4], // Email
			salary,    // Salary
		)
	}

	// Execute the batch
	br := conn.SendBatch(context.Background(), batch)
	if err := br.Close(); err != nil {
		return fmt.Errorf("could not execute batch: %v", err)
	}

	fmt.Printf("Successfully inserted %d records into the database\n", len(records))
	return nil
}

func main() {
	//// Generate 1 million employee records
	//err := GenerateEmployeeData(1000000, "employee_data.csv")
	//if err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//}

	dbURL := "postgres://postgres:Abir@123@localhost:5432/product_db"

	// Insert employee data from CSV
	err := InsertEmployeeData("employee_data.csv", dbURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
