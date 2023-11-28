
# cashier-service 

## Folder Structure Conventions

This project follows the hexagonal architecture pattern.

```plaintext

/cmd
    main.go           # Application entry point
/handler              # HTTP request handlers
/log                  # set Logger
/repository           # Interfaces for data storage & adapter sevice external
/service              #Business logic services

```

## Setup

1. Clone this repository:
    ```bash
   $ git clone https://github.com/MosChanvit/cashier.git
    ```
2. up: docker compose up:
    ```bash
   $ docker-compose up -d
    ```

3. Test API 1 with curl
    
    * api cal_xyz
    ```bash
    curl --location --request POST 'http://localhost:80/cal_xyz?numbers=1%2CX%2C8%2C17%2CY%2CZ%2C78%2C113'
    ```

4. Test API 2 with curl
    * api create cashier
    ```bash
        curl --location 'http://localhost:80/cashier' \
        --header 'Content-Type: application/json' \
        --data '{
            "name": "QC",
            "c1000": 10,
            "c500": 5,
            "c100": 5,
            "c50": 5,
            "c20": 10,
            "c10": 10,
            "c5": 10,
            "c1": 10,
            "c025": 10
        }'
    ```
        note--> 
        -set "name" required
        -"c1000": 5   means there are 5 1,000 baht bills


    * api Get cashier by name, Returns cashier information and payment transitions.
    ```bash
        curl --location 'http://localhost:80/cashier?name=QC'
    ```
    * api Get cashier All, Returns cashier information all.
    ```bash
        curl --location 'http://localhost:80/cashiers'
    ```
    * api calculate the change money when the customer pays
    ```bash
        curl --location 'http://localhost:80/pay' \
        --header 'Content-Type: application/json' \
        --data '{
            "name": "QC",
            "c1000": 1,
            "c500": 0,
            "c100": 1,
            "c50": 0,
            "c20": 2,
            "c10": 0,
            "c5": 0,
            "c1": 0,
            "c025": 0,
            "product_price": 1090,
            "customer_paid":1140
        }'
    ```
        note--> 
        -set "name" required
        -"c1000": 0   means there are 0 1,000 baht bills
        -product_price  = Set the customer_paid value to match the amount information entered.


