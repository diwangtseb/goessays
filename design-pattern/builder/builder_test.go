package builder

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	md := ManufacturingDirector{}
	carBuild := &CarBuild{}
	md.SetBuilder(carBuild)
	md.Construct()
	vp := carBuild.GetVehicleProduct()
	fmt.Println(vp)
}
