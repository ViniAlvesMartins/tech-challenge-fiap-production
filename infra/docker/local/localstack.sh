# -- > List SNS Topics
echo Listing topics ...
echo $(aws --endpoint-url=http://localhost:4566 sns list-topics --profile test-profile --region us-east-1 --output table | cat)

# -- > Send SNS order created event
echo Sending SNS order created event ...
echo $(aws --endpoint-url=http://localhost:4566 sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:order_created-topic --message "{\"id\": 1, \"client_id\": 1, \"order_status\": \"AWAITING_PAYMENT\", \"amount\": 123.4, \"created_at\": \"2016-11-01T20:44:39Z\", \"products\": [ { \"id\": 1, \"product_name\": \"bolinho\" } ]}" --profile test-profile --region us-east-1)


# -- > Send SNS payment confirmed
echo Sending SNS payment confirmed event ...
echo $(aws --endpoint-url=http://localhost:4566 sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:payment_status_updated-topic --message "{\"id\": 1, \"status\": \"CONFIRMED\"}" --profile test-profile --region us-east-1)
