package PE
import "fmt"

func (w *Wallet) Deposit(addcash Bitcoin){

	w.amount = w.amount + addcash
	fmt.Printf("Rs.%.1f cash Deposited\n",addcash)
	fmt.Printf("address of balance : %v\n",&w.amount)

}

func (w  *Wallet) Balance()Bitcoin{
	fmt.Printf("address of balance : %v\n",&w.amount)
	return w.amount
}