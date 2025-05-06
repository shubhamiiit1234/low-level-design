package main

/*

Question
	- question_id (PK) UUID
	- heading
	- asked_at
	- updated_at
	- views
	- description
	- tags
	- vote_count
	- asked_by
	- answer_count


POST /question/search
req body -
	{
		"data": text
	}
res body -
	{
		"response": [
						{
							"questionid": 1,
							"heading": "asdf"
						}
					],
		"questioncount": 2
	}

GET /question/{questionid}
res body -
	{
		"question": {
						"questionid": 1
					},
		"answer": [
					{
						"answerid": 1
					}
					],
		"count": 1
	}

POST /question/answer
request -
	{
		"questionid": 1,
		"answer": text
	}
response - success/failure

Answer
	- answer_id (PK)
	- description
	- vote_count
	- is_accepted
	- answered_at
	- updated_at
	- asnwered_by
	- question_id (FK)
	- accepted_at


POST /answer/upload/{questionid}



*/
