/*
** rss.go for RSS analyser in RSS_Analyser_wellcut/RSS
** 
** Made by King_a
** Login   <cros.arthur@epitech.eu>
** 
** Started on  Thu Jul  7 20:05:49 2016 King_a
** Last update Thu Jul  7 20:56:03 2016 King_a
*/

package RSS

import (
	"fmt"
	"net/http"
)

func	Parse(url string) (T_rss, error) {
	var rss T_rss
	reader := make([]byte, 15)
	resp, err := http.Get(url)
	if err != nil {
		return rss, err
	}
	resp.Body.Read(reader)
	fmt.Println(string(reader))
	return rss, nil
}
