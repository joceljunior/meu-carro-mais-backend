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

func GetCupomDestaqueByLojaID(lojaID uint) (*models.Cupom, error) {
	var cupom models.Cupom
	err := database.DB.Where("id_loja = ? AND destaque = ? AND data_exclusao IS NULL", lojaID, true).First(&cupom).Error
	if err != nil {
		return nil, err
	}
	return &cupom, nil
}

// GetCupons retorna todos os cupons com relacionamentos
func GetCupons() ([]models.Cupom, error) {
	var cupons []models.Cupom

	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("OfertaAutoMais").
		Where("data_exclusao IS NULL").
		Find(&cupons).Error

	if err != nil {
		return nil, err
	}

	return cupons, nil
}

// removeDestaqueDeOutrosCupons remove o destaque de todos os outros cupons da mesma loja
func removeDestaqueDeOutrosCupons(idLoja uint, idCupomExcluir uint) error {
	return database.DB.Model(&models.Cupom{}).
		Where("id_loja = ? AND destaque = ? AND data_exclusao IS NULL AND id != ?", idLoja, true, idCupomExcluir).
		Update("destaque", false).Error
}

// CreateCupom cria um novo cupom
func CreateCupom(req json.CupomRequest) (*models.Cupom, error) {
	if req.Destaque && req.IDLoja != nil && *req.IDLoja > 0 {
		if err := removeDestaqueDeOutrosCupons(*req.IDLoja, 0); err != nil {
			return nil, fmt.Errorf("erro ao remover destaque de outros cupons: %v", err)
		}
	}

	precoComDesconto := req.PrecoComDesconto
	if precoComDesconto == 0 && req.PorcentagemDesconto > 0 {
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
			precoOriginal = req.Preco
		}
		precoComDesconto = precoOriginal * (1 - req.PorcentagemDesconto/100)
	} else if precoComDesconto == 0 {
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

	cupom := models.Cupom{
		Titulo:              req.Titulo,
		Descricao:           req.Descricao,
		Preco:               req.Preco,
		Imagem:              req.Imagem,
		Destaque:            req.Destaque,
		Categoria:           req.Categoria,
		IDLoja:              req.IDLoja,
		IDProduto:           req.IDProduto,
		IDServico:           req.IDServico,
		IDVeiculo:           req.IDVeiculo,
		IDOfertaAutoMais:    req.IDOfertaAutoMais,
		TipoCupom:           req.TipoCupom,
		PorcentagemDesconto: req.PorcentagemDesconto,
		PrecoComDesconto:    precoComDesconto,
	}

	err := database.DB.Create(&cupom).Error
	if err != nil {
		return nil, err
	}

	return GetCupomByID(cupom.ID)
}

// GetCupomByID busca cupom por ID (apenas cupons não excluídos)
func GetCupomByID(id uint) (*models.Cupom, error) {
	var cupom models.Cupom
	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("OfertaAutoMais").
		Where("id = ? AND data_exclusao IS NULL", id).
		First(&cupom).Error
	if err != nil {
		return nil, err
	}
	return &cupom, nil
}

// GetAllCupons retorna todos os cupons ativos (não excluídos)
func GetAllCupons() ([]models.Cupom, error) {
	var cupons []models.Cupom
	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("OfertaAutoMais").
		Where("data_exclusao IS NULL").
		Order("data_cadastro DESC").
		Find(&cupons).Error
	if err != nil {
		return nil, err
	}
	return cupons, nil
}

// GetCuponsByLojaID retorna todos os cupons de uma loja específica
func GetCuponsByLojaID(lojaID uint) ([]models.Cupom, error) {
	var cupons []models.Cupom
	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("Servico").
		Preload("Veiculo").
		Preload("OfertaAutoMais").
		Where("id_loja = ? AND data_exclusao IS NULL", lojaID).
		Order("destaque DESC, data_cadastro DESC").
		Find(&cupons).Error
	if err != nil {
		return nil, err
	}
	return cupons, nil
}

