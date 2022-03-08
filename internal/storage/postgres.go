package storage

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

func (s *Storage) GetAll() (interface{}, error) {
	ctx := context.Background()

	cached, err := s.CacheGet(ctx, "todos")
	if err == nil {
		return cached, nil
	}

	rows, err := s.db.Query(ctx, GetAll)
	if err != nil {
		return nil, err
	}

	var todos []interface{}
	for rows.Next() {
		var todo TODO
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.DueDate)
		if err != nil {
			log.Println("Got error:", err)
		}

		todos = append(todos, todo)
	}

	json, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}

	if err := s.CacheSet(ctx, "todos", json); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *Storage) GetRemaining() (interface{}, error) {
	ctx := context.Background()

	cached, err := s.CacheGet(ctx, "todos-remaining")
	if err == nil {
		return cached, nil
	}

	rows, err := s.db.Query(ctx, GetRemaining)
	if err != nil {
		return nil, err
	}

	var todos []interface{}
	for rows.Next() {
		var todo TODO
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.DueDate)
		if err != nil {
			log.Println("Got error:", err)
		}

		todos = append(todos, todo)
	}

	json, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}

	if err := s.CacheSet(ctx, "todos-remaining", json); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *Storage) GetDone() (interface{}, error) {
	ctx := context.Background()

	cached, err := s.CacheGet(ctx, "todos-done")
	if err == nil {
		return cached, nil
	}

	rows, err := s.db.Query(ctx, GetDone)
	if err != nil {
		return nil, err
	}

	var todos []interface{}
	for rows.Next() {
		var todo TODO
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.DueDate)
		if err != nil {
			log.Println("Got error:", err)
		}

		todos = append(todos, todo)
	}

	json, err := json.Marshal(todos)
	if err != nil {
		return nil, err
	}

	if err := s.CacheSet(ctx, "todos-done", json); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *Storage) Get(key string) (interface{}, error) {
	ctx := context.Background()

	cached, err := s.CacheGet(ctx, key)
	if err == nil {
		return cached, nil
	}

	var todo TODO
	err = s.db.QueryRow(ctx, Get, key).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.DueDate)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(todo)
	if err != nil {
		return nil, err
	}

	err = s.CacheSet(ctx, key, data)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Storage) Create(id, title, desc, status string, created, due time.Time) (interface{}, error) {
	_, err := s.db.Exec(context.Background(), Insert, id, title, desc, status, created, due)
	if err != nil {
		return nil, err
	}

	s.CacheDelete(context.Background(), id)

	todo, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Storage) UpdateStatus(id, status string) (interface{}, error) {
	_, err := s.db.Exec(context.Background(), UpdateStatus, id, status)
	if err != nil {
		return nil, err
	}

	s.CacheDelete(context.Background(), id)

	todo, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *Storage) Delete(key string) error {
	ctx := context.Background()
	s.CacheDelete(ctx, key)
	_, err := s.db.Query(ctx, Delete, key)
	return err
}
