package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	numInserts = 20000000 // number of inserts to perform
	numWorkers = 50
)

var insertSize = numInserts / numWorkers

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/local_test?charset=utf8&timeout=5s&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	startTime := time.Now()

	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			stmt, err := db.Prepare("INSERT INTO my_table(id, first_name, last_name, age, email, address) VALUES (?, ?, ?, ?, ?, ?)")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			for j := 1; j <= insertSize; j++ {
				index := (workerID * insertSize) + j

				firstName := randString(10, rng)
				lastName := randString(10, rng)
				email := randString(10, rng) + "@example.com"
				address := randString(20, rng)
				age := rng.Intn(100)

				_, err = stmt.Exec(index, firstName, lastName, age, email, address)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				if index%1000 == 0 { // print status every 1000 inserts
					fmt.Printf("Worker %d: Inserted %d records\n", workerID, index)
				}
			}
		}(i)
	}

	wg.Wait()

	endTime := time.Now()

	fmt.Println("Inserts completed in", endTime.Sub(startTime).Seconds(), "seconds")
}

// helper function to generate random strings
func randString(n int, rng *rand.Rand) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rng.Intn(len(letterRunes))]
	}
	return string(b)
}
