package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.

func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok {
			if p.Out == false {
				options := ebiten.DrawImageOptions{}
				options.GeoM.Rotate(p.Rotation)
				options.GeoM.Scale(p.ScaleX, p.ScaleY)
				options.GeoM.Translate(p.PositionX, p.PositionY)
				options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
				screen.DrawImage(assets.ParticleImage, &options)
			}
		}
	}

	// affiche les ips, les nombre de particules, certaines valeurs...
	if config.General.Debug {
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("images par secondes : ", ebiten.CurrentTPS(), "\n", "nombre de particules : ", g.system.Content.Len(), "\n", "\n", "Spawn Rate : ", config.General.SpawnRate, "\n", "Speed Particles : ", config.General.SpeedParticles, "\n", "Gravity Acc : ", config.General.GravityAcc, "\n", "\n", "'tab' pour plus d'aide", "\n", "'esc' pour cacher ce menu"), 10, 10)
	}

	// affiche les touches pour changer les paramètres pendant l'exeecution.
	if config.General.Help {
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("espace : pause", "\n", "\n", "clic droit : active / désactive la génération des particules sur le pointeur de la souris", "\n", "r : active / désactive la génération aléatoire", "\n", "molette de la souris : augmente / diminue le SpawnRate", "\n", "g : active / désactive la gravité", "\n", "flèche du haut↑/ bas↓ : augmente / diminue la gravité", "\n", "d : active / désactive le dégradé", "\n", "clic gauche : fait une explosion de particule", "\n"), 10, config.General.WindowSizeY-160)
	}
}
