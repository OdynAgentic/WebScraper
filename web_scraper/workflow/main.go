package main

import (
	"fmt"

	"web_scraper/scripts"
)

func main() {
	fmt.Println("Starting workflow...")

	fmt.Println("Creating OpenRouter manual...")
	scripts.CreateOpenRouterManual()

	fmt.Println("Creating DeepSeek manual...")
	scripts.CreateDeepSeekManual()

	fmt.Println("Creating OpenAI manual...")
	scripts.CreateOpenAiManual()

	fmt.Println("Workflow finished successfully!")
}