Welcome to example case #2

In this example, we will fetch detailed information about products and account linked to a shop.

Steps:
1. Get 1 shop information
2. Without concurrency
    a. Fetch all products in the shop
    b. Fetch product info based on shop id and product id
    c. Fetch all accounts for the shop
    d. Fetch account info based on shop id and account id
3. With concurrency
    a. Fetch all products and account, concurrently
    b. Fetch detailed info for the product and account, conccurently
4. Output the time required for concurrent and non concurrent to finish the task
