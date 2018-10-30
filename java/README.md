# Java unit testing

It is preferable to use an IDE for Java since IDEs usually have testing frameworks bundled.

For example, IntelliJ has JUnit bundled with it. 

1. Right click on the project in the Project view.
2. Choose "Open Module Settings".
3. Choose "Libraries" on the left hand side.
4. Click "+" at the top and then "Java".
5. You need to add `junit-4.12.jar` and `hamcrest-core-1.3.jar`. These JARs are located in `<IntelliJ install directory>/lib`.
6. To run tests, you can click "Run CodeRetreatTest" under the "Run" tab.

**You must use JUnit 4.12.**

If you do not use IntelliJ, there are instructions for your chosen IDE online. If your IDE does not have JUnit 4.12 bundled, you can download the JARs at the links below.

* [JUnit 4.12](https://github.com/junit-team/junit4/releases/download/r4.12/junit-4.12.jar)
* [Hamcrest Core](https://search.maven.org/remotecontent?filepath=org/hamcrest/hamcrest-core/1.3/hamcrest-core-1.3.jar)
