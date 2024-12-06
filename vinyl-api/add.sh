curl http://localhost:8080/albums \
--include \
--header "Content-Type: application/json" \
--request "POST" \
--data '{"id":"4", "title":"mr morale", "artist" : "k." , "price" : 1.00}'

