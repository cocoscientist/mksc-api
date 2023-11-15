## Mario Kart Super Circuit
> A REST API to return different objects from *Mario Kart Super Circuit*

### Data Items
 - Characters
 ```
    {
        "ID": 1,
        "name": "Mario",
        "image": "https://mario.wiki.gallery/images/0/0f/MKSC_Mario.jpg",
        "speed": 3,
        "weight": 3
    }
 ```
 - Cups
 ```
    {
        "ID": 1,
        "name": "Mushroom Cup",
        "image": "https://mario.wiki.gallery/images/d/df/MKSC_Mushroom_Cup_Emblem.png"
    }
 ```
 - Tracks
 ```
    {
        "ID": 1,
        "name": "Peach Circuit",
        "icon": "https://mario.wiki.gallery/images/c/cd/Peach_Circuit_MKSC_icon.png",
        "map": "https://mario.wiki.gallery/images/d/df/Peach_Circuit.png",
        "extra": false,
        "CupID": 1
    }
 ```
 - Items
 ```
    {
        "ID": 1,
        "name": "Mushroom",
        "icon": "https://mario.wiki.gallery/images/c/c1/Mushroom_MKSC_item_slot_sprite.png",
        "description": "The Dash Mushroom can be used to get a short speed boost. It can also be used off-road to access a shortcut."
    }
 ```

### Description
 - This API was made using Gin and PostgreSQL
 - An API Key is required to perform any requests. The key is validated from the request header.