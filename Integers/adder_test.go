package integers
import "fmt"
import "testing"

func TestAdder(t *testing.T){
	expected := Add(2,2)
	want := 4
	if expected != want {
		t.Errorf("expected '%d' but got '%d'", want,expected)
	}
}
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}