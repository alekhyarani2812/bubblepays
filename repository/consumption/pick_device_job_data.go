package consumption

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
)

// PickDeviceData Move Consumption data from one directory to another
func PickDeviceData(ctx context.Context) {

	// Define source and destination directories
	sourceDir := "/home/satyanarayana/Alekhya/gofilesdirectories/machine_data/"
	destDir := "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"
	// Get list of files in the source directory
	files, err := filepath.Glob(filepath.Join(sourceDir, "*"))
	if err != nil {
		log.Errorft(ctx, "Error: %v ", err)
		return
	}
	dateval := time.Now()
	// Move each file to the destination directory
	for _, file := range files {
		// Extract filename from path
		_, fileName := filepath.Split(file)
		jobData := &models.Jobs{
			Filename:   fileName,
			Date:       &dateval,
			JobMessage: "Pending",
			Status:     utils.Inprogress,
		}
		SaveJobInfo(ctx, jobData)
		// Generate destination path
		destPath := filepath.Join(destDir, fileName)
		// Move the file
		err := os.Rename(file, destPath)
		if err != nil {
			log.Infof("Error moving file: %v ", err)
		} else {
			log.Infof("Moved file: , to  %v %v ", file, destPath)
		}
	}
}
