# GoGramm

This is a simple web service providing two endpoints:

## `POST /load`

Upload an array of words:

    $ curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]'
    
    // 200 OK

## `GET /get`

Find anagrams for a word provided in required query parameter `word`:

    $ curl 'localhost:8080/get?word=foobar' 
    
    ["foobar", "boofar"]
