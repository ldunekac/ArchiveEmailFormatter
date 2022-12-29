package main

import (
	"log"
	"os"
	"time"
)

func main() {
	time_stamp := time.Now().Local().Format("20060102_150405")
	log_file := "archive_mail_formatter_" + time_stamp + ".log"

	open_log_file, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		print("error opening log file file: %s with error %v", log_file, err)
		os.Exit(1)
	}
	defer open_log_file.Close()

	log.SetOutput(open_log_file)

	args := parse_command_line_arguments()

	if args.specific_pdf_file != "" {
		process_single_pdf_file(args.in_pdf_dir, args.out_pdf_dir, args.specific_pdf_file, false)
	} else {
		process_directory(args.in_pdf_dir, args.out_pdf_dir)
	}
}
