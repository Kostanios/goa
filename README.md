# goa
Simple use case of Goa framework (https://goa.design)
## Run server
```
iMac-2:goa jt$ go run *.go
2017/10/11 22:26:19 [INFO] mount ctrl=Hotels action=Search route=GET /hotel/search
2017/10/11 22:26:19 [INFO] listen transport=http addr=:8080
2017/10/11 22:26:57 [INFO] started req_id=Igf7A9OZRf-1 GET=/hotel/search from=::1 ctrl=HotelsController action=search
2017/10/11 22:26:57 [INFO] headers req_id=Igf7A9OZRf-1 User-Agent=curl/7.52.1 Accept=*/*
2017/10/11 22:26:57 [INFO] completed req_id=Igf7A9OZRf-1 status=200 bytes=2360 time=29.592371ms ctrl=HotelsController action=search
...
```
## GET /hotel/search
```
iMac-2:~ jt$ curl -i http://localhost:8080/hotel/search
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 11 Oct 2017 19:26:57 GMT
Transfer-Encoding: chunked

[
    {
        "id": "169173",
        "name": "The Crown Borneo"
    },
    {
        "id": "169173",
        "name": "The Crown Borneo"
    },
    {
        "id": "169173",
        "name": "The Crown Borneo"
    },
    {
        "id": "175947",
        "name": "Tune Hotel - 1Borneo Kota Kinabalu"
    },
    {
        "id": "151064",
        "name": "Kinabalu Daya Hotel"
    },
    {
        "id": "151064",
        "name": "Kinabalu Daya Hotel"
    },
    {
        "id": "151064",
        "name": "Kinabalu Daya Hotel"
    },
    {
        "id": "151064",
        "name": "Kinabalu Daya Hotel"
    },
    {
        "id": "151064",
        "name": "Kinabalu Daya Hotel"
    },
    {
        "id": "151064",
        "name": "Kinabalu Daya Hotel"
    },
    {
        "id": "169174",
        "name": "Megah D'Aru"
    },
    {
        "id": "124720",
        "name": "Cititel Express Kota Kinabalu"
    },
    {
        "id": "124720",
        "name": "Cititel Express Kota Kinabalu"
    },
    {
        "id": "124720",
        "name": "Cititel Express Kota Kinabalu"
    },
    {
        "id": "124720",
        "name": "Cititel Express Kota Kinabalu"
    },
    {
        "id": "74019",
        "name": "The Palace Hotel Kota Kinabalu"
    },
    {
        "id": "74019",
        "name": "The Palace Hotel Kota Kinabalu"
    },
    {
        "id": "187167",
        "name": "Winner Hotel"
    },
    {
        "id": "147333",
        "name": "Celyn Hotel City Mall"
    },
    {
        "id": "147333",
        "name": "Celyn Hotel City Mall"
    },
    {
        "id": "187474",
        "name": "Celyn City Hotel"
    },
    {
        "id": "187474",
        "name": "Celyn City Hotel"
    },
    {
        "id": "195390",
        "name": "Oceania Hotel"
    },
    {
        "id": "195390",
        "name": "Oceania Hotel"
    },
    {
        "id": "195390",
        "name": "Oceania Hotel"
    },
    {
        "id": "195390",
        "name": "Oceania Hotel"
    },
    {
        "id": "74020",
        "name": "Sabah Oriental"
    },
    {
        "id": "169997",
        "name": "Shangri-La"
    },
    {
        "id": "169997",
        "name": "Shangri-La"
    },
    {
        "id": "187230",
        "name": "Zara's Boutique Hotel @ Harbour City"
    },
    {
        "id": "187230",
        "name": "Zara's Boutique Hotel @ Harbour City"
    },
    {
        "id": "71680",
        "name": "Promenade Hotel Sabah"
    },
    {
        "id": "156613",
        "name": "Ming Garden Hotel & Residences"
    },
    {
        "id": "156613",
        "name": "Ming Garden Hotel & Residences"
    },
    {
        "id": "156613",
        "name": "Ming Garden Hotel & Residences"
    },
    {
        "id": "71684",
        "name": "Nexus Resort & Spa Karambunai"
    },
    {
        "id": "71684",
        "name": "Nexus Resort & Spa Karambunai"
    },
    {
        "id": "71684",
        "name": "Nexus Resort & Spa Karambunai"
    },
    {
        "id": "74769",
        "name": "Sutera Harbour - The Pacific Sutera"
    },
    {
        "id": "74769",
        "name": "Sutera Harbour - The Pacific Sutera"
    },
    {
        "id": "71682",
        "name": "Sutera Harbour Resort - Magellan Sutera"
    },
    {
        "id": "71685",
        "name": "Shangri-Laâ\u0080\u0099s Rasa Ria Resort & Spa"
    },
    {
        "id": "71685",
        "name": "Shangri-Laâ\u0080\u0099s Rasa Ria Resort & Spa"
    },
    {
        "id": "174326",
        "name": "Gaya Island Resort"
    },
    {
        "id": "74021",
        "name": "Gayana Eco Resort Kota Kinabalu"
    },
    {
        "id": "125457",
        "name": "Bunga Raya Island Resort"
    },
    {
        "id": "125457",
        "name": "Bunga Raya Island Resort"
    }
]
```
