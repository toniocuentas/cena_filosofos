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

		say("comiendo", f.id)
		time.Sleep(time.Second)

		f.tenedor_d.Unlock()
		f.tenedor_i.Unlock()
		say("terminó de comer", f.id)
	}
	waitgroup.Done()
}

func say(action string, id int) {
	fmt.Printf("Filosofo #%d %s\n", id+1, action)
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
