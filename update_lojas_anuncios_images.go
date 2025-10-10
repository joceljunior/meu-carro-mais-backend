package main

import (
	"fmt"
	"log"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
)

func main() {
	fmt.Println("🖼️ Atualizando imagens das lojas e anúncios...")

	// Inicializa o banco de dados
	database.InitDB()

	// URLs das imagens das lojas
	lojaImages := []string{
		"https://firebasestorage.googleapis.com/v0/b/meuautomais-63b00.firebasestorage.app/o/ic_moto10.png?alt=media&token=dacfadb2-4da9-4954-ae92-e437f0179a92",
		"https://firebasestorage.googleapis.com/v0/b/meuautomais-63b00.firebasestorage.app/o/ic_logo1.png?alt=media&token=52658d91-4d79-41ef-a777-37c5d8635158",
		"https://firebasestorage.googleapis.com/v0/b/meuautomais-63b00.firebasestorage.app/o/ic_eletrocar.png?alt=media&token=407448f9-0d6b-4815-9c15-a168a75c36f4",
	}

	// URL da imagem dos anúncios
	anuncioImage := "https://firebasestorage.googleapis.com/v0/b/meuautomais-63b00.firebasestorage.app/o/moto_ads.png?alt=media&token=b01a11cd-59de-4a06-8236-566250db811d"

	// 1. Atualizar imagens das lojas
	fmt.Println("\n1. Atualizando imagens das lojas...")
	var lojas []models.Loja
	err := database.DB.Where("data_exclusao IS NULL").Find(&lojas).Error
	if err != nil {
		log.Printf("❌ Erro ao buscar lojas: %v", err)
		return
	}

	for i, loja := range lojas {
		// Usa a imagem baseada no índice, repetindo se necessário
		imageIndex := i % len(lojaImages)
		loja.Imagem = lojaImages[imageIndex]

		err := database.DB.Model(&loja).Update("imagem", loja.Imagem).Error
		if err != nil {
			log.Printf("❌ Erro ao atualizar loja %d: %v", loja.ID, err)
		} else {
			fmt.Printf("✅ Loja '%s' (ID: %d) atualizada com imagem: %s\n", 
				loja.Nome, loja.ID, loja.Imagem)
		}
	}

	// 2. Atualizar imagens dos anúncios
	fmt.Println("\n2. Atualizando imagens dos anúncios...")
	var anuncios []models.Anuncio
	err = database.DB.Where("data_exclusao IS NULL").Find(&anuncios).Error
	if err != nil {
		log.Printf("❌ Erro ao buscar anúncios: %v", err)
		return
	}

	for _, anuncio := range anuncios {
		anuncio.Imagem = anuncioImage

		err := database.DB.Model(&anuncio).Update("imagem", anuncio.Imagem).Error
		if err != nil {
			log.Printf("❌ Erro ao atualizar anúncio %d: %v", anuncio.ID, err)
		} else {
			fmt.Printf("✅ Anúncio '%s' (ID: %d) atualizado com imagem: %s\n", 
				anuncio.Titulo, anuncio.ID, anuncio.Imagem)
		}
	}

	fmt.Println("\n✅ Atualização de imagens concluída!")
}
