/*
** strncmp.go for RSS analyser in /home/cros_a/Rendu/Go/src/RSS_Analyser_wellcut/utils
** 
** Made by King_a
** Login   <cros.arthur@epitech.eu>
** 
** Started on  Thu Jul  7 19:08:20 2016 King_a
** Last update Thu Jul  7 19:47:08 2016 King_a
*/

package utils


func My_strncmp(str_1, str_2 string, n int) (uint8) {
	if str_1 == "" || str_2 == "" {
		return 1
	}
	for i := 0; i < n; i += 1 {
		if (str_1[i] != str_2[i]) {
			return str_1[i] - str_2[i]
		}
	}
	return 0
}
