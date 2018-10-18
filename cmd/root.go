package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/iralution/tmignore/filescanner"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmignore",
	Short: "lol",
	Long:  `still lol`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if DebugFlag {
			log.SetLevel(log.DebugLevel)
		}
		/*
		- read file list starting at --rootDir (make it unique list)
		- if not --test run tmutil addexclusion for each path (check if excluded)
 		- write changefile (--- at the end for each change)
 		- restore command
  		*/
		log.WithFields(log.Fields{"args": args}).Debug("Entered arguments for command")
		run(args[0])
	},
}

var (
	RootDirFlag string
	TestFlag    bool
	DebugFlag   bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&RootDirFlag, "rootDir", "r", ".", "Define the path at which to start searching for given directory pattern. Default is current directory")
	rootCmd.PersistentFlags().BoolVarP(&TestFlag, "test", "t", false, "If set, it only prints the matching filepaths without setting them to timemachine ignore")
	rootCmd.PersistentFlags().BoolVarP(&DebugFlag, "debug", "d", false, "If set, it only prints the matching filepaths without setting them to timemachine ignore")
}

func run(pattern string) {
	log.WithFields(log.Fields{"pattern": pattern}).Debug("Running root command with pattern")

	fileList, _ := filescanner.FilePathWalkDir(RootDirFlag, pattern)

	if len(fileList) <= 0 {
		log.WithField("pattern", pattern).Warn("No directories found for pattern")
	}

	fmt.Print("\n")
	for file := range fileList {
		if TestFlag {
			fmt.Println(file)
		} else {
			isIncluded := strings.Contains(checkException(file), "[Included]")
			fmt.Println(isIncluded, file)
			if isIncluded {
				addException(file)
			}
		}
	}
}

func checkException(path string) string {
	tmutilCmd := exec.Command("tmutil", "isexcluded", path)

	output, err := tmutilCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.WithField("output", output).Debug("tmutil output")
	return string(output)
}

func addException(path string) {
	tmutilCmd := exec.Command("tmutil", "addexclusion", path)

	_, err := tmutilCmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
