## 1. Register customer with Auth
<!-- http://localhost:8000/api/v1/users/auth/register -->
```json
{
    "email": "test@gmail.com",
    "password": "password",
    "role": "CUSTOMER"
}

// 
{
	"authId": "6201d03e-0a76-4fe1-9452-60d8d8c0b373"
}
```

## 2. Verify Email (manually altered fied in db due to smtp constraints)

## 3. Login
<!-- http://localhost:8000/api/v1/users/auth/login -->
```json
{
    "email": "test@gmail.com",
    "password": "password"
}

// 
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNzQ3MzAxMjU1LCJpYXQiOjE3NDcyMTQ4NTUsInN1YiI6IjYyMDFkMDNlLTBhNzYtNGZlMS05NDUyLTYwZDhkOGMwYjM3MyIsInVzZXJJZCI6IiIsInVzZXJUeXBlIjoiQ1VTVE9NRVIifQ.wPxHSPrtdSe3gr-60ON6no76kIV-Rdb-AE4uhiAYwjs"
}

{
  "email": "test@gmail.com",
  "exp": 1747301255,
  "iat": 1747214855,
  "sub": "6201d03e-0a76-4fe1-9452-60d8d8c0b373", //authId
  "userId": "", // user is yet to add their details
  "userType": "CUSTOMER"
}
```

## 4. Get auth account (supply auth token from login)
<!-- http://localhost:8000/api/v1/users/auth/account -->
```json
{
	"data": {
		"authId": "6201d03e-0a76-4fe1-9452-60d8d8c0b373",
		"userId": "",
		"userType": "CUSTOMER",
		"email": "test@gmail.com",
		"verfied": true,
		"sessionData": {
			"lastLogin": "0001-01-01T00:00:00Z",
			"currentToken": "",
			"tokenExpiry": "0001-01-01T00:00:00Z",
			"loginAttempts": 0
		},
		"createdAt": "2025-05-14T09:08:13.078576Z",
		"updatedAt": "2025-05-14T09:08:13.078576Z"
	}
}
```

# CUSTOMERS

## 5. Register customer (supply token)
<!-- http://localhost:8000/api/v1/users/customers/register -->
```json
{
    "name": "first user"
}

//
{
	"customerId": "a90e2348-f8a6-4783-9771-47f08497f41a"
}
```

## 6. Refresh token (supply current token)
<!-- will deal later on with access token vs refresh token -->
<!-- http://localhost:8000/api/v1/users/auth/refresh -->
```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNzQ3MzA3MTkwLCJpYXQiOjE3NDcyMjA3OTAsInN1YiI6IjYyMDFkMDNlLTBhNzYtNGZlMS05NDUyLTYwZDhkOGMwYjM3MyIsInVzZXJJZCI6ImE5MGUyMzQ4LWY4YTYtNDc4My05NzcxLTQ3ZjA4NDk3ZjQxYSIsInVzZXJUeXBlIjoiQ1VTVE9NRVIiLCJ2ZXJpZmllZCI6dHJ1ZX0.P6II54ES76e8baOS_szfrbKhzi97GMUVobE3uuMmWuQ"
}
```

## 7. Get customer profile (supply the new current token)
```json
{
	"data": {
		"customerId": "a90e2348-f8a6-4783-9771-47f08497f41a",
		"name": "first user",
		"email": "test@gmail.com",
		"phone": "",
		"address": {
			"street": "",
			"city": "",
			"state": "",
			"postalCode": "",
			"country": ""
		},
		"status": "active",
		"preferences": {
			"currency": "",
			"language": ""
		},
		"createdAt": "2025-05-14T10:35:43.449792Z",
		"updatedAt": "2025-05-14T10:35:43.449792Z"
	}
}
```