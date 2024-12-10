/*
Copyright © 2024 Nabhdeep
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
	Long:  "Set and manage configuration options, including authentication details like COOKIE",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the 'cookie' flag value
		cookie, _ := cmd.Flags().GetString("cookie")
		if cookie == "" {
			fmt.Println("Error: --cookie flag is required")
			return
		}
		envStr := fmt.Sprintf("COOKIE_TOKEN = %s\n", cookie)
		fileFound, err := findEnvFile(".env")
		if err != nil {
			fmt.Printf("Error finding .env file: %s\n", err)
			return
		}
		if fileFound {
			err := readAndSave(".env", cookie)
			if err != nil {
				fmt.Printf("Error updating .env file: %s\n", err)
			} else {
				fmt.Println(".env file updated successfully.")
			}
		} else {
			err := createFile(".env", envStr)
			if err != nil {
				fmt.Printf("Error creating .env file: %s\n", err)
			} else {
				fmt.Println(".env file created successfully.")
			}
		}
	},
}

func readAndSave(fileName string, cookie string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	updated := false
	for i, line := range lines {
		if strings.HasPrefix(line, "COOKIE_TOKEN = ") {
			lines[i] = fmt.Sprintf("COOKIE_TOKEN = %s", cookie)
			updated = true
			break
		}
	}
	if !updated {
		lines = append(lines, fmt.Sprintf("COOKIE_TOKEN = %s", cookie))
	}

	newContent := strings.Join(lines, "\n")
	err = os.WriteFile(fileName, []byte(newContent), 0644)
	if err != nil {
		return err
	}
	return nil
}

func createFile(fileName string, content string) error {
	return os.WriteFile(fileName, []byte(content), 0644)
}

func findEnvFile(fileName string) (bool, error) {
	_, err := os.Stat(fileName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	configCmd.Flags().String("cookie", "", "Provide the cookie token for configuration")
	rootCmd.AddCommand(configCmd)
}
