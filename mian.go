package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// 1. 클라이언트 생성 (API 키 입력)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyAicO4RS3m1xM1QGzYBI8USb5gavfFrJYQ"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 2. 모델 설정 (성능이 좋은 Gemini 3 Flash 사용 권장)
	model := client.GenerativeModel("gemini-3-flash")

	// 3. 콘텐츠 생성 요청
	resp, err := model.GenerateContent(ctx, genai.Text("Go 언어로 Gemini API를 사용하는 장점 3가지를 알려줘."))
	if err != nil {
		log.Fatal(err)
	}

	// 4. 결과 출력
	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
}
