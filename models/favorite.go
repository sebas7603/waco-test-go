package models

import (
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
