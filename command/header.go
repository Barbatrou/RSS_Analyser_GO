/*
** header.go for RSS analyser in RSS_Analyser_wellcut/command
** 
** Made by King_a
** Login   <cros.arthur@epitech.eu>
** 
** Started on  Thu Jul  7 17:24:29 2016 King_a
** Last update Thu Jul  7 19:52:27 2016 King_a
*/

package command


const (
	eRROR_MSG	string = "[ERROR]: invalid command line!\n"
	eRROR_USG	string = "Usage: ./analyse [-u url] [--URL=url]"
	DFLT_ADDR	string = "http://www.dailymotion.com/rss"
)

/*
## ################# ##
## T_cline STRUCTURE ##
## ################# ##
*/

/*
** command line structure
** and methods
*/

type	T_cline struct {
	Url	string
	//Cache	bool
}

func (c_line *T_cline) Error() string {
	return eRROR_MSG + eRROR_USG
}
