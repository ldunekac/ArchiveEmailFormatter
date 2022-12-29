package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/pkg/errors"
)

func process_single_pdf_file(in_dir string, out_dir string, file_name string, create_sub_directory bool) {
	source := filepath.Join(in_dir, file_name)
	new_name, err := get_new_pdf_name(source)
	if err != nil {
		err = errors.Wrap(err, "Error: Converting "+source+" failed. Could not create new pdf name.\n")
		log.Print(err)
		return
	}

	destination := out_dir
	if create_sub_directory {
		destination = filepath.Join(destination, new_name)
		if err = os.Mkdir(destination, 0755); err != nil {
			log.Print("Error: Converting " + source + " failed. Could not make directory " + destination)
			return
		}
	}
	destination = filepath.Join(destination, new_name+".pdf")

	if err := copy(source, destination); err != nil {
		log.Print("Error: Converting " + source + " failed. Could not copy file from " + source + " to " + destination)
		log.Print(err)
		return
	}
	log.Print("Successfully processed: " + source)
}

func process_directory(in_dir string, out_dir string) {
	files, err := ioutil.ReadDir(in_dir)
	if err != nil {
		log.Print("Error: Could not open the input directory")
		log.Fatal(err)
		return
	}
	for _, file := range files {
		if !file.IsDir() {
			split_file_name := strings.Split(file.Name(), ".")
			file_type := split_file_name[len(split_file_name)-1]
			if file_type == "pdf" {
				process_single_pdf_file(in_dir, out_dir, file.Name(), true)
			}
		}
	}
}

func get_new_pdf_name(path_file_name string) (string, error) {
	content, err := readPdfByRows(path_file_name) // Read local pdf file
	if err != nil {
		return "", errors.Wrap(err, "Could not read "+path_file_name+"\n")
	} else if len(content) < 5 {
		return "", errors.New("Error: Could not read content in file " + path_file_name)
	}

	from_line := content[1]
	from_name, err := pasre_from_statment(from_line)
	if err != nil {
		return "", errors.Wrap(err, "Could not parse from statement in "+path_file_name+"\n")
	}

	time_sent_line := content[2]
	datetime, datetime_err := pase_date_time(time_sent_line)
	if datetime_err != nil {
		return "", errors.Wrap(err, "Could not parse datime in "+path_file_name+"\n")
	}

	subject := strings.Replace(content[5], " ", "_", -1)
	save_name := datetime + "_" + from_name + "_" + subject
	return save_name, nil
}

func pasre_from_statment(line string) (string, error) {
	name_line := strings.Split(line, "From:")[1]
	name_line = strings.Replace(name_line, ",", "", -1)
	names := strings.Split(name_line, " ")
	if len(names) < 2 {
		return "", errors.New("Could not find name to parse")
	}
	last_name := names[0]
	first_name := names[1]
	first_initial_last_name := string(first_name[0]) + last_name
	return first_initial_last_name, nil
}

func month_to_int(month string) (string, error) {
	switch strings.TrimSpace(month) {
	case "January":
		return "01", nil
	case "February":
		return "02", nil
	case "March":
		return "03", nil
	case "April":
		return "04", nil
	case "May":
		return "05", nil
	case "June":
		return "06", nil
	case "July":
		return "07", nil
	case "August":
		return "08", nil
	case "September":
		return "09", nil
	case "October":
		return "10", nil
	case "November":
		return "11", nil
	case "December":
		return "12", nil
	default:
		return "", errors.New("Error: Could not convert month: " + month + " to integer")
	}
}

func pase_date_time(line string) (string, error) {
	date := strings.Replace(line, ",", "", -1)
	date = strings.Replace(date, ":", " ", -1)
	date_list := strings.Split(date, " ")

	if len(date) < 5 {
		return "", errors.New("Could not pase date line")
	}
	month, month_err := month_to_int(date_list[1])
	if month_err != nil {
		return "", month_err
	}
	day := date_list[2]
	year := date_list[3]
	hour := date_list[4]
	minute := date_list[5]
	date_format := year + month + day + "_" + hour + minute
	return date_format, nil
}

func readPdfByRows(path string) ([]string, error) {
	pdf_file, pdf_reader, err := pdf.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer pdf_file.Close()
	lines := []string{}
	page1 := pdf_reader.Page(1)
	rows, _ := page1.GetTextByRow()
	for _, row := range rows {
		line := ""
		for _, word := range row.Content {
			line = line + word.S
		}
		lines = append(lines, line)
	}
	return lines, nil
}
