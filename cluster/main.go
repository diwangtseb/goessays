package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	m1 := NewMachine("1", Leader)
	m2 := NewMachine("2", Slave)
	m3 := NewMachine("3", Slave)
	cluster := NewCluster(m1, m2, m3)
	cluster.use(func(m []*Machine) {
		vote(m)
		loadbalance(m)
	})
	for _, v := range cluster.machines {
		fmt.Printf("%v", v)
	}
}

type EnumRole string

const (
	Leader EnumRole = "leader"
	Slave  EnumRole = "slave"
)

type Machine struct {
	ip      string
	role    EnumRole
	payload int
	weight  float64
}

func NewMachine(ip string, role EnumRole) *Machine {
	return &Machine{
		ip:      ip,
		role:    role,
		payload: 0,
		weight:  0,
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
	machines[0].setRole(Leader)
	maxPlayloadMachine := machines[0]
	for i := 1; i < len(machines); i++ {
		if machines[i].payload <= maxPlayloadMachine.payload {
			continue
		}
		maxPlayloadMachine.setRole(Slave)
		machines[i].setRole(Leader)
		maxPlayloadMachine = machines[i]
	}
}

func loadbalance(machines []*Machine) {
	total := 1.0
	for idx, v := range machines {
		v.weight = total - 0.25
		if idx == len(machines) {
			v.weight = total - v.weight
		}
	}
}

func (c *Cluster) trafficRequest(req interface{}) {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	r.Float64()
}
