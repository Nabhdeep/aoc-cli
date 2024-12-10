/*
Copyright © 2024 Nabhdeep
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type puzzleInfo struct {
	year string
	day  string
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the day's solution",
	Long: `Run the solution for the specified day. Usage:
	advent-cli run --year yyyy --day dd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		_y := args[0]
		_d, _ := cmd.Flags().GetString("d")
		_puzzleInfo := puzzleInfo{year: _y, day: _d}
		checkFile(_puzzleInfo)

	},
}

func checkFile(_dayInfo puzzleInfo) {
	_path := fmt.Sprintf(_dayInfo.year + "/day" + _dayInfo.day + "/day" + _dayInfo.day + ".go")
	packagePath := fmt.Sprintf("advent/%s/day%s", _dayInfo.year, _dayInfo.day)
	found, err := Findfile(_path)
	if err != nil {
		fmt.Printf("Err %v \n", err)
	}
	if !found {
		fmt.Printf("File not found: %s \n", _path)
		return
	}
	fileContent := fmt.Sprintf(`
		package main

		import (
			"%s"
		)

		func main() {
			day%s.Solve()
		}
	`, packagePath, _dayInfo.day)
	// make main.file
	f, e := os.Create("main.go")
	if e != nil {
		fmt.Printf("Error main file: %v\n", e)
	}
	f.WriteString(fileContent)
	f.Close()

	// Run the Go file
	cmd := exec.Command("go", "run", "main.go")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running file: %v\n", err)
		return
	}

	fmt.Println(string(output))

}

func init() {
	runCmd.Flags().String("d", "", "Day you want to run")
	rootCmd.AddCommand(runCmd)
}
