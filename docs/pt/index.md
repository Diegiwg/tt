# Time Tracker CLI

[Documentação em Inglês](https://diegiwg.github.io/tt)

## Visão Geral do Projeto 🚀

Este é um programa de linha de comando simples para rastrear o tempo gasto em atividades. Ele oferece funcionalidades básicas para iniciar, pausar, retomar, parar e mostrar o tempo registrado.

## Download

Para baixar este programa, acesse a página de [lançamentos](https://github.com/Diegiwg/tt/releases)

## Instalação

Para usar este programa, você precisa ter o Go instalado no seu sistema. Uma vez que o Go esteja configurado, você pode instalar o programa com o seguinte comando:

```bash
go install github.com/Diegiwg/tt@latest
```

Certifique-se de que o diretório `$GOPATH/bin` está adicionado ao seu `$PATH` para que você possa executar o comando `tt` de qualquer lugar no seu sistema.

## Uso

### Comandos Disponíveis

1. **start**: Inicia um novo registro de tempo.
   - **Descrição**: Limpa o banco de dados e inicia um novo registro de tempo.
   - **Uso**: `tt start`

2. **pause**: Adiciona uma pausa ao registro de tempo.
   - **Descrição**: Quando você quiser fazer uma pausa, use este comando.
   - **Uso**: `tt pause`

3. **resume**: Retoma a contagem do tempo após uma pausa.
   - **Descrição**: Quando você quiser retomar após uma pausa, use este comando.
   - **Uso**: `tt resume`

4. **stop**: Para a contagem do tempo, limpando o registro de tempo.
   - **Descrição**: Quando você quiser finalizar o registro de tempo, use este comando.
   - **Uso**: `tt stop`

5. **show**: Mostra o tempo passado no registro de tempo.
   - **Descrição**: Quando você quiser ver quanto tempo passou sem finalizar o registro, use este comando.
   - **Uso**: `tt show`

6. **list**: Lista todos os registros de tempo.
   - **Descrição**: Quando você quiser ver todos os registros de tempo, use este comando.
   - **Uso**: `tt list [--limit: int]`

### Exemplo de Uso

```bash
tt start
# Algum tempo se passou
tt pause
# Fez uma pausa
tt resume
# Trabalhou mais um pouco
tt show
# Verifique que em 10 minutos é hora de parar de trabalhar
tt stop
# Termina o trabalho de hoje e vê o resultado na tela de quantas horas se passaram
tt show
# Veja o registro anterior na lista de registros cadastrados
```

## Contribuindo

Este é um projeto de código aberto, e contribuições são bem-vindas! Sinta-se à vontade para fazer um fork deste repositório, implementar melhorias e enviar um pull request.

## Licença 📜

Este projeto é licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
