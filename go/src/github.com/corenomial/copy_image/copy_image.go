package main

import (
	"fmt"
	"github.com/corenomial/script_lib/ssh_connection_to_device"
	"log"
)

func main() {
	var work string
	var ip string

	fmt.Printf("enter workspace path (default=/work/v3): ")
	_, _ = fmt.Scanf("%s", &work)
	if work == "" {
		// my current turret address
		work = "/work/v3"
	}

	fmt.Printf("enter ip address (default=10.204.44.22): ")
	_, _ = fmt.Scanf("%s", &ip)
	if ip == "" {
		// my current turret address
		ip = "10.204.44.22"
	}

	var output1 = "sshpass -p \"apollor00t\" ssh root@" + ip + " mount -o rw,remount /rwboot"
	cmd := exec.Command(output1, "")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	fmt.Println("done rwboot")

	var output2 = "sshpass -p \"apollor00t\" scp " + work + "/apollo/target/mercury/images/mercury_fs0.img root@" + ip + ":/rwboot/images"
	cmd = exec.Command(output2, "")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	fmt.Println("done image")
}
