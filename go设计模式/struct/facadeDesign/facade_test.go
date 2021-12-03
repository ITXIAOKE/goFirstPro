package facadeDesign

import "testing"
//外观模式（门面设计模式）
func TestCarFacade_CreateCompleteCar(t *testing.T) {
	facade := NewCarFacade()
	facade.CreateCompleteCar()
}
