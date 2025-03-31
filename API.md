# E-commerce Platform API Endpoints

## User Service

### Customers

| Action              | Method  | Endpoint                                     | Description                           |
|--------------------|--------|---------------------------------------------|---------------------------------------|
| Register customer  | POST   | `/api/v1/customers`                         | Create a new customer account.       |
| Update customer    | PATCH  | `/api/v1/customers/{customerId}`            | Update customer details.             |
| Get customer       | GET    | `/api/v1/customers/{customerId}`            | Retrieve customer details.           |
| List customers     | GET    | `/api/v1/customers`                         | Get a paginated list of customers.   |
| Delete customer    | DELETE | `/api/v1/customers/{customerId}`            | Soft delete or remove a customer account. |
| Deactivate customer | PATCH  | `/api/v1/customers/{customerId}/deactivate` | Deactivate a customer account.       |

### Managers (Admins)

| Action            | Method  | Endpoint                                     | Description                           |
|------------------|--------|---------------------------------------------|---------------------------------------|
| Register manager | POST   | `/api/v1/managers`                         | Create a new manager (admin) account. |
| Update manager   | PATCH  | `/api/v1/managers/{managerId}`             | Update manager details.              |
| Get manager      | GET    | `/api/v1/managers/{managerId}`             | Retrieve manager details.            |
| List managers    | GET    | `/api/v1/managers`                         | Get a paginated list of managers.    |
| Delete manager   | DELETE | `/api/v1/managers/{managerId}`             | Remove a manager account.            |
| Update manager role | PATCH  | `/api/v1/managers/{managerId}/role`      | Change manager role (e.g., admin, supervisor). |

### Authentication & Authorization

| Action        | Method  | Endpoint                     | Description                                  |
|--------------|--------|-----------------------------|----------------------------------------------|
| Login       | POST   | `/api/v1/auth/login`        | Authenticate user and get a token.         |
| Logout      | POST   | `/api/v1/auth/logout`       | Logout the user and invalidate the session. |
| Refresh Token | POST   | `/api/v1/auth/refresh`     | Refresh authentication token.              |


## Product Service

| Action                | Method  | Endpoint                               | Description                                |
|----------------------|--------|--------------------------------------|--------------------------------------------|
| Create product      | POST   | `/api/v1/products`                   | Add a new product to the catalog.         |
| Update product      | PATCH  | `/api/v1/products/{productId}`       | Update product details.                   |
| Get product        | GET    | `/api/v1/products/{productId}`       | Retrieve product details.                 |
| List products      | GET    | `/api/v1/products`                   | Get a paginated list of products.         |
| Delete product     | DELETE | `/api/v1/products/{productId}`       | Remove a product from the catalog.        |
| Search products    | GET    | `/api/v1/products/search?query={term}` | Search for products based on keywords.   |
| Filter products    | GET    | `/api/v1/products?category={category}&price_min={min}&price_max={max}` | Filter products by category, price, etc. |

---

## Inventory Service

| Action                | Method  | Endpoint                                    | Description                                     |
|----------------------|--------|-------------------------------------------|-------------------------------------------------|
| Add inventory item  | POST   | `/api/v1/inventory`                        | Add a new stock item to inventory.             |
| Update inventory    | PATCH  | `/api/v1/inventory/{inventoryId}`          | Update stock levels or details.                |
| Get inventory item  | GET    | `/api/v1/inventory/{inventoryId}`          | Retrieve inventory item details.               |
| List inventory      | GET    | `/api/v1/inventory`                        | Get a paginated list of inventory items.       |
| Delete inventory item | DELETE | `/api/v1/inventory/{inventoryId}`          | Remove an item from inventory.                 |
| Check stock        | GET    | `/api/v1/inventory/check?productId={productId}` | Check stock availability for a product.       |
| Reserve stock      | POST   | `/api/v1/inventory/reserve`                | Reserve stock for an order.                    |
| Release stock      | POST   | `/api/v1/inventory/release`                | Release reserved stock if an order is canceled. |

---

## Order Service

### Order

