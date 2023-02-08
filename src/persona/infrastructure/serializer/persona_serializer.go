package persona_serializer

import (
	"fmt"

	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	"github.com/juanfer2/persona_5/src/shared/utilities"
)

type PersonaSerializer struct {
	ID              int                     `json:"id"`
	Name            string                  `json:"name"`
	Level           int                     `json:"level"`
	ImageUrl        string                  `json:"imageUrl"`
	ArcanaID        int                     `json:"aracanaId"`
	ArcanaName      string                  `json:"aracanaName"`
	Arcana          ArcanaSerializer        `json:"aracana"`
	PersonalityID   *int                    `json:"personalityId"`
	PersonalityName string                  `json:"personalityName"`
	Personality     PersonalitySerializer   `json:"personality"`
	Stats           []PersonaStatSerializer `json:"stats"`
}

type PersonaStatSerializer struct {
	Name  string `json:"name,omitempty"`
	Value int    `json:"value,omitempty"`
}

type ArcanaSerializer struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type PersonalitySerializer struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func PersonaStatsToSerializer(
	personaStats []persona_entities.PersonaStat,
) (personasStatSerializers []PersonaStatSerializer) {
	for _, personaStat := range personaStats {
		personasStatSerializers = append(personasStatSerializers, PersonaStatSerializer{
			Name: personaStat.Stat.Name, Value: personaStat.Value,
		})
	}
	return
}

func ArcanaToSerializer(arcana persona_entities.Arcana) ArcanaSerializer {
	return ArcanaSerializer{ID: arcana.ID, Name: arcana.Name}
}

func PersonalityToSerializer(personality persona_entities.Personality) PersonalitySerializer {
	return PersonalitySerializer{ID: personality.ID, Name: personality.Name}
}

func PersonaToSerializer(persona persona_entities.Persona) PersonaSerializer {
	return PersonaSerializer{ID: persona.ID, Name: persona.Name, Level: persona.Level,
		ImageUrl: persona.ImageUrl}
}

func PersonasToSerializer(personas []persona_entities.Persona) []PersonaSerializer {
	personaSerializer := []PersonaSerializer{}

	for _, persona := range personas {
		func(persona persona_entities.Persona) {
			personaSerializer = append(personaSerializer,
				PersonaSerializer{
					ID:              persona.ID,
					Name:            persona.Name,
					Level:           persona.Level,
					ImageUrl:        utilities.BuildImageUrl("personas/" + persona.ImageUrl),
					ArcanaID:        persona.ArcanaID,
					ArcanaName:      persona.Arcana.Name,
					Arcana:          ArcanaToSerializer(persona.Arcana),
					PersonalityID:   persona.PersonalityID,
					PersonalityName: persona.Personality.Name,
					Personality:     PersonalityToSerializer(persona.Personality),
					Stats:           PersonaStatsToSerializer(persona.PersonaStats),
				})
			fmt.Println("persona.Name: ", persona.Name)
		}(persona)
	}

	return personaSerializer
}
