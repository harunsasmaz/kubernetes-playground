#!/bin/sh

base_url=$1

insert_new_data() {
  payload=$( printf '{"id": "%d", "title": "title-%d", "description": "desc-%d", "due_date": %d}' $1 $1 $1 $(date +%s) )
  curl -s -X POST --url $base_url/todos --data "$payload" --header 'Content-Type: application/json' -o /dev/null 
}

for idx in {1..1000}
do
  insert_new_data $idx &
done