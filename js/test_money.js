const assert = require('assert');

class Dollar {
    constructor(amount) {
        this.amount = amount;
    }

    times(multipler) {
        return new Dollar(this.amount * multipler);
    }
}

class Money {
    constructor(amount, currency) {
        this.amount = amount;
        this.currency = currency
    }

    times(multipler) {
        return new Money(this.amount * multipler, this.currency);
    }
}

let fiver = new Dollar(5);
let tenner = fiver.times(2);
assert.strictEqual(tenner.amount, 10);

let tenEuros = new Money(10, "EUR");
let twentyEuroes = tenEuros.times(2);
assert.strictEqual(twentyEuroes.amount, 20);
assert.strictEqual(twentyEuroes.currency, "EUR");