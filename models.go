package main

func dbCreateArticle(article *Article) error {
	query, err := db.Prepare(`insert into articles(title, content) values (?, ?)`)

	if err != nil {
		return err
	}

	defer query.Close()

	_, err = query.Exec(article.Title, article.Content)

	if err != nil {
		return err
	}

	return nil
}

func dbGetAllArticles() ([]*Article, error) {
	query, err := db.Prepare(`select id, title, content from articles`)

	if err != nil {
		return nil, err
	}

	defer query.Close()

	result, err := query.Query()

	if err != nil {
		return nil, err
	}

	articles := make([]*Article, 0)

	for result.Next() {
		data := new(Article)
		err := result.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)

		if err != nil {
			return nil, err
		}

		articles = append(articles, data)
	}

	return articles, nil
}
