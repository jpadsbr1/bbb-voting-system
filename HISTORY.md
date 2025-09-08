# Histórico

Comecei já direto optando pela Clean Architecture e estruturando as pastas

Tive dúvidas de deixava "usecases" ou "services"

E também se utilizava a nomenclatura da arquitetura hexagonal "ports" e "adapters"

Optei por deixar os contratos de repositório dentro dos arquivos de domínio, a fim de condensar o entendimento

NoSQL ou Postgres? NoSQL seria bom pra dados brutos e em massa, útil para votos, Postgres poderia ser mais útil para rastreio, dados temporais e relatórios. Seria possível no futuro implementar os dois, um para dados em massa e outro pra ajudar nos relatórios. Vou começar com Postgres.

Começando a configurar contêiner. Primeiro o Dockerfile local, já com multibuild pra imagem ficar pequena.

Já vou iniciar desacoplando tudo no main.go, com arquivo específico de config na pasta internal para deixar tudo mais legível.

Eu poderia abstrair as configurações de handler e rotas pra uma função separada, mas, na minha opinião, já achei o server.go bem legível e as rotas separadas por domínio.

Sobre o cache, to em dúvida se deixo o domínio vote_cache embutido no vote.go.

Deixei separado, visto que são lógicas complementares, mas ainda separadas, uma pra interação do cache e outra do postgres. E em caso de troca de ferramenta, já está tudo mais fácil de organizar.

Vou instituir as migrations. Como são apenas 4 tabelas, sendo uma relacional, vou deixar uma migration só para instauração das 4 tabelas de uma vez. Em caso de crescimento de regras de negócio, vale instituir uma migration pra cada tabela, mas isso pode ser feito depois.

Não sei se deixo as migrations automáticas via docker compose. Como ainda estou desenvolvendo, vou deixar manual, pra fins de teste, mas fica aqui o TODO.

Tratei as possibilidades de erro que lembrei, é possível que eu tenha esquecido alguma, mas, nesse caso, é só adicionar a possibilidade no usecase de cada serviço.

Sistema de cache com flusher implementado. A cada voto, uma contagem é incrementada no cache. A cada 5 segundos um flusher faz a inserção desses votos no Postgres e depois reseta a contagem, isso ajuda a aumentar o throughput.

O Redis já foi configurado com mínima perda de dados com RDB e AOF.

Agora devo implementar uma fila com Kafka, usando a biblioteca open source Sarama. Muito boa.

O objetivo é: a cada voto, subir uma mensagem no publisher. O consumer vai ler a mensagem e rodar o mesmo processo de cache e flusher.

Existe a possibilidade de deixar o cache antes do Kafka, pra diminuir a latência e aproveita a rapidez do Redis, mas a fim de diminuir ao máximo a perda de votos e aumentar a confiabilidade do sistema.

Com tudo pronto e funcionando, só falta implementar o Kafka, de fato antes do processo de cache, é o próximo TODO.

Depois de tudo, o objetivo é começar a provisionar a infraestrutura usando Terraform. A plataforma de nuvem ideal seria a AWS, usando o próprio fluxo de CI/CD do GitHub em integração pra fazer o deploy a cada commit. Vale instituir branches de Desenvolvimento e Produção e um ambiente de Staging que imita Produção também.

É possível, num futuro, migrar o banco de dados para um AWS RDS talvez? Mas isso poderia acrescentar uma dependência de nuvem muito grande. Portanto, um serviço pra cada uma das Ferramentas é o ideal por enquanto. Com tudo configurado, o load balancer já estaria trafegando bem os requests e distribuindo carga para suportar milhares de votos por segundo.

Eu usaria, quando estivesse pronto, o k6s para fazer um teste de carga, pra saber o quanto o sistema aguenta.

Vale, por último, implementar um sistema de logs mais sofisticado com ELK, ElasticSearch, Logstash e Kibana, pra manter um controle maior dos logs.

E em conjunto, setar um Prometheus com Grafana pra monitoramento de hardware de todo o sistema.

Com isso o sistema estaria funcional e suportando uma enorme quantidade de dados ao mesmo tempo.

(Me diverti à beça fazendo, e devo continuar agora, só pelo entretenimento)