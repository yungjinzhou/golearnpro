package main

import "fmt"

type Player struct {
	face string
	arm  string
	leg  string
	body string
}

func (p Player) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", p.face, p.body, p.leg, p.arm)
}

type PlayerBuilder interface {
	buildBody()
	buildLeg()
	buildArm()
	buildFace()
	getPlayer() Player
}

type SexyGirlBuilder struct {
	player Player
}

func (s *SexyGirlBuilder) buildArm() {
	s.player.arm = "纤细胳膊"
}

func (s *SexyGirlBuilder) buildBody() {
	s.player.body = "苗条身材"
}

func (s *SexyGirlBuilder) buildLeg() {
	s.player.leg = "大长腿"
}

func (s *SexyGirlBuilder) buildFace() {
	s.player.face = "漂亮脸蛋"
}

func (s *SexyGirlBuilder) getPlayer() Player {
	return s.player
}

type MonsterBuilder struct {
	player Player
}

func (m *MonsterBuilder) buildArm() {
	m.player.arm = "粗壮胳膊"
}

func (m *MonsterBuilder) buildBody() {
	m.player.body = "粗壮身材"
}

func (m *MonsterBuilder) buildLeg() {
	m.player.leg = "粗壮大腿"
}

func (m *MonsterBuilder) buildFace() {
	m.player.face = "毛脸蛋"
}

func (m *MonsterBuilder) getPlayer() Player {
	return m.player
}

type PlayerDirector struct {
}

func (d PlayerDirector) buildPlayer(builder PlayerBuilder) Player {
	builder.buildFace()
	builder.buildBody()
	builder.buildLeg()
	builder.buildArm()
	return builder.getPlayer()
}

func main() {
	builder := &SexyGirlBuilder{}
	director := PlayerDirector{}
	p := director.buildPlayer(builder)
	fmt.Printf(p.String())
}
