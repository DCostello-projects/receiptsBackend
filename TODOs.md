Architecture:
<!-- https://drive.google.com/file/d/192OnCTn9nx6H9rGUzLhNE-LlKyQFIhyZ/view?usp=sharing -->

REST API Schema:

input: receipt photo + receipt information
- charity receipts
- purchase receipts (from groceries or any sort of spending that has a receipt)

Receipt photo
    - UUID
    - path to photo(s) in object storage
    - Image blob data from photos
        - uploaded photos
        - user takes photos
    - multiple redundant copies for data resiliency 
    - might want to parse data from the photo? Use OCR backend service?
        - OCR engine into client side application?
        - OCR engine in server side backend?
        - User manual input?
        - look into training AI/ML model to parse receipts?

Receipt information:
    - UUID
    - Date
    - where 
        - store name
        -
    - what
        - misc. list of goods purchased
    - amount spent
        - x amount
        - shipping, etc...
    - payment method
        - credit
        - debit
        - unknown
        - cash
        - paypal

user:
- UUID
- Country = USA ONLY (for now)
- state?
- username --> use AWS cognito user pool and/or federated login with google etc...
- password --> use AWS cognito user pool and/or federated login with google etc...
- name
    - first name
    - last name
- Date of Birth
- Date Joined
- Email
- Phone
- 2FA method(s)
- other linked accounts (amazon, costco etc...)
- multiple receipts
    - multiple receipt images


user -> many receipts 
one receipt -> many photos
            -> one receipt information

API function prototypes:

get_itemized_deductions
    - input: current date, current user 
    - will need: sql query for all receipt data for the past tax year
    - output: total itemized deductions + a presigned URL to download all images for user

create_new_receipt_information
    - input:
        - date
        - where/location (store name)
        - amount spent
        - list of what is purchased
        - payment method
        - path to receipt photo(s) in object storage bucket
    - output: HTTP status code
        - 200
        - 400 --> bad request (bad characters in request, invalid or missing data etc...)
        - 403 -->
        - 409 --> 
        - 429 --> this means we need to rate limit or we are getting too many requests to handle (DDOS)
        - 500 --> internal server error (if something on our side of processing crashes/doesn't work etc...)

delete_receipt_information
    - input:
        - user (UUID)
        - receipt (UUID)
        - path to receipt photo
    - output: HTTP status code

update_receipt_information
    - input:
        - user (UUID)
        - new data of the field we have to update (we have to determine which field?)
    - output: HTTP status code

get_all_user_receipts
    - input:
        - user (UUID)
    - output: all receipt objects for a user + HTTP status code

SPIKE:
- turn any image file into a JPEG (pdfs, docx, png, jpg, etc...)
- End user license agreement (EULA)
- federated login process + sign in with cognito

IRL stores (photos?)
online stores
- potentially look into Amazon API integration?
    - SSO/link accounts with Amazon to query purchase history/orders
- Doordash
- uber
- online deliveries


output: itemized deductions + 
- spending breakdown (how much $ is spent on groceries, eating out, car maintenance etc...)
    - user generated tag, can tag the receipt so the user can create their own spending breakdown


requirements:
- user database (login, persisted user data)


debug build
-> kubernetes cluster resources (backend server + databases)

test build
-> kubernetes cluster resources (backend server + databases)

test build
-> aws resources 

release build