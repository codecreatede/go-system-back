# go-system-back

- time based system backup application in golang. 
- native datastructures and structs implementation. 
- checks the current time with the previous time and execute the backup from the host to the system backup. 
- currenty implementing rsync, rclone, dd, cp and then a host to the system remote terminal execution. 

```
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » go run main.go
system date and time stage has been moved %s []
2024-10-09
 11:58:23.442250399
2024-10-09
 11:58:23.442250399
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » ls -la
total 20
drwxr-xr-x. 1 gauavsablok gauavsablok   96 Oct  9 11:58 .
drwxr-xr-x. 1 gauavsablok gauavsablok 1726 Oct  8 23:28 ..
drwxr-xr-x. 1 gauavsablok gauavsablok  164 Oct  9 11:57 .git
-rw-r--r--. 1 gauavsablok gauavsablok  200 Oct  9 10:34 go.mod
-rw-r--r--. 1 gauavsablok gauavsablok  896 Oct  9 10:34 go.sum
-rw-r--r--. 1 gauavsablok gauavsablok 2752 Oct  9 11:56 main.go
-rw-r--r--. 1 gauavsablok gauavsablok   17 Oct  9 11:57 README.md
-rw-r--r--. 1 gauavsablok gauavsablok   56 Oct  9 11:58 timeprevious.txt
```
Gaurav Sablok
