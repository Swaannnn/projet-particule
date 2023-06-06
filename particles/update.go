package particles

import (
	"project-particles/config"
)

// AJOUTER TAILLE MIN ET MAX DANS CONFIG.JSON

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de chacune des particules) à chaque pas de temps.
// Elle est appellée exactement 60 fois par seconde (de manière régulière) par la fonction principale du projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {
	for i := s.Content.Front(); i != nil; i = i.Next() {
		n := i.Value.(*Particle)

		// COMPTEUR
		n.Time += 1

		n.PositionX += n.VelocityX * float64(config.General.SpeedParticles)
		n.PositionY += n.VelocityY * float64(config.General.SpeedParticles)

		// 5.1 : GRAVITE
		if config.General.Gravity {
			n.Gravity()
		}

		// 5.2 : EXTERIEUR DE L'ECRAN
		if config.General.OutScreen {
			n.OutScreen()
		}

		// 5.3 : DUREE DE VIE
		if config.General.LifeTime {
			n.LifeTime()
		}

		// 5.4 : VARIATIONS DE COULEUR, D'ECHELLE, DE ROTATION, DE TRANSPARANCE
		if config.General.ChangeColor {
			n.ChangeColor()
		}
		if config.General.ChangeSize {
			n.ChangeSize()
		}
		if config.General.ChangeOpacity {
			n.ChangeOpacity()
		}
		if config.General.ChangeRotation {
			n.ChangeRotation()
		}

		// COLLISION SUR LES MURS
		if config.General.WallCollision {
			n.WallCollision()
		}

		// DUPLICATION DE PARTICULES
		if config.General.SplitParticles {
			SplitParticles(s, n)
			if n.Time%30 == 0 {
				n.Split = false
			}
		}

		// DEGRADE DE COULEUR
		if config.General.Degrade {
			n.Degrade()
		}
	}
	// 5.5 : OPTIMISATION DE LA MEMOIRE
	s.Optimisation1()

	s.Count += config.General.SpawnRate
	nbrspawn := s.Count - config.General.SpawnRate

	for i := 0; i < int(nbrspawn); i++ {
		if config.General.Degrade {
			s.AddParticleDegrade()
		} else {
			s.AddParticle()
		}
		s.Count = s.Count - 1.0
	}

}
