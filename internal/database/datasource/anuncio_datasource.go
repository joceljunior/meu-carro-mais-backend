package datasource

import (
	"errors"
	"fmt"
	"math"
	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
	"meu-carro-mais/internal/handlers/json"
	"sort"
	"time"
)

func GetAnuncioDestaqueByLojaID(lojaID uint) (*models.Anuncio, error) {
	var anuncio models.Anuncio
	err := database.DB.Where("id_loja = ? AND destaque = ? AND data_exclusao IS NULL", lojaID, true).First(&anuncio).Error
	if err != nil {
		return nil, err
	}
	return &anuncio, nil
}

// GetAnuncios retorna todos os anúncios com relacionamentos
func GetAnuncios() ([]models.Anuncio, error) {
	var anuncios []models.Anuncio

	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("OfertaAutoMais").
		Where("data_exclusao IS NULL").
		Find(&anuncios).Error

	if err != nil {
		return nil, err
	}

	return anuncios, nil
}

// removeDestaqueDeOutrosAnuncios remove o destaque de todos os outros anúncios da mesma loja
// Garante que apenas um anúncio por loja seja destaque
func removeDestaqueDeOutrosAnuncios(idLoja uint, idAnuncioExcluir uint) error {
	return database.DB.Model(&models.Anuncio{}).
		Where("id_loja = ? AND destaque = ? AND data_exclusao IS NULL AND id != ?", idLoja, true, idAnuncioExcluir).
		Update("destaque", false).Error
}

// CreateAnuncio cria um novo anúncio
func CreateAnuncio(req json.AnuncioRequest) (*models.Anuncio, error) {
	// Se o anúncio será criado como destaque e tem loja, remove o destaque de outros anúncios da mesma loja
	// Garante que apenas um anúncio por loja seja destaque
	if req.Destaque && req.IDLoja != nil && *req.IDLoja > 0 {
		if err := removeDestaqueDeOutrosAnuncios(*req.IDLoja, 0); err != nil {
			return nil, fmt.Errorf("erro ao remover destaque de outros anúncios: %v", err)
		}
	}

	// Calcula o preço com desconto se a porcentagem for fornecida mas o preço com desconto não
	precoComDesconto := req.PrecoComDesconto
	if precoComDesconto == 0 && req.PorcentagemDesconto > 0 {
		// Busca o preço original do produto/serviço/veículo
		var precoOriginal float64
		if req.IDProduto != nil {
			produto, err := GetProdutoByID(*req.IDProduto)
			if err == nil {
				precoOriginal = produto.Preco
			} else {
				precoOriginal = req.Preco
			}
		} else if req.IDServico != nil {
			servico, err := GetServicoByID(*req.IDServico)
			if err == nil {
				precoOriginal = servico.Preco
			} else {
				precoOriginal = req.Preco
			}
		} else {
			// Para veículos, usa o preço do anúncio como original
			precoOriginal = req.Preco
		}
		precoComDesconto = precoOriginal * (1 - req.PorcentagemDesconto/100)
	} else if precoComDesconto == 0 {
		// Se não houver desconto, o preço com desconto é igual ao preço original
		if req.IDProduto != nil {
			produto, err := GetProdutoByID(*req.IDProduto)
			if err == nil {
				precoComDesconto = produto.Preco
			} else {
				precoComDesconto = req.Preco
			}
		} else if req.IDServico != nil {
			servico, err := GetServicoByID(*req.IDServico)
			if err == nil {
				precoComDesconto = servico.Preco
			} else {
				precoComDesconto = req.Preco
			}
		} else {
			precoComDesconto = req.Preco
		}
	}

	anuncio := models.Anuncio{
		Titulo:             req.Titulo,
		Descricao:          req.Descricao,
		Preco:              req.Preco,
		Imagem:             req.Imagem,
		Destaque:           req.Destaque,
		Categoria:          req.Categoria,
		IDLoja:             req.IDLoja,
		IDProduto:          req.IDProduto,
		IDServico:          req.IDServico,
		IDVeiculo:          req.IDVeiculo,
		IDOfertaAutoMais:   req.IDOfertaAutoMais,
		TipoAnuncio:        req.TipoAnuncio,
		PorcentagemDesconto: req.PorcentagemDesconto,
		PrecoComDesconto:   precoComDesconto,
	}

	err := database.DB.Create(&anuncio).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o anúncio com os relacionamentos
	return GetAnuncioByID(anuncio.ID)
}

