# Tax-Calc

This API implemented using GraphQL and Pure Go.

## Depedencies
[Docker](https://docs.docker.com/engine/installation/) :whale: & [docker-compose](https://docs.docker.com/compose/install/).

## Developing

run locally run using docker-compose:

```bash
sudo docker-compose up
```

the app runs on `localhost:8080`

## GraphQL Mutation

This project is developed using GraphQL on `localhost:8080/query`. 
For ease-of-use, I recommend using Insomnia (https://insomnia.rest/), then select GraphQL Query on Body Type
Here are two Mutation for Store Tax and Calculate Tax:
``` 
mutation{
  inputProduct(product: {
    taxCode: 2,
    name: "Test 2",
    price: 2000
  }) {
    id
    name
    price
    taxCode
  }
}
```

```
mutation{
  calculateTax(productIDs: [1,2]){
    products{
      id
      name
      price
      taxCode
      refundable
      taxPrice
      productType
      total
    }
    priceSubTotal
    taxSubTotal
    grandTotal
  }
}
```
## Database Structure
I use 2 table on this project:
1. Product
    a. ID (PK, int, autogenerated)
    b. Name (varchar)
    c. Price (decimal)
    d. TaxID (int, FK)
2. Tax
    a. ID (PK, int, autogenerated)
    b. Code (int)