package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	alp = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
// nameMap will store list of robot names
type nameMap map[string]struct{}

var (
	maxNumOfRobot = 26*26*10*10*10 //  26 char (a-z) and 10 char (0-1)
	robotNames    nameMap
)


// Define the Robot type here.
type Robot struct {
	name string

}

func init(){
	// Init the robot names with
	robotNames = make(map[string]struct{}, maxNumOfRobot)
}

func (r *Robot) Name() (string, error) {
	// If name already present
	if r.name != ""{
		return r.name,nil
	}
	// If no space to generate new robot names
	if len(robotNames) == maxNumOfRobot {
		return "", fmt.Errorf("space constraint, no more robots can be created")
	}

	// Generate random string with fixed length (Ex: BB234, XY123)
	b := make([]byte, 2)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range b {
		b[i] = alp[rand.Intn(len(b))]
	}
	name := string(b) + fmt.Sprintf("%d",100 + rand.Intn(899))

	// Validate whether the name is  present in name list map
	if _, ok := robotNames[name]; ok {
		// If name is already present, create new
		return r.Name()
	}
	// Update robot name
	r.name = name
	// Update robot names map
	robotNames[name]= struct{}{}

	return name, nil
}

func (r *Robot) Reset() {
	// Validate if the robot name is present
	if _,ok := robotNames[r.name]; ok {
		// If present, remove from the map
		delete(robotNames,r.name)
	}
	// reset the name to empty
	r.name =""
}