// GetAnuncioByID busca anúncio por ID (apenas anúncios não excluídos)
func GetAnuncioByID(id uint) (*models.Anuncio, error) {
	var anuncio models.Anuncio
	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("OfertaAutoMais").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&anuncio).Error
	if err != nil {
		return nil, err
	}
	return &anuncio, nil
}

// GetAllAnuncios retorna todos os anúncios ativos (não excluídos)
func GetAllAnuncios() ([]models.Anuncio, error) {
	var anuncios []models.Anuncio
	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("OfertaAutoMais").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&anuncios).Error
	if err != nil {
		return nil, err
	}
	return anuncios, nil
}

// GetAnunciosByLojaID retorna todos os anúncios de uma loja específica
func GetAnunciosByLojaID(lojaID uint) ([]models.Anuncio, error) {
	var anuncios []models.Anuncio
	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("OfertaAutoMais").
		Where("id_loja = ? AND data_exclusao IS NULL", lojaID).
		Order("destaque DESC, data_cadastro DESC").
		Find(&anuncios).Error
	if err != nil {
		return nil, err
	}
	return anuncios, nil
}

// UpdateAnuncio atualiza um anúncio existente
func UpdateAnuncio(id uint, req json.AnuncioRequest) (*models.Anuncio, error) {
	// Verifica se o anúncio existe e não foi excluído
	anuncio, err := GetAnuncioByID(id)
	if err != nil {
		return nil, errors.New("anúncio não encontrado")
	}

	// Se o anúncio será atualizado como destaque e tem loja, remove o destaque de outros anúncios da mesma loja
	// Garante que apenas um anúncio por loja seja destaque
	if req.Destaque && req.IDLoja != nil && *req.IDLoja > 0 {
		// Sempre remove destaque dos outros anúncios da loja (seja loja nova ou antiga)
		// Isso garante que apenas um anúncio por loja seja destaque
		if err := removeDestaqueDeOutrosAnuncios(*req.IDLoja, id); err != nil {
			return nil, fmt.Errorf("erro ao remover destaque de outros anúncios: %v", err)
		}
	}

	// Calcula o preço com desconto se a porcentagem for fornecida mas o preço com desconto não
	precoComDesconto := req.PrecoComDesconto
	if precoComDesconto == 0 && req.PorcentagemDesconto > 0 {
		// Busca o preço original do produto/serviço/veículo
		var precoOriginal float64
		if req.IDProduto != nil {
			produto, err := GetProdutoByID(*req.IDProduto)
			if err == nil {
				precoOriginal = produto.Preco
			} else {
				precoOriginal = req.Preco
			}
		} else if req.IDServico != nil {
			servico, err := GetServicoByID(*req.IDServico)
			if err == nil {
				precoOriginal = servico.Preco
			} else {
				precoOriginal = req.Preco
			}
		} else {
			// Para veículos, usa o preço do anúncio como original
			precoOriginal = req.Preco
		}
		precoComDesconto = precoOriginal * (1 - req.PorcentagemDesconto/100)
	} else if precoComDesconto == 0 {
		// Se não houver desconto, o preço com desconto é igual ao preço original
		if req.IDProduto != nil {
			produto, err := GetProdutoByID(*req.IDProduto)
			if err == nil {
				precoComDesconto = produto.Preco
			} else {
				precoComDesconto = req.Preco
			}
		} else if req.IDServico != nil {
			servico, err := GetServicoByID(*req.IDServico)
			if err == nil {
				precoComDesconto = servico.Preco
			} else {
				precoComDesconto = req.Preco
			}
		} else {
			precoComDesconto = req.Preco
		}
	}

	// Atualiza os campos
	anuncio.Titulo = req.Titulo
	anuncio.Descricao = req.Descricao
	anuncio.Preco = req.Preco
	anuncio.Imagem = req.Imagem
	anuncio.Destaque = req.Destaque
	anuncio.Categoria = req.Categoria
	anuncio.IDLoja = req.IDLoja
	anuncio.IDProduto = req.IDProduto
	anuncio.IDServico = req.IDServico
	anuncio.IDVeiculo = req.IDVeiculo
	anuncio.IDOfertaAutoMais = req.IDOfertaAutoMais
	anuncio.TipoAnuncio = req.TipoAnuncio
	anuncio.PorcentagemDesconto = req.PorcentagemDesconto
	anuncio.PrecoComDesconto = precoComDesconto

	err = database.DB.Save(&anuncio).Error
	if err != nil {
		return nil, err
	}

	// Recarrega o anúncio com os relacionamentos
	return GetAnuncioByID(id)
}

