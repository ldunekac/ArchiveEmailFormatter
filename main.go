package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func process_single_pdf_file(in_dir string, out_dir string, file_name string) {
	source := filepath.Join(in_dir, file_name)
	new_name, err := get_new_pdf_name(source)
	if err != nil {
		panic(err)
	}
	destination := filepath.Join(out_dir, new_name)
	if err := copy(source, destination); err != nil {
		panic(err)
	}
}

func process_directory(in_dir string, out_dir string) {
	files, err := ioutil.ReadDir(in_dir)
	if err != nil {
		print(err)
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			split_file_name := strings.Split(file.Name(), ".")
			file_type := split_file_name[len(split_file_name)-1]
			if file_type == "pdf" {
				process_single_pdf_file(in_dir, out_dir, file.Name())
			}
		}
	}
}

func main() {
	args := parse_command_line_arguments()

	if args.specific_pdf_file != "" {
		process_single_pdf_file(args.in_pdf_dir, args.out_pdf_dir, args.specific_pdf_file)
	} else {
		process_directory(args.in_pdf_dir, args.out_pdf_dir)
	}
}
