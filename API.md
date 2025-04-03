# E-Commerce System API Endpoints

## Table of Contents

- [User Service](#user-service)
- [Product Service](#product-service)
- [Inventory Service](#inventory-service)
- [Order Service](#order-service)
- [Cart Service](#cart-service)
- [Payment Service](#payment-service)
- [Review Service](#review-service)
- [Recommendation Service](#recommendation-service)
- [Pricing Service](#pricing-service)
- [Shipping Service](#shipping-service)
- [Notification Service](#notification-service)
- [Discount Service](#discount-service)
- [Wishlist Service](#wishlist-service)
- [Return & Refund Service](#return--refund-service)
- [Manager Service](#manager-service)
- [Analytics Service](#analytics-service)
- [Auth Service](#auth-service)

## User Service

### Customers

| Action              | Method | Endpoint                                     | Description                         |
| ------------------- | ------ | -------------------------------------------- | ----------------------------------- |
| Register customer   | POST   | `/api/v1/customers`                          | Create a new customer account       |
| Update customer     | PATCH  | `/api/v1/customers/{customerId}`             | Update customer details             |
| Get customer        | GET    | `/api/v1/customers/{customerId}`             | Retrieve customer details           |
| List customers      | GET    | `/api/v1/customers`                          | Get a paginated list of customers   |
| Delete customer     | DELETE | `/api/v1/customers/{customerId}`             | Soft delete a customer account      |
| Deactivate customer | PATCH  | `/api/v1/customers/{customerId}/deactivate`  | Deactivate a customer account       |
| Update address      | PUT    | `/api/v1/customers/{customerId}/addresses`   | Update customer address information |
| Get preferences     | GET    | `/api/v1/customers/{customerId}/preferences` | Get customer preferences            |
| Update preferences  | PATCH  | `/api/v1/customers/{customerId}/preferences` | Update customer preferences         |

## Product Service

### Products

| Action                  | Method | Endpoint                                            | Description                        |
| ----------------------- | ------ | --------------------------------------------------- | ---------------------------------- |
| Create product          | POST   | `/api/v1/products`                                  | Create a new product               |
| Update product          | PATCH  | `/api/v1/products/{productId}`                      | Update product details             |
| Get product             | GET    | `/api/v1/products/{productId}`                      | Retrieve product details           |
| List products           | GET    | `/api/v1/products`                                  | Get a paginated list of products   |
| Delete product          | DELETE | `/api/v1/products/{productId}`                      | Remove a product                   |
| List product variants   | GET    | `/api/v1/products/{productId}/variants`             | Get all variants of a product      |
| Create product variant  | POST   | `/api/v1/products/{productId}/variants`             | Create a new variant for a product |
| Update product variant  | PATCH  | `/api/v1/products/{productId}/variants/{variantId}` | Update a product variant           |
| Delete product variant  | DELETE | `/api/v1/products/{productId}/variants/{variantId}` | Remove a product variant           |
| List product categories | GET    | `/api/v1/categories`                                | Get all product categories         |
| Get product category    | GET    | `/api/v1/categories/{categoryId}`                   | Get details of a product category  |
| Create product category | POST   | `/api/v1/categories`                                | Create a new product category      |
| Update product category | PATCH  | `/api/v1/categories/{categoryId}`                   | Update a product category          |
| Delete product category | DELETE | `/api/v1/categories/{categoryId}`                   | Delete a product category          |

### Media

| Action             | Method | Endpoint                                     | Description                       |
| ------------------ | ------ | -------------------------------------------- | --------------------------------- |
| Upload media       | POST   | `/api/v1/products/{productId}/media`         | Upload media for a product        |
| Update media       | PATCH  | `/api/v1/media/{mediaId}`                    | Update media metadata             |
| Get media          | GET    | `/api/v1/media/{mediaId}`                    | Retrieve media item               |
| List product media | GET    | `/api/v1/products/{productId}/media`         | Get all media for a product       |
| Delete media       | DELETE | `/api/v1/media/{mediaId}`                    | Remove media                      |
| Reorder media      | PATCH  | `/api/v1/products/{productId}/media/reorder` | Reorder media items for a product |

## Inventory Service

| Action                      | Method | Endpoint                                      | Description                             |
| --------------------------- | ------ | --------------------------------------------- | --------------------------------------- |
| Create inventory record     | POST   | `/api/v1/inventory`                           | Create new inventory record             |
| Update inventory            | PATCH  | `/api/v1/inventory/{inventoryId}`             | Update inventory information            |
| Get inventory               | GET    | `/api/v1/inventory/{inventoryId}`             | Get inventory details                   |
| Get product inventory       | GET    | `/api/v1/products/{productId}/inventory`      | Get inventory for a specific product    |
| List inventory              | GET    | `/api/v1/inventory`                           | Get paginated list of inventory records |
| Add inventory stock         | POST   | `/api/v1/inventory/{inventoryId}/stock`       | Add stock to inventory                  |
| Remove inventory stock      | POST   | `/api/v1/inventory/{inventoryId}/deduct`      | Remove stock from inventory             |
| Get inventory history       | GET    | `/api/v1/inventory/{inventoryId}/history`     | Get inventory adjustment history        |
| Create inventory adjustment | POST   | `/api/v1/inventory/{inventoryId}/adjustments` | Create an inventory adjustment          |
| Set low stock threshold     | PATCH  | `/api/v1/inventory/{inventoryId}/threshold`   | Set low stock alert threshold           |
| Get low stock items         | GET    | `/api/v1/inventory/low-stock`                 | Get items below threshold               |

## Order Service

| Action               | Method | Endpoint                                  | Description                        |
| -------------------- | ------ | ----------------------------------------- | ---------------------------------- |
| Create order         | POST   | `/api/v1/orders`                          | Create a new order                 |
| Update order         | PATCH  | `/api/v1/orders/{orderId}`                | Update order details               |
| Get order            | GET    | `/api/v1/orders/{orderId}`                | Retrieve order details             |
| List orders          | GET    | `/api/v1/orders`                          | Get a paginated list of orders     |
| List customer orders | GET    | `/api/v1/customers/{customerId}/orders`   | Get orders for a specific customer |
| Cancel order         | POST   | `/api/v1/orders/{orderId}/cancel`         | Cancel an order                    |
| Update order status  | PATCH  | `/api/v1/orders/{orderId}/status`         | Update order status                |
| Add order item       | POST   | `/api/v1/orders/{orderId}/items`          | Add an item to an order            |
| Update order item    | PATCH  | `/api/v1/orders/{orderId}/items/{itemId}` | Update an order item               |
| Remove order item    | DELETE | `/api/v1/orders/{orderId}/items/{itemId}` | Remove an item from an order       |
| Get order history    | GET    | `/api/v1/orders/{orderId}/history`        | Get history of order changes       |

## Cart Service

| Action              | Method | Endpoint                                  | Description                             |
| ------------------- | ------ | ----------------------------------------- | --------------------------------------- |
| Create cart         | POST   | `/api/v1/carts`                           | Create a new cart                       |
| Get cart            | GET    | `/api/v1/carts/{cartId}`                  | Retrieve cart details                   |
| Get cart by session | GET    | `/api/v1/carts/session/{sessionId}`       | Retrieve cart by session ID             |
| Get customer cart   | GET    | `/api/v1/customers/{customerId}/cart`     | Get active cart for a customer          |
| Add cart item       | POST   | `/api/v1/carts/{cartId}/items`            | Add an item to a cart                   |
| Update cart item    | PATCH  | `/api/v1/carts/{cartId}/items/{itemId}`   | Update a cart item                      |
| Remove cart item    | DELETE | `/api/v1/carts/{cartId}/items/{itemId}`   | Remove an item from a cart              |
| Apply discount      | POST   | `/api/v1/carts/{cartId}/discounts`        | Apply discount code to cart             |
| Remove discount     | DELETE | `/api/v1/carts/{cartId}/discounts/{code}` | Remove discount code from cart          |
| Merge carts         | POST   | `/api/v1/carts/{cartId}/merge`            | Merge anonymous cart with customer cart |
| Clear cart          | DELETE | `/api/v1/carts/{cartId}/items`            | Remove all items from cart              |
| Convert to order    | POST   | `/api/v1/carts/{cartId}/checkout`         | Convert cart to order                   |

## Payment Service

| Action                 | Method | Endpoint                                  | Description                    |
| ---------------------- | ------ | ----------------------------------------- | ------------------------------ |
| Create payment intent  | POST   | `/api/v1/payments/intent`                 | Create a payment intent        |
| Get payment details    | GET    | `/api/v1/payments/{paymentId}`            | Get payment details            |
| Capture payment        | POST   | `/api/v1/payments/{paymentId}/capture`    | Capture an authorized payment  |
| Refund payment         | POST   | `/api/v1/payments/{paymentId}/refund`     | Issue a refund                 |
| Update payment status  | PATCH  | `/api/v1/payments/{paymentId}/status`     | Update payment status          |
| List payments          | GET    | `/api/v1/payments`                        | Get paginated list of payments |
| Get order payment      | GET    | `/api/v1/orders/{orderId}/payment`        | Get payment for an order       |
| List customer payments | GET    | `/api/v1/customers/{customerId}/payments` | Get payments for a customer    |
| Create dispute         | POST   | `/api/v1/payments/{paymentId}/disputes`   | Create a payment dispute       |
| Update dispute         | PATCH  | `/api/v1/payments/disputes/{disputeId}`   | Update dispute details         |
| Get dispute            | GET    | `/api/v1/payments/disputes/{disputeId}`   | Get dispute details            |
| List disputes          | GET    | `/api/v1/payments/disputes`               | Get all payment disputes       |

## Review Service

| Action                | Method | Endpoint                                    | Description                           |
| --------------------- | ------ | ------------------------------------------- | ------------------------------------- |
| Create review         | POST   | `/api/v1/reviews`                           | Create a new product review           |
| Update review         | PATCH  | `/api/v1/reviews/{reviewId}`                | Update review details                 |
| Get review            | GET    | `/api/v1/reviews/{reviewId}`                | Get review details                    |
| Delete review         | DELETE | `/api/v1/reviews/{reviewId}`                | Delete a review                       |
| List reviews          | GET    | `/api/v1/reviews`                           | Get paginated list of reviews         |
| List product reviews  | GET    | `/api/v1/products/{productId}/reviews`      | Get reviews for a product             |
| List customer reviews | GET    | `/api/v1/customers/{customerId}/reviews`    | Get reviews by a customer             |
| Vote review helpful   | POST   | `/api/v1/reviews/{reviewId}/vote/helpful`   | Vote a review as helpful              |
| Vote review unhelpful | POST   | `/api/v1/reviews/{reviewId}/vote/unhelpful` | Vote a review as unhelpful            |
| Respond to review     | POST   | `/api/v1/reviews/{reviewId}/response`       | Add a response to a review            |
| Update review status  | PATCH  | `/api/v1/reviews/{reviewId}/status`         | Update review status (publish/reject) |
| Get product rating    | GET    | `/api/v1/products/{productId}/rating`       | Get aggregated rating for a product   |

## Recommendation Service

| Action                       | Method | Endpoint                                         | Description                                     |
| ---------------------------- | ------ | ------------------------------------------------ | ----------------------------------------------- |
| Get product recommendations  | GET    | `/api/v1/products/{productId}/recommendations`   | Get recommendations for a product               |
| Get customer recommendations | GET    | `/api/v1/customers/{customerId}/recommendations` | Get personalized recommendations for a customer |
| Get cart recommendations     | GET    | `/api/v1/carts/{cartId}/recommendations`         | Get recommendations based on cart items         |
| Get related products         | GET    | `/api/v1/products/{productId}/related`           | Get related products                            |
| Get trending products        | GET    | `/api/v1/recommendations/trending`               | Get trending products                           |
| Get new arrivals             | GET    | `/api/v1/recommendations/new-arrivals`           | Get new product arrivals                        |
| Get popular in category      | GET    | `/api/v1/categories/{categoryId}/popular`        | Get popular products in a category              |

## Pricing Service

| Action                    | Method | Endpoint                                     | Description                         |
| ------------------------- | ------ | -------------------------------------------- | ----------------------------------- |
| Create pricing rule       | POST   | `/api/v1/pricing-rules`                      | Create a new pricing rule           |
| Update pricing rule       | PATCH  | `/api/v1/pricing-rules/{ruleId}`             | Update pricing rule details         |
| Get pricing rule          | GET    | `/api/v1/pricing-rules/{ruleId}`             | Get pricing rule details            |
| List pricing rules        | GET    | `/api/v1/pricing-rules`                      | Get paginated list of pricing rules |
| Delete pricing rule       | DELETE | `/api/v1/pricing-rules/{ruleId}`             | Delete a pricing rule               |
| Activate pricing rule     | PATCH  | `/api/v1/pricing-rules/{ruleId}/activate`    | Activate a pricing rule             |
| Deactivate pricing rule   | PATCH  | `/api/v1/pricing-rules/{ruleId}/deactivate`  | Deactivate a pricing rule           |
| Get product pricing rules | GET    | `/api/v1/products/{productId}/pricing-rules` | Get pricing rules for a product     |

## Shipping Service

| Action                  | Method | Endpoint                                             | Description                            |
| ----------------------- | ------ | ---------------------------------------------------- | -------------------------------------- |
| Create shipping record  | POST   | `/api/v1/shipping`                                   | Create a new shipping record           |
| Update shipping record  | PATCH  | `/api/v1/shipping/{shippingId}`                      | Update shipping information            |
| Get shipping details    | GET    | `/api/v1/shipping/{shippingId}`                      | Get shipping details                   |
| List shipping records   | GET    | `/api/v1/shipping`                                   | Get paginated list of shipping records |
| Get order shipping      | GET    | `/api/v1/orders/{orderId}/shipping`                  | Get shipping information for an order  |
| Update tracking         | PATCH  | `/api/v1/shipping/{shippingId}/tracking`             | Update tracking information            |
| Get tracking info       | GET    | `/api/v1/shipping/{shippingId}/tracking`             | Get tracking information               |
| Add shipping package    | POST   | `/api/v1/shipping/{shippingId}/packages`             | Add a package to shipping record       |
| Update shipping package | PATCH  | `/api/v1/shipping/{shippingId}/packages/{packageId}` | Update a shipping package              |
| Calculate shipping cost | POST   | `/api/v1/shipping/calculate`                         | Calculate shipping cost                |
| Generate shipping label | POST   | `/api/v1/shipping/{shippingId}/label`                | Generate shipping label                |

## Notification Service

| Action                      | Method | Endpoint                                                | Description                             |
| --------------------------- | ------ | ------------------------------------------------------- | --------------------------------------- |
| Create notification         | POST   | `/api/v1/notifications`                                 | Create a new notification               |
| Get notification            | GET    | `/api/v1/notifications/{notificationId}`                | Get notification details                |
| List notifications          | GET    | `/api/v1/notifications`                                 | Get paginated list of notifications     |
| List customer notifications | GET    | `/api/v1/customers/{customerId}/notifications`          | Get notifications for a customer        |
| List manager notifications  | GET    | `/api/v1/managers/{managerId}/notifications`            | Get notifications for a manager         |
| Mark as read                | PATCH  | `/api/v1/notifications/{notificationId}/read`           | Mark notification as read               |
| Mark all as read            | PATCH  | `/api/v1/customers/{customerId}/notifications/read-all` | Mark all customer notifications as read |
| Delete notification         | DELETE | `/api/v1/notifications/{notificationId}`                | Delete a notification                   |
| Send test notification      | POST   | `/api/v1/notifications/test`                            | Send a test notification                |

## Discount Service

| Action                   | Method | Endpoint                                     | Description                          |
| ------------------------ | ------ | -------------------------------------------- | ------------------------------------ |
| Create discount code     | POST   | `/api/v1/discount-codes`                     | Create a new discount code           |
| Update discount code     | PATCH  | `/api/v1/discount-codes/{codeId}`            | Update discount code details         |
| Get discount code        | GET    | `/api/v1/discount-codes/{codeId}`            | Get discount code details            |
| Get discount by code     | GET    | `/api/v1/discount-codes/code/{code}`         | Get discount by code string          |
| List discount codes      | GET    | `/api/v1/discount-codes`                     | Get paginated list of discount codes |
| Delete discount code     | DELETE | `/api/v1/discount-codes/{codeId}`            | Delete a discount code               |
| Activate discount code   | PATCH  | `/api/v1/discount-codes/{codeId}/activate`   | Activate a discount code             |
| Deactivate discount code | PATCH  | `/api/v1/discount-codes/{codeId}/deactivate` | Deactivate a discount code           |
| Validate discount code   | POST   | `/api/v1/discount-codes/validate`            | Validate a discount code             |
| Get usage history        | GET    | `/api/v1/discount-codes/{codeId}/usage`      | Get usage history of a discount code |

## Wishlist Service

| Action                  | Method | Endpoint                                           | Description                      |
| ----------------------- | ------ | -------------------------------------------------- | -------------------------------- |
| Create wishlist         | POST   | `/api/v1/wishlists`                                | Create a new wishlist            |
| Update wishlist         | PATCH  | `/api/v1/wishlists/{wishlistId}`                   | Update wishlist details          |
| Get wishlist            | GET    | `/api/v1/wishlists/{wishlistId}`                   | Get wishlist details             |
| List customer wishlists | GET    | `/api/v1/customers/{customerId}/wishlists`         | Get all wishlists for a customer |
| Delete wishlist         | DELETE | `/api/v1/wishlists/{wishlistId}`                   | Delete a wishlist                |
| Add wishlist item       | POST   | `/api/v1/wishlists/{wishlistId}/items`             | Add an item to a wishlist        |
| Update wishlist item    | PATCH  | `/api/v1/wishlists/{wishlistId}/items/{productId}` | Update a wishlist item           |
| Remove wishlist item    | DELETE | `/api/v1/wishlists/{wishlistId}/items/{productId}` | Remove an item from a wishlist   |
| Share wishlist          | POST   | `/api/v1/wishlists/{wishlistId}/share`             | Share a wishlist                 |
| Move to cart            | POST   | `/api/v1/wishlists/{wishlistId}/move-to-cart`      | Move wishlist items to cart      |

## Return & Refund Service

| Action                 | Method | Endpoint                                 | Description                               |
| ---------------------- | ------ | ---------------------------------------- | ----------------------------------------- |
| Create return request  | POST   | `/api/v1/returns`                        | Create a new return request               |
| Update return request  | PATCH  | `/api/v1/returns/{returnId}`             | Update return request details             |
| Get return details     | GET    | `/api/v1/returns/{returnId}`             | Get return request details                |
| List returns           | GET    | `/api/v1/returns`                        | Get paginated list of returns             |
| List customer returns  | GET    | `/api/v1/customers/{customerId}/returns` | Get returns for a customer                |
| List order returns     | GET    | `/api/v1/orders/{orderId}/returns`       | Get returns for an order                  |
| Update return status   | PATCH  | `/api/v1/returns/{returnId}/status`      | Update return status                      |
| Process refund         | POST   | `/api/v1/returns/{returnId}/refund`      | Process refund for a return               |
| Generate return label  | POST   | `/api/v1/returns/{returnId}/label`       | Generate return shipping label            |
| Add inspection results | POST   | `/api/v1/returns/{returnId}/inspection`  | Add inspection results for returned items |

## Manager Service

| Action              | Method | Endpoint                                  | Description                    |
| ------------------- | ------ | ----------------------------------------- | ------------------------------ |
| Create manager      | POST   | `/api/v1/managers`                        | Create a new manager account   |
| Update manager      | PATCH  | `/api/v1/managers/{managerId}`            | Update manager details         |
| Get manager         | GET    | `/api/v1/managers/{managerId}`            | Get manager details            |
| List managers       | GET    | `/api/v1/managers`                        | Get paginated list of managers |
| Delete manager      | DELETE | `/api/v1/managers/{managerId}`            | Delete a manager account       |
| Update manager role | PATCH  | `/api/v1/managers/{managerId}/roles`      | Update manager roles           |
| Get audit logs      | GET    | `/api/v1/audit-logs`                      | Get system audit logs          |
| Get manager logs    | GET    | `/api/v1/managers/{managerId}/audit-logs` | Get audit logs for a manager   |

## Analytics Service

| Action                     | Method | Endpoint                                 | Description                            |
| -------------------------- | ------ | ---------------------------------------- | -------------------------------------- |
| Get sales report           | GET    | `/api/v1/analytics/sales`                | Get sales analytics                    |
| Get customer analytics     | GET    | `/api/v1/analytics/customers`            | Get customer analytics                 |
| Get product performance    | GET    | `/api/v1/analytics/products`             | Get product performance analytics      |
| Get specific product stats | GET    | `/api/v1/analytics/products/{productId}` | Get analytics for specific product     |
| Get inventory analytics    | GET    | `/api/v1/analytics/inventory`            | Get inventory analytics                |
| Get cart analytics         | GET    | `/api/v1/analytics/carts`                | Get cart analytics (abandonment, etc.) |
| Get review analytics       | GET    | `/api/v1/analytics/reviews`              | Get review analytics                   |
| Get dashboard stats        | GET    | `/api/v1/analytics/dashboard`            | Get main dashboard statistics          |
| Export reports             | POST   | `/api/v1/analytics/export`               | Export analytics reports               |

## Auth Service

| Action                 | Method | Endpoint                              | Description                           |
| ---------------------- | ------ | ------------------------------------- | ------------------------------------- |
| Login                  | POST   | `/api/v1/auth/login`                  | User login                            |
| Logout                 | POST   | `/api/v1/auth/logout`                 | User logout                           |
| Refresh token          | POST   | `/api/v1/auth/refresh`                | Refresh authentication token          |
| Request password reset | POST   | `/api/v1/auth/password-reset/request` | Request password reset                |
| Reset password         | POST   | `/api/v1/auth/password-reset/confirm` | Confirm password reset                |
| Change password        | POST   | `/api/v1/auth/password/change`        | Change user password                  |
| Enable 2FA             | POST   | `/api/v1/auth/2fa/enable`             | Enable two-factor authentication      |
| Disable 2FA            | POST   | `/api/v1/auth/2fa/disable`            | Disable two-factor authentication     |
| Verify 2FA             | POST   | `/api/v1/auth/2fa/verify`             | Verify two-factor authentication code |
| Get login history      | GET    | `/api/v1/auth/login-history`          | Get user login history                |
