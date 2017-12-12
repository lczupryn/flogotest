package picamera

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"os/exec"
	//"path/filepath"
	"strconv"
	"time"
)

var log = logger.GetLogger("activity-picamera")

//constant parameters for PI Camera for raspistill
const (
	STILL      = "raspistill" //constant parameters for PI Camera for raspistill
	HFLIP      = "-hf" //constant parameters for PI Camera for raspistill
	VFLIP      = "-vf" //constant parameters for PI Camera for raspistill
	TIMEEXP    = "-t" //constant parameters for PI Camera for raspistill
	OUTFLAG    = "-o" //constant parameters for PI Camera for raspistill
	FILETYPE  = ".jpg" // filetype
	TIMESTAMP = "2006-01-02_15_04_05" // timestamp
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
	timeExposition := context.GetInput("timeExposition").(int)
	//salutation := context.GetInput("salutation").(string)

	// Use the log object to log the greeting
	log.Debugf("Pi-camera time exposition is set to [%d]", timeExposition)

	//set timestamp for file
	fileName := time.Now().Format(TIMESTAMP) + FILETYPE
	//fullPath := filepath.Join(c.savePath, fileName)
	cmd := exec.Command(STILL, TIMEEXP, strconv.Itoa(timeExposition), OUTFLAG, fileName) //fullPath)
	_, err = cmd.StdoutPipe()
	if err != nil {
		log.Debugf("%s", err)
	}
	err = cmd.Start()
	if err != nil {
		log.Debugf("%s", err)
	}
	cmd.Wait()

	// Set the file in output
	context.SetOutput("file", fileName)

	// Signal to the Flogo engine that the activity is completed

	return true, nil
}
