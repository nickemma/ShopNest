# E-Commerce Models Documentation

## Table of Contents

1. [User Model](#1-user-model)
2. [Manager Model](#2-manager-model)
3. [Product Model](#3-product-model)
4. [Inventory Model](#4-inventory-model)
5. [Order Model](#5-order-model)
6. [Cart Model](#6-cart-model)
7. [Payment Model](#7-payment-model)
8. [Review Model](#8-review-model)
9. [Recommendation Model](#9-recommendation-model)
10. [Pricing Rule Model](#10-pricing-rule-model)
11. [Shipping Model](#11-shipping-model)
12. [Notification Model](#12-notification-model)
13. [Discount Code Model](#13-discount-code-model)
14. [Wishlist Model](#14-wishlist-model)

---

### 1. User Model

```json
{
  "userId": "string (uuid)",
  "name": "string",
  "email": "string (email)",
  "phone": "string (E.164)",
  "address": {
    "street": "string",
    "city": "string",
    "state": "string (ISO 3166-2)",
    "postalCode": "string",
    "country": "string (ISO 3166-1 alpha-2)"
  },
  "status": "active | inactive | suspended",
  "preferences": {
    "currency": "USD | EUR | GBP | JPY | CAD | AUD | CNY",
    "language": "en-US | es-ES | fr-FR | de-DE | ja-JP | zh-CN"
  },
  // NOTE:START->
  // This is a generic service, for now development can proceed
  // but needs to be factored out
  "authentication": {
    "passwordHash": "string",
    "twoFactorEnabled": "boolean",
    "twoFactorMethod": "email | sms | authenticatorApp"
  },
  // <-NOTE:END
  // Do not track orders in customer schema
  // "orderStats": {
  //   "totalOrders": "number",
  //   "totalSpent": "number",
  //   "averageOrderValue": "number"
  // },
  "createdAt": "string (ISO 8601)",
  "updatedAt": "string (ISO 8601)"
}
```

### 2. Manager Model

```json
{
  "managerId": "string (uuid)",
  "name": "string",
  "email": "string (email)",
  "roles": ["admin | supervisor | support"],
  // NOTESTART->
  // This is a generic service, for now development can proceed
  // but needs to be factored out
  "lastLogin": "string (ISO 8601)",
  "loginHistory": [
    {
      "timestamp": "string (ISO 8601)",
      "ip": "string (IPv4/IPv6)",
      "device": "string",
      "successful": "boolean"
    }
  ],
  "authentication": {
    "passwordHash": "string",
    "twoFactorEnabled": "boolean",
    "twoFactorMethod": "email | sms | authenticatorApp"
  },
  // <-NOTE:END
  "createdAt": "string (ISO 8601)",
  "updatedAt": "string (ISO 8601)",
  "createdBy": "string (manager UUID)"
}
```

### 3. Product Model

```json
{
  "productId": "string (uuid)",
  "sku": "string (unique)",
  "name": "string",
  "description": "string",
  "price": {
    "amount": "number (positive)",
    "currency": "USD | EUR | GBP"
  },
  "category": [
    "electronics | clothing | home | beauty | sports | food | books | toys | Others"
  ],
  "attributes": {
    "weightKg": "number",
    "dimensionsCm": {
      "length": "number",
      "width": "number",
      "height": "number"
    },
    "color": "string",
    "brand": "string"
  },
  "mediaIds": ["string (uuid)"],
  "status": "available | out_of_stock | discontinued | coming_soon | draft",
  // NOTE:START->
  // Since we have a recommendation service, we can create the rating system in it
  // or create a separate rating service. The rating would reference product by ID
  // Why ? This would allow aggreagation of rating across customers and alow effective
  // recommendation via user-product interactions
  // "rating": {
  //   "average": "number (1-5)",
  //   "count": "number"
  // },
  // <-NOTE:END
  // NOTE:START->
  // Having this here is redundant and inefficient. Which is the very reason we have
  // inventory service. Coordinating mutations across inventory and product service
  // for this is just an headache, let the source of truth be in the inventory service.
  // That is where it belongs.
  // <- NOTE:END
  // "inventory": {
  //   "stock": "number",
  //   "lowStockThreshold": "number",
  //   "reservedStock": "number",
  //   "backorderAllowed": "boolean",
  //   "expectedRestockDate": "string (ISO 8601 date, optional)"
  // },
  "tags": ["string"],
  "relatedProducts": ["string (product UUIDs)"],
  "variants": [
    {
      "variantId": "string (uuid)",
      "sku": "string (unique)",
      "attributes": {
        "color": "string",
        "size": "string"
      },
      "price": {
        "amount": "number (positive)",
        "currency": "USD | EUR | GBP"
      },
      // Read comment on inventory objet above. Variant is simply another product (similar)
      // "stock": {
      //   "available": "number (>=0)",
      //   "reserved": "number (>=0)",
      //   "imageUrl": ["string (URL)"]
      // }
    }
  ],
  "shipping": {
    "weight": "number",
    "dimensions": {
      "length": "number",
      "width": "number",
      "height": "number"
    },
    "freeShipping": "boolean"
  },
  "tax": {
    "taxable": "boolean",
    "taxClass": "standard | reduced | zero"
  },
  "createdAt": "string (ISO 8601)",
  "updatedAt": "string (ISO 8601)",
}
```

### Media Model

```json
{
  "mediaId": "string (uuid)",
  "productId": "string (uuid)",
  // we can have more than a single image url
  // changer from string to array of strings
  "imageUrl": ["string (url)"],
  "metadata": {
    "width": "number (pixels, optional)",
    "height": "number (pixels, optional)",
    "format": "string (jpg/png etc)",
    "size": "number (bytes)"
  },
  "position": "number (sort order)",
  "altText": "string (accessibility)",
  "status": "active | archived",
  "createdAt": "string (ISO 8601)"
}
```

### Inventory Model

```json
{
  "inventoryId": "string (uuid)",
  "productId": "string (reference)",
  // The current system is not managing warehouses, hence removal of id field.
  // Assunption is we have a single warehouse, but incase we have more than one,
  // specify the location only. 
  "warehouse": {
    // "id": "string",
    "location": "string"
  },
  "stock": {
    "available": "number (>=0)",
    "reserved": "number (>=0)",
    "safetyStock": "number (>=0)",
    "damagedStock": "number (>=0)"
  },
  "history": [
    {
      "date": "string (ISO 8601 datetime)",
      "adjustment": "number",
      "reason": "restock | sale | return | damage | inventory_check | transfer",
      "referenceId": "string (order/transfer ID)",
      "performedBy": "string (user/manager ID)"
    }
  ],
  "alertSettings": {
    "lowStockThreshold": "number",
    "alertSent": "boolean",
    "alertDate": "string (ISO 8601 datetime, optional)"
  },
  "lastInventoryCheck": "string (ISO 8601 datetime)",
  "createdAt": "string (ISO 8601 datetime)"
}
```

### Order model

```json
{
  "orderId": "string (uuid)",
  "ordernumber": "string (human-readable)",
  // NOTE:START->
  // Aside the id field, are any other immutable ? any mutable one should be removed.
  // Why ? Displaying the order with the customer details contained here when the details
  // has been updated in the user service results in stale data.
  // NB: Proceeding with this is easy and simplifies development workflow but data inconsistency
  // is guaranteed if we allow customer to change any of email, name or phone.
  // Professional Advice:
  // "customerId": "string (reference)",
  // <-NOTE:END
  "customer": {
    "id": "string (reference)",
    "email": "string (email)",
    "name": "string",
    "phone": "string (optional)"
  },
  "items": [
    {
      "productId": "string (UUID)",
      // NOTE:START->
      // variant is just a product. When creating the products,
      // other products can be selected (linked) to the currently created
      // product due to some features and similarities. They come into play when
      // displaying products to customers and more importantly, in recommendation service
      // in product-proudct recommendations
      // <-NOTE:END
      // "variantId": "string (UUID, optional)",
      "name": "string",
      "sku": "string",
      "quantity": "number",
      "unitPrice": "number",
      "lineTotal": "number",
      "discount": {
        "type": "percentage | fixed | none",
        "value": "number",
        "code": "string (optional)"
      },
      "tax": {
        "rate": "number",
        "amount": "number"
      }
    }
  ],
  "totals": {
    "subtotal": "number",
    "shipping": "number",
    "tax": "number",
    "discount": "number",
    "grandTotal": "number"
  },
    // NOTE:START->
    // Question: Why reference the transaction ID ? Note that it is must that the payment
    // service references the orderId. Which is why I included the payment intent. 
    // You can reference the payment intent but
    // Professional Advice:
    // Let payment intent be source of truth
    // <-NOTE:END
    // "payment": {
    //   "method": "credit_card | paypal | bank_transfer | crypto",
    //   "transactionId": "string",
    //   "status": ["pending", "captured", "refunded"]
    // },
  "fulfillment": {
    "priority": "standard | express | same_day",
    "carbonOffset": "boolean",
    "tracking": {
      "carrier": "string",
      "trackingNumber": "string",
      "status": "in_transit | out_for_delivery | delivered | delayed | returned",
      "estimatedDelivery": "string (ISO 8601 date)"
    }
  },
  "addresses": {
    "billing": {
      "street": "string",
      "city": "string",
      "state": "string",
      "postalCode": "string",
      "country": "string (ISO 3166-1 alpha-2)"
    },
    "shipping": {
      "street": "string",
      "city": "string",
      "state": "string",
      "postalCode": "string",
      "country": "string (ISO 3166-1 alpha-2)"
    }
  },
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)",
  "source": "web | mobile_app | marketplace | in_store | phone",
  "currency": "USD | EUR | GBP | JPY | CAD | AUD | CNY",
  "giftMessage": "string (optional)"
}
```

### Cart Model

```json
{
  "cartId": "string (UUID)",
  "sessionId": "string",
  "customerId": "string (UUID, optional)",
  "items": [
    {
      "productId": "string (UUID)",
      "variantId": "string (UUID, optional)",
      "quantity": "number",
      "priceSnapshot": "number",
      "addedAt": "string (ISO 8601 datetime)"
    }
  ],
  "totalItems": "number",
  "totalAmount": "number",
  "totalWeight": "number (optional)",
  "appliedDiscounts": [
    {
      "code": "string",
      "type": "percentage | fixed",
      "value": "number",
      "appliedTo": "cart | product",
      "productId": "string (UUID, optional)"
    }
  ],
  "expiresAt": "string (ISO 8601 datetime)",
  "metadata": {
    "userAgent": "string",
    "ipAddress": "string (IP address format)",
    "device": "mobile | desktop | tablet",
    "referrer": "string (URL, optional)"
  },
  "abandonedCartReminder": {
    "sent": "boolean",
    "sentAt": "string (ISO 8601 datetime, optional)"
  },
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)",
  "status": "active | abandoned | completed"
}
```

### Payment Intent Model

```json
{
  "paymentId": "string (UUID)",
  "orderId": "string (UUID)",
  "customerId": "string (UUID)",
  "amount": {
    "value": "number",
    "currency": "USD | EUR | GBP | JPY | CAD | AUD | CNY"
  },
  "status": "pending | authorized | succeeded | failed | refunded | partially_refunded | voided",
  "method": {
    "type": "credit_card | paypal | bank_transfer | apple_pay | google_pay | crypto | gift_card",
    "last4": "string (for cards)",
    "brand": "string (for cards)",
    "expiryDate": "string (MM/YY, never returned in public responses)",
    "cardholderName": "string (never returned in public responses)"
  },
  "riskAnalysis": {
    "score": "number (0-100)",
    "flags": [
      "high_value | unusual_location | multiple_attempts | address_mismatch"
    ]
  },
  "dispute": {
    "status": "none | open | under_review | resolved_merchant_favor | resolved_customer_favor",
    "reason": "fraudulent | unrecognized | duplicate | subscription_canceled | product_not_received | product_unacceptable | credit_not_processed | other",
    "amount": "number",
    "openedAt": "string (ISO 8601 datetime, optional)",
    "history": [
      {
        "timestamp": "string (ISO 8601 datetime)",
        "action": "opened | evidence_required | evidence_provided | decided | closed",
        "details": "string",
        "performedBy": "string (user/manager ID, optional)"
      }
    ]
  },
  // refund can only be one, not an array. It is applied to a single payment (the current payment)
  // "refunds": [
  //   {
  //     "refundId": "string (UUID)",
  //     "amount": "number",
  //     "reason": "customer_request | damaged | incorrect_item | fraud | other",
  //     "status": "pending | succeeded | failed",
  //     "createdAt": "string (ISO 8601 datetime)"
  //   }
  // ],
  "refundId": "string (reference)", // reference the refund model 
  "transactionFee": "number",
  "gateway": "stripe | paypal | adyen | braintree | square",
  "gatewayResponse": {
    "responseCode": "string",
    "message": "string",
    "transactionReference": "string"
  },
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)",
  "capturedAt": "string (ISO 8601 datetime, optional)"
}
```

### Review Model

```json
{
  "reviewId": "string (UUID)",
  "productId": "string (UUID)",
  // Question: Why are we reviewing an order ? Review is typically done by 
  // a CUSTOMER on a PRODUCT.
  // "orderId": "string (UUID, optional)",
  "customerId": "string (UUID)",
  // field is immutable, but specifying it here means choosing between ease of display
  // or data consistency.
  // Professional Advice: Do not specify.
  // Ease of development: Proceed with this and neccessary refactoring can be done later. 
  // "customerName": "string (optional, display name)", 
  "title": "string",
  "rating": "number (1-5)",
  "comment": "string",
  "recommendProduct": "boolean",
  "verifiedPurchase": "boolean",
  "helpfulVotes": "number",
  "unhelpfulVotes": "number",
  "status": "pending | published | rejected | flagged",
  "response": {
    "text": "string (optional)",
    "respondedBy": "string (manager ID, optional)",
    "respondedAt": "string (ISO 8601 datetime, optional)"
  },
  "media": [
    {
      "type": "image | video",
      "url": "string (URL)",
      "thumbnailUrl": "string (URL, optional)"
    }
  ],
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)",
  "publishedAt": "string (ISO 8601 datetime, optional)"
}
```

### Rating Model

```json
// This is the aggregate result ratings from the review model
{
    "productId": "string (UUID)",
    "averageRating": "number (float)",
    "ratingCount": "number",
    // include other fields as deemed necessary
}
```

### Pricing Rule:

```json
{
  "ruleId": "string (UUID)",
  "name": "string",
  "description": "string",
  "priority": "number",
  "conditions": {
    "products": ["string (UUID)"],
    "categories": ["string (UUID)"],
    "brands": ["string"],
    "customerGroups": ["string (UUID)"],
    "minQuantity": "number",
    "minOrderValue": "number",
    "validDates": {
      "start": "string (ISO 8601 datetime)",
      "end": "string (ISO 8601 datetime)"
    },
    "daysOfWeek": [
      "monday | tuesday | wednesday | thursday | friday | saturday | sunday"
    ],
    "timeOfDay": {
      "start": "string (HH:MM)",
      "end": "string (HH:MM)"
    }
  },
  "action": {
    "type": "percentage | fixed | buy_x_get_y | free_shipping",
    "value": "number",
    "maxDiscount": "number (optional)",
    "freeItemSku": "string (optional, for buy_x_get_y)",
    "freeItemQuantity": "number (optional, for buy_x_get_y)"
  },
  "combinable": "boolean",
  "couponRequired": "boolean",
  "couponCode": "string (optional)",
  "usageLimit": {
    "total": "number (optional)",
    "perCustomer": "number (optional)"
  },
  "status": "active | inactive | scheduled | expired",
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)",
  "createdBy": "string (manager ID)"
}
```

### Recommendation

```json
{
  "recommendationId": "string (UUID)",
  "engineVersion": "string",
  "context": {
    "source": "product_page | cart | order_history | homepage | category_page | search_results | email",
    "viewedProduct": "string (UUID, optional)",
    "customerId": "string (UUID, optional)",
    "sessionId": "string (optional)"
  },
  "products": [
    {
      "productId": "string (UUID)",
      "score": "number (0-1)",
      "reason": "frequently_bought_together | similar_items | trending | recently_viewed | new_arrival | popular_in_category | based_on_history | on_sale"
    }
  ],
  "createdAt": "string (ISO 8601 datetime)",
  "expiresAt": "string (ISO 8601 datetime)",
  "metadata": {
    "algorithm": "collaborative | content-based | hybrid",
    "conversionRate": "number (optional)",
    "impressions": "number (optional)",
    "clicks": "number (optional)"
  }
}
```

### Wishlist

```json
{
  "wishlistId": "string (UUID)",
  "customerId": "string (UUID)",
  "name": "string (default: 'Default')",
  "visibility": "private | public | shared",
  "items": [
    {
      "productId": "string (UUID)",
      // "variantId": "string (UUID, optional)", same as product
      "addedAt": "string (ISO 8601 datetime)",
      "notes": "string (optional)",
      "priority": "high | medium | low"
    }
  ],
  "sharedWith": ["string (email addresses, optional)"],
  "shareableLink": "string (URL, optional)",
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)"
}
```

### shipping model

```json
{
  "shippingId": "string (UUID)",
  "orderId": "string (UUID)",
  "carrier": "UPS | FedEx | DHL | USPS | Canada_Post | Australia_Post | Other",
  "service": "standard | express | overnight | two_day | ground | international",
  "trackingNumber": "string",
  "status": "pending | processing | info_received | in_transit | out_for_delivery | delivered | failed | returned | exception",
  "estimatedDelivery": {
    "start": "string (ISO 8601 date)",
    "end": "string (ISO 8601 date)"
  },
  "shippingAddress": {
    "name": "string",
    "street": "string",
    "city": "string",
    "state": "string",
    "postalCode": "string",
    "country": "string (ISO 3166-1 alpha-2)",
    "phone": "string (optional)"
  },
  "packages": [
    {
      "packageId": "string",
      "weight": "number",
      "dimensions": {
        "length": "number",
        "width": "number",
        "height": "number"
      },
      "items": [
        {
          "productId": "string (UUID)",
          "quantity": "number"
        }
      ]
    }
  ],
  "shippingLabel": "string (URL to label PDF, optional)",
  "trackingUrl": "string (URL, optional)",
  "cost": "number",
  "signature": {
    "required": "boolean",
    "name": "string (optional)",
    "timestamp": "string (ISO 8601 datetime, optional)"
  },
  "timeline": [
    {
      "status": "string",
      "location": "string",
      "timestamp": "string (ISO 8601 datetime)",
      "description": "string"
    }
  ],
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)"
}
```

### Notification Model

```json
{
  "notificationId": "string (UUID)",
  "customerId": "string (UUID, optional)",
  "managerId": "string (UUID, optional)",
  "type": "email | sms | push | in_app | webhook",
  "event": "order_placed | order_shipped | order_delivered | password_reset | back_in_stock | price_drop | account_update | review_request | abandoned_cart",
  "content": {
    "subject": "string",
    "body": "string",
    "templateId": "string (optional)",
    "templateData": "object (optional)"
  },
  "status": "sent | failed | pending | delivered | opened | clicked",
  "recipient": {
    "email": "string (optional)",
    "phone": "string (optional)",
    "deviceId": "string (optional)"
  },
  "metadata": {
    "ip": "string (optional)",
    "device": "string (optional)",
    "location": "string (optional)",
    "clickedLinks": ["string (URLs, optional)"]
  },
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)",
  "sentAt": "string (ISO 8601 datetime, optional)",
  "deliveredAt": "string (ISO 8601 datetime, optional)",
  "openedAt": "string (ISO 8601 datetime, optional)"
}
```

### Discount Code Model

```json
{
  "codeId": "string (UUID)",
  "code": "string (unique, case-insensitive)",
  "type": "percentage | fixed | free_shipping | buy_x_get_y",
  "value": "number",
  "minPurchase": "number",
  "maxDiscount": "number (optional)",
  "applicableTo": {
    "all": "boolean",
    "products": ["string (UUID)"],
    "categories": ["string (UUID)"],
    "collections": ["string (UUID)"]
  },
  "validDates": {
    "start": "string (ISO 8601 datetime)",
    "end": "string (ISO 8601 datetime)"
  },
  "customerEligibility": {
    "all": "boolean",
    "specificCustomers": ["string (UUID)"],
    "customerGroups": ["string"],
    "firstTimeOnly": "boolean"
  },
  "usageLimit": {
    "total": "number",
    "perCustomer": "number",
    "currentUsage": "number"
  },
  "status": "active | expired | used | inactive | scheduled",
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)",
  "createdBy": "string (manager ID)",
  "campaign": "string (optional)",
  "combinable": "boolean"
}
```

### Audit Log Model (Added)

```json
{
  "logId": "string (UUID)",
  "timestamp": "string (ISO 8601 datetime)",
  "performedBy": {
    "id": "string (UUID)",
    "type": "manager | customer | system",
    "name": "string (optional)"
  },
  "action": "create | update | delete | view | export | import | login | logout | checkout | refund | ship | cancel",
  "entityType": "product | order | customer | inventory | payment | discount | manager | settings",
  "entityId": "string (UUID)",
  "changes": [
    {
      "field": "string",
      "oldValue": "any",
      "newValue": "any"
    }
  ],
  "ipAddress": "string",
  "userAgent": "string",
  "notes": "string (optional)"
}
```

### Return/Refund Model (Added)

```json
{
  "returnId": "string (UUID)",
  "orderId": "string (UUID)",
  "customerId": "string (UUID)",
  "returnType": "return | exchange | repair | warranty",
  "status": "requested | approved | received | inspected | completed | rejected",
  "requestDate": "string (ISO 8601 datetime)",
  "items": [
    {
      // "orderItemId": "string", same as the productId
      "productId": "string (UUID)",
      "quantity": "number",
      "reason": "damaged | defective | wrong_item | not_as_described | unwanted | other",
      "condition": "new | used | damaged",
      "resellable": "boolean"
    }
  ],
  "refund": {
    "amount": "number",
    "method": "original_payment | store_credit | gift_card",
    "status": "pending | processed | failed",
    "transactionId": "string (optional)",
    "processedAt": "string (ISO 8601 datetime, optional)"
  },
  "exchange": {
    "items": [
      {
        "productId": "string (UUID)",
        "variantId": "string (UUID, optional)",
        "quantity": "number"
      }
    ],
    "priceDifference": "number",
    "priceDifferenceAction": "charge | refund | none"
  },
  "returnShipping": {
    "returnLabel": "string (URL)",
    "carrier": "string",
    "trackingNumber": "string",
    "cost": "number",
    "customerPaid": "boolean"
  },
  "inspectionResults": {
    "date": "string (ISO 8601 datetime, optional)",
    "condition": "as_described | damaged | used | missing_parts",
    "notes": "string",
    "performedBy": "string (manager ID)"
  },
  "notes": {
    "customer": "string (optional)",
    "internal": "string (optional, never returned to customer)"
  },
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)",
  "completedAt": "string (ISO 8601 datetime, optional)"
}
```

### Tax Configuration Model (Added)

```json
{
  "taxConfigId": "string (UUID)",
  "name": "string",
  "country": "string (ISO 3166-1 alpha-2)",
  "region": "string (state/province, optional)",
  "postalCodeRange": {
    "start": "string (optional)",
    "end": "string (optional)"
  },
  "taxRates": [
    {
      "category": "standard | reduced | zero | exempt",
      "name": "string",
      "rate": "number (percentage)",
      "priority": "number",
      "compound": "boolean",
      "appliesTo": {
        "products": ["string (UUID)"],
        "productCategories": ["string"]
      }
    }
  ],
  "taxExemptions": [
    {
      "customerId": "string (UUID)",
      "certificateNumber": "string",
      "certificateDocument": "string (URL)",
      "expiryDate": "string (ISO 8601 date)",
      "notes": "string"
    }
  ],
  "provider": "manual | avalara | taxjar | vertex",
  "providerSettings": {
    "apiKey": "string (never returned in responses)",
    "accountNumber": "string (never returned in responses)"
  },
  "calculationMode": "line_item | order_total",
  "shippingTaxable": "boolean",
  "active": "boolean",
  "createdAt": "string (ISO 8601 datetime)",
  "updatedAt": "string (ISO 8601 datetime)"
}
```
