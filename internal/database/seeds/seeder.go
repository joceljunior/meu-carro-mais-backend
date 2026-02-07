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

	if err := s.seedUsuario(); err != nil {
		return fmt.Errorf("erro ao executar seed usuario: %v", err)
	}

	if err := s.updateUsuarioCarlos(); err != nil {
		return fmt.Errorf("erro ao atualizar usuario carlos: %v", err)
	}

	if err := s.seedLoja(); err != nil {
		return fmt.Errorf("erro ao executar seed loja: %v", err)
	}

	if err := s.vincularUsuariosLojas(); err != nil {
		return fmt.Errorf("erro ao vincular usuarios às lojas: %v", err)
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

	if err := s.seedUsuarioComAnuncioDestaque(); err != nil {
		return fmt.Errorf("erro ao executar seed usuario com anuncio destaque: %v", err)
	}

	if err := s.seedProdutos(); err != nil {
		return fmt.Errorf("erro ao executar seed produtos: %v", err)
	}

	if err := s.seedVeiculos(); err != nil {
		return fmt.Errorf("erro ao executar seed veiculos: %v", err)
	}

	if err := s.seedUploads(); err != nil {
		return fmt.Errorf("erro ao executar seed uploads: %v", err)
	}

	if err := s.seedAvaliacoes(); err != nil {
		return fmt.Errorf("erro ao executar seed avaliacoes: %v", err)
	}

	if err := s.seedHistoricoPagamentos(); err != nil {
		return fmt.Errorf("erro ao executar seed historico pagamentos: %v", err)
	}

	if err := s.seedHistoricoResgates(); err != nil {
		return fmt.Errorf("erro ao executar seed historico resgates: %v", err)
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

	// Busca usuários para vincular às lojas
	var usuarios []models.Usuario
	if err := s.db.Limit(7).Find(&usuarios).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuarios: %v", err)
	}

	// Se não tiver usuários suficientes, usa o primeiro para todos
	getUsuarioID := func(index int) uint {
		if len(usuarios) == 0 {
			return 1
		}
		if index >= len(usuarios) {
			return usuarios[0].ID
		}
		return usuarios[index].ID
	}

	lojas := []models.Loja{
		{
			Nome:           "Auto Center São Paulo",
			CNPJ:           "12.345.678/0001-90",
			Imagem:         "https://via.placeholder.com/300x200",
			Latitude:       -23.5505,
			Longitude:      -46.6333,
			Rating:         5,
			IsMeuCarroMais: true,
			Categoria:      "Concessionária",
			IDUsuario:      getUsuarioID(0),
		},
		{
			Nome:           "Oficina do João",
			CNPJ:           "98.765.432/0001-10",
			Imagem:         "https://via.placeholder.com/300x200",
			Latitude:       -23.5489,
			Longitude:      -46.6388,
			Rating:         4,
			IsMeuCarroMais: false,
			Categoria:      "Oficina Mecânica",
			IDUsuario:      getUsuarioID(1),
		},
		{
			Nome:           "Carros Premium",
			CNPJ:           "11.222.333/0001-44",
			Imagem:         "https://via.placeholder.com/300x200",
			Latitude:       -23.5520,
			Longitude:      -46.6310,
			Rating:         5,
			IsMeuCarroMais: true,
			Categoria:      "Concessionária",
			IDUsuario:      getUsuarioID(2),
		},
		{
			Nome:           "Auto Center Porto Alegre Norte",
			CNPJ:           "22.333.444/0001-55",
			Imagem:         "https://via.placeholder.com/300x200",
			Latitude:       -30.01392952707345,
			Longitude:      -51.11168643659632,
			Rating:         4,
			IsMeuCarroMais: false,
			Categoria:      "Concessionária",
			IDUsuario:      getUsuarioID(3),
		},
		{
			Nome:           "Oficina Porto Alegre Centro",
			CNPJ:           "33.444.555/0001-66",
			Imagem:         "https://via.placeholder.com/300x200",
			Latitude:       -30.03004986742164,
			Longitude:      -51.12161634814089,
			Rating:         3,
			IsMeuCarroMais: false,
			Categoria:      "Oficina Mecânica",
			IDUsuario:      getUsuarioID(4),
		},
		{
			Nome:           "Loja de Peças Porto Alegre",
			CNPJ:           "44.555.666/0001-77",
			Imagem:         "https://via.placeholder.com/300x200",
			Latitude:       -30.02279847167631,
			Longitude:      -51.20306428600608,
			Rating:         4,
			IsMeuCarroMais: false,
			Categoria:      "Loja de Peças",
			IDUsuario:      getUsuarioID(5),
		},
		{
			Nome:           "Concessionária Porto Alegre Sul",
			CNPJ:           "55.666.777/0001-88",
			Imagem:         "https://via.placeholder.com/300x200",
			Latitude:       -30.04486628735914,
			Longitude:      -51.22406353079101,
			Rating:         5,
			IsMeuCarroMais: true,
			Categoria:      "Concessionária",
			IDUsuario:      getUsuarioID(6),
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

// vincularUsuariosLojas vincula usuários às lojas
func (s *Seeder) vincularUsuariosLojas() error {
	log.Println("📝 Vinculando usuários às lojas...")

	// Busca todos os usuários
	var usuarios []models.Usuario
	if err := s.db.Find(&usuarios).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuarios: %v", err)
	}

	if len(usuarios) == 0 {
		log.Println("⚠️ Nenhum usuário encontrado, pulando vinculação")
		return nil
	}

	// Busca todas as lojas
	var lojas []models.Loja
	if err := s.db.Find(&lojas).Error; err != nil {
		return fmt.Errorf("erro ao buscar lojas: %v", err)
	}

	if len(lojas) == 0 {
		log.Println("⚠️ Nenhuma loja encontrada, pulando vinculação")
		return nil
	}

	// Define vinculações específicas
	vinculacoes := map[string]string{
		"joao@email.com":        "12.345.678/0001-90", // João -> Auto Center São Paulo
		"maria@email.com":       "98.765.432/0001-10", // Maria -> Oficina do João
		"pedro@email.com":       "11.222.333/0001-44", // Pedro -> Carros Premium
		"carlos@email.com":      "22.333.444/0001-55", // Carlos -> Auto Center Porto Alegre Norte
		"ana.premium@email.com": "33.444.555/0001-66", // Ana -> Oficina Porto Alegre Centro
	}

	// Cria um mapa de lojas por CNPJ para facilitar a busca
	lojaPorCNPJ := make(map[string]models.Loja)
	for _, loja := range lojas {
		lojaPorCNPJ[loja.CNPJ] = loja
	}

	// Vincula usuários às lojas
	for _, usuario := range usuarios {
		if cnpjLoja, existe := vinculacoes[usuario.Email]; existe {
			if loja, lojaExiste := lojaPorCNPJ[cnpjLoja]; lojaExiste {
				// Atualiza o IDLoja do usuário
				usuario.IDLoja = &loja.ID
				if err := s.db.Save(&usuario).Error; err != nil {
					return fmt.Errorf("erro ao vincular usuario %s à loja %s: %v", usuario.Email, loja.Nome, err)
				}
				log.Printf("✅ Usuário %s vinculado à loja %s", usuario.Nome, loja.Nome)
			}
		}
	}

	return nil
}

// seedAnuncio popula a tabela anuncio
func (s *Seeder) seedAnuncio() error {
	log.Println("📝 Populando tabela anuncio...")

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
			Categoria:   "Carros",
			TipoAnuncio: "veiculo",
			IDLoja:      &loja.ID,
		},
		{
			Titulo:      "Toyota Corolla 2019",
			Descricao:   "Toyota Corolla XEi 2.0, automático, revisões em dia",
			Preco:       75000.00,
			Imagem:      "https://via.placeholder.com/400x300",
			Destaque:    false,
			Categoria:   "Carros",
			TipoAnuncio: "veiculo",
			IDLoja:      &loja.ID,
		},
		{
			Titulo:      "Volkswagen Golf GTI",
			Descricao:   "Golf GTI 2.0 TSI, manual, esportivo, baixa quilometragem",
			Preco:       95000.00,
			Imagem:      "https://via.placeholder.com/400x300",
			Destaque:    true,
			Categoria:   "Carros",
			TipoAnuncio: "veiculo",
			IDLoja:      &loja.ID,
		},
		{
			Titulo:      "Honda CB 500",
			Descricao:   "Honda CB 500F 2021, baixa quilometragem, revisões em dia",
			Preco:       32000.00,
			Imagem:      "https://via.placeholder.com/400x300",
			Destaque:    false,
			Categoria:   "Motos",
			TipoAnuncio: "veiculo",
			IDLoja:      &loja.ID,
		},
		{
			Titulo:      "Kit Pastilhas de Freio Premium",
			Descricao:   "Kit completo de pastilhas de freio cerâmicas para veículos de passeio",
			Preco:       250.00,
			Imagem:      "https://via.placeholder.com/400x300",
			Destaque:    false,
			Categoria:   "Peças",
			TipoAnuncio: "produto",
			IDLoja:      &loja.ID,
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

	servicos := []models.Servico{
		{
			Titulo:    "Troca de Óleo",
			Descricao: "Troca de óleo, filtro e óleo de filtro",
			Preco:     150.00,
			Imagem:    "https://via.placeholder.com/200x150",
			Categoria: "Manutenção",
			IDLoja:    loja.ID,
		},
		{
			Titulo:    "Alinhamento",
			Descricao: "Ajuste de direção, suspensão e direção",
			Preco:     250.00,
			Imagem:    "https://via.placeholder.com/200x150",
			Categoria: "Alinhamento",
			IDLoja:    loja.ID,
		},
		{
			Titulo:    "Revisão Completa",
			Descricao: "Revisão completa do motor, freios, suspensão e direção",
			Preco:     500.00,
			Imagem:    "https://via.placeholder.com/200x150",
			Categoria: "Revisão",
			IDLoja:    loja.ID,
		},
		{
			Titulo:    "Balanceamento de Rodas",
			Descricao: "Balanceamento completo das 4 rodas",
			Preco:     80.00,
			Imagem:    "https://via.placeholder.com/200x150",
			Categoria: "Balanceamento",
			IDLoja:    loja.ID,
		},
		{
			Titulo:    "Troca de Pastilhas de Freio",
			Descricao: "Troca de pastilhas de freio dianteiras ou traseiras",
			Preco:     200.00,
			Imagem:    "https://via.placeholder.com/200x150",
			Categoria: "Ajuste de Freios",
			IDLoja:    loja.ID,
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

	// Busca todos os usuários
	var usuarios []models.Usuario
	if err := s.db.Find(&usuarios).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuarios: %v", err)
	}

	if len(usuarios) == 0 {
		log.Println("⚠️ Nenhum usuário encontrado, pulando seed de carteiras")
		return nil
	}

	// Saldos variados para diferentes usuários (moedas do app - valores inteiros)
	saldos := []int{1000, 1500, 800, 2000, 1200, 3000}

	for i, usuario := range usuarios {
		// Usa saldo baseado no índice, ou um saldo padrão se não houver saldo específico
		saldo := 1000
		if i < len(saldos) {
			saldo = saldos[i]
		}

		carteira := models.Carteira{
			UsuarioID: usuario.ID,
			Saldo:     saldo,
		}

		var existing models.Carteira
		if err := s.db.Where("usuario_id = ?", carteira.UsuarioID).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&carteira).Error; err != nil {
				return fmt.Errorf("erro ao criar carteira para usuario %d: %v", carteira.UsuarioID, err)
			}
			log.Printf("✅ Carteira criada para usuario %s (ID: %d) com saldo %d moedas", usuario.Nome, carteira.UsuarioID, saldo)
		} else {
			log.Printf("⏭️ Carteira já existe para usuario %s (ID: %d)", usuario.Nome, carteira.UsuarioID)
		}
	}

	return nil
}

// seedUsuarioComAnuncioDestaque cria um usuário específico com um anúncio destaque vinculado
func (s *Seeder) seedUsuarioComAnuncioDestaque() error {
	log.Println("📝 Criando usuário com anúncio destaque...")

	// Busca o plano premium (ID = 3)
	var planoPremium models.TipoPlano
	if err := s.db.Where("nome = ?", "Premium").First(&planoPremium).Error; err != nil {
		return fmt.Errorf("erro ao buscar plano premium: %v", err)
	}

	// Busca uma loja para vincular o anúncio
	var loja models.Loja
	if err := s.db.First(&loja).Error; err != nil {
		return fmt.Errorf("erro ao buscar loja: %v", err)
	}

	// Cria o usuário com anúncio destaque
	usuarioComAnuncio := models.Usuario{
		Nome:     "Ana Silva Premium",
		Email:    "ana.premium@email.com",
		Senha:    "senha123",
		CPF:      "99988877766",
		Imagem:   "https://via.placeholder.com/150",
		Telefone: "(11) 99999-1111",
		Endereco: "Av. Faria Lima, 2000 - São Paulo/SP",
		Ativo:    true,
		IDPlano:  planoPremium.ID,
	}

	// Verifica se o usuário já existe
	var existingUsuario models.Usuario
	if err := s.db.Where("email = ?", usuarioComAnuncio.Email).First(&existingUsuario).Error; err != nil {
		// Se não existe, cria o usuário
		if err := s.db.Create(&usuarioComAnuncio).Error; err != nil {
			return fmt.Errorf("erro ao criar usuario %s: %v", usuarioComAnuncio.Email, err)
		}
		log.Printf("✅ Usuario criado: %s (%s)", usuarioComAnuncio.Nome, usuarioComAnuncio.Email)
	} else {
		usuarioComAnuncio = existingUsuario
		log.Printf("⏭️ Usuario já existe: %s (%s)", usuarioComAnuncio.Nome, usuarioComAnuncio.Email)
	}

	// Cria o anúncio destaque vinculado ao usuário
	anuncioDestaque := models.Anuncio{
		Titulo:      "BMW X5 2023 - Anúncio Premium",
		Descricao:   "BMW X5 xDrive40i 3.0 Turbo, automático, teto solar, bancos de couro, sistema de som premium, único dono, revisões na concessionária",
		Preco:       450000.00,
		Imagem:      "https://via.placeholder.com/400x300?text=BMW+X5+2023",
		Destaque:    true,
		Categoria:   "Carros",
		IDLoja:      &loja.ID,
		TipoAnuncio: "veiculo",
	}

	// Verifica se o anúncio já existe
	var existingAnuncio models.Anuncio
	if err := s.db.Where("titulo = ? AND id_loja = ?", anuncioDestaque.Titulo, anuncioDestaque.IDLoja).First(&existingAnuncio).Error; err != nil {
		// Se não existe, cria o anúncio
		if err := s.db.Create(&anuncioDestaque).Error; err != nil {
			return fmt.Errorf("erro ao criar anuncio %s: %v", anuncioDestaque.Titulo, err)
		}
		log.Printf("✅ Anuncio destaque criado: %s", anuncioDestaque.Titulo)
	} else {
		log.Printf("⏭️ Anuncio já existe: %s", anuncioDestaque.Titulo)
	}

	// Cria uma carteira para o usuário com saldo premium (moedas do app)
	carteira := models.Carteira{
		UsuarioID: usuarioComAnuncio.ID,
		Saldo:     5000, // Saldo maior para usuário premium (moedas do app)
	}

	var existingCarteira models.Carteira
	if err := s.db.Where("usuario_id = ?", carteira.UsuarioID).First(&existingCarteira).Error; err != nil {
		// Se não existe, cria a carteira
		if err := s.db.Create(&carteira).Error; err != nil {
			return fmt.Errorf("erro ao criar carteira para usuario %d: %v", carteira.UsuarioID, err)
		}
		log.Printf("✅ Carteira premium criada para usuario %s (ID: %d) com saldo %d moedas", usuarioComAnuncio.Nome, carteira.UsuarioID, carteira.Saldo)
	} else {
		log.Printf("⏭️ Carteira já existe para usuario %s (ID: %d)", usuarioComAnuncio.Nome, carteira.UsuarioID)
	}

	log.Printf("🎯 Cenário criado: Usuário %s (ID: %d) com anúncio destaque '%s' (ID: %d)",
		usuarioComAnuncio.Nome, usuarioComAnuncio.ID, anuncioDestaque.Titulo, anuncioDestaque.ID)

	return nil
}

// seedProdutos popula a tabela produtos
func (s *Seeder) seedProdutos() error {
	log.Println("📝 Populando tabela produtos...")

	// Busca algumas lojas
	var lojas []models.Loja
	if err := s.db.Limit(5).Find(&lojas).Error; err != nil {
		return fmt.Errorf("erro ao buscar lojas: %v", err)
	}

	if len(lojas) == 0 {
		log.Println("⚠️ Nenhuma loja encontrada, pulando seed de produtos")
		return nil
	}

	produtos := []models.Produto{
		{
			Nome:      "Óleo Motor 5W30",
			Descricao: "Óleo sintético para motor, 5W30, 4L",
			Preco:     89.90,
			Imagem:    "https://via.placeholder.com/300x200?text=Oleo+Motor",
			Estoque:   50,
			Categoria: "Óleos e Lubrificantes",
			Ativo:     true,
			IDLoja:    lojas[0].ID,
		},
		{
			Nome:      "Filtro de Óleo",
			Descricao: "Filtro de óleo original, compatível com diversos modelos",
			Preco:     25.50,
			Imagem:    "https://via.placeholder.com/300x200?text=Filtro+Oleo",
			Estoque:   100,
			Categoria: "Filtros",
			Ativo:     true,
			IDLoja:    lojas[0].ID,
		},
		{
			Nome:      "Pastilhas de Freio",
			Descricao: "Pastilhas de freio cerâmicas, alta performance",
			Preco:     180.00,
			Imagem:    "https://via.placeholder.com/300x200?text=Pastilhas+Freio",
			Estoque:   30,
			Categoria: "Freios",
			Ativo:     true,
			IDLoja:    lojas[1].ID,
		},
		{
			Nome:      "Bateria 60Ah",
			Descricao: "Bateria automotiva 60Ah, 12V, livre de manutenção",
			Preco:     350.00,
			Imagem:    "https://via.placeholder.com/300x200?text=Bateria+60Ah",
			Estoque:   15,
			Categoria: "Elétrica",
			Ativo:     true,
			IDLoja:    lojas[1].ID,
		},
		{
			Nome:      "Pneu Aro 15",
			Descricao: "Pneu 185/65 R15, banda de rodagem econômica",
			Preco:     280.00,
			Imagem:    "https://via.placeholder.com/300x200?text=Pneu+Aro+15",
			Estoque:   20,
			Categoria: "Pneus",
			Ativo:     true,
			IDLoja:    lojas[2].ID,
		},
		{
			Nome:      "Amortecedor Dianteiro",
			Descricao: "Amortecedor dianteiro, original de fábrica",
			Preco:     450.00,
			Imagem:    "https://via.placeholder.com/300x200?text=Amortecedor",
			Estoque:   10,
			Categoria: "Suspensão",
			Ativo:     true,
			IDLoja:    lojas[2].ID,
		},
		{
			Nome:      "Correia Dentada",
			Descricao: "Correia dentada de alta qualidade, kit completo",
			Preco:     120.00,
			Imagem:    "https://via.placeholder.com/300x200?text=Correia+Dentada",
			Estoque:   25,
			Categoria: "Motor",
			Ativo:     true,
			IDLoja:    lojas[3].ID,
		},
		{
			Nome:      "Lâmpada H7",
			Descricao: "Lâmpada halógena H7, 55W, luz branca",
			Preco:     15.90,
			Imagem:    "https://via.placeholder.com/300x200?text=Lampada+H7",
			Estoque:   200,
			Categoria: "Iluminação",
			Ativo:     true,
			IDLoja:    lojas[3].ID,
		},
	}

	for _, produto := range produtos {
		var existing models.Produto
		if err := s.db.Where("nome = ? AND id_loja = ?", produto.Nome, produto.IDLoja).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&produto).Error; err != nil {
				return fmt.Errorf("erro ao criar produto %s: %v", produto.Nome, err)
			}
			log.Printf("✅ Produto criado: %s", produto.Nome)
		} else {
			log.Printf("⏭️ Produto já existe: %s", produto.Nome)
		}
	}

	return nil
}