// UpdateCupom atualiza um cupom existente
func UpdateCupom(id uint, req json.CupomRequest) (*models.Cupom, error) {
	cupom, err := GetCupomByID(id)
	if err != nil {
		return nil, errors.New("cupom não encontrado")
	}

	if req.Destaque && req.IDLoja != nil && *req.IDLoja > 0 {
		if err := removeDestaqueDeOutrosCupons(*req.IDLoja, id); err != nil {
			return nil, fmt.Errorf("erro ao remover destaque de outros cupons: %v", err)
		}
	}

	precoComDesconto := req.PrecoComDesconto
	if precoComDesconto == 0 && req.PorcentagemDesconto > 0 {
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
			precoOriginal = req.Preco
		}
		precoComDesconto = precoOriginal * (1 - req.PorcentagemDesconto/100)
	} else if precoComDesconto == 0 {
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

	cupom.Titulo = req.Titulo
	cupom.Descricao = req.Descricao
	cupom.Preco = req.Preco
	cupom.Imagem = req.Imagem
	cupom.Destaque = req.Destaque
	cupom.Categoria = req.Categoria
	cupom.IDLoja = req.IDLoja
	cupom.IDProduto = req.IDProduto
	cupom.IDServico = req.IDServico
	cupom.IDVeiculo = req.IDVeiculo
	cupom.IDOfertaAutoMais = req.IDOfertaAutoMais
	cupom.TipoCupom = req.TipoCupom
	cupom.PorcentagemDesconto = req.PorcentagemDesconto
	cupom.PrecoComDesconto = precoComDesconto

	err = database.DB.Save(&cupom).Error
	if err != nil {
		return nil, err
	}

	return GetCupomByID(id)
}

// SoftDeleteCupom realiza soft delete do cupom
func SoftDeleteCupom(id uint) error {
	_, err := GetCupomByID(id)
	if err != nil {
		return errors.New("cupom não encontrado")
	}

	now := time.Now()
	err = database.DB.Model(&models.Cupom{}).
		Where("id = ?", id).
		Update("data_exclusao", now).Error
	if err != nil {
		return err
	}

	return nil
}

// RestoreCupom restaura um cupom que foi soft deleted
func RestoreCupom(id uint) error {
	var cupom models.Cupom
	err := database.DB.Where("id = ? AND data_exclusao IS NOT NULL", id).First(&cupom).Error
	if err != nil {
		return errors.New("cupom não encontrado ou não foi excluído")
	}

	err = database.DB.Model(&models.Cupom{}).
		Where("id = ?", id).
		Update("data_exclusao", nil).Error
	if err != nil {
		return err
	}

	return nil
}

// GetCuponsProdutos retorna todos os cupons de produtos ativos
func GetCuponsProdutos() ([]models.Cupom, error) {
	var cupons []models.Cupom
	err := database.DB.
		Preload("Loja").
		Preload("Produto").
		Preload("OfertaAutoMais").
		Where("tipo_cupom = ? AND data_exclusao IS NULL", "produto").
		Order("data_cadastro DESC").
		Find(&cupons).Error
	if err != nil {
		return nil, err
	}
	return cupons, nil
}

// CupomProdutoComDistancia representa um cupom de produto com sua distância calculada
type CupomProdutoComDistancia struct {
	Cupom     models.Cupom
	Distancia float64
}

// GetCuponsProdutosByProximidade retorna cupons de produtos ordenados por proximidade
func GetCuponsProdutosByProximidade(latitude, longitude float64) ([]CupomProdutoComDistancia, error) {
	cupons, err := GetCuponsProdutos()
	if err != nil {
		return nil, err
	}

	var cuponsComDistancia []CupomProdutoComDistancia
	for _, cupom := range cupons {
		distancia := calcularDistanciaCupom(latitude, longitude, cupom.Loja.Latitude, cupom.Loja.Longitude)
		cuponsComDistancia = append(cuponsComDistancia, CupomProdutoComDistancia{
			Cupom:     cupom,
			Distancia: distancia,
		})
	}

	sort.Slice(cuponsComDistancia, func(i, j int) bool {
		return cuponsComDistancia[i].Distancia < cuponsComDistancia[j].Distancia
	})

	return cuponsComDistancia, nil
}

// GetCuponsServicos retorna todos os cupons de serviços ativos
func GetCuponsServicos() ([]models.Cupom, error) {
	var cupons []models.Cupom
	err := database.DB.
		Preload("Loja").
		Preload("Servico").
		Preload("OfertaAutoMais").
		Where("tipo_cupom = ? AND data_exclusao IS NULL", "servico").
		Order("data_cadastro DESC").
		Find(&cupons).Error
	if err != nil {
		return nil, err
	}
	return cupons, nil
}

