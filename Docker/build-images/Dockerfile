# syntax=docker/dockerfile:1 

# Sempre aponta para a versão mais recente da sintaxe da versão 1. O BuildKit verifica automaticamente as atualizações da sintaxe antes de compilar, garantindo que você esteja usando a versão mais atual.
FROM golang:1.16-alpine

#O WORKDIR define o diretório de trabalho do container.
WORKDIR /app 

# Copia um arquivo ou diretório para o container.
COPY go.mod ./
COPY go.sum ./

# Isso funciona exatamente como se estivéssemos executando o go localmente em nossa máquina, mas desta vez esses módulos Go serão instalados em um diretório dentro da imagem.
RUN go mod download

# A próxima coisa que precisamos fazer é copiar nosso código-fonte na imagem. Usaremos o comando COPY assim como fizemos com nossos arquivos de módulo antes
COPY *.go ./

# Compila a aplicação.
# O resultado desse comando será um binário de aplicativo estático chamado docker-gs-ping
RUN go build -o /docker-gs-ping

# Diga ao Docker qual comando executar quando nossa imagem for usada para iniciar um contêiner
CMD [ "/docker-gs-ping" ]


# O comando docker build cria imagens do Docker a partir do Dockerfile

# O comando build opcionalmente recebe um sinalizador --tag. Esse sinalizador é usado para rotular a imagem 
# com um valor de string, que é fácil para os humanos lerem e reconhecerem. Se você não passar um --tag, o Docker usará o mais recente como o valor padrão

# exemplo:  docker build --tag docker-gs-ping .

# rodar a imagem: docker run docker-gs-ping
# Ao executar este comando curl http://localhost:8080/, você notará que não retornou ao prompt de comando. Isso ocorre porque nosso aplicativo é um servidor REST 
# e será executado em um loop aguardando solicitações recebidas sem retornar o controle de volta ao sistema operacional até interrompermos o contêiner.

#

# altera o nome da tag da imagem 
# docker image tag docker-gs-ping:latest docker-gs-ping:v1.0
# quando a imagem tem o mesmo id não tá criando uma nova imagem apenas alterando a tag


# Para publicar uma porta para nosso contêiner, usaremos o sinalizador --publish (-p para abreviar) no comando 
#docker run. O formato do comando --publish é [host_port]:[container_port]. Portanto, se quiséssemos expor a porta 8080 dentro do 
#contêiner para a porta 3000 fora do contêiner, passaríamos 3000:8080 para o sinalizador --publish.
# exemplo: docker run --publish 3000:8080 docker-gs-ping  onde 3000 é a porta do host e 8080 é a porta do container


# ======================= Run in detached mode

#nosso aplicativo de exemplo é um servidor web e não devemos ter nosso terminal conectado ao contêiner. O Docker pode executar seu contêiner no modo desanexado, 
#ou seja, em segundo plano. Para fazer isso, podemos usar o --detach ou -d para abreviar. O Docker iniciará seu contêiner da mesma forma que antes, mas desta 
#vez será “desanexado” do contêiner e retornará ao prompt do terminal
# exemplo : docker run --detach --publish 3000:8080 docker-gs-ping


# ======================= Run 

# docker run --detach --publish 3000:8080 --name rest-server docker-gs-ping



# ===================== List Containers

# docker ps
# Lista todos os contêiners que estão rodando no Docker

# docker ps -a
# Lista todos os contêiners que estão rodando no Docker, incluindo os que foram removidos

# docker ps -l
# Lista todos os contêiners que estão rodando no Docker, excluindo os que foram removidos

# docker ps -q
# Lista todos os IDs dos contêiners que estão rodando no Docker

# docker ps -f status=running
# Lista todos os contêiners que estão rodando no Docker e que estão em execução

# docker restart [container_name]
# Reinicia um contêiner
