# Go-BenderBeerAPI
Go-BenderBeerAPI

## Deployment

To deploy this project run

```bash
  docker-compose up 
```
## API Reference

#### Post Create Beer

```http
  POST /create-beer
{
    "name": "Balboa",
    "brewery": "Brewery",
    "country": "USA",
    "price": 5,
    "currency": 1
}
```
#### Get item

```http
  GET /get-beer/${id}
```
