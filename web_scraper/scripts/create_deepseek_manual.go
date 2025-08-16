package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func createDeepSeekManual() {
	// Use the web_scraper to get the data
	cmd := exec.Command("../web_scraper", "deepseek api documentation")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running web_scraper:", err)
		return
	}

	// Write the data to a file
	err = ioutil.WriteFile("../data/deepseek_api_manual.json", out, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully created deepseek_api_manual.json")
}