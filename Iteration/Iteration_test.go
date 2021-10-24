package iteration
import "testing"
import "fmt"

func TestIteration(t *testing.T){
	got := Iterate("a",5)
	want := "aaaaa"
	if got != want {
		t.Errorf("Expected %q , got %q ",want, got)
	}
}
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Iterate("a",4)
    }
}

func ExampleTest(){
	got := Iterate("a",5)
	fmt.Println(got)
}