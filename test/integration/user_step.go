package integration

import "github.com/cucumber/godog"

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^os usuarios salvos no banco de dados$`, provisionarCenario)
	ctx.Step(`^quando eu chamar a api na rota /users$`, chamarApi)
	ctx.Step(`^deve retornar os usuarios salvos no banco de dados$`, verificar)
}

func provisionarCenario() {

}

func chamarApi() {

}

func verificar() {

}