// seedVeiculos popula a tabela veiculos
func (s *Seeder) seedVeiculos() error {
	log.Println("📝 Populando tabela veiculos...")

	// Busca alguns usuários
	var usuarios []models.Usuario
	if err := s.db.Limit(5).Find(&usuarios).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuarios: %v", err)
	}

	if len(usuarios) == 0 {
		log.Println("⚠️ Nenhum usuário encontrado, pulando seed de veículos")
		return nil
	}

	veiculos := []models.Veiculo{
		{
			Marca:         "Honda",
			Modelo:        "Civic",
			AnoFabricacao: 2020,
			AnoModelo:     2020,
			Cor:           "Prata",
			Placa:         "ABC-1234",
			IDUsuario:     usuarios[0].ID,
			Ativo:         true,
		},
		{
			Marca:         "Toyota",
			Modelo:        "Corolla",
			AnoFabricacao: 2019,
			AnoModelo:     2019,
			Cor:           "Branco",
			Placa:         "DEF-5678",
			IDUsuario:     usuarios[1].ID,
			Ativo:         true,
		},
		{
			Marca:         "Volkswagen",
			Modelo:        "Golf",
			AnoFabricacao: 2021,
			AnoModelo:     2021,
			Cor:           "Preto",
			Placa:         "GHI-9012",
			IDUsuario:     usuarios[2].ID,
			Ativo:         true,
		},
		{
			Marca:         "Ford",
			Modelo:        "Focus",
			AnoFabricacao: 2018,
			AnoModelo:     2018,
			Cor:           "Azul",
			Placa:         "JKL-3456",
			IDUsuario:     usuarios[3].ID,
			Ativo:         true,
		},
		{
			Marca:         "Chevrolet",
			Modelo:        "Onix",
			AnoFabricacao: 2022,
			AnoModelo:     2022,
			Cor:           "Vermelho",
			Placa:         "MNO-7890",
			IDUsuario:     usuarios[0].ID,
			Ativo:         true,
		},
		{
			Marca:         "Fiat",
			Modelo:        "Argo",
			AnoFabricacao: 2020,
			AnoModelo:     2020,
			Cor:           "Cinza",
			Placa:         "PQR-1357",
			IDUsuario:     usuarios[1].ID,
			Ativo:         true,
		},
	}

	for _, veiculo := range veiculos {
		var existing models.Veiculo
		if err := s.db.Where("placa = ?", veiculo.Placa).First(&existing).Error; err != nil {
			// Se não existe, cria
		if err := s.db.Create(&veiculo).Error; err != nil {
			return fmt.Errorf("erro ao criar veiculo %s: %v", veiculo.Placa, err)
		}
		log.Printf("✅ Veiculo criado: %s %s %d (%s)", veiculo.Marca, veiculo.Modelo, veiculo.AnoFabricacao, veiculo.Placa)
	} else {
		log.Printf("⏭️ Veiculo já existe: %s %s %d (%s)", veiculo.Marca, veiculo.Modelo, veiculo.AnoFabricacao, veiculo.Placa)
	}
	}

	return nil
}

