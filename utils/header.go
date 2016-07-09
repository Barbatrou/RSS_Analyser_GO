/*
** header.go for RSS analyser in RSS_Analyser_wellcut/utils
** 
** Made by King_a
** Login   <cros.arthur@epitech.eu>
** 
** Started on  Sat Jul  9 17:58:58 2016 King_a
** Last update Sat Jul  9 19:30:55 2016 King_a
*/

package utils

type	Str_err string

const (
	TO_READ		int = 50
	ERROR_READ	Str_err = "[ERROR]: An error occured while reading!\n"
)

//func (str Str_err) Error() (string) {
//	return string(str)
//}
