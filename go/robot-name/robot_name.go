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
type nameMap map[string]struct{}

var (
	maxNumberOfRobot = 26*26*10*10*10
	robotNameList nameMap
)

func init(){
	robotNameList = make(map[string]struct{},maxNumberOfRobot)
}

func (r *Robot) Name() (string, error) {
	if r.name != ""{
		return r.name,nil
	}
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
	name := partOne + partTwo
	if _, ok := robotNameList[name]; ok {
		return r.Name()
	}
	r.name = name
	robotNameList[r.name]= struct{}{}
	return r.name, nil
}

func (r *Robot) Reset() {
	if _,ok := robotNameList[r.name]; ok {
		delete(robotNameList,r.name)
	}
	r.name =""
}
