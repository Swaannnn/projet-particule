package particles

import (
	"project-particles/config"
	"testing"
)

// test du nombre de particules initial

func TestInitNumParticles(t *testing.T) {
	list := NewSystem()
	if list.Content.Len() != config.General.InitNumParticles {
		t.Error("Il n'y a pas le bon nombre de particules initiales, il y en a", list.Content.Len(), "alors qu'il devrait y en avoir", config.General.InitNumParticles)
	}
}

// test de position des particules

func TestPositionParticles(t *testing.T) {
	for element := NewSystem().Content.Front(); element != nil; element = element.Next() {
		if config.General.RandomSpawn == false && element.Value.(*Particle).PositionX != float64(config.General.SpawnX) || element.Value.(*Particle).PositionY != float64(config.General.SpawnY) {
			t.Error("Une particule n'est pas apparue au bon endroit (RandomSpawn est désactivé).")
		} else if config.General.RandomSpawn == true && element.Value.(*Particle).PositionX == float64(config.General.SpawnX) || element.Value.(*Particle).PositionY == float64(config.General.SpawnY) {
			t.Error("Les particules ne devraient pas apparaître au centre (RandomSpawn est désactivé).")
		}
	}
}

// test de vitesse et gravité des particules

func TestVelocityParticles(t *testing.T) {
	config.General.InitNumParticles = 1000
	config.General.SpawnRate = 5
	config.General.SpeedParticles = 1
	s := NewSystem()
	var posX, posY float64
	for e := s.Content.Front(); e != nil; e = e.Next() {
		var element = e.Value.(*Particle)
		posX, posY = element.PositionX, element.PositionY
		s.Update()
		if config.General.Gravity {
			if posX+element.VelocityX != element.PositionX || posY+element.VelocityY+float64(config.General.GravityAcc) != element.PositionY {
				t.Error("les particules n'avancent pas à la bonne vitesse / gravité")
			}
		} else {
			if posX+element.VelocityX != element.PositionX || posY+element.VelocityY != element.PositionY {
				t.Error("les particules n'avancent pas à la bonne vitesse")
			}
		}
	}
}

// test de la durée de vie des particules

func TestLifeTime(t *testing.T) {
	for element := NewSystem().Content.Front(); element != nil; element = element.Next() {
		if element.Value.(*Particle).Time > config.General.DeadTime && element.Value.(*Particle).Out != false {
			t.Error("Une particule ne devrait plus être affichée.")
		}
	}
}

// test de l'extérieur de l'écran

func TestOutScreen(t *testing.T) {
	config.General.OutScreen = true
	config.General.WindowSizeX = 100
	config.General.WindowSizeY = 100
	config.General.OutScreenMaxX = 15
	config.General.OutScreenMaxY = 15
	for i := 0; i < 10; i++ {
		for element := NewSystem().Content.Front(); element != nil; element = element.Next() {
			if element.Value.(*Particle).PositionX > float64(config.General.WindowSizeX)-10+config.General.OutScreenMaxX || element.Value.(*Particle).PositionX < 0-config.General.OutScreenMaxX || element.Value.(*Particle).PositionY > float64(config.General.WindowSizeY)-10+config.General.OutScreenMaxY || element.Value.(*Particle).PositionY < 0-config.General.OutScreenMaxY {
				t.Errorf("Une particule ne peut pas être en dehors de l'écran")
			}
		}
	}
}

// test de la génération en forme de barre
func TestRodSpawn(t *testing.T) {
	config.General.RodSpawn = true
	config.General.WindowSizeY = 100
	config.General.SpawnY = 50
	config.General.RodSize = 10
	for i := 0; i < 10; i++ {
		for element := NewSystem().Content.Front(); element != nil; element = element.Next() {
			if element.Value.(*Particle).PositionY > float64(config.General.SpawnY)+float64(config.General.RodSize/2) && element.Value.(*Particle).PositionY < float64(config.General.SpawnY)-float64(config.General.RodSize/2) {
				t.Errorf("La particule n'a pas été générée dans la barre")
			}
		}
	}
}
