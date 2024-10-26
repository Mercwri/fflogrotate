package schedule

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func Schedule() {
	fflog_path, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	cmd_def := fmt.Sprintf(`%s\fflogrotate.exe`, fflog_path)
	task_action := fmt.Sprintf(`(New-ScheduledTaskAction -Execute '%s' -Argument rotate)`, cmd_def)
	trigger := `(New-ScheduledTaskTrigger -Daily -At '9:15 AM')`
	user, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	principal := fmt.Sprintf(`(New-ScheduledTaskPrincipal -UserId '%s' -RunLevel Highest)`, user.Username)
	settings := `(New-ScheduledTaskSettingsSet -RunOnlyIfNetworkAvailable -WakeToRun)`
	task := fmt.Sprintf(
		`New-ScheduledTask -Action %s -Principal %s -Trigger %s -Settings %s`,
		task_action,
		principal,
		trigger,
		settings,
	)
	reg_stask := fmt.Sprintf(`Register-ScheduledTask 'fflogrotate' -InputObject (%s)`, task)

	fmt.Println(reg_stask)
}
