import java.util.Arrays;

public class CodeRetreat {
  // isPalindrome is case-sensitive.
  public static boolean isPalindrome(String a) {
    return a.equals(new StringBuilder(a).reverse().toString());
  }

  public static int findMinimum(int[] a) {
    if (a.length == 0) {
      return -1;
    }
    Arrays.sort(a);
    return a[0];
  }
}