| Action                | Method  | Endpoint                              | Description                                     |
|----------------------|--------|-------------------------------------|-------------------------------------------------|
| Create order        | POST   | `/api/v1/orders`                    | Place a new customer order.                     |
| Update order        | PATCH  | `/api/v1/orders/{orderId}`          | Update order details (e.g., shipping address).  |
| Get order           | GET    | `/api/v1/orders/{orderId}`          | Retrieve details of a specific order.          |
| List orders        | GET    | `/api/v1/orders`                    | Get a paginated list of customer orders.       |
| Cancel order       | PATCH  | `/api/v1/orders/{orderId}/cancel`   | Cancel an order before shipment.               |
| Track order        | GET    | `/api/v1/orders/{orderId}/track`    | Get real-time tracking updates for an order.   |
| Process order      | POST   | `/api/v1/orders/{orderId}/process`  | Mark an order as processed.                    |
| Complete order     | POST   | `/api/v1/orders/{orderId}/complete` | Mark an order as completed and delivered.      |
| Return order       | POST   | `/api/v1/orders/{orderId}/return`   | Initiate a return request for an order.        |

---

### Cart

| Action                  | Method | Endpoint                              | Description                                               |
| ----------------------- | ------ | ------------------------------------- | --------------------------------------------------------- |
| Create Cart             | POST   | /api/v1/carts                        | Create a new empty cart for a customer.                   |
| Get Cart                | GET    | /api/v1/carts/{cartId}               | Retrieve a cart by its ID.                                |
| Add Item to Cart        | POST   | /api/v1/carts/{cartId}/items         | Add a product to the cart.                                |
| Update Item in Cart     | PATCH  | /api/v1/carts/{cartId}/items/{itemId}| Update the quantity of a product in the cart.             |
| Remove Item from Cart   | DELETE | /api/v1/carts/{cartId}/items/{itemId}| Remove a product from the cart.                            |
| Get Cart Total          | GET    | /api/v1/carts/{cartId}/total         | Retrieve the total amount of the cart.                     |
| Checkout Cart           | POST   | /api/v1/carts/{cartId}/checkout      | Convert cart to an order and initiate the checkout process.|


## Payment Service

| Action                  | Method  | Endpoint                                      | Description                                        |
|------------------------|--------|---------------------------------------------|----------------------------------------------------|
| Create payment intent | POST   | `/api/v1/payments/intents`                  | Create a new payment intent (prepare for payment). |
| Confirm payment       | POST   | `/api/v1/payments/{paymentId}/confirm`     | Confirm and capture the payment.                   |
| Cancel payment intent | POST   | `/api/v1/payments/{paymentId}/cancel`      | Cancel an existing payment intent.                 |
| Get payment status    | GET    | `/api/v1/payments/{paymentId}`             | Retrieve the status and details of a payment.     |
| List payments         | GET    | `/api/v1/payments`                         | Get a list of payments (supports pagination).     |
| Refund payment        | POST   | `/api/v1/payments/{paymentId}/refund`      | Refund a completed payment.                       |
| Get refund status     | GET    | `/api/v1/payments/{paymentId}/refunds/{refundId}` | Get status of a specific refund.                 |

---

## Pricing Service

| Action                | Method  | Endpoint                                         | Description                                  |
|----------------------|--------|------------------------------------------------|----------------------------------------------|
| Create price rule    | POST   | `/api/v1/pricing/rules`                        | Create a new pricing rule (e.g., discount, surcharge). |
| Update price rule    | PATCH  | `/api/v1/pricing/rules/{ruleId}`               | Update an existing pricing rule.             |
| Get price rule       | GET    | `/api/v1/pricing/rules/{ruleId}`               | Retrieve details of a specific pricing rule. |
| List price rules     | GET    | `/api/v1/pricing/rules`                        | List all pricing rules.                      |
| Delete price rule    | DELETE | `/api/v1/pricing/rules/{ruleId}`               | Delete a pricing rule.                       |
| Calculate price      | POST   | `/api/v1/pricing/calculate`                    | Calculate final price for a product or cart. |

---

## Recommendation Service

| Action                         | Method  | Endpoint                                  | Description                                     |
|-------------------------------|--------|-------------------------------------------|-------------------------------------------------|
| Get product recommendations  | GET    | `/api/v1/recommendations/products`       | Get recommended products for a customer.       |
| Get cart-based recommendations | GET    | `/api/v1/recommendations/cart`            | Get recommendations based on cart contents.   |
| Get order-based recommendations | GET    | `/api/v1/recommendations/orders/{orderId}` | Recommend products based on past order.       |
| Record user interaction       | POST   | `/api/v1/recommendations/interactions`   | Record user actions like views, clicks, etc.  |
| List recommendation models   | GET    | `/api/v1/recommendations/models`         | Get available recommendation models (if multiple exist). |
| Refresh recommendations      | POST   | `/api/v1/recommendations/refresh`        | Trigger manual refresh of recommendations.    |

---
