package main

import (
	"bytes"
	"encoding/json"
	"log"
)

type Problems struct {
	Problem1  [][]uint32
	Problem2  [][]uint32
	Problem3  uint32
	Problem4  uint32
	Problem5  string
	Problem6  uint32
	Problem7  uint32
	Problem8  []uint32
	Problem9  []uint32
	Problem10 Problem10
	Problem11 Problem11
	Problem12 []int
	Problem13 []string
	Problem14 Problem14
	Problem15 string
	Problem16 Problem16
}

type Problem10 struct {
	State         []uint32
	RoundConstant uint32 `json:"round_constant"`
	ScheduleWord  uint32 `json:"schedule_word"`
}

type Problem11 struct {
	State []int
	Block string
}

type Problem14 struct {
	OriginalInput string `json:"original_input"`
	ChosenSuffix  string `json:"chosen_suffix"`
}

type Problem16 struct {
	OriginalHash   string `json:"original_hash"`
	OriginalLength int    `json:"original_length"`
	ChosenSuffix   string `json:"chosen_suffix"`
}

type Solutions struct {
	Problem1  []uint32 `json:"problem1,omitempty"`
	Problem2  []uint32 `json:"problem2,omitempty"`
	Problem3  uint32   `json:"problem3,omitempty"`
	Problem4  uint32   `json:"problem4,omitempty"`
	Problem5  []uint32 `json:"problem5,omitempty"`
	Problem6  uint32   `json:"problem6,omitempty"`
	Problem7  uint32   `json:"problem7,omitempty"`
	Problem8  uint32   `json:"problem8,omitempty"`
	Problem9  uint32   `json:"problem9,omitempty"`
	Problem10 []uint32 `json:"problem10,omitempty"`
	Problem11 []int    `json:"problem11,omitempty"`
	Problem12 []string `json:"problem12,omitempty"`
	Problem13 []string `json:"problem13,omitempty"`
	Problem14 string   `json:"problem14,omitempty"`
	Problem15 []int    `json:"problem15,omitempty"`
	Problem16 string   `json:"problem16,omitempty"`
}

func pprint(problems any) {
	var pretty bytes.Buffer
	ugly, err := json.Marshal(problems)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Indent(&pretty, ugly, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pretty.String())
}
