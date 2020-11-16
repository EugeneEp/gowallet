# GoWallet
REST API для кошелька на Golang
____

# /auth/login
## POST
### application/json
{
    "phone":"7-111-111-11-12",
    "password":"11111111"
}
____
# /auth/registration
## POST
### application/json
{
    "phone":"7-111-111-13-12",
    "password":"11111111",
    "confirm":"11111111"
}
____
# /sms/send
## POST
### application/json
{
    "phone":"7-111-211-13-12",
    "action":"reg"
}
____
# /sms/check
## POST
### application/json
{
    "phone":"7-111-111-13-12",
    "action":"reg",
    "code":1111
}
____
# /transactions/get
## GET
### Query params
/transactions/get?size=2&asc&date_from=10.10.2019&date_ro=20.10.2019
____
# /transactions/csv
## GET
### Query params
/transactions/get?size=2&asc&date_from=10.10.2019&date_ro=20.10.2019
____
# /transactions/get/{id}
## GET
### Query params
/transactions/get/3
____
# /user/update
## PUT
### application/json
{
    "identity":{
        "fullname":"test test",
        "passport":124124124,
        "passportIssuedAt":"23.42.1223"
    }
}
____
# /user/upload/picture
## POST
### multipart/form-data
img: file
____
