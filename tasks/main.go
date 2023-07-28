package main

import (
	"fmt"
	"time"

	"github.com/maja42/goval"
)

type Task struct {
	ID         int
	Name       string
	Rule       *Rule
	St         time.Time
	Et         time.Time
	RtClassify RefreshClassify
	SRt        time.Time
	NRt        time.Time
	FixedTime  time.Time
	Week       int
	Day        int
	Hour       int
	Minute     int
	Sec        int
	Times      int
}

func (t *Task) initSRTNRT() {
	t.SRt = t.St
	switch t.RtClassify {
	case FixedPoint:
		t.NRt = t.FixedTime
	case Cyclicity:
		now := time.Now()
		daysUntilNextday := t.Day - int(now.Weekday()) + 7
		nextday := now.AddDate(0, 0, daysUntilNextday)
		next := time.Date(nextday.Year(), nextday.Month(), nextday.Day(), t.Hour, t.Minute, t.Sec, 0, nextday.Location())
		t.NRt = next
	default:
		t.NRt = time.Date(2099, 99, 0, 0, 0, 0, 0, time.Local)
	}
}

func (t *Task) RefreshSRTNRT() {
	t.SRt = t.NRt
	switch t.RtClassify {
	case FixedPoint:
		t.NRt = t.Et
	case Cyclicity:
		now := time.Now()
		daysUntilNextday := t.Day - int(now.Weekday()) + 7
		nextday := now.AddDate(0, 0, daysUntilNextday)
		next := time.Date(nextday.Year(), nextday.Month(), nextday.Day(), t.Hour, t.Minute, t.Sec, 0, nextday.Location())
		t.NRt = next
	default:
		t.NRt = time.Date(2099, 99, 0, 0, 0, 0, 0, time.Local)
	}
}

func NewTask() *Task {
	task := &Task{
		ID:         1,
		Name:       "累计事件大于等于100",
		Rule:       &Rule{ID: 1, Content: "event_acc >= 100"},
		St:         time.Date(2023, 7, 27, 2, 0, 0, 0, time.Local),
		Et:         time.Date(2023, 7, 29, 2, 0, 0, 0, time.Local),
		RtClassify: Cyclicity,
		SRt:        time.Date(2023, 7, 27, 2, 0, 0, 0, time.Local),
		NRt:        time.Date(2023, 7, 27, 19, 0, 0, 0, time.Local),
		FixedTime:  time.Time{},
		Week:       1,
		Day:        1,
		Times:      3,
	}
	task.initSRTNRT()
	task.SRt = time.Date(2023, 7, 27, 1, 0, 0, 0, time.Local)
	task.NRt = time.Date(2023, 7, 27, 2, 0, 0, 0, time.Local)
	return task
}

type Rule struct {
	ID      int
	Content string
}

type RefreshClassify int

const (
	FixedPoint RefreshClassify = 1
	Cyclicity  RefreshClassify = 2
)

type UserEvent struct {
	UID int
	Key string
	Val interface{}
}

func main() {
	// newTask := NewTask()
	// fmt.Println(newTask)
	// create a user
	ue := &UserEvent{
		UID: 1,
		Key: "event_acc",
		Val: 120,
	}
	doTask(ue)
	ue2 := &UserEvent{
		UID: 1,
		Key: "event_acc",
		Val: 20,
	}
	doTask(ue2)
	ue3 := &UserEvent{
		UID: 1,
		Key: "event_acc",
		Val: 20,
	}
	doTask(ue3)
	ue4 := &UserEvent{
		UID: 1,
		Key: "event_acc",
		Val: 20,
	}
	doTask(ue4)

	ue5 := &UserEvent{
		UID: 1,
		Key: "event_acc",
		Val: 40,
	}
	doTask(ue5)
	ue6 := &UserEvent{
		UID: 1,
		Key: "event_acc",
		Val: 40,
	}
	doTask(ue6)
	for k, v := range userTaskRecordMap {
		for _, v1 := range v {
			fmt.Println(k, v1)
		}
	}
}

