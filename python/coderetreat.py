def is_palindrome(str):
  return str == str[::-1]

def find_min(arr):
  if not arr:
    return -1
  n = sorted(arr)
  return n[0]
