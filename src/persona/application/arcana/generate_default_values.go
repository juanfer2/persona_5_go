package persona_application

import (
	persona_entities "github.com/juanfer2/persona_5/src/persona/domain/entities"
	persona_repository "github.com/juanfer2/persona_5/src/persona/domain/repositories"
)

type GenerateDefaultValues struct {
	repository persona_repository.ArcanaRepository
}

func NewGenerateDefaultValues(
	repository persona_repository.ArcanaRepository,
) *GenerateDefaultValues {
	return &GenerateDefaultValues{
		repository: repository,
	}
}

func (generateDefaultValues GenerateDefaultValues) Call() {
	listArcana := []persona_entities.Arcana{
		{Name: "The Fool", CharacterName: "Igor"},
		{Name: "The Magician", CharacterName: "Morgana"},
		{Name: "Priestess", CharacterName: "Makoto Niijima"},
		{Name: "Empress", CharacterName: "Haru Okumura"},
		{Name: "Emperor", CharacterName: "Yusuke Kitagawa"},
		{Name: "Heirophant", CharacterName: "Sojiro Sakura"},
		{Name: "Lovers", CharacterName: "Ann Takamaki"},
		{Name: "Chariot", CharacterName: "Ryuji Sakamoto"},
		{Name: "Justice", CharacterName: "Goro Akechi"},
		{Name: "Hermit", CharacterName: "Futaba Sakura"},
		{Name: "Wheel of Fortune", CharacterName: "Mifune Chihaya"},
		{Name: "Strength", CharacterName: "Caroline/Justine"},
		{Name: "Hanged Man", CharacterName: "Iwai Munehisa"},
		{Name: "Death", CharacterName: "Tae Takemi"},
		{Name: "Temperance", CharacterName: "Sadayo Kawakami"},
		{Name: "Devil", CharacterName: "Ohya Ichiko"},
		{Name: "Tower", CharacterName: "Shinya Oda"},
		{Name: "Star", CharacterName: "Hifumi Togo"},
		{Name: "Moon", CharacterName: "Yuki Mishima"},
		{Name: "Sun", CharacterName: "Toranosuke Yoshida"},
		{Name: "Judgement", CharacterName: "Sae Niijima"},
	}

	generateDefaultValues.repository.CreateInBatches(listArcana)
}
