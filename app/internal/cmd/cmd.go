package cmd

import (
	"fmt"
	"log"

	"net/http"

	"github.com/kh411d/goshrpac/pkg/sqx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goshrpac",
	Short: "goshrpac",
	Long:  `goshrpac`,
}

func init() {
	cobra.OnInitialize(viper.AutomaticEnv)
}

// New Initialize registered cli commands
func New() *cobra.Command {

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		log.Println("hello world")
	}

	cmdHTTPServer()

	return rootCmd
}

func cmdHTTPServer() {

	c := &cobra.Command{
		Use:   "httpd",
		Short: "Run a http server",
		Long:  `Run a http server for REST API`,
		Run: func(cmd *cobra.Command, args []string) {
			http.HandleFunc("/hello", hello)
			http.ListenAndServe(":3000", nil)
		},
	}

	rootCmd.AddCommand(c)
}

func hello(w http.ResponseWriter, req *http.Request) {
	x := sqx.Eq{"hello": "world"}
	a, b, c := x.ToSql()
	y := sqx.NewEq()

	s := fmt.Sprintf("%v %v %v \n%v", a, b, c, y)

	w.Write([]byte(s))

}
