package seeds

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"meu-carro-mais/internal/database"
	"meu-carro-mais/internal/database/models"
)

// Seeder estrutura para gerenciar seeds
type Seeder struct {
	db *gorm.DB
}

// NewSeeder cria uma nova instância do seeder
func NewSeeder() *Seeder {
	return &Seeder{
		db: database.GetDB(),
	}
}

// Run executa todos os seeds
func (s *Seeder) Run() error {
	log.Println("🌱 Iniciando seeds...")

	// Executa os seeds em ordem de dependência
	if err := s.seedTipoPlano(); err != nil {
		return fmt.Errorf("erro ao executar seed tipo_plano: %v", err)
	}

	if err := s.seedCategoriaLojista(); err != nil {
		return fmt.Errorf("erro ao executar seed categoria_lojista: %v", err)
	}

	if err := s.seedCategoriaAnuncio(); err != nil {
		return fmt.Errorf("erro ao executar seed categoria_anuncio: %v", err)
	}

	if err := s.seedCategoriaServico(); err != nil {
		return fmt.Errorf("erro ao executar seed categoria_servico: %v", err)
	}

	if err := s.seedUsuario(); err != nil {
		return fmt.Errorf("erro ao executar seed usuario: %v", err)
	}

	if err := s.updateUsuarioCarlos(); err != nil {
		return fmt.Errorf("erro ao atualizar usuario carlos: %v", err)
	}

	if err := s.seedLoja(); err != nil {
		return fmt.Errorf("erro ao executar seed loja: %v", err)
	}

	if err := s.seedAnuncio(); err != nil {
		return fmt.Errorf("erro ao executar seed anuncio: %v", err)
	}

	if err := s.seedServico(); err != nil {
		return fmt.Errorf("erro ao executar seed servico: %v", err)
	}

	if err := s.seedCarteira(); err != nil {
		return fmt.Errorf("erro ao executar seed carteira: %v", err)
	}

	log.Println("✅ Seeds executados com sucesso!")
	return nil
}

// seedTipoPlano popula a tabela tipo_plano
func (s *Seeder) seedTipoPlano() error {
	log.Println("📝 Populando tabela tipo_plano...")

	tiposPlano := []models.TipoPlano{
		{Nome: "Gratuito"},
		{Nome: "Básico"},
		{Nome: "Premium"},
		{Nome: "Enterprise"},
	}

	for _, tipo := range tiposPlano {
		var existing models.TipoPlano
		if err := s.db.Where("nome = ?", tipo.Nome).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&tipo).Error; err != nil {
				return fmt.Errorf("erro ao criar tipo_plano %s: %v", tipo.Nome, err)
			}
			log.Printf("✅ TipoPlano criado: %s", tipo.Nome)
		} else {
			log.Printf("⏭️ TipoPlano já existe: %s", tipo.Nome)
		}
	}

	return nil
}

// seedCategoriaLojista popula a tabela categoria_lojista
func (s *Seeder) seedCategoriaLojista() error {
	log.Println("📝 Populando tabela categoria_lojista...")

	categorias := []models.CategoriaLojista{
		{Nome: "Concessionária"},
		{Nome: "Loja de Peças"},
		{Nome: "Oficina Mecânica"},
		{Nome: "Lava Jato"},
		{Nome: "Seguros"},
		{Nome: "Financiamento"},
	}

	for _, categoria := range categorias {
		var existing models.CategoriaLojista
		if err := s.db.Where("nome = ?", categoria.Nome).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&categoria).Error; err != nil {
				return fmt.Errorf("erro ao criar categoria_lojista %s: %v", categoria.Nome, err)
			}
			log.Printf("✅ CategoriaLojista criada: %s", categoria.Nome)
		} else {
			log.Printf("⏭️ CategoriaLojista já existe: %s", categoria.Nome)
		}
	}

	return nil
}

// seedCategoriaAnuncio popula a tabela categoria_anuncio
func (s *Seeder) seedCategoriaAnuncio() error {
	log.Println("📝 Populando tabela categoria_anuncio...")

	categorias := []models.CategoriaAnuncio{
		{Nome: "Carros"},
		{Nome: "Motos"},
		{Nome: "Caminhões"},
		{Nome: "Peças"},
		{Nome: "Serviços"},
		{Nome: "Acessórios"},
	}

	for _, categoria := range categorias {
		var existing models.CategoriaAnuncio
		if err := s.db.Where("nome = ?", categoria.Nome).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&categoria).Error; err != nil {
				return fmt.Errorf("erro ao criar categoria_anuncio %s: %v", categoria.Nome, err)
			}
			log.Printf("✅ CategoriaAnuncio criada: %s", categoria.Nome)
		} else {
			log.Printf("⏭️ CategoriaAnuncio já existe: %s", categoria.Nome)
		}
	}

	return nil
}

