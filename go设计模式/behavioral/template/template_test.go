package template

import "testing"

//模版模式
func TestWorker_Daily(t *testing.T) {

	worker:=NewWork(&Corder{})
	worker.Daily()

}
