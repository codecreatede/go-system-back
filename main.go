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
	}
	os.Exit(1)
}

var (
	hostpath        string
	hostdestination string
	exclude         string
	include         string
	foriegnaddress  string
	foriegnpath     string
	deletedir       bool
	backupdir       bool
)

var rootCmd = &cobra.Command{
	Use:  "option",
	Long: "system backup application and uses cp, dd and rsync for the system back from local to the remote application development",
}

var cpCmd = &cobra.Command{
	Use:  "system back",
	Long: "This is the system back up configuration and it uses three main type, dd, cp and rsync",
	Run:  systemBack,
}

var rsyncCmd = &cobra.Command{
	Use:  "ryncdir",
	Long: "recursive syncing of the directories on the host system",
	Run:  rsyncHost,
}

var rsyncArchCmd = &cobra.Command{
	Use:  "rsyncArch",
	Long: "archiving of the system directories",
	Run:  rsyncArch,
}

var rsyncHostPushCmd = &cobra.Command{
	Use:  "rynscHostpull",
	Long: "syncing the transfer files from the host to remote",
	Run:  hostPush,
}

var rsyncHostPullCmd = &cobra.Command{
	Use:  "rsyncHostpull",
	Long: "syncing of the transfer file from the remote to the host",
	Run:  hostPull,
}

var rsyncHostRemoteCmd = &cobra.Command{
	Use:  "rsyncBack",
	Long: "end to end syncing of the rsync between host and remote",
	Run:  rsyncEnd,
}

var tunnelCmd = &cobra.Command{
	Use:  "tunnel",
	Long: "this is establishing a tunneling system and is equivalent to rsync -azP",
	Run:  tunnelFunc,
}

var tunnelDCmd = &cobra.Command{
	Use:  "tunnelD",
	Long: "establishing a delete tunnel and this is rsync -an --delete",
	Run:  tunnelDFunc,
}

func init() {
	cpCmd.Flags().
		StringVarP(&hostpath, "path to the file/folder", "p", "path to the file/folder which needs to be backed up", "system init path")
	cpCmd.Flags().
		StringVarP(&hostdestination, "destination path", "d", "input the destination path", "system init destination path")
	rsyncCmd.Flags().
		StringVarP(&hostpath, "path to the file/folder", "p", "path to the file/folder which needs to be backed up", "system init path")
	rsyncCmd.Flags().
		StringVarP(&hostdestination, "destination path", "d", "input the destination path", "system init destination path")
	rsyncArchCmd.Flags().
		StringVarP(&hostpath, "path to the file/folder", "p", "path to the file/folder which needs to be backed up", "system init path")
	rsyncArchCmd.Flags().
		StringVarP(&hostdestination, "destination path", "d", "input the destination path", "system init destination path")
	rsyncHostPushCmd.Flags().StringVarP(&hostpath, "hostpath", "H", "path on the host", "host path")
	rsyncHostPushCmd.Flags().
		StringVarP(&foriegnpath, "path on the destination", "D", "destination path", "define path")
	rsyncHostPushCmd.Flags().
		StringVarP(&exclude, "anyfiles to exclude", "e", "exclusive", "exclude")
	rsyncHostPushCmd.Flags().
		StringVarP(&include, "include these files", "I", "inclusive", "include")
	rsyncHostPushCmd.Flags().
		StringVarP(&foriegnaddress, "address of the server", "A", "server address", "ip route")
	rsyncHostPullCmd.Flags().StringVarP(&hostpath, "hostpath", "H", "path on the host", "host path")
	rsyncHostPullCmd.Flags().
		StringVarP(&foriegnpath, "path on the destination", "D", "destination path", "path on the remote")
	rsyncHostPullCmd.Flags().
		StringVarP(&exclude, "anyfiles to exclude", "e", "exclusive", "exclude")
	rsyncHostPullCmd.Flags().
		StringVarP(&include, "include these files", "I", "inclusive", "include")
	rsyncHostPullCmd.Flags().
		StringVarP(&foriegnaddress, "address of the server", "A", "server address", "ip route")

	rootCmd.AddCommand(cpCmd)
	rootCmd.AddCommand(rsyncCmd)
	rootCmd.AddCommand(rsyncArchCmd)
	rootCmd.AddCommand(rsyncHostPullCmd)
	rootCmd.AddCommand(rsyncHostPushCmd)
	rootCmd.AddCommand(rsyncHostRemoteCmd)
	rootCmd.AddCommand(tunnelCmd)
	rootCmd.AddCommand(tunnelDCmd)
}

/*
   the time function and the embedded struct and the iter will be called in each function and will evaluate
	 whether tot ake back up or not and this is reading directly from the remote server or the host server.

*/

func systemBack(cmd *cobra.Command, args []string) {
	cpout, err := exec.Command("cp", hostpath, hostdestination).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cpout))
}

func rsyncHost(cmd *cobra.Command, args []string) {
	rsyncerr, err := exec.Command("rsync", "-r", hostpath, hostdestination).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rsyncerr))
}

func rsyncArch(cmd *cobra.Command, args []string) {
	rsyncArchA, err := exec.Command("rsync", "-az", hostpath, hostdestination).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rsyncArchA))
}

func hostPush(cmd *cobra.Command, args []string) {
	hostpushAdd, err := exec.Command("rsync", "-a", hostpath, foriegnaddress, ":", foriegnpath).
		Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(hostpushAdd))
}

func hostPull(cmd *cobra.Command, args []string) {
	hostpullAdd, err := exec.Command("rsync", "-a", foriegnaddress, ":", foriegnpath, hostdestination).
		Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(hostpullAdd))
}

func timeCheck(cmd *cobra.Command, args []string) {
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
