// main.go
package main

import (
	"github.com/corenomial/script_lib/automationd"
	"github.com/corenomial/script_lib/logger"
	"github.com/corenomial/script_lib/turret_connection"
	x "github.com/corenomial/script_lib/xml_creator"
)

func main() {
	/***************************/
	/* start syslog equivalent */
	/***************************/
	mylogger := logger.LoggerFactory("test_log")
	err := mylogger.Start()
	if err != nil {
		panic("could not start the logger")
	}
	defer mylogger.Stop()
	/******************************/
	/* create the message to send */
	/******************************/
	msg := &x.XML_Message{}
	msg.
		/*** jump from left to right to left and then press release ***/
		/*Header().
		EnableAutomation(x.Attr_ea_autoFormat()+x.Attr_ea_cpus("0")).
		PressButton(x.CPU_PTU, x.B12          ,x.TenthSec(2)).Pause(x.TenthSec(2)).
		PressButton(x.CPU_PTU, x.B11          ,x.TenthSec(2)).Pause(x.TenthSec(2)).
		PressButton(x.CPU_PTU, x.RIGHT_RELEASE,x.TenthSec(5)).Pause(x.TenthSec(2)).
		DisableAutomation().
		ReturnQueuedMessages().
		Footer()*/

		/*** make a call from my left to right ***/
		Header().
		EnableAutomation(x.Attr_ea_autoFormat()+x.Attr_ea_cpus("0 1")).
		PressButton(x.CPU_PTU, x.MENU, x.TenthSec(2)).Pause(x.TenthSec(8)).
		PressButton(x.CPU_PTU, x.LEFT_RELEASE, x.TenthSec(2)).Pause(x.TenthSec(8)).
		PressButton(x.CPU_PTU, x.RIGHT_RELEASE, x.TenthSec(2)).Pause(x.TenthSec(8)).
		PressButton(x.CPU_BU1_TOP, x.FAVORITES, x.TenthSec(2)).Pause(x.TenthSec(8)).
		PressButton(x.CPU_BU1_TOP, x.B1, x.TenthSec(2)).Pause(x.TenthSec(8)).
		PressButton(x.CPU_PTU, x.DIGIT_3, x.TenthSec(2)).Pause(x.TenthSec(3)).
		PressButton(x.CPU_PTU, x.DIGIT_0, x.TenthSec(2)).Pause(x.TenthSec(3)).
		PressButton(x.CPU_PTU, x.DIGIT_0, x.TenthSec(2)).Pause(x.TenthSec(3)).
		PressButton(x.CPU_PTU, x.DIGIT_2, x.TenthSec(2)).Pause(x.TenthSec(3)).
		PressButton(x.CPU_PTU, x.DIGIT_POUND, x.TenthSec(2)).Pause(x.TenthSec(3)).
		Pause(2000).
		PressButton(x.CPU_PTU, x.B12, x.TenthSec(2)).Pause(x.TenthSec(8)).
		PressButton(x.CPU_BU1_TOP, x.B2, x.TenthSec(2)).Pause(x.TenthSec(8)).
		Pause(10000).
		PressButton(x.CPU_PTU, x.LEFT_RELEASE, x.TenthSec(2)).Pause(x.TenthSec(8)).
		DisableAutomation().
		ReturnQueuedMessages().
		Footer()

	/*** put turret on page 10, these are all OBD's ***/
	/*Header().
	EnableAutomation(x.Attr_ea_autoFormat()+x.Attr_ea_cpus("1 2")+x.Attr_ea_appTypes("directory")).
	PressButton(x.CPU_PTU    , x.MENU         ,x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_PTU    , x.LEFT_RELEASE ,x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_PTU    , x.RIGHT_RELEASE,x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.FAVORITES    ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B2           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B3           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B4           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B5           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B6           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B7           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B8           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B9           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B10          ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B11          ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_TOP, x.B12          ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_BOT, x.B1           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_BOT, x.B2           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	PressButton(x.CPU_BU1_BOT, x.B3           ,  x.TenthSec(2)).Pause(x.TenthSec(2)).
	Pause(x.TenthSec(10)).
	PressButton(x.CPU_PTU    , x.LEFT_RELEASE ,  x.TenthSec(2)).Pause(  x.TenthSec(8)).
	DisableAutomation().
	ReturnQueuedMessages().
	Footer()*/

	/************************/
	/* create sample turret */
	/************************/
	creds := turret_connection.PasswordBasedCredentialsFactory("ipctech", "a123456")
	tconn := turret_connection.ConnectionFactory("10.204.45.168", "myV3Turret", creds)
	tconn.Start()
	defer tconn.Stop()
	/********************************/
	/* and execute the script on it */
	/********************************/
	tconn.TCP.SendCommand(automationd.ConvertString(msg.String()), true)
}
