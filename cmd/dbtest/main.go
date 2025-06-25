package main

import (
	"fmt"
	"log"
	"romantic-situation/db"
	"romantic-situation/models"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// データベース初期化
	db.Init()

	// テストユーザーの作成
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	result := db.DB.Create(user)
	if result.Error != nil {
		log.Printf("Error creating user: %v\n", result.Error)
	} else {
		fmt.Printf("Created user with ID: %d\n", user.ID)
	}

	// テストシチュエーションの作成
	situation := &models.Situation{
		UserID:      user.ID,
		Title:       "テストシチュエーション",
		Location:    "静かな図書館",
		TimeSetting: "夕暮れ時",
		Character:   "知的な司書さん",
		Situation:   "本を取ろうとして手が触れ合う",
		Mood:        "ドキドキ",
		StoryText:   "夕暮れ時の図書館で、偶然手が触れ合った瞬間、時間が止まったかのような感覚...",
		IsPublic:    true,
	}

	result = db.DB.Create(situation)
	if result.Error != nil {
		log.Printf("Error creating situation: %v\n", result.Error)
	} else {
		fmt.Printf("Created situation with ID: %d\n", situation.ID)
	}

	// お気に入りの作成
	favorite := &models.Favorite{
		UserID:      user.ID,
		SituationID: situation.ID,
	}

	result = db.DB.Create(favorite)
	if result.Error != nil {
		log.Printf("Error creating favorite: %v\n", result.Error)
	} else {
		fmt.Printf("Created favorite with ID: %d\n", favorite.ID)
	}

	// データの取得テスト
	var retrievedSituation models.Situation
	db.DB.Preload("User").First(&retrievedSituation, situation.ID)
	fmt.Printf("\nRetrieved Situation:\n")
	fmt.Printf("Title: %s\n", retrievedSituation.Title)
	fmt.Printf("Created by: %s\n", retrievedSituation.User.Username)
	fmt.Printf("Created at: %s\n", retrievedSituation.CreatedAt.Format(time.RFC3339))

	// お気に入り一覧の取得テスト
	var favorites []models.Favorite
	db.DB.Where("user_id = ?", user.ID).Preload("Situation").Find(&favorites)
	fmt.Printf("\nUser's Favorites:\n")
	for _, fav := range favorites {
		fmt.Printf("- %s\n", fav.Situation.Title)
	}
} 