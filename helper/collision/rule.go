package collision

import (
	"fmt"
	"math"
	"time"

	"github.com/grzesl/ship/component"
	"github.com/grzesl/ship/helper/enum"
)

// Reaction is the way to resolve the collision.
type Reaction func(e Event, c Contact)

// Rule is the condition under which collisions are allowed and how they are resolved.
type Rule struct {
	A, B enum.CollisionGroup
	Reaction
}

// Slide is a way to resolve collisions in which one object will crawl on the surface of another.
func Slide(a, b enum.CollisionGroup) Rule {
	return Rule{a, b, func(e Event, c Contact) {
		e.VelA.L += c.NormalX * math.Abs(e.VelA.L) * (1 - c.Time)
		e.VelA.M += c.NormalY * math.Abs(e.VelA.M) * (1 - c.Time)

		if e.ControlA != nil {
			e.ControlA.VolumeSpeed = -e.ControlA.VolumeSpeed

			if e.GoodsA != nil {
				if time.Since(e.GoodsA.Visited) < time.Second*30 {
					e.GoodsA.Visited = time.Now()
					return
				}
				e.GoodsA.Visited = time.Now()
			}
			if e.GoodsB != nil {
				if time.Since(e.GoodsB.Visited) < time.Second*30 {
					e.GoodsB.Visited = time.Now()
					return
				}
				e.GoodsB.Visited = time.Now()
			}

			if e.ControlA.Carry {
				e.ControlA.Carry = false
			} else {
				e.ControlA.Carry = true
			}
			e.ControlA.VolumeSpeed = 0
		}
		if e.ControlB != nil {
			e.ControlB.VolumeSpeed = -e.ControlB.VolumeSpeed

			if e.GoodsA != nil {
				if time.Since(e.GoodsA.Visited) < time.Second*30 {
					e.GoodsA.Visited = time.Now()
					return
				}
				e.GoodsA.Visited = time.Now()
			}
			if e.GoodsB != nil {
				if time.Since(e.GoodsB.Visited) < time.Second*30 {
					e.GoodsB.Visited = time.Now()
					return
				}
				e.GoodsB.Visited = time.Now()
			}

			if e.ControlB.Carry {
				e.ControlB.Carry = false
			} else {
				e.ControlB.Carry = true
			}

			e.ControlB.VolumeSpeed = 0
		}

	}}
}

// Event is all the additional data needed to resolve a particular collision.
type Event struct {
	PosA, PosB         *component.Pos
	VelA, VelB         *component.Vel
	SizeA, SizeB       *component.Size
	SolidA, SolidB     *component.Solid
	ControlA, ControlB *component.Control
	GoodsA, GoodsB     *component.Goods
	Reaction           Reaction
	Time               float64
}

func NewEvent(
	aPos *component.Pos, aVel *component.Vel, aSize *component.Size, aSolid *component.Solid, aControl *component.Control, aGoods *component.Goods,
	bPos *component.Pos, bVel *component.Vel, bSize *component.Size, bSolid *component.Solid, bControl *component.Control, bGoods *component.Goods,
	reaction Reaction, time float64,
) Event {
	return Event{aPos, bPos, aVel, bVel, aSize, bSize, aSolid, bSolid, aControl, bControl, aGoods, bGoods, reaction, time}
}

func (e Event) String() string {
	return fmt.Sprintf(
		"collision %s, %s, %s -> %s, %s, %s at time: %.2f",
		e.PosA, e.VelA, e.SizeA, e.PosB, e.VelB, e.SizeB, e.Time,
	)
}

// FindRule finds a collision resolution rule in the list that is similar to or reversed from the given one.
func FindRule(rules []Rule, a, b enum.CollisionGroup) (rule Rule, inverted, ok bool) {
	for _, r := range rules {
		if a == r.A && b == r.B {
			return r, false, true
		}

		if a == r.B && b == r.A {
			return r, true, true
		}
	}

	return Rule{}, false, false
}