// seedUploads popula a tabela uploads
func (s *Seeder) seedUploads() error {
	log.Println("📝 Populando tabela uploads...")

	// Busca alguns veículos
	var veiculos []models.Veiculo
	if err := s.db.Limit(3).Find(&veiculos).Error; err != nil {
		return fmt.Errorf("erro ao buscar veiculos: %v", err)
	}

	// Busca alguns produtos
	var produtos []models.Produto
	if err := s.db.Limit(3).Find(&produtos).Error; err != nil {
		return fmt.Errorf("erro ao buscar produtos: %v", err)
	}

	// Busca algumas lojas
	var lojas []models.Loja
	if err := s.db.Limit(3).Find(&lojas).Error; err != nil {
		return fmt.Errorf("erro ao buscar lojas: %v", err)
	}

	uploads := []models.Upload{}

	// Uploads de imagens de veículos
	for i, veiculo := range veiculos {
		uploads = append(uploads, models.Upload{
			IDVeiculo:    &veiculo.ID,
			TipoEntidade: "veiculo",
			Tipo:         "Imagem",
			URL:          fmt.Sprintf("https://via.placeholder.com/800x600?text=%s+%d", veiculo.Modelo, veiculo.AnoFabricacao),
			NomeArquivo:  fmt.Sprintf("veiculo_%d_principal.jpg", veiculo.ID),
			Tamanho:      2048000, // 2MB
			TipoMime:     "image/jpeg",
			Principal:    true,
			Ordem:        1,
		})

		if i < 2 { // Adiciona upload adicional para os primeiros 2 veículos
			uploads = append(uploads, models.Upload{
				IDVeiculo:    &veiculo.ID,
				TipoEntidade: "veiculo",
				Tipo:         "Imagem",
				URL:          fmt.Sprintf("https://via.placeholder.com/800x600?text=%s+%d+Interior", veiculo.Modelo, veiculo.AnoFabricacao),
				NomeArquivo:  fmt.Sprintf("veiculo_%d_interior.jpg", veiculo.ID),
				Tamanho:      1856000, // 1.8MB
				TipoMime:     "image/jpeg",
				Principal:    false,
				Ordem:        2,
			})
		}
	}

	// Uploads de imagens de produtos
	for i, produto := range produtos {
		uploads = append(uploads, models.Upload{
			IDProduto:    &produto.ID,
			TipoEntidade: "produto",
			Tipo:         "Imagem",
			URL:          fmt.Sprintf("https://via.placeholder.com/400x400?text=%s", produto.Nome),
			NomeArquivo:  fmt.Sprintf("produto_%d.jpg", produto.ID),
			Tamanho:      1024000, // 1MB
			TipoMime:     "image/jpeg",
			Principal:    true,
			Ordem:        1,
		})

		if i == 0 { // Adiciona upload adicional para o primeiro produto
			uploads = append(uploads, models.Upload{
				IDProduto:    &produto.ID,
				TipoEntidade: "produto",
				Tipo:         "Imagem",
				URL:          fmt.Sprintf("https://via.placeholder.com/400x400?text=%s+Detalhe", produto.Nome),
				NomeArquivo:  fmt.Sprintf("produto_%d_detalhe.jpg", produto.ID),
				Tamanho:      896000, // 896KB
				TipoMime:     "image/jpeg",
				Principal:    false,
				Ordem:        2,
			})
		}
	}

	// Uploads de imagens de lojas
	for _, loja := range lojas {
		uploads = append(uploads, models.Upload{
			IDLoja:       &loja.ID,
			TipoEntidade: "loja",
			Tipo:         "Imagem",
			URL:          fmt.Sprintf("https://via.placeholder.com/600x400?text=%s", loja.Nome),
			NomeArquivo:  fmt.Sprintf("loja_%d_fachada.jpg", loja.ID),
			Tamanho:      1536000, // 1.5MB
			TipoMime:     "image/jpeg",
			Principal:    true,
			Ordem:        1,
		})
	}

	for _, upload := range uploads {
		var existing models.Upload
		if err := s.db.Where("url = ?", upload.URL).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&upload).Error; err != nil {
				return fmt.Errorf("erro ao criar upload: %v", err)
			}
			log.Printf("✅ Upload criado: %s", upload.NomeArquivo)
		} else {
			log.Printf("⏭️ Upload já existe: %s", upload.NomeArquivo)
		}
	}

	return nil
}

