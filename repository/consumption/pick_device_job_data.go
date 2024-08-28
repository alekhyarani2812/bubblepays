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
func PickDeviceData(ctx context.Context) error{
	// Define source and destination directories
	sourceDir := "/home/satyanarayana/Alekhya/gofilesdirectories/machine_data/"
	destDir := "/home/satyanarayana/Alekhya/gofilesdirectories/start_consume/"
	// Get list of files in the source directory
	files, err := filepath.Glob(filepath.Join(sourceDir, "*"))
	if err != nil {
		log.Errorft(ctx, "Error: %v ", err)
		return err
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
		err := SaveJobInfo(ctx, jobData)
		if err != nil {
			log.Errorft(ctx, "Error while saving the job data %v", err)
			return err
		}
		log.Infof("Job file saved into %v directory", destDir)
		// Generate destination path
		destPath := filepath.Join(destDir, fileName)
		// Move the file
		err1 := os.Rename(file, destPath)
		if err1 != nil {
			log.Errorft(ctx, "Error moving file: %v ", err1)
			return err1
		} else {
			log.Infof("Moved file: %v to directory %v ", fileName, destDir)
		}
	}
	return nil
}

