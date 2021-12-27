import unittest

class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency
    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)


class TestDollarMoney(unittest.TestCase):
    def testMultiplication(self):
        fiveDollars = Money(5,"USD")
        tenDollars = fiveDollars.times(2)
        self.assertEqual(10, tenDollars.amount)
        self.assertEqual("USD", tenDollars.currency)

class TestEuroMoney(unittest.TestCase):
    def testMultiplication(self):
        tenEuro = Money(10,'EUR')
        twentyEuro = tenEuro.times(2)
        self.assertEqual(20, twentyEuro.amount)
        self.assertEqual("EUR", twentyEuro.currency)


if __name__ == '__main__':
    unittest.main()
