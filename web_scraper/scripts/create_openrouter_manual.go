package scripts

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func CreateOpenRouterManual() {
	// Use the web_scraper to get the data
	cmd := exec.Command("/Users/lpm/Repo/ODYN Agentic/project/WebScraper/web_scraper/web_scraper", "https://openrouter.ai/api/v1/models")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running web_scraper:", err)
		return
	}

	// Write the data to a file
	err = ioutil.WriteFile("/Users/lpm/Repo/ODYN Agentic/project/WebScraper/web_scraper/data/openrouter_api_manual.json", out, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully created openrouter_api_manual.json")
}