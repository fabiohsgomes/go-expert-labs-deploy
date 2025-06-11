# go-expert-labs-deploy
Laboratório que tem objetivo de publicar uma Api Rest usando o Google Claud Run

## Executando o projeto
- Para testar as urls estão disponíveis na no arquivo api/api.http
- Para subir o container de desenvolvimento localmente, basta executar o comando abaixo:
```bash
docker compose up -d
```

## Link do deploy no CloudRun

 - cep válido : https://deploy-cloudrun-previsao-237975055910.us-west1.run.app/cidades/05893130/temperaturas
 - cep não encontrado : https://deploy-cloudrun-previsao-237975055910.us-west1.run.app/cidades/00000000/temperaturas
 - cep inválido : https://deploy-cloudrun-previsao-237975055910.us-west1.run.app/cidades/000000a0/temperaturas