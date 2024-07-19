package database

import "fmt"

type Image struct {
	ID        int64  `json:"id"`
	Prompt    string `json:"prompt"`
	Path      string `json:"path"`
	UserID    string `json:"user_id"`
	Public    bool   `json:"public"`
	CreatedAt string `json:"created_at"`
}

func GetImages() ([]Image, error) {
	var images []Image
	rows, err := DB.Query("SELECT * FROM images ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("GetImages: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var i Image
		if err := rows.Scan(&i.ID, &i.Prompt, &i.Path, &i.UserID, &i.Public, &i.CreatedAt); err != nil {
			return nil, fmt.Errorf("GetImages: %v", err)
		}
		images = append(images, i)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetImages: %v", err)
	}
	return images, nil
}

func GenerateImage(prompt, path, userID string, public bool) (int64, error) {
	result, err := DB.Exec(`
  INSERT INTO images
  (prompt, path, user_id, public) 
  VALUES (?, ?, ?, ?)`,
		prompt, path, userID, public)
	if err != nil {
		return 0, fmt.Errorf("CreateImage: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateImage: %v", err)
	}

	return id, nil
}
