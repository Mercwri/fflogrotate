# fflogrotate

A simple tool for rotating logfiles produced by [ACT](https://advancedcombattracker.com/home.php) or [IINACT](https://www.iinact.com/). Built this in a few hours after finding out just how bloated the log folder had grown and how janky it was to get Logrotate on windows.

```
PS C:\Users\Mercwri\fflogrotate> .\fflogrotate.exe --help
Rotate, Archive, Delete, and Backup Logfiles from ACT and IINACT

Usage:
  fflogrotate.exe [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  rotate      compress and delete logs and archives
  schedule    create a task to run on a schedule
  version     Print the version number of generated code example

Flags:
  -h, --help   help for fflogrotate.exe
  ```

Rotate handles the lifecycle of your ffxiv log files, archives after a short-term, and deletes archives after a longterm.

```
PS C:\Users\Mercwri\fflogrotate> .\fflogrotate.exe rotate --help
Rotates short-term logs into archives and deletes long-lived archives

Usage:
  fflogrotate.exe rotate [flags]

Flags:
  -h, --help             help for rotate
      --long-term int    time in hours before deleting an archive (default 120)
      --short-term int   time in hours before archiving (default 24)
```
Schedule is a WIP to scaffold a powershell command to generate a scheduled task
```
PS C:\Users\Mercwri\fflogrotate> .\fflogrotate.exe schedule --help
Rotates short-term logs into archives and deletes long-lived archives

Usage:
  fflogrotate.exe schedule [flags]

Flags:
  -h, --help   help for schedule
PS C:\Users\Mercwri\fflogrotate> .\fflogrotate.exe schedule
Register-ScheduledTask 'fflogrotate' -InputObject (New-ScheduledTask -Action (New-ScheduledTaskAction -Execute 'C:\Users\Mercwri\fflogrotate\fflogrotate.exe' -Argument rotate) -Principal (New-ScheduledTaskPrincipal -UserId 'GAMING\Mercwri' -RunLevel Highest) -Trigger (New-ScheduledTaskTrigger -Daily -At '9:15 AM') -Settings (New-ScheduledTaskSettingsSet -RunOnlyIfNetworkAvailable -WakeToRun))
```