package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	alp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var (
	maxNumberOfRobot = 26*26*10*10*10
	robotNameList = make(map[string]struct{}, maxNumberOfRobot)
)

func (r *Robot) Name() (string, error) {
	if len(robotNameList) == maxNumberOfRobot {
		return "", fmt.Errorf("space constraint, no more robots can be created")
	}
	b := make([]byte, 2)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range b {
		b[i] = alp[rand.Intn(len(b))]
	}
	partOne := string(b)
	partTwo := fmt.Sprintf("%d",100 + rand.Intn(899))
	r.name = partOne + partTwo
	if _, ok := robotNameList[r.name]; ok {
		r.Name()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	if _,ok := robotNameList[r.name]; ok {
		delete(robotNameList,r.name)
	}
}
