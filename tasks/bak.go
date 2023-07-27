package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/maja42/goval"
// )

// func init() {
// 	var task = &Task{
// 		ID: 1,
// 		Rule: &Rule{
// 			ID:      1,
// 			Content: "event_acc >= 100",
// 		},
// 		Times:  3,
// 		ReTime: time.Second * 5,
// 	}
// 	userTask = make(map[int]*Task)
// 	userTask[1] = task
// }

// var userTask map[int]*Task

// func main() {
// 	PushEvent(&Event{
// 		Uid:      1,
// 		Key:      "event_acc",
// 		Val:      300,
// 		Classify: 1,
// 	})
// 	PushEvent(&Event{
// 		Uid:      1,
// 		Key:      "event_acc",
// 		Val:      20,
// 		Classify: 1,
// 	})
// 	PushEvent(&Event{
// 		Uid:      1,
// 		Key:      "event_acc",
// 		Val:      30,
// 		Classify: 1,
// 	})
// 	PushEvent(&Event{
// 		Uid:      1,
// 		Key:      "event_acc",
// 		Val:      40,
// 		Classify: 1,
// 	})
// 	for k, v := range globalUserTaskRecord {
// 		for _, v1 := range v {
// 			fmt.Println(k, v1)
// 		}
// 	}
// }

// func getUserTask(uid int) *Task {
// 	return userTask[uid]
// }

// func PushEvent(e *Event) {
// 	task := getUserTask(e.Uid)
// 	switch e.Classify {
// 	case 1:
// 		records := getUserRecord(e.Uid)
// 		userTimes := 0
// 		latestVal := 0
// 		for _, v := range records {
// 			if v.Status {
// 				userTimes++
// 			} else {
// 				latestVal = v.TaskVal
// 			}
// 		}
// 		if userTimes >= task.Times {
// 			return
// 		}
// 		e.Val += latestVal
// 	}
// 	eval := goval.NewEvaluator()
// 	m := map[string]interface{}{
// 		e.Key: e.Val,
// 	}
// 	complete := e.Val / 100
// 	for i := 0; i < complete; i++ {
// 		result, err := eval.Evaluate(task.Rule.Content, m, nil)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if v, ok := result.(bool); ok && v {
// 			recordUser(1, task, e, true)
// 		} else {
// 			recordUser(1, task, e, false)
// 		}
// 		val := (m[e.Key].(int))
// 		m[e.Key] = val - 100
// 		e.Val -= 100
// 	}
// }

// func getUserRecord(uid int) []*UserTaskRecord {
// 	return globalUserTaskRecord[uid]
// }

// var globalUserTaskRecord map[int][]*UserTaskRecord

// func recordUser(uid int, t *Task, e *Event, status bool) {
// 	if globalUserTaskRecord == nil {
// 		globalUserTaskRecord = make(map[int][]*UserTaskRecord)
// 	}
// 	globalUserTaskRecord[uid] = append(globalUserTaskRecord[uid], &UserTaskRecord{
// 		Uid:     uid,
// 		TaskId:  t.ID,
// 		TaskVal: e.Val,
// 		Status:  status,
// 	})

// }

// type Task struct {
// 	ID     int
// 	Rule   *Rule
// 	Times  int
// 	ReTime time.Duration
// }

// type Rule struct {
// 	ID      int
// 	Content string
// }

// type Event struct {
// 	Uid      int
// 	Key      string
// 	Val      int
// 	Classify int
// }

// type UserTaskRecord struct {
// 	Uid     int
// 	TaskId  int
// 	TaskVal int
// 	Status  bool
// }
