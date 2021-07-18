package models

import "time"

type AttackPlan struct {
	ID         int64     `json:"id" db:"attack_plan_id"`
	Name       string    `json:"name" db:"attack_plan_name"`
	Type       int       `json:"type" db:"attack_plan_type"`
	Content    string    `json:"content" db:"attack_plan_content"`
	UseCount   int       `json:"used_count" db:"attack_plan_use_count"`
	UseTime    time.Time `json:"use_time" db:"attack_plan_use_time"`
	FinishTime time.Time `json:"finish_time" db:"attack_plan_finish_time"`
	Status     int       `json:"status" db:"attack_plan_status"`
	Remark     string    `json:"remark" db:"attack_plan_remark"`
}