func doTask(ue *UserEvent) {
	saveUserEventRecord(ue.UID, &UserEventRecord{
		UserEvent: ue,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	})
	task := getTaskByEventKey(ue.Key)
	if time.Since(task.St) <= 0 || time.Since(task.Et) >= 0 {
		return
	}

	if task.NRt.Sub(task.SRt) <= 0 {
		panic("Refresh Time Err")
	}

	if time.Since(task.NRt) > 0 {
		task.RefreshSRTNRT()
	}

	// get user records > rt
	userTaskRecord := userTaskRecordMap[ue.UID]
	if len(userTaskRecord) == 0 {
		eval := goval.NewEvaluator()
		m := map[string]interface{}{
			ue.Key: ue.Val,
		}
		result, err := eval.Evaluate(task.Rule.Content, m, nil)
		if err != nil {
			return
		}
		if v, ok := result.(bool); ok {
			status := false
			if v {
				status = true
			}
			saveUserTaskRecord(ue.UID, &UserTaskRecord{
				TaskId:   task.ID,
				TaskVal:  ue.Val.(int),
				Status:   status,
				CreateAt: time.Now(),
				UpdateAt: time.Now(),
			})
		}
	} else {
		userTaskRecord = filter(userTaskRecord, func(i int) bool {
			return userTaskRecord[i].CreateAt.Sub(task.SRt) > 0
		})
		times := 0
		borrow := 0
		for _, v := range userTaskRecord {
			if v.Status {
				times++
				borrow += (v.TaskVal - 100)
				v.TaskVal -= borrow
			}
		}
		ueValInt := ue.Val.(int)
		ueValInt += borrow
		if times >= task.Times {
			return
		}

		userTaskRecord = filter(userTaskRecord, func(i int) bool {
			return !userTaskRecord[i].Status
		})

		// userEventRecord = filter(userEventRecord, func(i int) bool {
		// 	return userEventRecord[i].CreateAt.Sub(task.SRt) > 0
		// })
		if len(userTaskRecord) == 0 {
			eval := goval.NewEvaluator()
			m := map[string]interface{}{
				ue.Key: ueValInt,
			}
			result, err := eval.Evaluate(task.Rule.Content, m, nil)
			if err != nil {
				return
			}
			if v, ok := result.(bool); ok {
				status := false
				if v {
					status = true
				}
				saveUserTaskRecord(ue.UID, &UserTaskRecord{
					TaskId:   ue.UID,
					TaskVal:  ueValInt,
					Status:   status,
					CreateAt: time.Now(),
					UpdateAt: time.Now(),
				})
			}
			return
		}
		latest := userTaskRecord[len(userTaskRecord)-1]
		eval := goval.NewEvaluator()
		total := ueValInt + latest.TaskVal
		m := map[string]interface{}{
			ue.Key: total,
		}
		result, err := eval.Evaluate(task.Rule.Content, m, nil)
		if err != nil {
			return
		}
		if v, ok := result.(bool); ok {
			status := false
			if v {
				status = true
			}
			latest.Status = status
			latest.TaskVal = total
		}
	}
}

func filter[T any](utrs []*T, f func(i int) bool) (data []*T) {
	for i, utr := range utrs {
		if f(i) {
			data = append(data, utr)
		}
	}
	return
}

type UserTaskRecord struct {
	TaskId   int
	TaskVal  int
	Status   bool
	CreateAt time.Time
	UpdateAt time.Time
}

var userTaskRecordMap = make(map[int][]*UserTaskRecord)

func saveUserTaskRecord(uid int, utr *UserTaskRecord) {
	userTaskRecordMap[uid] = append(userTaskRecordMap[uid], utr)
}

type UserEventRecord struct {
	UserEvent *UserEvent
	CreateAt  time.Time
	UpdateAt  time.Time
}

var userEventRecordMap = make(map[int][]*UserEventRecord)

func saveUserEventRecord(uid int, uer *UserEventRecord) {
	userEventRecordMap[uid] = append(userEventRecordMap[uid], uer)
}

func getTaskByEventKey(eventKey string) *Task {
	switch eventKey {
	case "event_acc":
		return NewTask()
	default:
		return nil
	}
}
