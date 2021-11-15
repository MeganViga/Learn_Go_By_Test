package PE
import "testing"
import "fmt"
type Bitcoin float64
type Wallet struct{
	amount Bitcoin
}
func TestWallet(t *testing.T){
	wallet := Wallet{}
	wallet.Deposit(11.0)
	got := wallet.Balance()
	fmt.Printf("address of balance : %v\n",&wallet.amount)
	want := Bitcoin(11.0)
	fmt.Println(wallet.amount)
	if got != want{
		t.Errorf("want %.1f, got %.1f",want,got)
	}
}