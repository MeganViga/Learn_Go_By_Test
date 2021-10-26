package SMI
import "math"
func (r rectangle) Perimeter()float64{
	return 2*(r.width + r.height)
}
func (r rectangle) Area()float64{
	return r.width * r.height
}
func (c circle) Perimeter()float64{
	return 2 * math.Pi *c.radius
}
func (c circle) Area()float64{
	return math.Pi *(c.radius * c.radius)
}