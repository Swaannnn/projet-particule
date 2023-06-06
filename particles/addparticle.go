package particles

import (
	"math/rand"
	"project-particles/config"
	"time"
)

func (s *System) AddParticle() {
	var PosX float64 = 0
	var PosY float64 = 0
	e := rand.New(rand.NewSource(time.Now().UnixNano()))

	if config.General.RandomSpawn {
		PosX = float64(e.Intn(config.General.WindowSizeX - 10))
		PosY = float64(e.Intn(config.General.WindowSizeY - 10))

		// GENERATEUR EN FORME DE TRAIT
	} else if config.General.RodSpawn {
		PosX = rand.Float64() * float64(config.General.WindowSizeX)
		PosY = rand.Float64()*float64(config.General.RodSize) + (float64(config.General.SpawnY) - float64(config.General.RodSize/2))

	} else {
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
	}
	taille := rand.Float64()
	s.Content.PushFront(&Particle{
		PositionX: PosX, PositionY: PosY,
		Rotation: 0,
		ScaleX:   taille, ScaleY: taille,
		ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
		Opacity:   1,
		VelocityX: rand.Float64() - rand.Float64(), VelocityY: rand.Float64() - rand.Float64(),
		GravityY: config.General.GravityAcc,
		Out:      false,
		Split:    false,
	})
}

func (s *System) AddParticleDegrade() {
	var PosX float64 = 0
	var PosY float64 = 0
	e := rand.New(rand.NewSource(time.Now().UnixNano()))

	if config.General.RandomSpawn {
		PosX = float64(e.Intn(config.General.WindowSizeX - 10))
		PosY = float64(e.Intn(config.General.WindowSizeY - 10))

		// GENERATEUR EN FORME DE TRAIT
	} else if config.General.RodSpawn {
		PosX = rand.Float64() * float64(config.General.WindowSizeX)
		PosY = rand.Float64()*float64(config.General.RodSize) + (float64(config.General.SpawnY) - float64(config.General.RodSize/2))
	} else {
		PosX = float64(config.General.SpawnX)
		PosY = float64(config.General.SpawnY)
	}
	taille := rand.Float64()

	HorizontalColor := float64(PosX) / float64(config.General.WindowSizeX)
	VerticalColor := float64(PosY) / float64(config.General.WindowSizeY)

	var Blue float64 = 0
	var Red float64 = 0
	var Green float64 = 0

	if config.General.HorizontalColor == "blue" {
		Blue = HorizontalColor
		if config.General.VerticalColor == "red" {
			Red = VerticalColor
		} else {
			Green = VerticalColor
		}
	} else if config.General.HorizontalColor == "red" {
		Red = HorizontalColor
		if config.General.VerticalColor == "blue" {
			Blue = VerticalColor
		} else {
			Green = VerticalColor
		}
	} else if config.General.HorizontalColor == "green" {
		Green = HorizontalColor
		if config.General.VerticalColor == "red" {
			Red = VerticalColor
		} else {
			Blue = VerticalColor
		}
	}
	s.Content.PushFront(&Particle{
		PositionX: PosX, PositionY: PosY,
		Rotation: 0,
		ScaleX:   taille, ScaleY: taille,
		Opacity:   1,
		VelocityX: rand.Float64() - rand.Float64(), VelocityY: rand.Float64() - rand.Float64(),
		ColorBlue: Blue, ColorGreen: Green, ColorRed: Red,
		GravityY: config.General.GravityAcc,
		Out:      false,
		Split:    false,
	})
}
