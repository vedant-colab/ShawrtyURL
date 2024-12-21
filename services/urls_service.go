package services

import (
	"log"
	"src/database"
)

type URL struct {
	ID         int    `json:"id"`
	ShortenURL string `json:"shorten_url"`
	ActualURL  string `json:"actual_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// SaveShortenURL inserts a new shorten URL into the database
func SaveShortenURL(shortenURL string, actualURL string) error {
	// Prepare the SQL query
	query := "INSERT INTO urls (shorten_url, actual_url) VALUES ($1, $2)"

	// Execute the query with parameters
	_, err := database.DB.Exec(query, shortenURL, actualURL)
	if err != nil {
		// Log and return the error
		log.Printf("Error inserting URL: %v\n", err)
		return err
	}

	log.Println("URL successfully saved to database")
	return nil
}

func FetchAllURL() ([]URL, error) {
	query := "select * from urls;"
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatalf("Error fetching all urls from database : %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var urls []URL

	for rows.Next() {
		var url URL
		if err := rows.Scan(&url.ID, &url.ShortenURL, &url.ActualURL, &url.CreatedAt, &url.UpdatedAt); err != nil {
			log.Printf("Error scanning row: %v\n", err)
			return nil, err
		}
		urls = append(urls, url)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error during row iteration: %v\n", err)
		return nil, err
	}

	return urls, err

}
