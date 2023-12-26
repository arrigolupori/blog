package main

func dbCreatePost(post *BlogPost) error {
	query, err := db.Prepare(`insert into posts(title, content) values (?, ?)`)

	if err != nil {
		return err
	}

	defer query.Close()

	_, err = query.Exec(post.Title, post.Content)

	if err != nil {
		return err
	}

	return nil
}

func dbGetAllPosts() ([]*BlogPost, error) {
	query, err := db.Prepare(`select id, title, content from posts`)

	if err != nil {
		return nil, err
	}

	defer query.Close()

	result, err := query.Query()

	if err != nil {
		return nil, err
	}

	posts := make([]*BlogPost, 0)

	for result.Next() {
		data := new(BlogPost)
		err := result.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, data)
	}

	return posts, nil
}