// seedAvaliacoes popula a tabela avaliacoes
func (s *Seeder) seedAvaliacoes() error {
	log.Println("📝 Populando tabela avaliacoes...")

	// Busca alguns usuários
	var usuarios []models.Usuario
	if err := s.db.Limit(4).Find(&usuarios).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuarios: %v", err)
	}

	// Busca algumas lojas
	var lojas []models.Loja
	if err := s.db.Limit(4).Find(&lojas).Error; err != nil {
		return fmt.Errorf("erro ao buscar lojas: %v", err)
	}

	if len(usuarios) == 0 || len(lojas) == 0 {
		log.Println("⚠️ Usuários ou lojas não encontrados, pulando seed de avaliações")
		return nil
	}

	avaliacoes := []models.Avaliacao{
		{
			IDUsuario:  usuarios[0].ID,
			IDLoja:     &lojas[0].ID,
			Nota:       5,
			Comentario: "Excelente atendimento! Serviço de qualidade e preço justo. Recomendo!",
		},
		{
			IDUsuario:  usuarios[1].ID,
			IDLoja:     &lojas[0].ID,
			Nota:       4,
			Comentario: "Muito bom serviço, apenas demorou um pouco mais que o esperado.",
		},
		{
			IDUsuario:  usuarios[2].ID,
			IDLoja:     &lojas[1].ID,
			Nota:       5,
			Comentario: "Profissionais qualificados e equipamentos modernos. Super recomendo!",
		},
		{
			IDUsuario:  usuarios[3].ID,
			IDLoja:     &lojas[1].ID,
			Nota:       3,
			Comentario: "Serviço ok, mas poderia melhorar na comunicação com o cliente.",
		},
		{
			IDUsuario:  usuarios[0].ID,
			IDLoja:     &lojas[2].ID,
			Nota:       4,
			Comentario: "Boa qualidade dos produtos e preços competitivos.",
		},
		{
			IDUsuario:  usuarios[1].ID,
			IDLoja:     &lojas[2].ID,
			Nota:       5,
			Comentario: "Atendimento excepcional! Encontraram exatamente o que eu precisava.",
		},
		{
			IDUsuario:  usuarios[2].ID,
			IDLoja:     &lojas[3].ID,
			Nota:       4,
			Comentario: "Loja bem organizada e funcionários atenciosos.",
		},
		{
			IDUsuario:  usuarios[3].ID,
			IDLoja:     &lojas[3].ID,
			Nota:       2,
			Comentario: "Demorou muito para ser atendido e o produto não estava como descrito.",
		},
	}

	for _, avaliacao := range avaliacoes {
		var existing models.Avaliacao
		if err := s.db.Where("id_usuario = ? AND id_loja = ?", avaliacao.IDUsuario, avaliacao.IDLoja).First(&existing).Error; err != nil {
			// Se não existe, cria
			if err := s.db.Create(&avaliacao).Error; err != nil {
				return fmt.Errorf("erro ao criar avaliacao: %v", err)
			}
			log.Printf("✅ Avaliação criada: %d estrelas para loja ID %d", avaliacao.Nota, avaliacao.IDLoja)
		} else {
			log.Printf("⏭️ Avaliação já existe para usuário %d e loja %d", avaliacao.IDUsuario, avaliacao.IDLoja)
		}
	}

	return nil
}

