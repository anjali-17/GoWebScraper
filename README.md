# GoWebScraper

a service that will make a background request, fetch the respective web site, parse it and give back a valid result to the client. When requesting an ID from this API, we want to get back accessible and meaningful results in the JSON format provided below.
Do not use a database to store or query intermediate results. The result should be fetched and calculated in real time.
The scraper should act as a REST API, listening on http://localhost:8080. 
Example URLs to be fetched:
- http://www.amazon.de/gp/product/B00KY1U7GM 
- http://www.amazon.de/gp/product/B00K19SD8Q
The last part of the URL is called an Amazon ID: e.g. `B00KY1U7GM`
When requesting an Amazon ID at the route http://localhost:8080/movie/amazon/{amazon_id}, the API needs to return the following JSON, e.g. when using amazon_id `B00K19SD8Q`:
{
    "title":"Um Jeden Preis", 
    "release_year":2013, 
    "actors":[
        "Dennis Quaid",
        "Zac Efron"
        ]
}


- go run server.go
