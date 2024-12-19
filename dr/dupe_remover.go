package dr

//I know am able to get all the files now I just need to check for duplicates
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Find_dupes(folder_path string) []string {
	files := get_files_from_dir(folder_path)
	for _, file := range files {
		fmt.Println(file)
	}
	return files
}

func get_files_from_dir(folder_path string) []string {
	files, err := os.ReadDir(folder_path)
	return_files := make([]string, 0)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		dir_path := filepath.Join(folder_path, file.Name())
		file, err := os.Open(dir_path)
		if err != nil {
			log.Fatal(err)
		}

		fileInfo, err := file.Stat()
		if fileInfo.IsDir() {

			new_files := get_files_from_dir(dir_path)
			return_files = append(return_files, new_files...)
		} else {
			return_files = append(return_files, file.Name())
		}

	}
	return return_files
}
