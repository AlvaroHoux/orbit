# Especificação Técnica

## Conceitos

- **Peer:** Qualquer dispositivo rodando o Orbit.
- **Seed:** Um Peer que possui a cópia **completa e atualizada** de um cofre.
- **Vault:** A pasta do Obsidian que estamos sincronizando. É a unidade principal de dados.
- **Handshake:** O processo inicial onde dois Peers se conectam, trocam chaves de criptografia e verificam se têm permissão para conversar.

---

## Detalhes dos Comandos

### `orbit serve`

**Descrição**: Inicia o processo em segundo plano (daemon) que escuta a rede e sincroniza os arquivos. Sem isso, a sincronização não acontece.

### `orbit list`

**Descrição**: Exibe todos os vaults e o estado.
**Output esperado**:

```text
ID       NAME            PATH                            STATUS
8f4a2c   Brain           ~/Documents/Obsidian/Brain      🟢 Online (Syncing)
a1c3     Work            ~/Documents/Work                🔴 Offline
```

### `orbit add [path]`
**Flags**
- `-n, --name`: Nome amigável para o vault.

**Descrição**: Inicializa um diretório como um vault Orbit.
**Comportamento:**
- Cria a pasta .orbit.
- Solicita senha. Se vazia, gera uma chave aleatória.
- Exibe o ID único.

**Output Esperado**:
```
🔒 Protegendo vault "Brain"...
Digite uma senha secreta (Deixe vazio para gerar auto): ***********
✨ Vault Criado!
   ID Público: 8f4a2c
   Status: Aguardando Peers...
```

### `orbit join [id]`

**Descrição**: Descrição: Conecta a um vault existente na rede local.
**Comportamento**: Escuta broadcasts UDP procurando pelo ID. Se achar, tenta o handshake criptográfico com a senha.
**Output Esperado**:

```
Conectando ao Vault ID 8f4a2c...
Digite a senha do vault: ***********
🔍 Buscando peers na rede...
✅ Handshake realizado com sucesso com "Notebook-Dell"!
🔄 Iniciando sincronização...
```

### `orbit stop [id]`

**Descrição**: Interrompe a sincronização de um vault com outros dispositivos.
**Comportamento**:
1. O sistema para de monitorar alterações locais (não envia nada).
2. O sistema ignora solicitações de peers externos para este vault (não recebe nada).
3. O status muda para "Stopped".

### `orbit start [id]`

**Descrição**: Retoma a sincronização de um vault que foi parado anteriormente.
**Comportamento**: Reativa o watcher de arquivos e volta a anunciar o ID na rede.

### `orbit remove [id]`

**Descrição**: Remove o vault da lista de sincronização e apaga as configurações locais do Orbit (não apaga os arquivos do usuário).
