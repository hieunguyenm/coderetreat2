#include <algorithm>
#include <vector>
#include <string>

class CodeRetreat {
  public:
    static bool is_palindrome(std::string a) {
      std::string b = a;
      std::reverse(b.begin(), b.end());
      return a.compare(b) == 0;
    }
    static int find_min(std::vector<int>& a) {
      if (a.size() == 0) {
        return -1;
      }
      std::sort(a.begin(), a.end());
      return a[0];
    }
};
