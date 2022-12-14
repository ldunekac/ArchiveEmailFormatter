package main

import (
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/pkg/errors"
)

func get_new_pdf_name(path_file_name string) (string, error) {
	content, err := readPdfByRows(path_file_name) // Read local pdf file
	if err != nil {
		return "", errors.Wrap(err, "could not read "+path_file_name)
	}

	from_line := content[1]
	from_name, err := pasre_from_statment(from_line)
	if err != nil {
		return "", errors.Wrap(err, "could not parse from statement in "+path_file_name)
	}

	time_sent_line := content[2]
	datetime, datetime_err := pase_date_time(time_sent_line)
	if datetime_err != nil {
		return "", errors.Wrap(err, "could not parse datime in "+path_file_name)
	}

	subject := strings.Replace(content[5], " ", "_", -1)
	save_name := datetime + "_" + from_name + "_" + subject + ".pdf"
	return save_name, nil
}

func pasre_from_statment(line string) (string, error) {
	a := strings.Split(line, "From:")[1]
	a = strings.Replace(a, ",", "", -1)
	b := strings.Split(a, " ")
	if len(b) < 2 {
		return "", errors.New("could not find name to parse")
	}
	last_name := b[0]
	first_name := b[1]
	f_initial_last := string(first_name[0]) + last_name
	return f_initial_last, nil
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
		return "", errors.New("could not convert month: " + month + " to integer")
	}
}

func pase_date_time(line string) (string, error) {
	date := strings.Replace(line, ",", "", -1)
	date = strings.Replace(date, ":", " ", -1)
	date_list := strings.Split(date, " ")

	if len(date) < 5 {
		return "", errors.New("could not pase date line")
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
	f, r, err := pdf.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	lines := []string{}
	p := r.Page(1)
	rows, _ := p.GetTextByRow()
	for _, row := range rows {
		line := ""
		for _, word := range row.Content {
			line = line + word.S
		}
		lines = append(lines, line)
	}
	return lines, nil
}
