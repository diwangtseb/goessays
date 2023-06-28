package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	m1 := NewMachine("192.168.0.1", Leader, 3)
	m2 := NewMachine("192.168.0.2", Slave, 2)
	m3 := NewMachine("192.168.0.3", Leader, 1)
	cluster := NewCluster(m1, m2, m3)
	cluster.use(func(m []*Machine) {
		vote(m)
	})
	// make a break
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("re vote !!!!")
		m1.status = DISABLE
		cluster.use(func(m []*Machine) {
			vote(m)
		})
	}()
	for i := 0; i < 100; i++ {
		go cluster.trafficRequest(struct{}{})
		if i > 50 {
			time.Sleep(time.Second * 1)
		}
	}

	// define os signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, os.Interrupt)
	<-ch
}

type EnumRole string

const (
	Leader EnumRole = "leader"
	Slave  EnumRole = "slave"
)

type EnumStatus int

const (
	ABLE    EnumStatus = 0
	DISABLE EnumStatus = -1
)

type Machine struct {
	ip            string
	role          EnumRole
	payload       int
	weight        float64
	currentWeight float64
	status        EnumStatus
}

func NewMachine(ip string, role EnumRole, weight float64) *Machine {
	return &Machine{
		ip:     ip,
		role:   role,
		weight: weight,
		status: ABLE,
	}
}

func (m *Machine) setRole(role EnumRole) {
	m.role = role
}

type Cluster struct {
	machines []*Machine
}

func NewCluster(machines ...*Machine) *Cluster {
	return &Cluster{
		machines: machines,
	}
}

type machineFunc func([]*Machine)

func (c *Cluster) use(mf ...machineFunc) {
	for _, f := range mf {
		f(c.machines)
	}
}

func vote(machines []*Machine) {
	var maxPlayloadMachine *Machine
	var disbaleCount int
	for i := 0; i < len(machines); i++ {
		if maxPlayloadMachine == nil {
			if machines[i].status == DISABLE {
				disbaleCount++
				continue
			}
			machines[i].setRole(Leader)
			maxPlayloadMachine = machines[i]
			continue
		}
		if machines[i].payload > maxPlayloadMachine.payload && machines[i].status == ABLE {
			maxPlayloadMachine.setRole(Slave)
			maxPlayloadMachine = machines[i]
			maxPlayloadMachine.setRole(Leader)
		} else {
			machines[i].setRole(Slave)
		}
	}
	if disbaleCount == len(machines) {
		panic("all machines breakdown")
	}
}

func (c *Cluster) getSelectedMachine() *Machine {
	var totalWeight float64
	var selectedMachine *Machine
	for i := 0; i < len(c.machines); i++ {
		if c.machines[i].status == DISABLE {
			continue
		}
		totalWeight += c.machines[i].weight
		c.machines[i].currentWeight += c.machines[i].weight
		if selectedMachine == nil || c.machines[i].currentWeight > selectedMachine.currentWeight {
			selectedMachine = c.machines[i]
		}
	}
	selectedMachine.currentWeight -= totalWeight
	return selectedMachine
}

func (c *Cluster) trafficRequest(req interface{}) {
	machine := c.getSelectedMachine()
	defer func() {
		fmt.Printf("current machine ip %s payload %d currentweight %f \n", machine.ip, machine.payload, machine.currentWeight)
	}()
	machine.payload++
}
