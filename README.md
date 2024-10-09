# go-system-back

- time based system backup application in golang. 
- native datastructures and structs implementation. 
- checks the current time with the previous time and execute the backup from the host to the system backup. 
- currenty implementing rsync, rclone, dd, cp and then a host to the system remote terminal execution. 
- implemented the go env files so that you dont have to remember, just execute and it will check how old the back up is and will run the execution. 

```
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » \
go run main.go
system backup application and uses cp, dd and rsync for the system back from local to the remote application development

Usage:
  option [command]

Available Commands:
  completion    Generate the autocompletion script for the specified shell
  help          Help about any command
  rsyncArch
  rsyncBack
  rsyncHostpull
  ryncdir
  rynscHostpush
  system
  tunnel

Flags:
  -h, --help   help for option

Use "option [command] --help" for more information about a command.
exit status 1
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » \
go run main.go rsyncArch -h                                                                     
archiving of the system directories

Usage:
  option rsyncArch [flags]

Flags:
  -d, --destination path string          system init destination path (default "input the destination path")
  -h, --help                             help for rsyncArch
  -p, --path to the file/folder string   system init path (default "path to the file/folder which needs to be backed up")
exit status 1
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » \
go run main.go rsyncBack -h                                                                     
end to end syncing of the rsync between host and remote

Usage:
  option rsyncBack [flags]

Flags:
  -B, --backup dir string                backupdrive for tunnel (default "backup drive")
  -h, --help                             help for rsyncBack
  -H, --path on the host string          host init directory (default "path to the host folder")
  -D, --path to the destination string   destination directory (default "destination drive")
exit status 1
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » \
go run main.go rsyncHostpull -h                                                                 
syncing of the transfer file from the remote to the host

Usage:
  option rsyncHostpull [flags]

Flags:
  -A, --address of the server string     ip route (default "server address")
  -e, --anyfiles to exclude string       exclude (default "exclusive")
  -h, --help                             help for rsyncHostpull
  -H, --hostpath string                  host path (default "path on the host")
  -I, --include these files string       include (default "inclusive")
  -D, --path on the destination string   path on the remote (default "destination path")
exit status 1
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » \
go run main.go system -h                                                                       
This is the system back up configuration and it uses three main type, dd, cp and rsync

Usage:
  option system back [flags]

Flags:
  -d, --destination path string          system init destination path (default "input the destination path")
  -h, --help                             help for system
  -p, --path to the file/folder string   system init path (default "path to the file/folder which needs to be backed up")
exit status 1
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » \
go run main.go tunnel -h                                                                  
this is establishing a tunneling system and is equivalent to rsync -anzP

Usage:
  option tunnel [flags]

Flags:
  -h, --help                help for tunnel
  -S, --host drive string   host path to the drive (default "drive on the host")
  -O, --hostpath string     host base drive (default "path on the host")
exit status 1
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » \
go run main.go ryncdir  -h                                                                      
recursive syncing of the directories on the host system

Usage:
  option ryncdir [flags]

Flags:
  -d, --destination path string          system init destination path (default "input the destination path")
  -h, --help                             help for ryncdir
  -p, --path to the file/folder string   system init path (default "path to the file/folder which needs to be backed up")
exit status 1
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-system-back ±main⚡ » \
go run main.go rsyncHostpush -h
syncing the transfer files from the host to remote

Usage:
  option rsyncHostpush [flags]

Flags:
  -A, --address of the server string     ip route (default "server address")
  -e, --anyfiles to exclude string       exclude (default "exclusive")
  -h, --help                             help for rsyncHostpush
  -H, --hostpath string                  host path (default "path on the host")
  -I, --include these files string       include (default "inclusive")
  -D, --path on the destination string   define path (default "destination path")
exit status 1

```
Gaurav Sablok
