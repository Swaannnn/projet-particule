package main

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.

func (g *game) Update() error {

	g.system.Update()

	// permet de suivre la souris
	if config.General.CursorSpawn {
		var mX, mY = ebiten.CursorPosition()
		config.General.SpawnX, config.General.SpawnY = mX, mY
	}

	// permet d'activer / désactiver la génération sur la souris
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		if config.General.RandomSpawn {
			config.General.RandomSpawn = false
		}
		config.General.CursorSpawn = !config.General.CursorSpawn
	}

	// permet d'activer / désactiver le Random Spawn
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		config.General.RandomSpawn = !config.General.RandomSpawn
	}

	// permet d'activer / désactiver la gravité
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		config.General.Gravity = !config.General.Gravity
	}

	// permet d'activer / désactiver le dégradé
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		config.General.Degrade = !config.General.Degrade
		g.system.RandomColor()
	}

	// permet d'activer / désactiver la génération en forme de cercle
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		config.General.RodSpawn = !config.General.RodSpawn
		config.General.SpawnRate = 100
		config.General.LifeTime = true
		config.General.DeadTime = 10
		config.General.DisparitionTime = 10
	}

	// permet d'augmenter le Spawn Rate avec la molette de la souris
	var _, test = ebiten.Wheel()
	if test < 0 {
		if config.General.SpawnRate > 0 {
			config.General.SpawnRate += test
		}
	} else {
		config.General.SpawnRate += test
	}

	// SJLDHFKJSHFKJGHSLJFGS
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		config.General.DeadTime += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		config.General.DeadTime -= 1
	}

	// permet d'augmenter / diminuer la force de gravité
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		config.General.GravityAcc += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		config.General.GravityAcc -= 1
	}

	// permet de mettre pause (fige les particules)
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.PauseNumber%2 == 0 {
			g.SpeedSave = config.General.SpeedParticles
			g.SpawnSave = config.General.SpawnRate
			g.LifeTimeSave = config.General.DeadTime
			g.DisparitionTimeSave = int(config.General.DeadTime)
			config.General.SpeedParticles = 0
			config.General.SpawnRate = 0
			config.General.DeadTime = 0
			config.General.DisparitionTime = 0
			g.PauseNumber += 1
		} else {
			config.General.SpeedParticles = g.SpeedSave
			config.General.SpawnRate = g.SpawnSave
			config.General.DeadTime = g.LifeTimeSave
			config.General.DisparitionTime = float64(g.DisparitionTimeSave)
			g.PauseNumber += 1
		}
	}

	// permet de faire une explosion de particules
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		for i := 0.0; i <= config.General.SpawnRate*100; i++ {
			if config.General.Degrade {
				g.system.AddParticleDegrade()
			} else {
				g.system.AddParticle()
			}
		}
	}

	// permet d'augmenter / diminuer la vitesse des particules
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadAdd) {
		config.General.SpeedParticles += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyNumpadSubtract) {
		config.General.SpeedParticles -= 1
	}

	// afiche / cache le premier menu
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		config.General.Help = !config.General.Help
	}

	// affiche / cache le deuxième menu
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		config.General.Debug = !config.General.Debug
	}
	return nil
}
