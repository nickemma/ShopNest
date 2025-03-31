# **E-commerce Platform Models**

## **User Schema**

### 1. Customer Model
```json
{
  "customerId": "string",
  "name": "string",
  "email": "string",
  "phone": "string",
  "address": {
    "street": "string",
    "city": "string",
    "state": "string",
    "postalCode": "string",
    "country": "string"
  },
  "status": "string",
  "createdAt": "string",
  "updatedAt": "string"
}
```

##### Example

```json
{
  "customerId": "12345",
  "name": "John Doe",
  "email": "johndoe@example.com",
  "phone": "+1234567890",
  "address": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postalCode": "10001",
    "country": "USA"
  },
  "status": "active",
  "createdAt": "2025-03-31T12:00:00Z",
  "updatedAt": "2025-03-31T12:30:00Z"
}
```

### 2. Manager Model

```json
{
  "managerId": "string",
  "name": "string",
  "email": "string",
  "role": "string",
  "createdAt": "string",
  "updatedAt": "string"
}
```

##### Example

```json
{
  "managerId": "MGR001",
  "name": "Alice Johnson",
  "email": "alice@example.com",
  "role": "admin",
  "createdAt": "2025-03-31T09:00:00Z",
  "updatedAt": "2025-03-31T10:00:00Z"
}
```

### 3. Product Model

```json
{
  "productId": "string",
  "name": "string",
  "description": "string",
  "price": "number",
  "category": "string",
  "imageUrl": "string",
  "status": "string",
  "createdAt": "string",
  "updatedAt": "string"
}
```

##### Example

```json
{
  "productId": "98765",
  "name": "Wireless Headphones",
  "description": "Noise-canceling over-ear headphones.",
  "price": 199.99,
  "category": "Electronics",
  "imageUrl": "https://example.com/images/headphones.jpg",
  "status": "available",
  "createdAt": "2025-03-31T10:00:00Z",
  "updatedAt": "2025-03-31T11:00:00Z"
}
```

### 4. Inventory Model

```json
{
  "inventoryId": "string",
  "productId": "string",
  "warehouseLocation": "string",
  "stockQuantity": "number",
  "reservedQuantity": "number",
  "updatedAt": "string"
}
```

##### Example

```json
{
  "inventoryId": "INV001",
  "productId": "98765",
  "warehouseLocation": "Warehouse A - Section B",
  "stockQuantity": 50,
  "reservedQuantity": 10,
  "updatedAt": "2025-03-31T11:30:00Z"
}
```

### 5. Order Model

```json
{
  "orderId": "string",
  "customerId": "string",
  "items": [
    {
      "productId": "string",
      "quantity": "number",
      "price": "number"
    }
  ],
  "totalAmount": "number",
  "orderStatus": "string",
  "paymentStatus": "string",
  "shippingAddress": "object",
  "createdAt": "string",
  "updatedAt": "string"
}
```

##### Example

```json
{
  "orderId": "ORD12345",
  "customerId": "12345",
  "items": [
    {
      "productId": "98765",
      "quantity": 2,
      "price": 199.99
    }
  ],
  "totalAmount": 399.98,
  "orderStatus": "pending",
  "paymentStatus": "unpaid",
  "shippingAddress": {
    "street": "123 Main St",
    "city": "New York",
    "state": "NY",
    "postalCode": "10001",
    "country": "USA"
  },
  "createdAt": "2025-03-31T12:15:00Z",
  "updatedAt": "2025-03-31T12:15:00Z"
}
```

### 6. Cart Model

```json
{
  "cartId": "string",
  "customerId": "string",
  "items": [
    {
      "productId": "string",
      "quantity": "number",
      "price": "number",
      "productName": "string"
    }
  ],
  "totalAmount": "number",
  "status": "string",
  "createdAt": "string",
  "updatedAt": "string"
}
```

##### Example

```json
{
  "cartId": "CART12345",
  "customerId": "12345",
  "items": [
    {
      "productId": "98765",
      "quantity": 2,
      "price": 199.99,
      "productName": "Wireless Headphones"
    }
  ],
  "totalAmount": 399.98,
  "status": "active",
  "createdAt": "2025-03-31T12:15:00Z",
  "updatedAt": "2025-03-31T12:30:00Z"
}
```


### 7. Payment Intent Model

```json
{
  "paymentId": "string",
  "orderId": "string",
  "customerId": "string",
  "amount": "number",
  "currency": "string",
  "status": "string",
  "createdAt": "string",
  "updatedAt": "string"
}
```

##### Example

```json
{
  "paymentId": "PAY67890",
  "orderId": "ORD12345",
  "customerId": "12345",
  "amount": 399.98,
  "currency": "USD",
  "status": "pending",
  "createdAt": "2025-03-31T12:20:00Z",
  "updatedAt": "2025-03-31T12:20:00Z"
}
```

### 8. Pricing Rule Model

```json
{
  "ruleId": "string",
  "productId": "string",
  "discountPercentage": "number",
  "minQuantity": "number",
  "validFrom": "string",
  "validTo": "string",
  "status": "string"
}
```

##### Example

```json
{
  "ruleId": "RULE123",
  "productId": "98765",
  "discountPercentage": 10,
  "minQuantity": 2,
  "validFrom": "2025-03-01T00:00:00Z",
  "validTo": "2025-03-31T23:59:59Z",
  "status": "active"
}
```

### 9. Recommendation Model

```json
{
  "recommendationId": "string",
  "customerId": "string",
  "recommendedProducts": ["string"],
  "generatedAt": "string"
}
```

##### Example

```json
{
  "recommendationId": "REC98765",
  "customerId": "12345",
  "recommendedProducts": ["98765", "54321"],
  "generatedAt": "2025-03-31T13:00:00Z"
}
```