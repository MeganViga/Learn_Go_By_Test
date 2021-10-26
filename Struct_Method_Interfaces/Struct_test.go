package SMI
import "testing"
import "math"

type rectangle struct{
	width float64;
	height float64

}
type circle struct{
	radius float64
}
type Shape interface{
	Area() float64
	Perimeter() float64
}

/*var checkArea := func(t *testing.T,shape Shape,want float64){
	got := shape.Area()
	if got != want {
		t.Errorf("want %.1f,got %.1f",want,got)
	}
}*/

func TestPerimeter(t *testing.T){
	checkPerimeter := func(t testing.TB,shape Shape,want float64){
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("want %.1f,got %.1f",want,got)
		}
	}

	t.Run("RectanglePerimeter",func(t *testing.T){
		Rectangle := rectangle{10.0,10.0}
		//got := Rectangle.Perimeter()
		want := 40.0
		checkPerimeter(t,Rectangle,want)
		/*if got != want {
			t.Errorf("want %.1f,got %.1f",want,got)
		}*/

	})
	t.Run("CirclePerimeter",func(t *testing.T){
		Circle := circle{5.0}
		//got := Circle.Perimeter()
		want := 2 * math.Pi * 5.0
		checkPerimeter(t,Circle,want)
		/* got != want {
			t.Errorf("want %.1f,got %.1f",want,got)
		}*/

	})

	
}

func TestArea(t *testing.T){
	checkArea := func(t testing.TB,shape Shape,want float64){
		got := shape.Area()
		if got != want {
			t.Errorf("want %.1f,got %.1f",want,got)
		}
	}
	t.Run("RectangleArea",func(t *testing.T){
		Rectangle := rectangle{12.0,6.0}
		//got := Rectangle.Area()
		want := 72.0
		checkArea(t,Rectangle,want)
		/*if got != want {
			t.Errorf("want %.1f,got %.1f",want,got)
		}*/

	})
	t.Run("CircleArea",func(t *testing.T){
		Circle := circle{5.0}
		//got := Circle.Area()
		want := math.Pi *(5.0 * 5.0)
		checkArea(t,Circle,want)
		/*if got != want {
			t.Errorf("want %.1f,got %.1f",want,got)
		}*/

		
	})
}