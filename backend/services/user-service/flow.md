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



## MANAGER FLOW
<!-- http://localhost:8000/api/v1/users/auth/register -->
```json
{
    "email": "test2@gmail.com",
    "password": "password",
    "role": "MANAGER"
}
// 

{
	"authId": "821b3f18-1d4f-45d9-bedc-6b157dfbc2c3"
}
```


## Login
<!-- http://localhost:8000/api/v1/users/auth/login -->
```json
{
        "email": "test2@gmail.com",
    "password": "password"
}
// 
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQGdtYWlsLmNvbSIsImV4cCI6MTc0NzY4MTU5OCwiaWF0IjoxNzQ3NTk1MTk4LCJzdWIiOiI4MjFiM2YxOC0xZDRmLTQ1ZDktYmVkYy02YjE1N2RmYmMyYzMiLCJ1c2VySWQiOiIiLCJ1c2VyVHlwZSI6Ik1BTkFHRVIiLCJ2ZXJpZmllZCI6ZmFsc2V9.Hcr3MkK-fX6nHSgzAtPDt0_ZZmJauGLBwpsdV9Rihbg"
}
// --
{
  "email": "test2@gmail.com",
  "exp": 1747681598,
  "iat": 1747595198,
  "sub": "821b3f18-1d4f-45d9-bedc-6b157dfbc2c3",
  "userId": "",
  "userType": "MANAGER",
  "verified": false
}
```

## Manually verified (dev workaround around smtp)

## Manager registration (provide token)
<!-- http://localhost:8000/api/v1/users/managers/register -->
```json
{
    "name": "second user"
}
// --
{
	"userId": "26b39c93-120c-4e14-b087-740a44e1ce3c"
}
```

## Refresh token (supply current token)
<!-- http://localhost:8000/api/v1/users/auth/refresh -->
```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QyQGdtYWlsLmNvbSIsImV4cCI6MTc0NzY4MjgxMiwiaWF0IjoxNzQ3NTk2NDEyLCJzdWIiOiI4MjFiM2YxOC0xZDRmLTQ1ZDktYmVkYy02YjE1N2RmYmMyYzMiLCJ1c2VySWQiOiIyNmIzOWM5My0xMjBjLTRlMTQtYjA4Ny03NDBhNDRlMWNlM2MiLCJ1c2VyVHlwZSI6Ik1BTkFHRVIiLCJ2ZXJpZmllZCI6dHJ1ZX0._lRAvwBXNfUvxaylhesgWR6RGOL8QgriZ4ehJse4yRs"
}
```

## get manager profile (supply new token)

```json
{
	"data": {
		"managerId": "26b39c93-120c-4e14-b087-740a44e1ce3c",
		"name": "second user",
		"email": "test2@gmail.com",
		"role": "ADMIN",
		"approved": false,
		"createdAt": "2025-05-18T19:18:04.980882Z",
		"updatedAt": "2025-05-18T19:18:04.980882Z"
	}
}
```

## approve manager registration
since we are yet to implement the full permission,the workarround was to 
encode the superadmin into the code.
Later refactor would require us bootstrapping a superadmin with the defined
permissinon and perhaps a custom role not exposed to the frontend.
```json
{
    "email": "superadmin@gmail.com",
    "password": "password",
    "role": "MANAGER"
}
// --
{
	"authId": "0f399e1a-ef15-4a09-a61b-b647ae5e7fd1"
}
```
log into superadmin account
```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cGVyYWRtaW5AZ21haWwuY29tIiwiZXhwIjoxNzQ3Njg0MDk4LCJpYXQiOjE3NDc1OTc2OTgsInN1YiI6IjBmMzk5ZTFhLWVmMTUtNGEwOS1hNjFiLWI2NDdhZTVlN2ZkMSIsInVzZXJJZCI6IiIsInVzZXJUeXBlIjoiTUFOQUdFUiIsInZlcmlmaWVkIjp0cnVlfQ.4VqjfvJnI1N5hPN0Yn_3IJ96DQU_OcVsOm-kLW0_uvA"
}
```

approve a previously createdd manager as superadmin (supply superadmin token)
<!-- http://localhost:8000/api/v1/users/managers/approve -->
```json
// input
{
	"authId": "821b3f18-1d4f-45d9-bedc-6b157dfbc2c3"
}
// output
{
	"message": "user activated",
	"userId": "manager with email address test2@gmail.com has been approved\n"
}
```

## TODO
1. Implementation of permisions so that there is fine grained control
2. using the created permissions, bootstrap a super admin and create managers that can approve 
other managers via permissions set
3. remove the workaround of removing superadmin email encoded in manager `ApproveRegistration` handler
4. add more input data for user and manager during their registration. currently we supplied only name.