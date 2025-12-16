package main

import (
	robot "RobotTelemetry/sensorData"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type Sensor struct {
	RobotId     int     `json:"RobotId"`
	Sensor_1    float64 `json:"Sensor_1"`
	Motor_left  int     `json:"Motor_left"`
	Motor_right int     `json:"Motor_right"`
	timestamp   time.Time
}

func (sl *Sensor) NewSensor(robotId int,
	sensor_1 float64,
	motor_left int,
	motor_right int,
) Sensor {
	if robotId != 0 || sensor_1 != 0 || motor_left != 0 || motor_right != 0 {
		return Sensor{
			RobotId:     robotId,
			Sensor_1:    sensor_1,
			Motor_left:  motor_left,
			Motor_right: motor_right,
			timestamp:   time.Now(),
		}
	}
	return Sensor{}
}

func main() {

	for range 10 {
		var sensor robot.Sensor
		sensor_1 := rand.Float64()
		motor_left := rand.Intn(360)
		motor_right := rand.Intn(360)

		newRobot := sensor.NewSensor(1, sensor_1, motor_left, motor_right)
		data := map[string]any{
			"robot_id":    newRobot.RobotId,
			"sensor_1":    newRobot.Sensor_1,
			"motor_right": newRobot.Motor_right,
			"motor_left":  newRobot.Motor_left,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		conn, _ := net.Dial("udp", "localhost:1235")
		defer conn.Close()
		conn.Write(jsonData)
		fmt.Printf("Sent: %v\n", string(jsonData))
	}
}
