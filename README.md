##  Thali Price Index ##
A cost of living index for cities across India using user contributed / owned data. TPI v1 focuses on the price of a thali (meal) while v2 will focus on apartment rentals. In theory the platform can be extended to cover all sorts of price data which is otherwise difficult to obtain. 

## Contents     
- [Methodology](#methodology)
- [JSONAPI](#jsonapi)
    * [Validate](#validate)
    * [Create](#create)
    * [Retrieve](#retrieve)
    * [Update](#update)
    * [Delete](#delete)
- [Services / Middleware](#services)
- [Backend](#back)
- [Data](#data)
- [URLs](#urls)
- [App](#app) 
- [Incentives](#incentives)
- [Gotchas](#gotchas)
- [References](#references)  


## Methodology ##
We're a highly stratified society so our thalis are also stratified. A thali can be broadly classified based on the target customer:

1. Blue Collar (Unorganised Labour Workers)
2. Yellow Collar (Retail & Organized Labour Workers)
3. White Collar (Office Workers)
4. Leisure

In addition the thali can also be classified on it's characteristics:

1. Limited
2. Unlimited

Or

1. South Indian
2. North Indian
3. Other Regional

After filtering outliers, the price index will be based on a weighted average of the data collected with a Yellow Collar Limited South Indian Thali in Bengaluru in 2016 being the benchmark of 100.  

## JSONAPI ##

Data contribution, edit & retrieval is done via a simple HTTP/S REST JSON API. 

### VALIDATE ###

POST : Response codes 200 OK or 401 Unauthorized

/auth_token : POST request body must contain {"email":"jan@uary.com", "password":"allurbase") : Response body contains {"token": "1f7e44..."}

DELETE : Response code 200 OK 

/auth_token : DELETE request header should contain valid authorization (Authorization : Bearer 1f7e44...) token which will then be added to blacklist so can't be used further even if not expired. Always returns 200 OK even if POSTed token in invalid. 

GET : Response codes 200 OK or 401 Unauthorized

/hello (test URI for token auth) : GET Request header must contain valid authorization token (Authorization : Bearer 1f7e44...)

PUT : Response codes 202 Accepted or 401 Unauthorized

/auth_token : PUT request header must contain valid unexpired or recently (less than 1hr ago) expired token which gets extended. 200 response body contains refreshed JSON AuthToken {"token": "1e8432..."}. 401 must login again

### CREATE ###

POST : Response 201 Created or 4XX Error. Response body contains Id of created entity

/user : Request body must consist of json formatted User {"name": "Bob", "email": "bob@tpi.org", "password":"allurbase"}

/venue : Request body must consist of json formatted Venue with Name, appengine.Location

/thali : Request body must consist of json with 

/thali/{id} : Request body contains multipart file data

### RETRIEVE ###

GET : Response 200 OK or 404 Not Found. Response in json format unless specified otherwise

/users?offset=20

/user/{id} : Response 

/venues?offset=20

/venue/{id}/thalis?offset=20 - NOT IMPLEMENTED YET

OR

/thalis?offset=20&venue={id} - PREFERRED

/thalis?offset=20

/thali/{id}/image : Response in jpeg encoded bytes

### UPDATE ###

PUT : Response 202 Accepted if update successful or 4XX Error

/user/{id} : Request body contains JSON formatted User : Response body contains Id of updated User

/venue/{id} : Request body contains JSON formatted Venue : Response body contains Id of updated Venue

/thali/{id} : Request body contains JSON formatted Thali : Response body contains Id of updated Thali

### DELETE ###

DELETE : Response 202 Accepted if successful

/user/{id}?email={email}&fullname={name} : GAE ignores request body in DELETE requests so need to use URL parameters

/venue/{id}?name={name}

/thali/{id}?name={name}


## URLs ##

HTML templates for logs and list of users/venues/thalis are available at:

/logs

/list/users

/list/venues

/list/thalis

HTML forms for user/venue/thali creation are available at:

/getform/user

/getform/venue

/getform/thali : using this form directly will assign VenueId field to 0. Use /list/venues and 'Add Thali' instead.

HTML form for upload of image is available at

/list/thalis : select Upload

/image/{id} : sends GET request. HTML template response with base64 encoded image string

## Services ##

### jwt-go ###
Used as authorization middleware to protect some handlers. 

## Back ##

Data provider to backends should be made configurable - currently hard coded. 

## Data ##
In v1 there's three data structures of interest:

+ User
    + Name string
    + Email string
    + Confirmed bool
    + Thalis []Thali // thalis contributed
    + Venues []int64 // venues contributed - []int64 due to datastore restriction of no nested slices
    + Rep int
    + Submitted time.Time

+ Venue
    + Name string
    + Latitude float64 // can be replaced with Location appengine.GeoPoint
    + Longitude float64 // can be replaced with Location appengine.GeoPoint
    + Thalis []int64
    + Submitted time.Time

+ Thali
    + Name string
    + Target int // 1-4 target customer profile
    + Limited bool
    + Region int // 1-3 target cuisine
    + Price float64 //
    + Photo PhotoUrl
    + Venueid int64  // available at venue with id
    + Userid int64 // contributing by user with id
    + Verified bool
    + Accepted bool
    + Submitted time.Time

Along with some other structs

+ Counter
    + Users int
    + Venues int
    + Thalis int

+ PhotoUrl 
    + Url string
    + Size Rect

+ Rect
    + W int
    + H int

User -> Thali = One-to-many

We need a appengine datastore access structure and also a Postgres and/or Mongo access structure for deployment in case of move away from Appengine. All in Go.

In the appengine datastore version, Thali is slightly modified to include Id of Venue rather than a Venue (see appengine datastore reference). 

## App  ##

Mobile app needs to be very simple. Dagger v2 for DI is optional. 

Must have basic modules (networking, storage etc.) and Activities/Fragments to allow User to login/register/logou. Simple unobtrusive drop down menu in top right/left corner with option to logout at any time. 

Must have List/Recyclerview of Venues & Thalis. Each entry in list of Venues to be selectable to activate new Thali entry Activity for selected Venue. Each entry in list of Thalis to be selectable to either show an image or signal Intent to Camera to take a photo and upload.

In v2 use of ConnectivityManager along with the local data persistence model may be included. This will allow creation of local data even in absence of data and later sync with server in presence of network.  

Use JobScheduler 


Preferable to avoid any and all javascript in browser version. Need to consider data contributors with older phones/computers so app needs to be very basic. Should have some basic user validationinterface for Google/Facebook oAuth2, some basic data input functionality and ability to post that data to a server. Ability to get data and display tpi at user's location is secondary. 

Responsiveness of app is of primary importance rather than bells & whistles.

+Android
+ IOS app 

## Incentives ##
The starting group of users will be a small number - 30 colleagues, friends, families, willing acquaintances. So no real need to have a super scalable backend. Other users and spammers will hopefully contribute. 

As soon as a user contributes 10 verified/accepted data points, they get access to the entire data set via the JSON API. 

Spammers should gain negative reputation for every unverified/unaccepted data point and after 10 such points unable to contribute.

## Gotchas ##



## References ##
+ [Writing images to templates](http://www.sanarias.com/blog/1214PlayingwithimagesinHTTPresponseingolang)
+ [Appengine datastore api](https://godoc.org/google.golang.org/appengine/datastore)
+ [GCP Appengine Console](https://console.cloud.google.com/appengine?project=tpi)
+ [Method: apps.repair](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps/repair)
+ [Google Cloud Platform Datastore Reference](https://cloud.google.com/appengine/docs/go/datastore/reference)


