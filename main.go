package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-9


a backup utility that is written in golang for the system managment and for the automatic backup
of the system. the time function is implemented which returns the year and then a dotenv is implemented
which allows to run this from the time of the installation. You can define the time of the days when you want
to create a backup and then it will create the backup automatically without you need to remember anything. Your
doenv file should look like this
.env
content
year = "2024-10-10"
backup = 10

you can change the backup time to any number of  days while invoking this and then it will automatically check the interval
and if the system time is not greater than the dotenv file then it will not backup the system and if the system time is
less than the backup time, then it will not backup the system.

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
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
	backupdir       string
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
	Use:  "rsyncdir",
	Long: "recursive syncing of the directories on the host system",
	Run:  rsyncHost,
}

var rsyncArchCmd = &cobra.Command{
	Use:  "rsyncArch",
	Long: "archiving of the system directories",
	Run:  rsyncArch,
}

var rsyncHostPushCmd = &cobra.Command{
	Use:  "rsyncHostpush",
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
	Long: "this is establishing a tunneling system and is equivalent to rsync -anzP",
	Run:  tunnelFunc,
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
	rsyncHostRemoteCmd.Flags().
		StringVarP(&hostpath, "path on the host", "H", "path to the host folder", "host init directory")
	rsyncHostRemoteCmd.Flags().
		StringVarP(&hostdestination, "path to the destination", "D", "destination drive", "destination directory")
	rsyncHostRemoteCmd.Flags().
		StringVarP(&backupdir, "backup dir", "B", "backup drive", "backupdrive for tunnel")
	tunnelCmd.Flags().StringVarP(&hostpath, "hostpath", "O", "path on the host", "host base drive")
	tunnelCmd.Flags().
		StringVarP(&hostdestination, "host drive", "S", "drive on the host", "host path to the drive")

	rootCmd.AddCommand(cpCmd)
	rootCmd.AddCommand(rsyncCmd)
	rootCmd.AddCommand(rsyncArchCmd)
	rootCmd.AddCommand(rsyncHostPullCmd)
	rootCmd.AddCommand(rsyncHostPushCmd)
	rootCmd.AddCommand(rsyncHostRemoteCmd)
	rootCmd.AddCommand(tunnelCmd)
}

/*
added several new functions such as fileTime check and others and now it is calling
the dotenv and the systemTime and the dotsystemTime function in each of the same and
if the difference of the date is not greater than 10 then it will not do the back up
otherwise will do the backup and automatic backup and you dont have to do anything else.
*/
func dotEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	installYear := os.Getenv("Year")

	return installYear
}

func systemBack(cmd *cobra.Command, args []string) {
	yearDot, monthDot, dateDot := dotsystemTime()
	yearTag, monthTag, dateTag := systemTime()
	if (yearDot == yearTag) && (monthDot == monthTag) && (dateTag-dateDot > 10) {
		cpout, err := exec.Command("cp", hostpath, hostdestination).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(cpout))
	}
}

func rsyncHost(cmd *cobra.Command, args []string) {
	yearDot, monthDot, dateDot := dotsystemTime()
	yearTag, monthTag, dateTag := systemTime()
	if (yearDot == yearTag) && (monthDot == monthTag) && (dateTag-dateDot > 10) {
		rsyncerr, err := exec.Command("rsync", "-r", hostpath, hostdestination).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(rsyncerr))
	}
}

func rsyncArch(cmd *cobra.Command, args []string) {
	yearDot, monthDot, dateDot := dotsystemTime()
	yearTag, monthTag, dateTag := systemTime()
	if (yearDot == yearTag) && (monthDot == monthTag) && (dateTag-dateDot > 10) {
		rsyncArchA, err := exec.Command("rsync", "-az", hostpath, hostdestination).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(rsyncArchA))
	}
}

func hostPush(cmd *cobra.Command, args []string) {
	yearDot, monthDot, dateDot := dotsystemTime()
	yearTag, monthTag, dateTag := systemTime()
	if (yearDot == yearTag) && (monthDot == monthTag) && (dateTag-dateDot > 10) {
		hostpushAdd, err := exec.Command("rsync", "-a", hostpath, foriegnaddress, ":", foriegnpath).
			Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(hostpushAdd))
	}
}

func hostPull(cmd *cobra.Command, args []string) {
	yearDot, monthDot, dateDot := dotsystemTime()
	yearTag, monthTag, dateTag := systemTime()
	if (yearDot == yearTag) && (monthDot == monthTag) && (dateTag-dateDot > 10) {
		hostpullAdd, err := exec.Command("rsync", "-a", foriegnaddress, ":", foriegnpath, hostdestination).
			Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(hostpullAdd))
	}
}

func rsyncEnd(cmd *cobra.Command, args []string) {
	yearDot, monthDot, dateDot := dotsystemTime()
	yearTag, monthTag, dateTag := systemTime()
	if (yearDot == yearTag) && (monthDot == monthTag) && (dateTag-dateDot > 10) {
		rsyncend, err := exec.Command("rsync", "-a", "--delete", "--backup-dir=", backupdir, hostpath, hostdestination).
			Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(rsyncend))
	}
}

func tunnelFunc(cmd *cobra.Command, args []string) {
	yearDot, monthDot, dateDot := dotsystemTime()
	yearTag, monthTag, dateTag := systemTime()
	if (yearDot == yearTag) && (monthDot == monthTag) && (dateTag-dateDot > 10) {
		tunnelAdd, err := exec.Command("rsync", "-anzP", "--delete", hostpath, hostdestination).
			Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tunnelAdd)
	}
}

func timeCheck() []string {
	time := time.Now()
	writetime := time.String()
	file, err := os.Create("currenttimefile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(writetime + "\n")

	year := []string{}
	date := []string{}
	yearext := []string{}
	monthext := []string{}
	dateext := []string{}

	fOpen, err := os.Open("currenttimefile.txt")
	if err != nil {
		log.Fatal(err)
	}

	fRead := bufio.NewScanner(fOpen)

	for fRead.Scan() {
		line := fRead.Text()
		year = append(year, strings.Split(string(line), " ")[0])
		date = append(date, strings.Split(string(line), " ")[1])
		yearext = append(
			yearext,
			strings.Split(string(strings.Split(string(line), " ")[0]), "-")[0],
		)
		monthext = append(
			monthext,
			strings.Split(string(strings.Split(string(line), " ")[0]), "-")[1],
		)
		dateext = append(
			dateext,
			strings.Split(string(strings.Split(string(line), " ")[0]), "-")[2],
		)
	}

	return year
}

func dotsystemTime() (int, int, int) {
	dotYear := dotEnv()
	var dateD int
	var monthD int
	var yearD int
	dotStore := strings.Split(dotYear, "-")
	for i := range dotStore {
		dateD, _ = strconv.Atoi(string(dotStore[i][0]))
		monthD, _ = strconv.Atoi(string(dotStore[i][1]))
		yearD, _ = strconv.Atoi(string(dotStore[i][2]))
	}
	return dateD, monthD, yearD
}

func systemTime() (int, int, int) {
	dotStore := timeCheck()
	var date int
	var month int
	var year int
	for i := range dotStore {
		dotStore := strings.Split(dotStore[i], "-")
		date, _ = strconv.Atoi(dotStore[0])
		month, _ = strconv.Atoi(dotStore[1])
		year, _ = strconv.Atoi(dotStore[2])
	}
	return date, month, year
}