// seedCategoriaServico popula a tabela categoria_servico
func (s *Seeder) seedCategoriaServico() error {
	log.Println("📝 Populando tabela categoria_servico...")

	categorias := []models.CategoriaServico{
		{Nome: "Manutenção"},
		{Nome: "Revisão"},
		{Nome: "Troca de Óleo"},
		{Nome: "Alinhamento"},
		{Nome: "Balanceamento"},
		{Nome: "Ajuste de Freios"},
	}

	for _, categoria := range categorias {
		var existing models.CategoriaServico
		if err := s.db.Where("nome = ?", categoria.Nome).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&categoria).Error; err != nil {
				return fmt.Errorf("erro ao criar categoria_servico %s: %v", categoria.Nome, err)
			}
			log.Printf("✅ CategoriaServico criada: %s", categoria.Nome)
		} else {
			log.Printf("⏭️ CategoriaServico já existe: %s", categoria.Nome)
		}
	}

	return nil
}

// seedUsuario popula a tabela usuario
func (s *Seeder) seedUsuario() error {
	log.Println("📝 Populando tabela usuario...")

	// Busca o plano básico (ID = 2)
	var planoBásico models.TipoPlano
	if err := s.db.Where("nome = ?", "Básico").First(&planoBásico).Error; err != nil {
		return fmt.Errorf("erro ao buscar plano básico: %v", err)
	}

	usuarios := []models.Usuario{
		{
			Nome:     "João Silva",
			Email:    "joao@email.com",
			Senha:    "senha123",
			CPF:      "12345678901",
			Imagem:   "https://via.placeholder.com/150",
			Telefone: "(11) 99999-9999",
			Endereco: "Rua das Flores, 123 - São Paulo/SP",
			Ativo:    true,
			IDPlano:  planoBásico.ID,
		},
		{
			Nome:     "Maria Santos",
			Email:    "maria@email.com",
			Senha:    "senha123",
			CPF:      "98765432100",
			Imagem:   "https://via.placeholder.com/150",
			Telefone: "(11) 88888-8888",
			Endereco: "Av. Paulista, 1000 - São Paulo/SP",
			Ativo:    true,
			IDPlano:  planoBásico.ID,
		},
		{
			Nome:     "Pedro Costa",
			Email:    "pedro@email.com",
			Senha:    "senha123",
			CPF:      "11122233344",
			Imagem:   "https://via.placeholder.com/150",
			Telefone: "(11) 77777-7777",
			Endereco: "Rua Augusta, 500 - São Paulo/SP",
			Ativo:    true,
			IDPlano:  planoBásico.ID,
		},
		{
			Nome:      "Carlos Porto Alegre",
			Email:     "carlos@email.com",
			Senha:     "senha123",
			CPF:       "55566677788",
			Imagem:    "https://via.placeholder.com/150",
			Telefone:  "(51) 99999-9999",
			Endereco:  "Av. Borges de Medeiros, 500 - Porto Alegre/RS",
			Ativo:     true,
			Latitude:  &[]float64{-29.99627328048075}[0],
			Longitude: &[]float64{-51.14104068890408}[0],
			IDPlano:   planoBásico.ID,
		},
	}

	for _, usuario := range usuarios {
		var existing models.Usuario
		if err := s.db.Where("email = ?", usuario.Email).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&usuario).Error; err != nil {
				return fmt.Errorf("erro ao criar usuario %s: %v", usuario.Email, err)
			}
			log.Printf("✅ Usuario criado: %s (%s)", usuario.Nome, usuario.Email)
		} else {
			log.Printf("⏭️ Usuario já existe: %s (%s)", usuario.Nome, usuario.Email)
		}
	}

	return nil
}

// updateUsuarioCarlos atualiza as coordenadas do usuário Carlos Porto Alegre
func (s *Seeder) updateUsuarioCarlos() error {
	log.Println("📝 Atualizando coordenadas do usuário Carlos Porto Alegre...")

	var usuario models.Usuario
	if err := s.db.Where("email = ?", "carlos@email.com").First(&usuario).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuario Carlos Porto Alegre: %v", err)
	}

	usuario.Latitude = &[]float64{-29.99627328048075}[0]
	usuario.Longitude = &[]float64{-51.14104068890408}[0]

	if err := s.db.Save(&usuario).Error; err != nil {
		return fmt.Errorf("erro ao atualizar coordenadas do usuario Carlos Porto Alegre: %v", err)
	}

	log.Printf("✅ Coordenadas do usuário Carlos Porto Alegre atualizadas.")
	return nil
}

