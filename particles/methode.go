package particles

import (
	"math/rand"
	"project-particles/config"
)

// optimisation de mémoire 1 (5.5)
func (s *System) Optimisation1() {
	for i := s.Content.Front(); i != nil; {
		n := i.Value.(*Particle)
		next := i.Next()
		if n.Out || n.Opacity == 0 {
			s.Content.Remove(i)
		}
		i = next
	}
}

// extention n°1 : Gravité (5.1)
func (p *Particle) Gravity() {
	if p.Time%30 == 0 {
		p.Bound = true
	}
	p.VelocityY += 0.01 * config.General.GravityAcc
}

// extention n°2 : Extérieur de l'écran (5.2)
func (p *Particle) OutScreen() {
	if p.PositionX > float64(config.General.WindowSizeX)+config.General.OutScreenMaxX || p.PositionX < -config.General.OutScreenMaxX || p.PositionY < -config.General.OutScreenMaxY || p.PositionY > float64(config.General.WindowSizeY)+config.General.OutScreenMaxY {
		p.Out = true
	}
}

// extention n°3 : Durée de vie (5.3)
func (p *Particle) LifeTime() {
	if p.Time >= config.General.DeadTime {
		p.Opacity -= 0.01 * config.General.DisparitionTime
		if p.Opacity <= 0.01 {
			p.Out = true
		}
	}
}

// extention n°4 : Variation de couleur, d'échelle, de rotation, de transparance (5.4)

// variation de couleur
func (p *Particle) ChangeColor() {
	if p.Time%20 == 0 {
		p.ColorRed = rand.Float64()
		p.ColorGreen = rand.Float64()
		p.ColorBlue = rand.Float64()
	}
}

// variation de taille
func (p *Particle) ChangeSize() {
	if p.Time%20 == 0 {
		taille := rand.Float64()
		p.ScaleX = taille
		p.ScaleY = taille
	}
}

// variation d'opacité
func (p *Particle) ChangeOpacity() {
	if p.Time%120 == 0 {
		p.Opacity = rand.Float64()
	}
}

// variation de rotation
func (p *Particle) ChangeRotation() {
	p.Rotation += (rand.Float64() - rand.Float64()) * 0.02 * config.General.RotationAcc
}

// Voici quelques extentions que nous avons choisit d'ajouter :

// 1 - collision sur les murs
func (p *Particle) WallCollision() {
	if p.PositionX <= 0 || p.PositionX+p.ScaleX*10 >= float64(config.General.WindowSizeX) {
		p.VelocityX = -p.VelocityX
	} else if p.PositionY <= 0 {
		p.VelocityY = -p.VelocityY
	} else if p.PositionY+p.ScaleY*10 >= float64(config.General.WindowSizeY) {
		p.VelocityY = -p.VelocityY
		p.Bound = true
		p.ColorBlue = rand.Float64()
		p.ColorGreen = rand.Float64()
		p.ColorRed = rand.Float64()
		p.NbrBound++
	}

	// Pour l'extention GRAVITE
	if config.General.Gravity {
		if p.Bound {
			p.VelocityY = p.VelocityY / 1.2
			p.Bound = false
			p.VelocityY += 0.2
		}
		if p.NbrBound == 5 {
			p.Opacity = 0
		}
	}
}

// 2 - duplication d'une particule après collision sur un mur (ne marche pas totalement)
func SplitParticles(s *System, p *Particle) {
	config.General.WallCollision = true
	taille := rand.Float64()
	if !p.Split {
		if p.PositionX <= 0 {
			s.Content.PushFront(&Particle{
				PositionX: p.PositionX + 1, PositionY: p.PositionY,
				ScaleX: taille, ScaleY: taille,
				ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
				Opacity:   1,
				VelocityX: rand.Float64(), VelocityY: rand.Float64() - rand.Float64(),
				GravityY: config.General.GravityAcc,
				Out:      false,
				Split:    true,
			})
			p.Split = true
		} else if p.PositionY <= 0 {
			s.Content.PushFront(&Particle{
				PositionX: p.PositionX, PositionY: p.PositionY + 1,
				ScaleX: taille, ScaleY: taille,
				ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
				Opacity:   1,
				VelocityX: -p.VelocityX, VelocityY: -p.VelocityY,
				GravityY: config.General.GravityAcc,
				Out:      false,
				Split:    true,
			})
			p.Split = true
		} else if p.PositionY+10*p.ScaleY >= float64(config.General.WindowSizeY) {
			s.Content.PushFront(&Particle{
				PositionX: p.PositionX, PositionY: p.PositionY - 1,
				ScaleX: taille, ScaleY: taille,
				ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
				Opacity:   1,
				VelocityX: -p.VelocityX, VelocityY: -p.VelocityY,
				GravityY: config.General.GravityAcc,
				Out:      false,
				Split:    true,
			})
			p.Split = true
		} else if p.PositionX+10*p.ScaleX >= float64(config.General.WindowSizeX) {
			s.Content.PushFront(&Particle{
				PositionX: p.PositionX - 1, PositionY: p.PositionY,
				ScaleX: taille, ScaleY: taille,
				ColorRed: rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
				Opacity:   1,
				VelocityX: -rand.Float64(), VelocityY: rand.Float64() - rand.Float64(),
				GravityY: config.General.GravityAcc,
				Out:      false,
				Split:    true,
			})
			p.Split = true
		}
	}

	// pour éviter des bugs
	if p.PositionX <= 0 || p.PositionY <= 0 || p.PositionX+10*p.ScaleX >= float64(config.General.WindowSizeX) || p.PositionY+10*p.ScaleY >= float64(config.General.WindowSizeY) {
		p.Opacity -= 0.05
		if p.Opacity <= 0.5 {
			p.Out = true
		}
	}
}

// 3 - changement de couleur selon la position des particules
func (p *Particle) Degrade() {
	HorizontalColor := float64(p.PositionX) / float64(config.General.WindowSizeX)
	VerticalColor := float64(p.PositionY) / float64(config.General.WindowSizeY)

	if config.General.HorizontalColor == "blue" {
		p.ColorBlue = HorizontalColor
		if config.General.VerticalColor == "red" {
			p.ColorRed = VerticalColor
			p.ColorGreen = 0
		} else {
			p.ColorRed = 0
			p.ColorGreen = VerticalColor
		}
	} else if config.General.HorizontalColor == "red" {
		p.ColorRed = HorizontalColor
		if config.General.VerticalColor == "blue" {
			p.ColorBlue = VerticalColor
			p.ColorGreen = 0
		} else {
			p.ColorBlue = 0
			p.ColorGreen = VerticalColor
		}
	} else if config.General.HorizontalColor == "green" {
		p.ColorGreen = HorizontalColor
		if config.General.VerticalColor == "red" {
			p.ColorRed = VerticalColor
			p.ColorBlue = 0
		} else {
			p.ColorRed = 0
			p.ColorBlue = VerticalColor
		}
	}
}

// Fonction supplémentaire pour éviter un problème de couleur quand les particules prennent la couleur du dégradé
func (s *System) RandomColor() {
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		p.ColorBlue, p.ColorRed, p.ColorGreen = rand.Float64(), rand.Float64(), rand.Float64()
	}
}
