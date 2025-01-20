package dr

//I know am able to get all the files now I just need to check for duplicates
import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/exp/slices"
)

func Find_dupes(folder_path string) []string {
	files := get_files_from_dir(folder_path)
	// for _, file := range files {
	// 	fmt.Println(file)
	// }
	dupes := check_for_dupes(files)
	fmt.Println(dupes)
	return_string := fmt.Sprintf("Size of Dupes: %s", int(get_size_of_files(dupes)))
	fmt.Println(return_string)
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

func check_for_dupes(files []string) []string {
	same_files := make([]string, 0)

	for _, file_name := range files {
		file, err := os.Open(file_name)
		if err != nil {
			log.Fatal(err)
		}
		current_file_info, err := file.Stat()
		for _, check_file_name := range files {
			check_file, err := os.Open(check_file_name)
			if err != nil {
				log.Fatal(err)
			}
			check_file_info, err := check_file.Stat()
			if os.SameFile(current_file_info, check_file_info) && !slices.Contains(same_files, file_name) {
				same_files = append(same_files, file_name)
			}
		}
	}

	return same_files
}

func get_size_of_files(files []string) int64 {
	var files_size int64
	for _, file_name := range files {
		file, err := os.Open(file_name)
		if err != nil {
			log.Fatal(err)
		}
		file_info, err := file.Stat()
		files_size += file_info.Size()
		if err != nil {
			log.Fatal(err)
		}
	}
	return files_size
}
