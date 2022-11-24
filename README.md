# swapi

##### _Projeto de interface para a api pública de Star Wars (https://swapi.dev/)._
##### _Cobertura de Testes atual_: 70,7% (24/11/2022)
###### Para gerar a cobertura de teste execute os comandos abaixo ou utilize o ./skaffold.sh que será explicada mais abaixo neste documento.
- `swapi/$ go test -v -coverpkg=./... -coverprofile=profile.cov ./...; go tool cover -func profile.cov`

![Fluxo Básico](https://raw.githubusercontent.com/marcosArruda/swapi/main/imgs/yoda.webp "Do. Or do not. There is no try")
> "No! Do. Or do not. There is no try." - Yoda

## Pré Requisitos

- **Golang 1.19+**: Acredito que funciona na 1.15+ pois não utilizei nenhuma 'feature' nova dessas novas versões como 'Generics' e 'Fuzzing Tests", por exemplo.
- **Docker & Docker-Compose**: Para rodar o ambiente completo
- **Linux**: Recomendo fortemente rodar o projeto em um ambiente **linux**. Se você for utilizar windows, vai precisar lidar com algumas questões envolvendo o **SWL2** e os espelhamentos de portas. Se for usar **MacOS** também irá funcionar, mas também precisará lidar com os espelhamentos de portas se quiser subir o ambiente completo(docker-compose). Windows assim como o MacOS, no docker, utiliza uma estrutura de virtualização de um ambiente nativo Linux(filesystem) para instanciar as imagens docker e por isso problemas como _"quero acessar diretamente uma porta do container em execução"_ se tornam mais complexos de resolver nesses ambientes. Com bom entendimento, é possivel rodar tranquilamente nos dois, mas o suporte dessa documentação é **_exclusivamente no Linux_**.
- **curl**: Para rodar as requisições manuais, mas fique a vontade para usar o **wget** se preferir.

## Arquitetura

Seguindo os princípios **SOLID**, o código segue uma estrutura em camadas com o diferencial dos padrões **Inversão de Controle**(Inspirado no ApplicationContext do velho Spring..) e **"NoOps"**(**_No Operation_**, padrão _"importado"_ da engenharia civil e mecânica) implementados. Olharemos esses dois padrões mais a fundo no decorrer desta documentação mas à principio, com o uso desses padrões, qualquer "camada"(serviço/componente) pode ser acessada de qualquer lugar do código sem bloqueios. Esta estrutura dá a este código uma _flexibilidade quase **infinita**_. No sentido "infinito" relativo à _performance_, todas as chamadas externas(**https://swapi.dev**) são feitas de maneira **_concorrente/paralela_** e de maneira **_não blocantes_**, reduzindo drasticamente a latência das requisições experimentadas pelo usuário final durante as requisições Http. No sentido "infinito" relativo à _velocidade_ de inclusão de novas features, com o ServiceManager gerenciando o ciclo de vida de qualquer componente, basta adicionarmos a nova feature como um componente novo, implementar o NoOps dela, adiciona-la ao ServiceManager e então ela ja poderá ser utilizada por qualquer outro componente de maneira livre e principalmente desacoplada. 

![Fluxo Básico](/imgs/Diagram.png?raw=true "Fluxo Básico")
Vamos agora destrinchar esses dois padrões de design principais utilizados:

### Inversão de Controle

O padrão _Inversão de Controle_ consiste em permitir que outra entidade se encarregue de gerenciar o ciclo de vida de **TODAS** as dependências(objetos/instâncias) de um componente específico. Por exemplo, o componente _PlanetFinderService_ precisa dos componentes _SwApiService_ e do componente _PersistenceService_ para executar suas funções. Numa aplicação que não utilize _Inversão de Controle_, o _PlanetFinderService_ seria o responsável por **instânciar** os objetos dos quais ele depende e assim o mesmo acaba sendo responsável por controlar todo o ciclo de vida desses componentes, fazendo-se necessário dessa forma implementar dentro do _PlanetFinderService_ todo código de controle(boilerplate code) para os componentes dependentes.

No nosso caso, o **ServiceManager** é a entidade _responsável por controlar o ciclo de vida de TODOS os outos componentes da aplicação e essa é **a única responsabilidade dessa entidade(SOLID)**_. Dessa forma, se o +PlanetFinderService_ precisar utilizar o _SwApiService_ para realizar a request para a API pública, com uma única linha de código o _PlanetFinderService_ recebe do _ServiceManager_ a instância do _SwApiService_ e em seguida ja tem acesso à interface pública dela. Este mesmo comportamente existe no _"diálogo"_ entre TODAS os _serviços/componentes_ do sistema, **isolando-os e desacoplando-os**.

O ServiceManager utiliza de 'api fluente' para possibilitar um fácil uso de todas as funções do ciclo de vida, como exemplificado no cmd/main/main.go da aplicação.

### NoOps (No Operation)

_No Operation_ é um nome pouco conhecido na industria do software, porém, é bastante utilizado. Inspirado na construção civil, um famoso exemplo do padrão é a existência das _"bolas de aço de equilibrio"_ utilizadas na construção de predios muito grandes em locais onde existe muito vento. Com a pressão do vento, todo prédio muito alto enverga e desenverga naturalmente. No centro desses prédios existe SEMPRE uma grande bola de aço presa por uma corda de aço no teto e pendurada há uma certa altura(metade do prédio normalmente) suspensa no ar. Esta bola balança conforme o prédio _"inclina-se"_ fazendo o papel de ajuste do centro de equilibrio do prédio.

O conceito de No Operation do caso da "bola de aço" vem do fato de que essa bola não exige **_NENHUMA_** manutenção. Qualquer modificação de engenharia feita no prédio(tirando modificação da altura do mesmo) não resulta em mudanças na posição da bola. A bola apenas existirá lá, suspensa no centro do prédio, fazendo seu papel. Não exigindo nenhuma manutenção, é a isso que damos o nome do padrão NoOps (No Operation).

Trazendo o exemplo para o nosso mundo do software, o padrão NoOps, consiste em estruturas de código que são necessárias para o sistema funcionar de acordo como que foi desenhado porém não exigindo nenhuma manutenção :). Por exemplo, a struct **noOpsPlanetFinderService** é exatamente a implementação **_NoOps_** da interface PlanetFinderService. Essa struct implementa a interface e provê comportamentos básicos usados em **TODOS** os testes unitários desse projeto. **Basicamente, as instâncias _NoOps_, fazem todo o papel de base de retorno de "mocks"** utilizados nos testes unitários de maneira que muito menos código é necessário para escrever um teste unitário que precise de Mocks da interface **PlanetFinderService** por exemplo, que é exatamente o caso dos testes unitários da entidade **HttpService**. Por sua vez, a interface **HttpService** também tem a sua implementação **NoOps**, o que permite que qualquer outro componente que precise de mocks de HttpService utilize dessa outra implementação sem burocracias.

### Logs

Utilizei a lib Zap do Uber (https://go.uber.org/zap) para fazer os logs estruturados contendo as seguintes informações:
```
{
    "level":"info", # nível de log (info, warn, error, debug)
    "ts":1669308163.428072, # timestamp do horario
    "caller":"planetfinder/planetfinder.go:29", # "package/filename:line"
    "msg":"PlanetFinder Service Started!", # a mensagem de log
    "AppEnv":"PROD", # ambiente que está rodando.
    "service":"swapiapp", #nome do projeto
    "version":"1.0" # versão 
}
```

**DISCLAIMER: Não implementei escrita dos logs em arquivo texto. Eu poderia implementar tranquilamente e implementarei se for pedido sem problema nenhum. Porém, com o modelo de aplicações rodando dentro dos containers Docker, os logs já têm um caminho muito bem definido pela estrutura do Docker que contém, no seu driver de logs (Docker Log Driver) toda a mecânica para escrever o stdOut e stdErr do processo executado dentro do container para um arquivo texto e para o stdOut stdError do proprio container. Este excelente padrão implicou no continuo desuso de arquivos texto(dentro do container) pois se torna uma PÉSSIMA prática um outro processo precisar entrar dentro do container para ler um determinado arquivo em um determinado 'path' para saber o que estaria acontecendo com a aplicação. A conclusão é que se uma aplicação pretende ser utilizada com containers Docker, definitivamente não é uma boa prática escrever os logs em arquivo texto. Se um Filebeat da vida for realmente utilizado na infraestrutura para ler os logs de arquivo texto, que sejam então pegos DE FORA do container, utilizando o Docker Log Driver padrão do Docker que já escreve esses logs em arquivo texto no filesystem do Host.**

### Colocando o código para rodar

Você pode rodar o código com um simples `>$ go mod tidy; go run main.go` porém, sem uma instância do mysql no ar, escutando o host **_db:3306_** você receberá erros. Por esse motivo, é um dos pré requisitos a utilização do Docker e do Docker Compose para subirmos o projeto.

Para facilitar a vida de todos, criei um shell script chamado `scaffold.sh` que possui os seguintes comandos:
```
swapi/$ ./scaffold.sh full-rebuild -prune -runtests #compilação e subida com testes e limpeza dos volumes(docker).
swapi/$ ./scaffold.sh runtests #roda todos os testes e mostra a cobertura.
swapi/$ ./scaffold.sh build #apenas compila tudo usando docker.
swapi/$ ./scaffold.sh down #destroi as imagens desse projeto que estão rodando atualmente.
swapi/$ ./scaffold.sh up #sobe todos os containers que estão ja compilados, compilando-os caso não estejam.
swapi/$ ./scaffold.sh logs -app #appenda o tail nos logs do swapiapp.
swapi/$ ./scaffold.sh logs -db #appenda o tail nos logs da instancia do mysql(docker).
```
ps.: Os comandos com -prune e -runtests são opcionais.

Dado esses comandos:

- execute `./scaffold.sh full-rebuild -prune -runtests` para subir o projeto do zero.
- Aguarde a finalização de todo o script.
- Confirme que a base de dados já subiu e está escutando na porta 3306 verificando se a linha **[Server] /usr/sbin/mysqld: ready for connections. Version: '8.0.31'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server - GPL.** foi impressa nos logs do mysql executando: `./scaffold.sh logs -db`.
- Execute **CTRL + C** para voltar ao terminal.
- Execute `./scaffold.sh restartapp` ou `docker-compose restart swapiapp` para reiniciar o container da app depois que o mysql esteja disponível.
- Siga os passos abaixo para executar as chamadas que preferir.

### GET /planet/:id

Retorna o planeta identificado pelo **:id** (int). Segue os padrões de **idempotência** portanto, o fluxo irá verificar se esse planeta ja existe na base de dados. Se existir, irá retorna-lo prontamente. Se não existir, irá carrega-lo do https://swapi.dev/api/planets/1/, retornar na response e em seguinda salva-lo na base de dados (**seguindo os conselhos do Yoda, toda a persistência é feita de maneira assíncrona.**).

Dado que toda a persistência é feita de maneira assíncrona (golang channels) o sistema, uma vez que possui os dados do https://swapi.dev ou mysql, o sistema ja retorna a response e continua executando o que for necessário em _background_:
Ex:
```
curl http://localhost:8080/planet/1 --header "Content-Type: application/json" --request "GET"
```

### DELETE /planet/:id

Remove da base de dados local o planeta com o **:id**(int) informado.
Ex:
```
curl http://localhost:8080/planet/1 --header "Content-Type: application/json" --request "DELETE"
```

### GET /planet?search=<NOME_PARCIAL>

Retorna todos os planetas que contenham o nome parcial infomado. O fluxo irá **diretamente** carregar do https://swapi.dev/api/planets?search=<NOME_PARCIAL>, retornar a response e salvar todos esses dados de maneira assíncrona seguindo as dicas do querido Yoda.
```
curl http://localhost:8080/planet?search=h --header "Content-Type: application/json" --request "GET"
```

### DELETE /planet?name=<NOME_EXATO>
Remove da base de dados local o planeta com o nome exato informado. Segue os padrões de **idempotência** portanto sempre retorna sucesso mesmo não encontrando o planeta na base de dados. Por exemplo, se você não lembra exatamente o nome do planeta que quer remover, por segurança a base de dados não irá remove-lo se você não informar o nome exato.

Ex:
```
curl http://localhost:8080/planet?name=saleucami --header "Content-Type: application/json" --request "DELETE" -vvv
```

### GET /planets

Retorna todos os planetas já existentes na base de dados.
```
curl http://localhost:8080/planets --header "Content-Type: application/json" --request "GET"
```

### Desligue todos os containers

Execute `docker-compose down`(remove todos os containers em execução no context do docker-compose.yml, `docker volume prune -f`(remove todos os volumes 'dangling') e se quiser `docker system prune -f`(remove todos os 'dangling' containers)