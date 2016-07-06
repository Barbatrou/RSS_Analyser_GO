/*
** test.go for RSS-Analyser in /home/cros_a/Rendu/Go/src/RSS_Analyser_wellcut
** 
** Made by King_a
** Login   <cros.arthur@epitech.eu>
** 
** Started on  Wed Jul  6 20:15:41 2016 King_a
** Last update Wed Jul  6 21:07:19 2016 King_a
*/

package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	tO_READ = 15
)

func main() () {
	reader := make([]byte, tO_READ)
	var nb_read int

	resp, error := http.Get("http://www.dailymotion.com/rss")
	if error != nil {
		fmt.Printf("lol error is here!\n")
		return
	}
	defer resp.Body.Close()
	nb_read, error = resp.Body.Read(reader)
	for nb_read != 0 {
		fmt.Printf("%v", string(reader))
		nb_read, error = resp.Body.Read(reader)
	}
	if error != io.EOF {
		fmt.Printf("ho lala an error occured : %v", error)
	}
	return
}
