# Golang Rate Limiter

Uma implementação robusta de rate limiter em Go que utiliza Redis como backend para armazenamento e gerenciamento de limites de requisições.

## Funcionalidades

- 🔐 Rate limiting por chave de API
- 🔄 Redis como backend para persistência
- 📦 Suporte a Docker e Docker Compose
- ⚙️ Configuração via variáveis de ambiente
- 📊 Três níveis de rate limiting:
  - Default: para usuários comuns
  - Admin: para usuários administradores
  - Tester: para usuários de teste

## Requisitos

- Go 1.21 ou superior
- Redis 7.2 ou superior
- Docker e Docker Compose (opcional, para ambiente de desenvolvimento)

## Instalação

### Usando Go Modules

```bash
go get github.com/sunnygosdk/rate-limiter
```

### Usando Docker

```bash
docker-compose up -d
```

## Configuração

Todas as configurações são feitas através de variáveis de ambiente. O arquivo `.env` de exemplo contém todas as configurações necessárias:

### Variáveis de Ambiente

#### Configuração Geral
- `APP_ENV`: Ambiente de execução (DEV/TEST/PROD)
- `APP_PORT`: Porta da aplicação
- `CACHE_CLIENT`: Cliente de cache (atualmente apenas REDIS suportado, mas já configurado para suporte a outros clientes)

#### Configuração Redis
- `REDIS_HOST`: Host do Redis
- `REDIS_PORT`: Porta do Redis
- `REDIS_PASSWORD`: Senha do Redis (opcional)
- `REDIS_DB`: Número do banco Redis

#### Configuração Rate Limiter Default
- `DEFAULT_LIMIT`: Limite de requisições (padrão: 10)
- `DEFAULT_WINDOW`: Janela de tempo em segundos (padrão: 60)

#### Configuração Rate Limiter Admin
- `ADMIN_LIMIT`: Limite de requisições para admin (padrão: 100)
- `ADMIN_WINDOW`: Janela de tempo para admin (padrão: 30)
- `ADMIN_API_KEY`: Chave de API para admin

#### Configuração Rate Limiter Tester
- `TESTER_LIMIT`: Limite de requisições para tester (padrão: 50)
- `TESTER_WINDOW`: Janela de tempo para tester (padrão: 30)
- `TESTER_API_KEY`: Chave de API para tester

## Uso

### Configuração novos Clients

No arquivo `internal/infrastructure/persistence/cache_client.go` é possível adicionar novos clients de cache, basta que ele implemente a interface `CacheClient`.

```go
// CacheClient represents a cache client
type CacheClient interface {
	CloseCacheClient() error
	CheckCacheKeysOnWindow(key string, context context.Context, window time.Duration) (int64, error)
}

// ConfigureCacheClient configures the cache client
func ConfigureCacheClient() CacheClient {
	log.Println("App environment:", config.AppEnvConfig.APP_ENV)
	if config.AppEnvConfig.APP_ENV == "TEST" {
		return fixture.NewRedisClientFixture()
	}

	log.Println("Cache client:", config.AppEnvConfig.CACHE_CLIENT)
	if config.AppEnvConfig.CACHE_CLIENT == "REDIS" {
		return NewRedisClient(config.AppEnvConfig)
	}

	// TODO: Add other cache clients
	log.Println("Cache client not found")
	return nil
}
```

### Como Servidor

O servidor está configurado para rodar na porta 8080 (configurável via variável de ambiente `APP_PORT`).

## Docker

### Iniciar Serviço

```bash
docker-compose up -d
```

### Parar Serviço

```bash
docker-compose down
```

## Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob licença MIT. Veja o arquivo LICENSE para mais detalhes.

## Suporte

Para suporte ou dúvidas, abra uma issue no repositório ou entre em contato com os mantenedores.