// CupomServicoComDistancia representa um cupom de serviço com sua distância calculada
type CupomServicoComDistancia struct {
	Cupom     models.Cupom
	Distancia float64
}

// GetCuponsServicosByProximidade retorna cupons de serviços ordenados por proximidade
func GetCuponsServicosByProximidade(latitude, longitude float64) ([]CupomServicoComDistancia, error) {
	cupons, err := GetCuponsServicos()
	if err != nil {
		return nil, err
	}

	var cuponsComDistancia []CupomServicoComDistancia
	for _, cupom := range cupons {
		distancia := calcularDistanciaCupom(latitude, longitude, cupom.Loja.Latitude, cupom.Loja.Longitude)
		cuponsComDistancia = append(cuponsComDistancia, CupomServicoComDistancia{
			Cupom:     cupom,
			Distancia: distancia,
		})
	}

	sort.Slice(cuponsComDistancia, func(i, j int) bool {
		return cuponsComDistancia[i].Distancia < cuponsComDistancia[j].Distancia
	})

	return cuponsComDistancia, nil
}

// calcularDistanciaCupom calcula a distância entre dois pontos usando a fórmula de Haversine
func calcularDistanciaCupom(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371

	lat1Rad := lat1 * math.Pi / 180
	lng1Rad := lng1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lng2Rad := lng2 * math.Pi / 180

	dlat := lat2Rad - lat1Rad
	dlng := lng2Rad - lng1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// GetCuponsVeiculos retorna todos os cupons de veículos ativos
func GetCuponsVeiculos() ([]models.Cupom, error) {
	var cupons []models.Cupom
	err := database.DB.
		Preload("Loja").
		Preload("Veiculo").
		Preload("Veiculo.Usuario").
		Preload("OfertaAutoMais").
		Where("tipo_cupom = ? AND data_exclusao IS NULL", "veiculo").
		Order("data_cadastro DESC").
		Find(&cupons).Error
	if err != nil {
		return nil, err
	}
	return cupons, nil
}

// CupomVeiculoComDistancia representa um cupom de veículo com sua distância calculada
type CupomVeiculoComDistancia struct {
	Cupom     models.Cupom
	Distancia float64
}

// GetCuponsVeiculosByProximidade retorna cupons de veículos ordenados por proximidade
func GetCuponsVeiculosByProximidade(latitude, longitude float64) ([]CupomVeiculoComDistancia, error) {
	cupons, err := GetCuponsVeiculos()
	if err != nil {
		return nil, err
	}

	var cuponsComDistancia []CupomVeiculoComDistancia
	for _, cupom := range cupons {
		var distancia float64
		if cupom.Loja != nil {
			distancia = calcularDistanciaCupom(latitude, longitude, cupom.Loja.Latitude, cupom.Loja.Longitude)
		} else if cupom.Veiculo != nil && cupom.Veiculo.Usuario.ID != 0 &&
			cupom.Veiculo.Usuario.Latitude != nil && cupom.Veiculo.Usuario.Longitude != nil {
			distancia = calcularDistanciaCupom(latitude, longitude, *cupom.Veiculo.Usuario.Latitude, *cupom.Veiculo.Usuario.Longitude)
		} else {
			distancia = 999999
		}

		cuponsComDistancia = append(cuponsComDistancia, CupomVeiculoComDistancia{
			Cupom:     cupom,
			Distancia: distancia,
		})
	}

	sort.Slice(cuponsComDistancia, func(i, j int) bool {
		return cuponsComDistancia[i].Distancia < cuponsComDistancia[j].Distancia
	})

	return cuponsComDistancia, nil
}

// GetCupomByVeiculoID busca cupom ativo por ID do veículo
func GetCupomByVeiculoID(veiculoID uint) (*models.Cupom, error) {
	var cupom models.Cupom
	err := database.DB.
		Where("id_veiculo = ? AND data_exclusao IS NULL", veiculoID).
		First(&cupom).Error
	if err != nil {
		return nil, err
	}
	return &cupom, nil
}
