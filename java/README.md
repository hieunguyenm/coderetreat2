# Java unit testing

It is preferable to use an IDE for Java since IDEs usually have testing frameworks bundled.

**NOTE: You must use JUnit 4.12.**

For example, IntelliJ has JUnit bundled with it. 

1. Right click on the project in the Project view.
2. Choose "Open Module Settings".
3. Choose "Libraries" on the left hand side.
4. Click "+" at the top and then "Java".
5. You need to add `junit-4.12.jar` and `hamcrest-core-1.3.jar`. These JARs are located in `<IntelliJ install directory>/lib`.
6. To run tests, you can click "Run CodeRetreatTest" under the "Run" tab.

If you do not use IntelliJ, there are instructions for your chosen IDE online. If your IDE does not have JUnit 4.12 bundled, the necessary JARs are in this folder.

If you prefer to run tests from the command line, do the following from inside this folder:

1. Compile classes with `javac -cp hamcrest-core-1.3.jar:junit-4.12.jar:. CodeRetreatTest.java`.
2. Run the test class with `java -cp hamcrest-core-1.3.jar:junit-4.12.jar:. org.junit.runner.JUnitCore CodeRetreatTest`.

