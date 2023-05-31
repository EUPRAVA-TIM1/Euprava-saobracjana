# SSOService
Road police service as part of e_uprava class project assigment. Developed by [petar__k](https://www.linkedin.com/in/petar-komord%C5%BEi%C4%87-23765a233/)
## Getting started
To run this project you will need:
- [Docker+Docker compose](https://www.docker.com/)
- [GoLang](https://go.dev/dl/) version: 1.18 

To start your service you can see attached docker-compose.yaml that I wrote for all services(including this one)
in [this repo](https://github.com/EUPRAVA-TIM1/DockerCompose). Ports can be seen in `.env` there or in `config.go` file of this project

## Endpoints(for other services)

- `GET /saobracajna/Gradjanin/Nalozi/{jmbg}` Gets all user prekrsajniNalozi.\
**Expects** `Authorization` header with JWT token(with or without Barrer) \
**Returns** Json arrat of nalozi in this format:
```
{
    "id": int,
    "datum": date,
    "opis": string,
    "izdatoOdStrane": string,
    "izdatoZa": string,
    "JMBGZapisanog": string,
    "tipPrekrsaja" : string,
    "jedinicaMere" : string | null,
    "vrednost": string | null,
    slike: []string,
    kaznaIzvrsena bool
}
```
- `GET /saobracajna/Gradjanin/Nalozi/{jmbg}` Gets all user prekrsajniNalozi.\
**Expects** `Authorization` header with JWT token(with or without Barrer) and JMBG of a gradjanin as a url param \
**Returns** Json arrat of nalozi in this format:
```
{
"id": int,
"datum": date,
"opis": string,
"izdatoOdStrane": string,
"izdatoZa": string,
"JMBGZapisanog": string,
"tipPrekrsaja" : string,
"jedinicaMere" : string | null,
"vrednost": string | null,
slike: []string,
kaznaIzvrsena bool
}
```
- `PUT /saobracajna/Policajac/Sud/Nalozi/Status/{id}` Upadtes status of user Sud Nalog.\
  **Expects** `Authorization` header with JWT token(with or without Barrer) id of nalog as url param and json with status in this format:
```
{
"status": string
}
```
**Returns** Status code OK or Err