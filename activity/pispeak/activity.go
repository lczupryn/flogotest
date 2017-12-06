package pispeak

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os/exec"		
)

var log = logger.GetLogger("activity-pispeak")

//constant parameters for espeak library
const (
	ESPEAK      = "espeak" //constant parameters for espeak library
	OUTFLAG    = "2>/dev/null" //constant parameters for espeak library
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	espeakArgument, ok := context.GetInput("espeakArgument").(string)
	if !ok {
		log.Error("Espeak espeakArgument argument wrong type")
	}
	textToSpeak, ok := context.GetInput("textToSpeak").(string)
	if !ok {
		log.Error("Espeak textToSpeak argument wrong type")
	}
	//salutation := context.GetInput("salutation").(string)

	// Use the log object to log the greeting
	log.Debugf("Pi-speak text : [%s]", textToSpeak)
	
	status := 1
	
	cmd := exec.Command(ESPEAK, espeakArgument, textToSpeak, OUTFLAG)
	_, err = cmd.StdoutPipe()
	if err != nil {
		log.Debugf("%s", err)
		status = 0
	}
	err = cmd.Start()
	if err != nil {
		log.Debugf("%s", err)
		status = 0
	}
	cmd.Wait()

	// Set the file in output
	context.SetOutput("status", status)

	// Signal to the Flogo engine that the activity is completed

	return true, nil
}
