package sensorData

import "time"

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
