# Product-Rating-API

In this project you will be asked to make a command to calculate product reviews.
- Make 2 command that can list reviews:
  - `review:summary` to calculate all submitted reviews.
  - `review:product <product_id>` to calculate review for the selected product.
- Result of the commands consists of:
  - Total count of reviews.
  - Average from all reviews.
  - Total count from each review rating.

In addition to command, you can make it to next level with:
- Add cache layer to your logic when calculating summary (global or for every product), so that every time there's a query for a review. The program does not need to recalculate the review.
- Add unit tests for your code.

 ## Example Command
 ````
 $ application review:summary
{
  "total_reviews": 500,
  "average_ratings": 2.9,
  "5_star": 62,
  "4_star": 117,
  "3_star": 126,
  "2_star": 123,
  "1_star": 72
}
$ application review:product 1
{
    "total_reviews": 100,
    "average_ratings": 4.9,
    "5_star": 40,
    "4_star": 34,
    "3_star": 16,
    "2_star": 13,
    "1_star": 7
}
 ````
 
 
