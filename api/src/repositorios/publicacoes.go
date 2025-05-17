package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type publicacoes struct {
	db *sql.DB
}

func NovoRepositorioPublicacao(db *sql.DB) *publicacoes {
	return &publicacoes{db}
}

func (repositorio publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio publicacoes) BuscarPublicacao(ID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
	SELECT p.*
	,u.nick 
	FROM publicacoes p
		INNER JOIN usuarios u 
		ON u.id = p.autor_id
	WHERE p.id = ?
	`, ID)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}
	return publicacao, nil
}

func (repositorio publicacoes) BuscarPublicacoes(ID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	SELECT DISTINCT p.*, u.nick
	FROM publicacoes p
		INNER JOIN usuarios u 
			ON u.id = p.autor_id
		INNER JOIN seguidores s 
			ON p.autor_id = s.usuario_id
	WHERE u.id = ? 
	OR s.seguidor_id = ?
	`, ID, ID)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (repositorio publicacoes) Atualizar(ID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?")
	if erro != nil {
		return erro
	}

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio publicacoes) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM publicacoes  WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio publicacoes) BuscarPorUsuario(ID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		SELECT p.*, 
		u.nick 
		FROM publicacoes p
			JOIN usuarios u 
				ON u.id = p.autor_id
		WHERE p.autor_id = ?
		`, ID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {

		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}
