# Checkout Kata

Implement the code for a checkout system that handles pricing schemes such as "pineapples cost 50, three pineapples cost 130."

Implement the code for a supermarket checkout that calculates the total price of a number of items. In a normal supermarket, things are identified using Stock Keeping Units, or SKUs. In our store, we’ll use individual letters of the alphabet (A, B, C, and so on). Our goods are priced individually. In addition, some items are multi-priced: buy n of them, and they’ll cost you y pence. For example, item A might cost 50 individually, but this week we have a special offer: buy three As and they’ll cost you 130. In fact the prices are:

| SKU  | Unit Price | Special Price |
| ---- | ---------- | ------------- |
| A    | 50         | 3 for 130     |
| B    | 30         | 2 for 45      |
| C    | 20         |               |
| D    | 15         |               |

The checkout accepts items in any order, so that if we scan a B, an A, and another B, we’ll recognize the two Bs and price them at 45 (for a total price so far of 95). **The pricing changes frequently, so pricing should be independent of the checkout.**

The interface to the checkout could look like:

```cs
interface ICheckout
{
    void Scan(string item);
    int GetTotalPrice();
}
```

# Approach
This implementation follows the rules of the cata and implements the required interface.
The data and the rules are provded to the checkout function externally, avoiding coupling with specific rules.

There is a very basic pipeline with Github Actions that builds the app and runs the tests.

## Running the application
Pull down the repo and run `go run main.go`. \
To run the tests, navigate to the `checkout` folder and run `go test`

The above commands can abstracted away in a local `bash` script or a `Make` file to improve the commands usage.

## Data

The data and the rules for the application are read from local files, but they can come fom other external sources - api, cache, CDN etc.

## Rules
The checkout is implemented to be agnostic of the rules and to accept these rules as arguments. \
This allows the checkout be used with different rules if we have this requirement and scaled independently.

## Testing
The checkout application contains a single test, making sure the `GetTotalPrice()` function returns the correct amount from the provided data.

Other tests that I thought about to add:
- tsting scanning and totaling with no special prices
- testing behaviour with items not in the pricing rules (thinking about how we want to handle them i.e errors)
- testing the loading of the pricing rules from the JSON file
- edge cases - testing with unusual or unexpected inputs such as negative prices or counts in the pricing rules to see how the app performs

## Observations and further thoughts for improvements
The current application reads data from a single file, but that data can be much bigger and also we might want to initialise a checkout with multiple sets of data of `n` size.

What I was thinking about as a next step to make sure this can handle bigger sets of data is to look into utilising `go routines`, where we can load multiple pricing rules in parallel and serve checkouts in parallel. It is not important for this implementation but it was something that I thought about when I asked myself what if I need to handle mutple rules for multiple checkouts.


