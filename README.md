# swapi

Projeto que faz uma interface com a api pública de Star Wars (https://swapi.dev/).

## Pré Requisitos

- Golang 1.19+: Acredito que funciona na 1.15+ pois não utilizei nenhuma 'feature' nova dessas novas versões como 'Generics' e 'Fuzzing Tests", por exemplo.

- Docker & Docker-Compose: Para rodar o ambiente completo

- Linux:
    - Recomendo fortemente rodar o projeto em um ambiente linux, porém, se você utilizar windows você vai precisar lidar com algumas questões envolvendo o SWL2 e os espelhamentos de portas. Se for usar MacOS também irá funcionar, mas para testar corretamente o ambiente todo precisará lidar com os espelhamentos de portas. Windows assim como o MacOS, no docker, utilizam uma estrutura de virtualização de um ambiente nativo Linux(filesystem), por isso, problemas como "quero acessar diretamente uma porta do container em execução" se tornam mais complexos de resolver nesses ambientes, mas com bom entendimento, é possivel rodar tranquilamente nos dois.

- curl: para rodar as requisições manuais, mas fique a vontade para usar o wget se preferir.

## Arquitetura

Seguindo os princípios SOLID, o código segue uma estrutura em camadas, porém, com o diferencial dos padrões "Inversão de Controle"(Inspirado no ApplicationContext do velho Spring..) e "NoOps"(No Operation, padrão "importado" da engenharia civil e mecânica) implementados, qualquer "camada"(serviço/componente) pode ser acessada de qualquer lugar do código sem bloqueios. Esta estrutura dá a este código em específico uma flexibilidade quase infinita. No sentido "infinito" relativo à performance, todas as chamadas externas(swapi.dev) são feitas de maneira concorrente/paralela e não blocantes, reduzindo drasticamente a latência das requisições percebidas pelo usuário final.

Vamos agora destrinchar esses dois padrões de design principais utilizados:

### Inversão de Controle

O padrão Inversão de Controle consiste em permitir que outra entidade se encarregue de gerenciar o ciclo de vida de TODOS dependências(objetos/instâncias) de um componente específico. Por exemplo, o componente PlanetFinderService precisa dos componentes SwApiService e do componente PersistenceService. Numa aplicação que não utiliza Inversão de Controle, o PlanetFinderService seria o responsável por instânciar os objetos os quais ele depende e assim, sendo responsável por controlar todo o ciclo de vida desses componentes, fazendo-se necessário dessa forma implementar uma serie de controles(boilerplate code) para esse controle.

No nosso caso, o ServiceManager é a entidade responsável por controlar o ciclo de vida de todos os outos componentes e essa é a única responsabilidade dessa entidade(SOLID). Dessa forma, se o PlanetFinderService precisa utilizar o SwApiService para realizar a request para a API pública, com uma única linha de código o PlanetFinderService "pergunta" ao ServiceManager pela instância do SwApiService e em seguida ja tem acesso à interface pública dela. Este mesmo comportamente existe no "diálogo" entre TODAS os serviços/componentes do sistema, isolando-os e desacoplando-os.

### NoOps (No Operation)

No Operation é um nome pouco conhecido na industria do software, porém, é bastante utilizado. Inspirado na construção civil, um famoso exemplo do padrão é a existência das "bolas de aço de equilibrio" utilizadas na construção de predios muito grandes em locais onde existe muito vento. Com a pressão do vento, todo prédio muito alto enverga e desenverga naturalmente. No centro desses prédios existe SEMPRE uma grande bola de aço presa por uma corda de aço no teto e pendurada há uma certa altura(metade do prédio normalmente) suspensa no ar. Esta bola balança conforme o prédio "inclina-se" fazendo o papel de ajuste do centro de equilibrio do prédio.

O conceito de No Operation do caso da "bola de aço" vem do fato de que essa bola não exige NENHUMA manutenção. Qualquer modificação de engenharia feita no prédio(tirando modificação da altura do mesmo) não resultam em mudanças na posição da bola. A bola apenas existirá lá, no centro do prédio, fazendo seu papel. Não exigindo nenhuma manutenção na "bola", é isso a que damos o nome do padrão NoOps (No Operation).

Trazendo o exemplo para o nosso mundo do software, o padrão NoOps, consiste em estruturas de código que são necessárias para o sistema funcionar de acordo como foi desenhado, porém, não exigem nenhuma manutenção :). Por exemplo, a struct noOpsPlanetFinderService é exatamente a implementação NoOps da interface PlanetFinderService. Essa struct implementa a interface e provê comportamentos básicos usados em TODOS os testes unitários desse projeto. Basicamente, as instâncias NoOps, fazem todo o papel de base de retorno de "mocks" utilizados nos testes unitários de maneira que MUITO menos código precisa ser escrito para escrever um teste unitário que precisa de Mocks da interface PlanetFinderService por exemplo, que é exatamente o caso dos testes unitários da entidade HttpService. Por sua vez, a interface HttpService também tem a sua implementação NoOps, o que permite que qualquer outro componente que precisa de mocks de HttpService utilize dessa implementação sem muitas complexidades.

TODO: Working on..

