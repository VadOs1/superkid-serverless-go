1. serverless create -t aws-go-dep 
2. serverless create -t aws-go-dep -p service
3. make clean
4. make build 
5. make deploy
6. sls deploy function -f send-new-order-to-sns
7. sls deploy function -f products
8. sls deploy function -f products-by-article
9. sls deploy function -f featured-products
10. sls invoke local --function send-new-order-to-sns --data '{"body": "{\"message\": \"OK\"}"}'
11. sls invoke local --function products
12. sls invoke local --function products-by-article --data '{ "pathParameters": {"article":"300101"}}'
13. sls invoke local --function featured-products --data '{ "pathParameters": {"article":"300101"}}'