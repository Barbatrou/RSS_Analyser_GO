/*
** analyse.go for RSS analyser in RSS_Analyser_wellcut/analyse
** 
** Made by King_a
** Login   <cros.arthur@epitech.eu>
** 
** Started on  Thu Jul  7 16:57:13 2016 King_a
** Last update Thu Jul  7 20:57:53 2016 King_a
*/

package main

import (
	"os"
	"fmt"
	"RSS_Analyser_wellcut/command"
	"RSS_Analyser_wellcut/RSS"
)

func main() () {
	var rss RSS.T_rss
	c_line, err := command.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c_line.Url)
	rss, err = RSS.Parse(c_line.Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rss.Title)
	return
}
