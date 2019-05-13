package main
import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)
type tenedor struct{ sync.Mutex }
type filosofo struct {
	id int
	tenedor_i, tenedor_d *tenedor
}
func (f filosofo) servir() {
	for j := 0; j < 3; j++ {
		f.tenedor_i.Lock()
		f.tenedor_d.Lock()

		say("comiendo", f.id)
		time.Sleep(60)

		f.tenedor_d.Unlock()
		f.tenedor_i.Unlock()

		say("terminó de comer", f.id)
		time.Sleep(60)
	}
	waitgroup.Done()
}

func say(action string, id int) {
	fmt.Printf("Filosofo #%d es %s\n", id+1, action)
}

var waitgroup sync.WaitGroup

func main() {
	
	count := 5

	forks := make([]*tenedor, count)
	for i := 0; i < count; i++ {
		forks[i] = new(tenedor)
	}

	//
	filosofos := make([]*filosofo, count)
	for i := 0; i < count; i++ {
		filosofos[i] = &filosofo{
			id: i, tenedor_i: forks[i], tenedor_d: forks[(i+1)%count]}
		waitgroup.Add(1)
		go filosofos[i].servir()
	}
	waitgroup.Wait()
}