// SoftDeleteAnuncio realiza soft delete do anúncio (marca como excluído)
func SoftDeleteAnuncio(id uint) error {
	// Verifica se o anúncio existe e não foi excluído
	_, err := GetAnuncioByID(id)
	if err != nil {
		return errors.New("anúncio não encontrado")
	}

	// Atualiza a data de exclusão
	now := time.Now()
	err = database.DB.Model(&models.Anuncio{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreAnuncio restaura um anúncio que foi soft deleted
func RestoreAnuncio(id uint) error {
	var anuncio models.Anuncio
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&anuncio).Error
	if err != nil {
		return errors.New("anúncio não encontrado ou não foi excluído")
	}

	// Remove a data de exclusão
	err = database.DB.Model(&models.Anuncio{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}

// GetAnunciosProdutos retorna todos os anúncios de produtos ativos
func GetAnunciosProdutos() ([]models.Anuncio, error) {
	var anuncios []models.Anuncio
	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("OfertaAutoMais").
		Where("tipo_anuncio = ? AND data_exclusao IS NULL", "produto").
		Order("data_cadastro DESC").
		Find(&anuncios).Error
	if err != nil {
		return nil, err
	}
	return anuncios, nil
}

// AnuncioProdutoComDistancia representa um anúncio de produto com sua distância calculada
type AnuncioProdutoComDistancia struct {
	Anuncio   models.Anuncio
	Distancia float64
}

// GetAnunciosProdutosByProximidade retorna anúncios de produtos ordenados por proximidade
func GetAnunciosProdutosByProximidade(latitude, longitude float64) ([]AnuncioProdutoComDistancia, error) {
	anuncios, err := GetAnunciosProdutos()
	if err != nil {
		return nil, err
	}

	var anunciosComDistancia []AnuncioProdutoComDistancia
	for _, anuncio := range anuncios {
		// Calcula a distância usando a fórmula de Haversine
		distancia := calcularDistanciaAnuncio(latitude, longitude, anuncio.Loja.Latitude, anuncio.Loja.Longitude)
		anunciosComDistancia = append(anunciosComDistancia, AnuncioProdutoComDistancia{
			Anuncio:   anuncio,
			Distancia: distancia,
		})
	}

	// Ordena por distância (menor primeiro)
	sort.Slice(anunciosComDistancia, func(i, j int) bool {
		return anunciosComDistancia[i].Distancia < anunciosComDistancia[j].Distancia
	})

	return anunciosComDistancia, nil
}

// GetAnunciosServicos retorna todos os anúncios de serviços ativos
func GetAnunciosServicos() ([]models.Anuncio, error) {
	var anuncios []models.Anuncio
	err := database.DB.
		Preload("Loja").
		Preload("Servico").
		Preload("OfertaAutoMais").
		Where("tipo_anuncio = ? AND data_exclusao IS NULL", "servico").
		Order("data_cadastro DESC").
		Find(&anuncios).Error
	if err != nil {
		return nil, err
	}
	return anuncios, nil
}

// AnuncioServicoComDistancia representa um anúncio de serviço com sua distância calculada
type AnuncioServicoComDistancia struct {
	Anuncio   models.Anuncio
	Distancia float64
}

// GetAnunciosServicosByProximidade retorna anúncios de serviços ordenados por proximidade
func GetAnunciosServicosByProximidade(latitude, longitude float64) ([]AnuncioServicoComDistancia, error) {
	anuncios, err := GetAnunciosServicos()
	if err != nil {
		return nil, err
	}

	var anunciosComDistancia []AnuncioServicoComDistancia
	for _, anuncio := range anuncios {
		// Calcula a distância usando a fórmula de Haversine
		distancia := calcularDistanciaAnuncio(latitude, longitude, anuncio.Loja.Latitude, anuncio.Loja.Longitude)
		anunciosComDistancia = append(anunciosComDistancia, AnuncioServicoComDistancia{
			Anuncio:   anuncio,
			Distancia: distancia,
		})
	}

	// Ordena por distância (menor primeiro)
	sort.Slice(anunciosComDistancia, func(i, j int) bool {
		return anunciosComDistancia[i].Distancia < anunciosComDistancia[j].Distancia
	})

	return anunciosComDistancia, nil
}

// calcularDistanciaAnuncio calcula a distância entre dois pontos usando a fórmula de Haversine
func calcularDistanciaAnuncio(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371 // Raio da Terra em km

	// Converte para radianos
	lat1Rad := lat1 * math.Pi / 180
	lng1Rad := lng1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lng2Rad := lng2 * math.Pi / 180

	// Diferenças
	dlat := lat2Rad - lat1Rad
	dlng := lng2Rad - lng1Rad

	// Fórmula de Haversine
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// GetAnunciosVeiculos retorna todos os anúncios de veículos ativos
func GetAnunciosVeiculos() ([]models.Anuncio, error) {
	var anuncios []models.Anuncio
	err := database.DB.
		Preload("Loja").
		Preload("Veiculo").
		Preload("Veiculo.Usuario"). // Preload do usuário dono do veículo
		Preload("OfertaAutoMais").
		Where("tipo_anuncio = ? AND data_exclusao IS NULL", "veiculo").
		Order("data_cadastro DESC").
		Find(&anuncios).Error
	if err != nil {
		return nil, err
	}
	return anuncios, nil
}

// AnuncioVeiculoComDistancia representa um anúncio de veículo com sua distância calculada
type AnuncioVeiculoComDistancia struct {
	Anuncio   models.Anuncio
	Distancia float64
}

// GetAnunciosVeiculosByProximidade retorna anúncios de veículos ordenados por proximidade
func GetAnunciosVeiculosByProximidade(latitude, longitude float64) ([]AnuncioVeiculoComDistancia, error) {
	anuncios, err := GetAnunciosVeiculos()
	if err != nil {
		return nil, err
	}

	var anunciosComDistancia []AnuncioVeiculoComDistancia
	for _, anuncio := range anuncios {
		var distancia float64
		// Se o anúncio tem loja, calcula a distância
		// Caso contrário (veículo de usuário), usa a localização do usuário dono do veículo
		if anuncio.Loja != nil {
			distancia = calcularDistanciaAnuncio(latitude, longitude, anuncio.Loja.Latitude, anuncio.Loja.Longitude)
		} else if anuncio.Veiculo != nil && anuncio.Veiculo.Usuario.ID != 0 &&
			anuncio.Veiculo.Usuario.Latitude != nil && anuncio.Veiculo.Usuario.Longitude != nil {
			// Usa a localização do usuário dono do veículo
			distancia = calcularDistanciaAnuncio(latitude, longitude, *anuncio.Veiculo.Usuario.Latitude, *anuncio.Veiculo.Usuario.Longitude)
		} else {
			// Sem localização disponível, coloca no final (distância muito grande)
			distancia = 999999
		}
		
		anunciosComDistancia = append(anunciosComDistancia, AnuncioVeiculoComDistancia{
			Anuncio:   anuncio,
			Distancia: distancia,
		})
	}

	// Ordena por distância (menor primeiro)
	sort.Slice(anunciosComDistancia, func(i, j int) bool {
		return anunciosComDistancia[i].Distancia < anunciosComDistancia[j].Distancia
	})

	return anunciosComDistancia, nil
}

// GetAnuncioByVeiculoID busca anúncio ativo por ID do veículo
func GetAnuncioByVeiculoID(veiculoID uint) (*models.Anuncio, error) {
	var anuncio models.Anuncio
	err := database.DB.
		Where("id_veiculo = ? AND data_exclusao IS NULL", veiculoID).
		First(&anuncio).Error
	if err != nil {
		return nil, err
	}
	return &anuncio, nil
}