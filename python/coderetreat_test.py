import coderetreat
import unittest

class CodeRetreat(unittest.TestCase):
  def test_palindrome(self):
    self.assertFalse(coderetreat.is_palindrome('notracecar'))
    self.assertTrue(coderetreat.is_palindrome('kayak'))

  def test_find_min(self):
    self.assertEqual(-1, coderetreat.find_min([]), "Find minimum of empty array")
    self.assertEqual(-5, coderetreat.find_min([9, -5, 1, 4]), "Find minimum in [9, -5, 1, 4]")

if __name__ == '__main__':
  unittest.main()
