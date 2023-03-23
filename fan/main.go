package main

func main() {

}

type Msg struct {
	req interface{}
	rsp interface{}
}

var proch chan *Msg

func produer() {
	for i := 0; i < 100; i++ {
		proch <- &Msg{
			req: i,
			rsp: nil,
		}
	}
}

func distribute() {
	// req := <-proch
}

func worker(req <-chan interface{}) {

}
