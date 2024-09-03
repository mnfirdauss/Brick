# Brick
------------------
**How To Run**
* migrate database
  `make migrate:up`
* run app
  `make start`

**Save Account**

  `Create a new account`

* **URL**

  `/add-account`

* **Method:**

  `POST`

* **Body**

	```
	{
		"account_name": "daus",
		"account_number": 12345,
		"bank_name": "bsi"
	}
	```

* **Success Response:**

  * **Code:** 200
    **Content:**
		```
		{
      "account_number": 12345,
      "account_name": "daus",
      "bank_name": "bsi"
    }
		```
 
------------------
**Validate Account**

  `Check if account is valid`

* **URL**

  `/validate-account`

* **Method:**

  `POST`   

* **Body**

  ```
		{
      "account_name": "daus",
      "account_number": 12345,
      "bank_name": "bsi"
    }
	```


* **Success Response:**

  * **Code:** 200 OK
    **Content:**
	```
	  {
      "account_name": "daus",
      "account_number": 12345,
      "bank_name": "bsi"
    }
		```

------------------
**Send Money**

  `To send money from an account to another account`

* **URL**

  `/transaction/transfer`

* **Method:**

  `POST`   

* **Body**

  ```
  {
    "source_account": {
        "account_name": "daus",
        "account_number": 12345,
        "bank_name": "mandiri"
    },
    "destination_account": {
        "account_name": "daus",
        "account_number": 12345,
        "bank_name": "bsi"
    },
    "amount": 1000
  }
  ```

* **Success Response:**

  * **Code:** 200 OK
    **Content:**
	 	```
	 	{
      "id": "31712f3b-2497-499c-ae84-60a51f4a28d6",
      "source_account": {
          "account_number": 12345,
          "account_name": "daus",
          "bank_name": "mandiri"
      },
      "destination_account": {
          "account_number": 12345,
          "account_name": "daus",
          "bank_name": "bsi"
      },
      "amount": 1000,
      "status": "PENDING",
      "created_at": "2024-09-03T08:11:17.669204Z",
      "updated_at": "2024-09-03T08:11:17.669205Z"
    }
		```

------------------
**Callback Transaction**

  `Receive callback from bank to validate if transfer is success or not`

* **URL**

  `/transaction/callback`

* **Method:**

  `POST`   

* **Body**

  ```
		{
       "id": "31712f3b-2497-499c-ae84-60a51f4a28d6",
       "status": "SUCCESS"
    }
	```

* **Success Response:**

  * **Code:** 200 OK
    **Content:**
	 	```
	 	{
        "id": "31712f3b-2497-499c-ae84-60a51f4a28d6",
        "source_account": {
            "account_number": 12345,
            "account_name": "daus",
            "bank_name": "mandiri"
        },
        "destination_account": {
            "account_number": 12345,
            "account_name": "daus",
            "bank_name": "bsi"
        },
        "amount": 1000,
        "status": "SUCCESS",
        "created_at": "2024-09-03T08:11:17.669204Z",
        "updated_at": "2024-09-03T08:14:43.45506Z"
    }
		```
