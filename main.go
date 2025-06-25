package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"romantic-situation/db"
	"time"

	"github.com/joho/godotenv"
)

type Situations struct {
	Locations  []string `json:"locations"`
	Times      []string `json:"times"`
	Situations []string `json:"situations"`
	Characters []string `json:"characters"`
	Moods      []string `json:"moods"`
}

type Response struct {
	Location   string `json:"location"`
	Time       string `json:"time"`
	Character  string `json:"character"`
	Situation  string `json:"situation"`
	Mood       string `json:"mood"`
	StoryText  string `json:"story_text"`
}

var situationsData Situations

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// データベースの初期化
	db.Init()

	// シード値を設定してランダム性を確保
	rand.Seed(time.Now().UnixNano())

	// JSONファイルを読み込む
	data, err := os.ReadFile("situations.json")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// JSONをパース
	if err := json.Unmarshal(data, &situationsData); err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	// ルートハンドラー（静的ファイル配信用）
	http.Handle("/", http.FileServer(http.Dir("static")))

	// API エンドポイント
	http.HandleFunc("/api/generate", generateHandler)

	// ポート番号を環境変数から取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	// CORSヘッダーを設定
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")

	// OPTIONSリクエストの場合は早期リターン
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// GETメソッド以外は許可しない
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// ランダムな要素を選択
	location := situationsData.Locations[rand.Intn(len(situationsData.Locations))]
	timeStr := situationsData.Times[rand.Intn(len(situationsData.Times))]
	situation := situationsData.Situations[rand.Intn(len(situationsData.Situations))]
	character := situationsData.Characters[rand.Intn(len(situationsData.Characters))]
	mood := situationsData.Moods[rand.Intn(len(situationsData.Moods))]

	// 物語形式のテキストを生成
	storyText := fmt.Sprintf("ある%sのこと。%sで%sと二人きりになった。\nそんな時、突然%s…\n雰囲気は%sに包まれていった…",
		timeStr, location, character, situation, mood)

	// レスポンスを構築
	response := Response{
		Location:  location,
		Time:      timeStr,
		Character: character,
		Situation: situation,
		Mood:      mood,
		StoryText: storyText,
	}

	// JSONとしてレスポンスを返す
	json.NewEncoder(w).Encode(response)
} 