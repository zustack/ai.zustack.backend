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

func GetUserImages(user_id int64, limit, cursor int) ([]Image, error) {
	var images []Image
	rows, err := DB.Query("SELECT * FROM images WHERE user_id = ? ORDER BY id DESC LIMIT ? OFFSET ?",
	user_id, limit, cursor)
	if err != nil {
		return nil, fmt.Errorf("GetUserImages: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var i Image
		if err := rows.Scan(&i.ID, &i.Prompt, &i.Path, &i.UserID, &i.Public, &i.CreatedAt); err != nil {
      return nil, fmt.Errorf("GetUserImages: %v", err)
		}
		images = append(images, i)
	}
	if err := rows.Err(); err != nil {
    return nil, fmt.Errorf("GetUserImages: %v", err)
	}
	return images, nil
}

func GetUserImagesCount(userID int64) (int, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM images WHERE user_id = ?", userID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("GetUserImagesCount: %v", err)
	}
	return count, nil
}

func GetImagesCount() (int, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM images").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("GetImagesCount: %v", err)
	}
	return count, nil
}

func GetImages(searchParam string, limit, cursor int) ([]Image, error) {
	var images []Image
	rows, err := DB.Query("SELECT * FROM images WHERE prompt LIKE ? ORDER BY id DESC LIMIT ? OFFSET ?",
		searchParam, limit, cursor)
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

func GenerateImage(prompt, path string, userID int64, public bool) (int64, error) {
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
