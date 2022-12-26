package main

import "context"

func main() {
	// moneTrayA := NewMoneTrayActor(&Account{
	// 	id:     "montrayA",
	// 	amount: 0,
	// })
	// for i := 0; i < 1000; i++ {
	// 	go func() {
	// 		moneTrayA.TryTrans(context.TODO(), "a", "b", 100)
	// 		moneTrayA.RollTrans(context.TODO(), "b", "a", 100)
	// 		moneTrayA.TryTrans(context.TODO(), "a", "b", 100)
	// 		moneTrayA.RollTrans(context.TODO(), "b", "a", 100)
	// 	}()
	// }
	// time.Sleep(time.Second * 5)
	// go func() {
	// 	r := moneTrayA.LookupAccountAmount(context.TODO())
	// 	for _, v := range r {
	// 		fmt.Println("account:", v.id, v.amount)
	// 	}
	// }()
	// time.Sleep(time.Second * 5)
	bam := NewBankAccountManager()
	uam := NewUserAccountManager()
	mta := NewMoneTaryTransferActor(bam, uam)
	bankId := "工伤"
	ba := bam.GetAccountByBid(bankId)
	userId := "工伤-1号"
	ua := uam.GetAccountByUid(userId)
	err := mta.TryTrans(context.TODO(), ba.id, ua.id, 100)
	if err != nil {
		panic(err)
	}
}
