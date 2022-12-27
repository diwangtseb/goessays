package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type Comment struct {
	Uid         string
	Cid         string
	Content     string
	ReplyToCid  string
	Interaction []*Interaction
}

type Interaction struct {
	Iid     string
	Content string
}

func (i *Interaction) String() string {
	v, _ := json.Marshal(i)
	return string(v)
}

type CommentManager interface {
	Create(ctx context.Context, commet *Comment) error
	Get(ctx context.Context, cid string) (*Comment, error)
	List(ctx context.Context) ([]*Comment, error)
	Interact(ctx context.Context, cid string, uid string, interaciton *Interaction) error
	Reply(ctx context.Context, cid string, comment *Comment) error
}

type commentManager struct {
}

func NewCommentManager() CommentManager {
	return &commentManager{}
}

// Create implements CommentManager
func (*commentManager) Create(ctx context.Context, commet *Comment) error {
	return nil
}

// Get implements CommentManager
func (*commentManager) Get(ctx context.Context, cid string) (*Comment, error) {
	return nil, nil
}

// List implements CommentManager
func (*commentManager) List(ctx context.Context) ([]*Comment, error) {
	return nil, nil
}

// Interact implements CommentManager
func (*commentManager) Interact(ctx context.Context, cid string, uid string, interaciton *Interaction) error {
	return nil
}
func (*commentManager) Reply(ctx context.Context, cid string, comment *Comment) error {
	return nil
}

var _ CommentManager = (*commentManager)(nil)

// var _ BarrageManager = (*barrageManager)(nil)

type User struct {
	Uid string
}

const EMPTY_STR = ""

func main() {
	cm := NewCommentManager()
	// create comment
	c := &Comment{
		Uid:        "1",
		Cid:        "1",
		Content:    "你真帅",
		ReplyToCid: EMPTY_STR,
	}
	err := cm.Create(context.TODO(), c)
	if err != nil {
		panic(err)
	}
	interraction := &Interaction{
		Iid:     "1",
		Content: "赞",
	}
	for i := 0; i < 10; i++ {
		if i == 0 {
			c.Interaction = append(make([]*Interaction, 0), interraction)
		} else {
			c.Interaction = append(c.Interaction, interraction)
		}
	}
	err = cm.Interact(context.TODO(), c.Cid, c.Uid, interraction)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", c)
	rc := &Comment{
		Uid:         "2",
		Cid:         "2",
		Content:     "你真不帅",
		ReplyToCid:  "1",
		Interaction: []*Interaction{},
	}
	cm.Reply(context.TODO(), c.Cid, rc)
	fmt.Printf("%+v,%+v", c, rc)
}
