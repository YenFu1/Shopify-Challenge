### Instructions

To access the app, use a program like postman to send requests to the following endpoints using baseURL: `https://shopify-challenge.yenfu1.repl.co`

1. To create an inventory:
    - url: `POST: {baseURL}/api/inventories`
    - payload:
    ```
    {
        "name": string
    }
    ```

2. To get a list of all inventories:
    - url: `GET: {baseURL}/api/inventories`

3. To delete an inventory:
    - url: `DELETE: {baseURL}/api/inventories/{inventoryId}`

4. To update an inventory:
    - url: `PUT: DELETE: {baseURL}/api/inventories/{inventoryId}`
    - payload:
    ```
    {
        "name": string
    }
    ```

5. To get all items belonging to an inventory:
    - url: `GET: {baseURL}/api/inventories/{inventoryId}/items`

6. To create an item:
    - url: `POST: {baseURL}/api/inventories/{inventoryId}/items`
    - payload:
    ```
    {
        "name": string
        "count": int
    }
    ```

7. To delete an item: 
    - url: `DELETE: {baseURL}/api/inventories/{inventoryId}/items/{itemId}`

8. To update an item:
    - url: `PUT: {baseURL}/api/inventories/{inventoryId}/items/{itemId}`
    - payload: 
    ```
    {
        "name": string
        "count" int
    }
    ```

9. To create a shipment:
    - url: `POST: {baseURL}/api/shipments`
    - payload:
    ```
    {
        "name": string
        "inventoryId": string // must be an actual inventoryId in the database
        "items": {
            "id": int //must be an actual itemId in the database
            "count" int //amount to be moved to shipment, returns status code 500 if not enough items in stock
        }
    }
    ```

10. To get a list of all shipments:
    - url: `GET: {baseURL}/api/shipments`
 