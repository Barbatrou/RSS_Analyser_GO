/*
** command.go for RSS analyser in RSS_Analyser_wellcut/command
** 
** Made by King_a
** Login   <cros.arthur@epitech.eu>
** 
** Started on  Thu Jul  7 17:17:33 2016 King_a
** Last update Thu Jul  7 19:46:55 2016 King_a
*/

package command

import (
	"RSS_Analyser_wellcut/utils"
)

func init_c_line(c_line *T_cline) () {
	c_line.Url = DFLT_ADDR
}

func Parse(arg []string) (c_line T_cline, err error) {
	init_c_line(&c_line)
	err = nil
	for i:= 0; i < len(arg); i+= 1 {
		if arg[i] == "-u"{
			if (i == len(arg)) {
				err = &c_line
				return
			}
			i += 1
			c_line.Url = arg[i]
		} else if utils.My_strncmp(arg[i], "--URL=", 5) == 0 {
			if len(arg[i]) > 6 {
				c_line.Url = arg[i][6:]
			} else {
				err = &c_line
				return
			}
		} else {
			err = &c_line
			return
		}
	}
	return
}