// seedHistoricoPagamentos popula a tabela historico_pagamentos
func (s *Seeder) seedHistoricoPagamentos() error {
	log.Println("📝 Populando tabela historico_pagamentos...")

	// Busca alguns usuários
	var usuarios []models.Usuario
	if err := s.db.Limit(3).Find(&usuarios).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuarios: %v", err)
	}

	if len(usuarios) == 0 {
		log.Println("⚠️ Nenhum usuário encontrado, pulando seed de histórico de pagamentos")
		return nil
	}

	historicos := []models.HistoricoPagamento{
		{
			Valor:           150.00,
			StripeSessionID: "cs_test_session_001",
			StripePaymentID: "pi_test_payment_001",
			Status:          "completed",
			TipoPlano:       "monthly",
			Moeda:           "BRL",
			IDUsuario:       usuarios[0].ID,
		},
		{
			Valor:           89.90,
			StripeSessionID: "cs_test_session_002",
			StripePaymentID: "pi_test_payment_002",
			Status:          "completed",
			TipoPlano:       "monthly",
			Moeda:           "BRL",
			IDUsuario:       usuarios[1].ID,
		},
		{
			Valor:           350.00,
			StripeSessionID: "cs_test_session_003",
			StripePaymentID: "pi_test_payment_003",
			Status:          "pending",
			TipoPlano:       "yearly",
			Moeda:           "BRL",
			IDUsuario:       usuarios[2].ID,
		},
		{
			Valor:           25.50,
			StripeSessionID: "cs_test_session_004",
			StripePaymentID: "pi_test_payment_004",
			Status:          "completed",
			TipoPlano:       "monthly",
			Moeda:           "BRL",
			IDUsuario:       usuarios[0].ID,
		},
		{
			Valor:           500.00,
			StripeSessionID: "cs_test_session_005",
			StripePaymentID: "pi_test_payment_005",
			Status:          "completed",
			TipoPlano:       "yearly",
			Moeda:           "BRL",
			IDUsuario:       usuarios[1].ID,
		},
	}

	for _, historico := range historicos {
		if err := s.db.Create(&historico).Error; err != nil {
			return fmt.Errorf("erro ao criar historico de pagamento: %v", err)
		}
		log.Printf("✅ Histórico de pagamento criado: R$ %.2f - %s", historico.Valor, historico.Status)
	}

	return nil
}

