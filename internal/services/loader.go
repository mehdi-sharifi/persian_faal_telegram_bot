package services

import (
	"encoding/csv"
	"math/rand"
	"os"
	"persian_faal_bot/internal/models"
	"strconv"
	"time"
)

func LoadPoems(filePath string) ([]models.Poem, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var poems []models.Poem
	for _, line := range lines[1:] { // Skip header row
		id, _ := strconv.Atoi(line[0])
		poems = append(poems, models.Poem{
			ID:             id,
			Text:           line[1],
			Interpretation: line[2],
		})
	}
	return poems, nil
}

func GetRandomPoem(poems []models.Poem) models.Poem {
	rand.Seed(time.Now().UnixNano())
	return poems[rand.Intn(len(poems))]
}
