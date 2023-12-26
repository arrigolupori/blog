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

func dbGetPost(postID string) (*BlogPost, error) {
	query, err := db.Prepare("select id, title, content from posts where id = ?")

	if err != nil {
		return nil, err
	}

	defer query.Close()

	result := query.QueryRow(postID)
	data := new(BlogPost)
	err = result.Scan(&data.ID, &data.Title, &data.Content)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func dbUpdatePost(id string, post *BlogPost) error {
	query, err := db.Prepare("update posts set (title, content) = (?,?) where id=?")
	defer query.Close()

	if err != nil {
		return err
	}
	_, err = query.Exec(post.Title, post.Content, id)

	if err != nil {
		return err
	}

	return nil
}

func dbDeletePost(id string) error {
	query, err := db.Prepare("delete from posts where id=?")

	if err != nil {
		return err
	}

	defer query.Close()

	_, err = query.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
