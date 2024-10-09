package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-9


a backup utility that is written in golang for the system managment and for the automatic backup
of the system.

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var (
	path        string
	destination string
)

var rootCmd = &cobra.Command{
	Use:  "system back",
	Long: "This is the system back up configuration and it uses three main type, dd, cp and rsync",
	Run:  systemBack,
}

func init() {
	rootCmd.Flags().
		StringVarP(&path, "path to the local file", "p", "path to the file which needs to be backed up", "system init path")
	rootCmd.Flags().
		StringVarP(&destination, "destination path", "d", "input the destination path", "system init destination path")
}

func systemBack(cmd *cobra.Command, args []string) {
	time := time.Now()
	writetime := time.String()
	file, err := os.Create("currenttimefile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(writetime + "\n")

	move, err := exec.Command("mv", "currenttimefile.txt", "timeprevious.txt").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("system date and time stage has been moved %s", move)

	type timeNow struct {
		yearNow     string
		dateNow     string
		yearextNow  string
		monthextNow string
		datextNow   string
	}

	currentNow := []timeNow{}
	currentNow = append(currentNow, timeNow{
		yearNow:     strings.Split(writetime, " ")[0],
		dateNow:     strings.Split(writetime, " ")[1],
		yearextNow:  strings.Split(string(strings.Split(writetime, " ")[0]), "-")[0],
		monthextNow: strings.Split(string(strings.Split(writetime, " ")[0]), "-")[1],
		datextNow:   strings.Split(string(strings.Split(writetime, " ")[0]), "-")[2],
	})

	type timeDate struct {
		year     string
		date     string
		yearext  string
		monthext string
		dateext  string
	}

	timeStore := []timeDate{}
	fOpen, err := os.Open("timeprevious.txt")
	if err != nil {
		log.Fatal(err)
	}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		timeStore = append(timeStore, timeDate{
			year:     strings.Split(string(line), " ")[0],
			date:     strings.Split(string(line), " ")[1],
			yearext:  strings.Split(string(strings.Split(string(line), " ")[0]), "-")[0],
			monthext: strings.Split(string(strings.Split(string(line), " ")[0]), "-")[1],
			dateext:  strings.Split(string(strings.Split(string(line), " ")[0]), "-")[2],
		})
	}

	for i := range timeStore {
		fmt.Println(timeStore[i].year, "\n", timeStore[i].date)
	}

	for i := range currentNow {
		fmt.Println(currentNow[i].yearNow, "\n", currentNow[i].dateNow)
	}
}
