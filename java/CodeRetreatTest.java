// Use JUnit 4 and Hamcrest

import org.junit.Test;

import static org.junit.Assert.*;

public class CodeRetreatTest {
  @Test
  public void testNotPalindrome() {
    assertFalse("isPalindrome: notracecar", CodeRetreat.isPalindrome("notracecar"));
  }

  @Test
  public void testIsPalindrome() {
    assertTrue("isPalindrome: kayak", CodeRetreat.isPalindrome("kayak"));
  }

  @Test
  public void testFindMinEmpty() {
    assertEquals("Find minimum of empty array", -1, CodeRetreat.findMinimum(new int[]{}));
  }

  @Test
  public void testFindMinimumNonEmpty() {
    assertEquals("Find minimum in [9, -5, 1, 4]", -5, CodeRetreat.findMinimum(new int[]{9, -5, 1, 4}));
  }
}
