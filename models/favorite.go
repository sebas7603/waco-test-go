package models

import (
	"database/sql"
	"fmt"

	"github.com/sebas7603/waco-test-go/pkg/db"
)

var tableNameFavorites = "favorites"

type Favorite struct {
	ID     int64
	UserID int64
	RefAPI string
}

func GetFavoritesStringByUserID(user_id int64) (string, error) {
	var favString string
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ?", tableNameFavorites)
	rows, err := db.GetDB().Query(query, user_id)
	if err != nil {
		fmt.Println("Query error:", err)
		return "", err
	}

	for rows.Next() {
		var favorite Favorite
		if err := rows.Scan(&favorite.ID, &favorite.UserID, &favorite.RefAPI); err != nil {
			fmt.Println("Scan error:", err)
			return "", err
		}

		if favString == "" {
			favString = favorite.RefAPI
			continue
		}

		favString = fmt.Sprintf("%s,%s", favString, favorite.RefAPI)
	}

	return favString, nil
}

func AddFavorite(favorite *Favorite) error {
	insertQuery := fmt.Sprintf("INSERT INTO %s (user_id, ref_api) VALUES (?, ?)", tableNameFavorites)
	result, err := db.GetDB().Exec(insertQuery, favorite.UserID, favorite.RefAPI)
	if err != nil {
		fmt.Println("Insert error:", err)
		return err
	}

	id, _ := result.LastInsertId()
	favorite.ID = id

	return nil
}

func CheckFavoriteExists(favorite *Favorite) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id = ? AND ref_api = ? LIMIT 1", tableNameFavorites)

	var id int64
	err := db.GetDB().QueryRow(query, favorite.UserID, favorite.RefAPI).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("Error checking favorite in DB: %v", err)
	}

	return true, nil
}
