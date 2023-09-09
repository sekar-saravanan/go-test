package main
import (
 "fmt"
 "github.com/VictoriaMetrics/metricsql"
)
func main(){
 ql := "sum(rate(foo{bar='baz', foo='zar'}[5m])) by (job)"
 expr, err := metricsql.Parse(ql)
 if err != nil {
  panic(err)
 }

 fmt.Println(expr)
}
