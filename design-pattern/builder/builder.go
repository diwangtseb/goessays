package builder

type BuildProcesser interface {
	SetWheel() BuildProcesser
	SetSeats() BuildProcesser
}

type ManufacturingDirector struct {
	buildProcesser BuildProcesser
}

func (md *ManufacturingDirector) SetBuilder(builder BuildProcesser) {
	md.buildProcesser = builder
}

func (md *ManufacturingDirector) Construct() {
	md.buildProcesser.SetSeats().SetWheel()
}

type VehicleProduct struct {
	Wheels int
	Seats  int
}

type CarBuild struct {
	VehicleProduct
}

// SetSeats implements BuildProcesser.
func (c *CarBuild) SetSeats() BuildProcesser {
	c.Seats = 1
	return c
}

// SetWheel implements BuildProcesser.
func (c *CarBuild) SetWheel() BuildProcesser {
	c.Wheels = 1
	return c
}

func (c *CarBuild) GetVehicleProduct() VehicleProduct {
	return c.VehicleProduct
}

var _ (BuildProcesser) = (*CarBuild)(nil)