// seedLoja popula a tabela loja
func (s *Seeder) seedLoja() error {
	log.Println("📝 Populando tabela loja...")

	// Busca categoria concessionária
	var categoriaConcessionaria models.CategoriaLojista
	if err := s.db.Where("nome = ?", "Concessionária").First(&categoriaConcessionaria).Error; err != nil {
		return fmt.Errorf("erro ao buscar categoria concessionária: %v", err)
	}

	// Busca categoria oficina
	var categoriaOficina models.CategoriaLojista
	if err := s.db.Where("nome = ?", "Oficina Mecânica").First(&categoriaOficina).Error; err != nil {
		return fmt.Errorf("erro ao buscar categoria oficina: %v", err)
	}

	lojas := []models.Loja{
		{
			Nome:        "Auto Center São Paulo",
			CNPJ:        "12.345.678/0001-90",
			Imagem:      "https://via.placeholder.com/300x200",
			Latitude:    -23.5505,
			Longitude:   -46.6333,
			IDCategoria: categoriaConcessionaria.ID,
		},
		{
			Nome:        "Oficina do João",
			CNPJ:        "98.765.432/0001-10",
			Imagem:      "https://via.placeholder.com/300x200",
			Latitude:    -23.5489,
			Longitude:   -46.6388,
			IDCategoria: categoriaOficina.ID,
		},
		{
			Nome:        "Carros Premium",
			CNPJ:        "11.222.333/0001-44",
			Imagem:      "https://via.placeholder.com/300x200",
			Latitude:    -23.5520,
			Longitude:   -46.6310,
			IDCategoria: categoriaConcessionaria.ID,
		},
		{
			Nome:        "Auto Center Porto Alegre Norte",
			CNPJ:        "22.333.444/0001-55",
			Imagem:      "https://via.placeholder.com/300x200",
			Latitude:    -30.01392952707345,
			Longitude:   -51.11168643659632,
			IDCategoria: categoriaConcessionaria.ID,
		},
		{
			Nome:        "Oficina Porto Alegre Centro",
			CNPJ:        "33.444.555/0001-66",
			Imagem:      "https://via.placeholder.com/300x200",
			Latitude:    -30.03004986742164,
			Longitude:   -51.12161634814089,
			IDCategoria: categoriaOficina.ID,
		},
		{
			Nome:        "Loja de Peças Porto Alegre",
			CNPJ:        "44.555.666/0001-77",
			Imagem:      "https://via.placeholder.com/300x200",
			Latitude:    -30.02279847167631,
			Longitude:   -51.20306428600608,
			IDCategoria: categoriaOficina.ID,
		},
		{
			Nome:        "Concessionária Porto Alegre Sul",
			CNPJ:        "55.666.777/0001-88",
			Imagem:      "https://via.placeholder.com/300x200",
			Latitude:    -30.04486628735914,
			Longitude:   -51.22406353079101,
			IDCategoria: categoriaConcessionaria.ID,
		},
	}

	for _, loja := range lojas {
		var existing models.Loja
		if err := s.db.Where("cnpj = ?", loja.CNPJ).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&loja).Error; err != nil {
				return fmt.Errorf("erro ao criar loja %s: %v", loja.Nome, err)
			}
			log.Printf("✅ Loja criada: %s", loja.Nome)
		} else {
			log.Printf("⏭️ Loja já existe: %s", loja.Nome)
		}
	}

	return nil
}

