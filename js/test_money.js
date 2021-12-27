const assert = require('assert');

class Money {
    constructor(amount, currency) {
        this.amount = amount;
        this.currency = currency
    }

    times(multipler) {
        return new Money(this.amount * multipler, this.currency);
    }
}

let fiveDollars = new Money(5,"USD");
let tenDollars = fiveDollars.times(2);
assert.strictEqual(tenDollars.amount, 10);
assert.strictEqual(tenDollars.currency, "USD");

let tenEuros = new Money(10, "EUR");
let twentyEuroes = tenEuros.times(2);
assert.strictEqual(twentyEuroes.amount, 20);
assert.strictEqual(twentyEuroes.currency, "EUR");