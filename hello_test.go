package main
import "testing"

func TestHello(t *testing.T){

	assertionCorrectMessage := func(t testing.TB,got, want string){
		t.Helper()
		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people",func(t *testing.T){
		got :=Hello("Chris")
		want :="Hello,Chris"
		assertionCorrectMessage(t,got, want)
	})
	t.Run("saying 'Hello,World' to people",func(t *testing.T){
		got :=Hello("")
		want :="Hello,World"
		assertionCorrectMessage(t,got, want)
	})
}