/*
** gnl.go for RSS_Analyser in RSS_Analyser_wellcut/utils
** 
** Made by King_a
** Login   <cros.arthur@epitech.eu>
** 
** Started on  Sat Jul  9 16:21:23 2016 King_a
** Last update Sat Jul  9 19:31:21 2016 King_a
*/

package	utils

import (
	"io"
	"bytes"
)

var (
	next_line string = ""
)

func	str_append(str_1, str_2 string) (string) {
	var buffer bytes.Buffer

	buffer.WriteString(str_1)
	buffer.WriteString(str_2)
	return buffer.String()
}

func	is_newline(line *string) (bool) {
	var str_line = *line

	for i, c:= range str_line {
		if str_line[i] == '\n' {
			if i + 1 < len(str_line) {
				next_line = str_line[i + 1 :]
			} else {
				next_line = ""
			}
				str_line = str_line[: i]
				return true
		}
	}
	return false
}

func (str Str_err) Error() (string) {
	return string(str)
	
}

func	Get_next_line(stream io.ReadCloser) (line string, err error) {
	var nb_read int
	tmp_line := make([]byte, TO_READ)
	err = nil

	line = next_line
	for !is_newline(&line) {
		nb_read, err = stream.Read(tmp_line)
		if (nb_read == 0 && err != io.EOF) {
			err = ERROR_READ.Error
			line = str_append(line, string(tmp_line))
			return
		}
		line = str_append(line, string(tmp_line))
	}
	return
}
