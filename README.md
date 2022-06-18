# stan-api

Simple JSON service that handles POST requests of json data, filters the request and returns a json response

The endpoint accepts POST json data on the root index `/`

ENV Variables required can be found in `.env.template`
APIPORT - the port that the api will run on

These can be populated into a .env file and exported using `export $(grep -v '^#' .env | xargs)`