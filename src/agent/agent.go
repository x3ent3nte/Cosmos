package agent

import (
	"vec"
	"fmt"
	"math"
	"encoding/json"
	"sync"
	"strings"
)

type Agent struct {
	sync.RWMutex
	odin *Odin
	Type string `json:"type"`
	Id int64 `json:"id"`

	Pos vec.Vec3 `json:"pos"`
	Target vec.Vec3 `json:"target"`

	Forward vec.Vec3 `json:"forward"`
	Up vec.Vec3 `json:"up"`
	Right vec.Vec3 `json:"Right"`

	velocity vec.Vec3
	angular_velocity vec.Vec3
	Euler vec.Vec3 `json:"euler"`

	mass float64
	radius float64

	alive bool
	age int64
	lifespan int64

	rocket Rocket
}

func (agent *Agent) Orientate() {
	agent.Forward, agent.Up, agent.Right = vec.FURFromPYR(agent.Euler)
}

func (agent *Agent) findClosestPlant(ents []Entity) Entity {
	var closest_dist float64 = math.MaxFloat64
	var closest Entity = nil

	for _, ent := range ents {
		if agent != ent {
			if strings.Compare(ent.GetType(), "plant") == 0 {
				dist := vec.Vec3DistanceBetween(agent.GetPos(), ent.GetPos())
				if dist < closest_dist {
					dist = closest_dist
					closest = ent
				}
			}
		}
	}
	return closest
}

func (agent *Agent) calculateMovement(time_delta float64) {
	agent.turn(time_delta)
	agent.thrustForward(time_delta)
	agent.stabilize(time_delta)

	if vec.Vec3DistanceBetween(agent.Target, agent.Pos) < 2000 {
		agent.Target = vec.Vec3Add(agent.Pos, vec.Vec3Random(35000))
		//agent.Target = vec.Vec3Random(35000)
	}
}

func (agent *Agent) turn(time_delta float64) {
	var course = vec.Vec3Sub(agent.Target, agent.Pos)
	var course_normal = vec.Vec3Normal(course)

	var angle_diff = vec.Vec3AngleBetween(agent.Forward, course_normal)
	if angle_diff > math.Pi {
		angle_diff = math.Pi
	}
	var axis = vec.Vec3Cross(agent.Forward, course_normal)

	var delta_turn = time_delta
	agent.Forward = vec.AxisAngleRotation(agent.Forward, delta_turn, axis)
}

func (agent *Agent) thrustForward(time_delta float64) {
	impulse := agent.rocket.Thrust(vec.Vec3Normal(agent.Forward), 0.4, time_delta)
	agent.applyImpulse(impulse)
}

func (agent *Agent) stabilize(time_delta float64) {
	var velocity_normal = vec.Vec3Normal(agent.velocity)
	var course = vec.Vec3Sub(agent.Target, agent.Pos)
	var course_relative = vec.Vec3Sub(course, agent.Pos)
	var course_normal = vec.Vec3Normal(course_relative)
	var force_dir = vec.AxisAngleRotation(velocity_normal, math.Pi, course_normal)

	impulse := agent.rocket.Thrust(force_dir, 0.2, time_delta)
	agent.applyImpulse(impulse)
}

func (agent *Agent) applyForce(force vec.Vec3, time_delta float64) {
	agent.Lock()
	agent.applyImpulse(vec.Vec3Scale(force, time_delta))
	agent.Unlock()
}

func (agent *Agent) applyImpulse(impulse vec.Vec3) {
	agent.Lock()
	velo_delta := vec.Vec3Scale(impulse, 1 / agent.mass)
	agent.velocity = vec.Vec3Add(agent.velocity, velo_delta)
	agent.Unlock()
}

func (agent *Agent) applyTorque(force vec.Vec3, point vec.Vec3, time_delta float64) {
	agent.Lock()
	lever := vec.Vec3Sub(agent.Pos, point)
	lever_dist := vec.Vec3Mag(lever)
	cos_sim := vec.Vec3CosineSimilarity(lever, force)

	linear_force := vec.Vec3Scale(force, cos_sim)
	linear_impulse := vec.Vec3Scale(linear_force, time_delta)
	agent.applyImpulse(linear_impulse)

	force_perpendicular := vec.Vec3Scale(force, 1 - math.Abs(cos_sim))
	_= vec.Vec3Scale(force_perpendicular, lever_dist) // torque
	agent.Unlock()
}

func (agent *Agent) Move(time_delta float64) {
	agent.Pos = vec.Vec3Add(agent.Pos, vec.Vec3Scale(agent.velocity, time_delta))
	agent.Euler = vec.Vec3Add(agent.Euler, vec.Vec3Scale(agent.angular_velocity, time_delta))
}

func (agent *Agent) Act(time_delta float64) {
	fmt.Println("agent acted")
}

func (agent *Agent) Alive() bool {
	return agent.alive
}

func (agent *Agent) GetPos() vec.Vec3 {
	return agent.Pos
}

func (agent *Agent) SetPos(new_pos vec.Vec3) {
	agent.Lock()
	agent.Pos = new_pos
	agent.Unlock()
}

func (agent *Agent) GetVelocity() vec.Vec3 {
	return agent.velocity
}

func (agent *Agent) AddVelocity(delta vec.Vec3) {
	agent.Lock()
	agent.velocity = vec.Vec3Add(agent.velocity, delta)
	agent.Unlock()
}

func (agent *Agent) AddAngularVelocity(delta vec.Vec3) {
	agent.Lock()
	agent.angular_velocity = vec.Vec3Add(agent.angular_velocity, delta)
	agent.Unlock()
}

func (agent *Agent) GetMass() float64 {
	return agent.mass
}

func (agent *Agent) GetRadius() float64 {
	return agent.radius
}

func (agent *Agent) GetType() string {
	return agent.Type
}

func (agent *Agent) GetJSON() string {
	json, _ := json.Marshal(agent)
	return string(json)
}


