package main

import (
	"flag"
	"os"
	"path/filepath"
)

type CommandLineArgs struct {
	in_pdf_dir        string
	out_pdf_dir       string
	specific_pdf_file string
}

func parse_command_line_arguments() CommandLineArgs {
	var in_pdf_dir = flag.String("in", "", "Specifies the directory that houses the pdfs that need processing.")
	var out_pdf_dir = flag.String("out", "", "Specifies the directory to save the renamed pdf files.")
	var specific_pdf_file = flag.String("pdf", "", "Specify a specific pdf to convert. Specifying this will save the newly named pdf in the same location as the origional pdf.")
	flag.Parse()

	if *specific_pdf_file != "" {
		// If a specific pdf file is entered,
		// the in/out dirs must not be specified
		if *out_pdf_dir != "" || *in_pdf_dir != "" {
			panic("cannot specify 'pdf' with 'in' or 'out' at the same time")
		}
		if _, err := os.Stat(*specific_pdf_file); err != nil {
			panic(*specific_pdf_file + " does not exist")
		}

		directory, err := filepath.Abs(*specific_pdf_file)
		if err != nil {
			panic("cold not get absolute path of " + *specific_pdf_file)
		}
		return CommandLineArgs{
			filepath.Dir(directory),
			filepath.Dir(directory),
			filepath.Base(*specific_pdf_file),
		}

	} else {
		// 'in' and 'out' dirs must be specified
		// since the specific pdf file was not specified
		if *in_pdf_dir == "" && *out_pdf_dir != "" {
			panic("The command line argument 'in' was specified but 'out' was not. Both need to be specified.")
		}
		if *in_pdf_dir != "" && *out_pdf_dir == "" {
			panic("The command line argument 'out' was specified but 'in' was not. Both need to be specified.")
		}
		if _, err := os.Stat(*in_pdf_dir); err != nil {
			panic(*in_pdf_dir + " does not exist")
		}
		if _, err := os.Stat(*out_pdf_dir); err != nil {
			panic(*out_pdf_dir + " does not exist")
		}
		return CommandLineArgs{
			*in_pdf_dir,
			*out_pdf_dir,
			*specific_pdf_file,
		}
	}
}