// seedHistoricoResgates popula a tabela historico_resgates
func (s *Seeder) seedHistoricoResgates() error {
	log.Println("📝 Populando tabela historico_resgates...")

	// Busca alguns usuários
	var usuarios []models.Usuario
	if err := s.db.Limit(3).Find(&usuarios).Error; err != nil {
		return fmt.Errorf("erro ao buscar usuarios: %v", err)
	}

	if len(usuarios) == 0 {
		log.Println("⚠️ Nenhum usuário encontrado, pulando seed de histórico de resgates")
		return nil
	}

	// Busca algumas lojas para os resgates
	var lojas []models.Loja
	if err := s.db.Limit(2).Find(&lojas).Error; err != nil {
		return fmt.Errorf("erro ao buscar lojas: %v", err)
	}

	if len(lojas) == 0 {
		log.Println("⚠️ Nenhuma loja encontrada, pulando seed de histórico de resgates")
		return nil
	}

	resgates := []models.HistoricoResgate{
		{
			Valor:       100.00,
			TipoResgate: "produto",
			Status:      "confirmado",
			IDUsuario:   usuarios[0].ID,
			IDLoja:      lojas[0].ID,
		},
		{
			Valor:       250.00,
			TipoResgate: "servico",
			Status:      "pendente",
			IDUsuario:   usuarios[1].ID,
			IDLoja:      lojas[1].ID,
		},
		{
			Valor:       50.00,
			TipoResgate: "produto",
			Status:      "confirmado",
			IDUsuario:   usuarios[2].ID,
			IDLoja:      lojas[0].ID,
		},
		{
			Valor:       300.00,
			TipoResgate: "veiculo",
			Status:      "pendente",
			IDUsuario:   usuarios[0].ID,
			IDLoja:      lojas[1].ID,
		},
	}

	for _, resgate := range resgates {
		if err := s.db.Create(&resgate).Error; err != nil {
			return fmt.Errorf("erro ao criar historico de resgate: %v", err)
		}
		log.Printf("✅ Histórico de resgate criado: R$ %.2f - %s", resgate.Valor, resgate.Status)
	}

	return nil
}
