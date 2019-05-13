package main
import (
	"fmt"
	"sync"
	"time"
)
type tenedor struct{ sync.Mutex }
type filosofo struct {
	id int
	tenedor_i, tenedor_d *tenedor
}
func (f filosofo) servir() {
	for j := 0; j < 3; j++ {
		time.Sleep(time.Second)

		f.tenedor_i.Lock()
		f.tenedor_d.Lock()

		say("comenzando a comer", f.id)
		time.Sleep(time.Second)

		say("terminando de comer", f.id)
		f.tenedor_d.Unlock()
		f.tenedor_i.Unlock()
		
	}
	waitgroup.Done()
}

func say(action string, id int) {
	fmt.Printf("%s #%d\n",action,id+1)
}

var waitgroup sync.WaitGroup

func main() {
	
	count := 5

	tenedores := make([]*tenedor, count)
	for i := 0; i < count; i++ {
		tenedores[i] = new(tenedor)
	}
	// 
	filosofos := make([]*filosofo, count)
	for i := 0; i < count; i++ {
		filosofos[i] = &filosofo{
			id: i, tenedor_i: tenedores[i], tenedor_d: tenedores[(i+1)%count]}
		waitgroup.Add(1)
		go filosofos[i].servir()
	}
	waitgroup.Wait()
}