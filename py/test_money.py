import unittest

class Dollar:
    def __init__(self, amount):
        self.amount = amount
    def times(self, multiplier):
        return Dollar(self.amount * multiplier)

class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency
    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)


class TestMoney(unittest.TestCase):
    def testMultiplication(self):
        fiver = Dollar(5)
        tenner = fiver.times(2)
        self.assertEqual(10, tenner.amount)

class TestEuroMoney(unittest.TestCase):
    def testMultiplicationInEuro(self):
        tenEuro = Money(10,'EUR')
        twentyEuro = tenEuro.times(2)
        self.assertEqual(20, twentyEuro.amount)
        self.assertEqual("EUR", twentyEuro.currency)


if __name__ == '__main__':
    unittest.main()
