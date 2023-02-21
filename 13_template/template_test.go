package _3_template

import "testing"

func TestTemplate(t *testing.T) {

	// 做西红柿炒蛋
	tomatoScrambledEggs := &TomatoScrambledEggs{}
	DoCook(tomatoScrambledEggs)

	// 做麻婆豆腐
	mapoTofu := &MapoTofu{}
	DoCook(mapoTofu)

}
