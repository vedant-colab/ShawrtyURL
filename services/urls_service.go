package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"src/database"
	"src/utils"
)

type URL struct {
	ID         int    `json:"id"`
	ShortenURL string `json:"shorten_url"`
	ActualURL  string `json:"actual_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func SaveActualURL(actualURL string) (int, error) {
	var id int

	query := "INSERT INTO urls (actual_url) VALUES ($1) RETURNING id"
	err := database.DB.QueryRow(query, actualURL).Scan(&id)
	fmt.Println(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no rows returned")
		}
		return 0, fmt.Errorf("failed to insert URL: %v", err)
	}

	return id, nil
}

// SaveShortenURL inserts a new shorten URL into the database
func SaveShortenURL(actualURL string) (string, error) {
	baseID, err := SaveActualURL(actualURL)
	fmt.Println(baseID)
	fmt.Println(err)
	if err != nil {
		return "", err
	}
	shortenedURL := utils.GenerateRandom(baseID)

	query := "UPDATE urls SET shorten_url = $1 WHERE id = $2"
	_, err = database.DB.Exec(query, shortenedURL, baseID)
	if err != nil {
		return "", fmt.Errorf("failed to save shortened URL: %v", err)
	}

	return shortenedURL, nil
}

// Fetches all the urls saved in database
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

func FetchActualURL(shortenURL string) (string, error) {
	var actualURL string

	query := "SELECT actual_url FROM urls WHERE shorten_url = $1"
	err := database.DB.QueryRow(query, shortenURL).Scan(&actualURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("no rows returned")
		}
		return "", fmt.Errorf("failed to fetch actual URL: %v", err)
	}

	return actualURL, nil
}
