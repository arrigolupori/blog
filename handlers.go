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
