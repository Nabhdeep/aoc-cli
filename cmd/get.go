/*
Copyright © 2024 Nabhdeep
*/
package cmd

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type info struct {
	Year string
	Day  string
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download the day's puzzle.",
	Long:  `Fetch the puzzle input for the current day or a specified year and day from the Advent of Code website.`,
	Run: func(cmd *cobra.Command, args []string) {
		_y := args[0]
		_d, _ := cmd.Flags().GetString("d")
		if _d == "" {
			fmt.Printf("-d must be passed \n")
			return
		}
		fmt.Printf("DAY FALG IS %s\n", _d)
		arguments := info{
			Year: _y,
			Day:  _d,
		}
		getCurrDay(arguments)
	},
}

func getTokenFromEnv(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "COOKIE_TOKEN = ") {
			_cookie := strings.Split(line, " = ")
			return _cookie[len(_cookie)-1], nil
		}
	}
	return "", errors.New(".env file not present: try creating .env file from config --cookie")
}

func getCurrDay(paras info) {

	var baseUrl string = "https://adventofcode.com"
	cookie, err := getTokenFromEnv(".env")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	d, err := getDay(baseUrl, cookie, paras)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	err = makeDayFile(paras, d)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	return
}

func makeDayFile(params info, content string) error {
	_path := fmt.Sprintf("%s/day%s", params.Year, params.Day)
	_file, _ := Findfile(_path)
	fmt.Println(_file)
	if !_file {
		// mkdir
		err := os.Mkdir(_path, 0755)
		if err != nil {
			return err
		}
		//go file
		var _f *os.File
		_f, err = os.Create(_path + "/day" + params.Day + ".go")
		if err != nil {
			return err
		}
		_f.Close()
		// input
		err = os.WriteFile(_path+"/input.txt", []byte(content), 0755)
		if err != nil {
			return err
		}
	} else {
		var _opt string
		fmt.Println("File Exist Do you want to rewrite the puzzle input? (y) , (n)")
		fmt.Println(_path + "/input.txt")
		fmt.Scan(&_opt)
		switch _opt {
		case "y":
			err := os.WriteFile(_path+"/input.txt", []byte(content), 0755)
			if err != nil {
				return err
			}
			return nil
		default:
			return nil
		}
	}

	return nil
}
func Findfile(fileName string) (bool, error) {
	_, err := os.Stat(fileName)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func getDay(baseUrl, _cookie string, params info) (string, error) {
	url := fmt.Sprintf("%s/%s/day/%s/input", baseUrl, params.Year, params.Day)
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		err_str := fmt.Sprintf("Request failed - %v", err)
		return "", errors.New(err_str)
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: _cookie,
	}
	request.AddCookie(cookie)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		err_str := fmt.Sprintf("Response failed with - %v\n", err)
		return "", errors.New(err_str)
	}
	if response.StatusCode != http.StatusOK {
		err_str := fmt.Sprintf("Request failed with - %d \n", response.StatusCode)
		return "", errors.New(err_str)
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		err_str := fmt.Sprintf("Body Parsing failed with  - %v \n", err)
		return "", errors.New(err_str)
	}
	defer response.Body.Close()

	return string(bodyBytes), nil

}

// Add flags

func init() {
	getCmd.Flags().String("d", "", "Need Day of the puzzle")
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
