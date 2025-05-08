package main

/*

Movie
	- movie_id			int (PK)		//will be UUID in real
	- movie_name		varchar
	- type				varchar
	- rating			float
	- review			varchar
	- vote_count		int
	- image				varchar
	- release_date		TimeStamp
	- time_duration		varchar
	- language			varchar
	- about				varchar
	- cast				varchar

POST /movies/upload
Request
	{
		"data": [
					{
						"movie_id": 1,
						"movie_name": "insidious",
						.
						.
					},
					{
						"movie_id": 2,
						"movie_name": "insidious-2",
						.
						.
					},
					.
					.
					.
				]

	}

Response - success/failure

GET movies/{cityName}
Request
Response
	{
		"data": [
					{
						"movie_id": 1,
						"movie_name": "insidious"
						.
						.
					},
					{
						"movie_id": 2,
						"movie_name": "insidious-2",
						.
						.
					},
					.
					.
					.
				]
	}

City
	- city_id		int (PK)
	- city_name		varchar
	- theatre_id	int (FK)
	- movie_id		int (FK)

Theatre
	- theatre_id	int (PK)
	- show_id	int (FK)
	- location
	- type		varchar (PVR, CINEPOLIS)

JOIN City and Theatre

Show
	- show_id	int (PK)	-	will be UUID in real
	- date		TimeStamp
	- time		Time
	- theatre_id	int	(FK)
	- price			float
	- movie_id		int	(FK)
	- audi			int
	- available_seats	int

POST /shows/addShow
Request
	{
		"data": [
					{
						"show_id": 1,	// will be passed from frontend
						"date": "25/05/2025",	// will be passed from frontend
						"time": 9:00,
						"theatre_id": 1,
						"price": 250.0,
						"movie_id": 1,
					},
					{
						"show_id": 2,
						"date": "25/05/2025",
						"time": 12:00,
						"theatre_id": 1,
						"price": 250.0,
						"movie_id": 1,
					},
					.
					.
					.
				]
	}

GET /shows/showAll/date={date}&movie={movie_id}
Request
Response
	{
		"data": [
					{
						"show_id": 1,
						"time": 9:00,
						"theatre_id": 1,
						"price": 250.0,
					},
					{
						"show_id": 2,
						"time": 12:00,
						"theatre_id": 1,
						"price": 250.0,
					},
					{
						"show_id": 3,
						"time": 9:00,
						"theatre_id": 2,
						"price": 300.0,
					},
					.
					.
					.
				]
	}



Booking
	- booking_id 	int (PK)
	- show_id	int (FK)
	- total_amount	float


POST /booking/create
Request
	{
		"data": {
					"booking_id": 1,
					"show_id": 1, // will be sent from frontend
					"total_amount": 1000.0,
					"total_seats": 4,
					"seat_type": "Recliner",
				}
	}

Response - success/failure

Transaction
	- transaction_id 	int (PK)
	- booking_id 		int (FK)
	- payment_mode		varchar
	- transaction_status	Boolean
	- amount			float

POST /transaction/make
Request
	{
		"data": {
					"transaction_id": 1,
					"booking_id": 1, 	// will be sent from frontend
					"payment_mode": "UPI",
					"amount": 1000.0,
				}
	}

GET /transaction/status/booking_id={booking_id}
Request
Response
	{
		"data": [
					{
						"transaction_id": 1,
						"booking_id": 1, // from url
						"payment_mode": "UPI",
						"transaction_status": "failed",
					},
					{
						"transaction_id": 1,
						"booking_id": 1, // from url
						"payment_mode": "UPI",
						"transaction_status": "success",
					}
				]
	}









*/
