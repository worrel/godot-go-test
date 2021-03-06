package main

import (
	"math"

	"github.com/shadowapex/godot-go/godot"
)

func NewMyGodot() godot.Class {
	return &MyGodotClass{
		velocity: godot.NewVector2(1.0, 0.0),
	}
}

type MyGodotClass struct {
	godot.Area2D
	velocity       *godot.Vector2
	screenSize     *godot.Rect2
	animatedSprite *godot.AnimatedSprite
}

const SPEED = 100.0

var ANIMATED_SPRITE_PATH = godot.NewNodePath("AnimatedSprite")

func (p *MyGodotClass) X_ready() {
	p.screenSize = p.GetViewportRect()
	// This works
	asNode := p.GetNode(ANIMATED_SPRITE_PATH)
	as := godot.AnimatedSprite{}
	as.Object = asNode.Object
	p.animatedSprite = &as
	// This works
	p.animatedSprite.Stop()
}

func (p *MyGodotClass) X_process(delta float64) {
	var deltaX, deltaY float64

	// All these work
	if godot.Input.IsActionPressed("ui_right") {
		deltaX += 1
	}
	if godot.Input.IsActionPressed("ui_left") {
		deltaX -= 1
	}
	if godot.Input.IsActionPressed("ui_down") {
		deltaY += 1
	}
	if godot.Input.IsActionPressed("ui_up") {
		deltaY -= 1
	}

	if deltaX != 0 || deltaY != 0 {
		godot.Log.Println("Got input: X=", deltaX, ", Y=", deltaY)
	}

	p.velocity = godot.NewVector2(deltaX, deltaY).Normalized()

	if p.velocity.Length() > 0 {
		p.velocity = p.velocity.OperatorMultiplyScalar(SPEED)

		screensize := p.screenSize.GetSize()

		position := p.GetPosition().OperatorAdd(*p.velocity.OperatorMultiplyScalar(delta))
		position.SetX(math.Max(0, math.Min(position.GetX(), screensize.GetX())))
		position.SetY(math.Max(0, math.Min(position.GetY(), screensize.GetY())))

		if p.velocity.GetX() != 0 {
			// This fails
			p.animatedSprite.SetAnimation("right")
			p.animatedSprite.SetFlipV(false)
			p.animatedSprite.SetFlipH(p.velocity.GetX() < 0)
		} else if p.velocity.GetY() != 0 {
			// This fails
			p.animatedSprite.SetAnimation("up")
			p.animatedSprite.SetFlipV(p.velocity.GetY() > 0)
		}

		// This also fails
		p.animatedSprite.Play("")
	} else {
		p.animatedSprite.Stop()
	}

}

func init() {
	// Register will register the given class constructor with Godot.
	godot.Register(NewMyGodot)
}

func main() {}