// seedAnuncio popula a tabela anuncio
func (s *Seeder) seedAnuncio() error {
	log.Println("📝 Populando tabela anuncio...")

	// Busca categoria carros
	var categoriaCarros models.CategoriaAnuncio
	if err := s.db.Where("nome = ?", "Carros").First(&categoriaCarros).Error; err != nil {
		return fmt.Errorf("erro ao buscar categoria carros: %v", err)
	}

	// Busca primeira loja
	var loja models.Loja
	if err := s.db.First(&loja).Error; err != nil {
		return fmt.Errorf("erro ao buscar loja: %v", err)
	}

	anuncios := []models.Anuncio{
		{
			Titulo:      "Honda Civic 2020",
			Descricao:   "Honda Civic EXL 2.0, automático, completo, único dono",
			Preco:       85000.00,
			Imagem:      "https://via.placeholder.com/400x300",
			Destaque:    true,
			IDLoja:      loja.ID,
			IDCategoria: categoriaCarros.ID,
		},
		{
			Titulo:      "Toyota Corolla 2019",
			Descricao:   "Toyota Corolla XEi 2.0, automático, revisões em dia",
			Preco:       75000.00,
			Imagem:      "https://via.placeholder.com/400x300",
			Destaque:    false,
			IDLoja:      loja.ID,
			IDCategoria: categoriaCarros.ID,
		},
		{
			Titulo:      "Volkswagen Golf GTI",
			Descricao:   "Golf GTI 2.0 TSI, manual, esportivo, baixa quilometragem",
			Preco:       95000.00,
			Imagem:      "https://via.placeholder.com/400x300",
			Destaque:    true,
			IDLoja:      loja.ID,
			IDCategoria: categoriaCarros.ID,
		},
	}

	for _, anuncio := range anuncios {
		var existing models.Anuncio
		if err := s.db.Where("titulo = ? AND id_loja = ?", anuncio.Titulo, anuncio.IDLoja).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&anuncio).Error; err != nil {
				return fmt.Errorf("erro ao criar anuncio %s: %v", anuncio.Titulo, err)
			}
			log.Printf("✅ Anuncio criado: %s", anuncio.Titulo)
		} else {
			log.Printf("⏭️ Anuncio já existe: %s", anuncio.Titulo)
		}
	}

	return nil
}

// seedServico popula a tabela servico
func (s *Seeder) seedServico() error {
	log.Println("📝 Populando tabela servico...")

	// Busca primeira loja
	var loja models.Loja
	if err := s.db.First(&loja).Error; err != nil {
		return fmt.Errorf("erro ao buscar loja: %v", err)
	}

	// Busca primeira categoria de serviço
	var categoriaServico models.CategoriaServico
	if err := s.db.First(&categoriaServico).Error; err != nil {
		return fmt.Errorf("erro ao buscar categoria de serviço: %v", err)
	}

	servicos := []models.Servico{
		{
			Titulo:      "Troca de Óleo",
			Descricao:   "Troca de óleo, filtro e óleo de filtro",
			Preco:       150.00,
			Imagem:      "https://via.placeholder.com/200x150",
			IDLoja:      loja.ID,
			IDCategoria: categoriaServico.ID,
		},
		{
			Titulo:      "Alinhamento",
			Descricao:   "Ajuste de direção, suspensão e direção",
			Preco:       250.00,
			Imagem:      "https://via.placeholder.com/200x150",
			IDLoja:      loja.ID,
			IDCategoria: categoriaServico.ID,
		},
		{
			Titulo:      "Revisão Completa",
			Descricao:   "Revisão completa do motor, freios, suspensão e direção",
			Preco:       500.00,
			Imagem:      "https://via.placeholder.com/200x150",
			IDLoja:      loja.ID,
			IDCategoria: categoriaServico.ID,
		},
	}

	for _, servico := range servicos {
		var existing models.Servico
		if err := s.db.Where("titulo = ? AND id_loja = ?", servico.Titulo, servico.IDLoja).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&servico).Error; err != nil {
				return fmt.Errorf("erro ao criar servico %s: %v", servico.Titulo, err)
			}
			log.Printf("✅ Servico criado: %s", servico.Titulo)
		} else {
			log.Printf("⏭️ Servico já existe: %s", servico.Titulo)
		}
	}

	return nil
}

// seedCarteira popula a tabela carteira
func (s *Seeder) seedCarteira() error {
	log.Println("📝 Populando tabela carteira...")

	// Busca primeiro usuário
	var usuario models.Usuario
	if err := s.db.First(&usuario).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuario: %v", err)
	}

	carteiras := []models.Carteira{
		{
			UsuarioID: usuario.ID,
			Saldo:     1000.00,
		},
	}

	for _, carteira := range carteiras {
		var existing models.Carteira
		if err := s.db.Where("usuario_id = ?", carteira.UsuarioID).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&carteira).Error; err != nil {
				return fmt.Errorf("erro ao criar carteira para usuario %d: %v", carteira.UsuarioID, err)
			}
			log.Printf("✅ Carteira criada para usuario ID: %d", carteira.UsuarioID)
		} else {
			log.Printf("⏭️ Carteira já existe para usuario ID: %d", carteira.UsuarioID)
		}
	}

	return nil
}
