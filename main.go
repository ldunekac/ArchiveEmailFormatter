package main

import (
	"path/filepath"
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

func main() {
	args := parse_command_line_arguments()

	if args.specific_pdf_file != "" {
		process_single_pdf_file(args.in_pdf_dir, args.out_pdf_dir, args.specific_pdf_file)
	} else {
		// rename all pdf files
		print("Still need to implement")
	}
}
