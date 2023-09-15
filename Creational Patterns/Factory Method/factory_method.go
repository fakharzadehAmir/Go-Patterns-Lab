package Factory_Method

// IVehicle is the main interface representing all types of vehicles.
type IVehicle interface {
	StartEngine()
	StopEngine()
	GetBrand() string
	GetModel() string
	GetMaxSpeed() float64
}

// Car represents a type of vehicle.
type Car struct {
	Brand         string
	Model         string
	MaxSpeed      float64
	EngineRunning bool
}

// NewCar creates a new Car instance.
func NewCar(brand string, model string, maxSpeed float64) *Car {
	return &Car{
		Brand:    brand,
		Model:    model,
		MaxSpeed: maxSpeed,
	}
}

// GetBrand returns the brand of the car.
func (c *Car) GetBrand() string {
	return c.Brand
}

// GetModel returns the model of the car.
func (c *Car) GetModel() string {
	return c.Model
}

// GetMaxSpeed returns the maximum speed of the car.
func (c *Car) GetMaxSpeed() float64 {
	return c.MaxSpeed
}

// StartEngine starts the car's engine.
func (c *Car) StartEngine() {
	c.EngineRunning = true
}

// StopEngine stops the car's engine.
func (c *Car) StopEngine() {
	c.EngineRunning = false
}

// Motorcycle represents a type of vehicle.
type Motorcycle struct {
	Brand         string
	Model         string
	MaxSpeed      float64
	EngineRunning bool
}

// NewMotorcycle creates a new Motorcycle instance.
func NewMotorcycle(brand string, model string, maxSpeed float64) *Motorcycle {
	return &Motorcycle{
		Brand:    brand,
		Model:    model,
		MaxSpeed: maxSpeed,
	}
}

// GetBrand returns the brand of the motorcycle.
func (m *Motorcycle) GetBrand() string {
	return m.Brand
}

// GetModel returns the model of the motorcycle.
func (m *Motorcycle) GetModel() string {
	return m.Model
}

// GetMaxSpeed returns the maximum speed of the motorcycle.
func (m *Motorcycle) GetMaxSpeed() float64 {
	return m.MaxSpeed
}

// StartEngine starts the motorcycle's engine.
func (m *Motorcycle) StartEngine() {
	m.EngineRunning = true
}

// StopEngine stops the motorcycle's engine.
func (m *Motorcycle) StopEngine() {
	m.EngineRunning = false
}

// Truck represents a type of vehicle.
type Truck struct {
	Brand         string
	Model         string
	MaxSpeed      float64
	EngineRunning bool
}

// NewTruck creates a new Truck instance.
func NewTruck(brand string, model string, maxSpeed float64) *Truck {
	return &Truck{
		Brand:    brand,
		Model:    model,
		MaxSpeed: maxSpeed,
	}
}

// GetBrand returns the brand of the truck.
func (t *Truck) GetBrand() string {
	return t.Brand
}

// GetModel returns the model of the truck.
func (t *Truck) GetModel() string {
	return t.Model
}

// GetMaxSpeed returns the maximum speed of the truck.
func (t *Truck) GetMaxSpeed() float64 {
	return t.MaxSpeed
}

// StartEngine starts the truck's engine.
func (t *Truck) StartEngine() {
	t.EngineRunning = true
}

// StopEngine stops the truck's engine.
func (t *Truck) StopEngine() {
	t.EngineRunning = false
}

// IVehicleFactory is the main interface representing concrete factory for creating all types of vehicles.
type IVehicleFactory interface {
	CreateVehicle(brand string, model string, maxSpeed float64) IVehicle
}

// CarFactory is a concrete factory for creating cars.
type CarFactory struct{}

// CreateVehicle creates a Car instance with the provided attributes.
func (cf *CarFactory) CreateVehicle(brand string, model string,
	maxSpeed float64) IVehicle {
	return NewCar(brand, model, maxSpeed)
}

// TruckFactory is a concrete factory for creating trucks.
type TruckFactory struct{}

// CreateVehicle creates a Truck instance with the provided attributes.
func (tf *TruckFactory) CreateVehicle(brand string, model string,
	maxSpeed float64) IVehicle {
	return NewTruck(brand, model, maxSpeed)
}

// MotorcycleFactory is a concrete factory for creating motorcycles.
type MotorcycleFactory struct{}

// CreateVehicle creates a Motorcycle instance with the provided attributes.
func (mf *MotorcycleFactory) CreateVehicle(brand string, model string,
	maxSpeed float64) IVehicle {
	return NewMotorcycle(brand, model, maxSpeed)
}
