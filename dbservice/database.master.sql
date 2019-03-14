CREATE DATABASE IF NOT EXISTS cs53 COLLATE = utf8_general_ci;

USE cs53;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;

--
-- Database: `cs53`
--

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `adminfaq`
--

CREATE TABLE `adminfaq`
(
  `id`        int(11)                     NOT NULL,
  `timestamp` datetime                    NOT NULL,
  `questions` text COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `adminfaq`
--

INSERT INTO `adminfaq` (`id`, `timestamp`, `questions`)
VALUES (1, '2018-07-21 17:45:00',
        'Q: How do I make a course?\n--------------------------------\n**A:** Dashboard -> Courses -> new. And create the course there\n\nQ: How do I invite students to course?\n--------------------------------\n**A:** You go to [admin/course](/admin/course) or [admin/](/admin) and on the course card, click the copy button to get the `join course through link` and send that to all students in preferred way (ex: email)\n\nQ: How do I make an assignment?\n--------------------------------\n**A:** Dashboard -> Assignments-> new. And create the assignment there');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `assignments`
--

CREATE TABLE `assignments`
(
  `id`            int(11)     NOT NULL,
  `name`          varchar(64) NOT NULL,
  `description`   text,
  `created`       timestamp   NOT NULL,
  `publish`       datetime    NOT NULL,
  `deadline`      datetime    NOT NULL,
  `course_id`     int(11)     NOT NULL,
  `submission_id` int(11) DEFAULT NULL,
  `review_id`     int(11) DEFAULT NULL,
  `validation_id` int(11)     NULL,
  `reviewers`     int(11)     NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `assignments`
--

INSERT INTO `assignments` (`id`, `name`, `description`, `created`, `publish`, `deadline`, `course_id`, `submission_id`,
                           `review_id`, `validation_id`, `reviewers`)
VALUES (1, 'Lab 1: in-memory banking app',
        '# Lab 1: in-memory banking app\r\n\r\nFor this lab, we will create an application that mocks a simple mobile banking app. The app consists of 3 activities, that allow a user to check current account balance, make a payment to a friend, and see the list of recent transactions. Note: the app is in-memory stateful, but, it **should not** use persistence or any networking facilities. All the data within the app is mocked. When the app is started, the app stores the internal state in-memory, until the app is \"restarted\" afresh, ie. killed by the Android OS and re-initialized upon requesting the entry-level Activity. \r\n\r\n\r\n## MainActivity\r\n\r\nThe first, entry activity will be called MainActivity. It will be the main entry point for the app, and the \"main menu\" for the app. It will consist of:\r\n\r\n* A fixed label with text: \"Balance [EUR]\"\r\n* Another fixed label with a text representing the actual account balance. The activity should generate a random integer that represents the current account balance, and it should be between 90 and 110. We will refer to this field as `lbl_balance`\r\n* A button with the label: \"Transactions\", referred as `btn_transactions`\r\n* A button with the label \"Transfer\", referred as `btn_transfer`\r\n\r\nBehaviour:\r\n* upon starting up the activity, the field `lbl_balance` should have the newly generated random value between 90..110 euros. \r\n* pressing btn_transfer moves the user to the `TransferActivity`\r\n* pressing btn_transactions moves the user to the `TransactionsActivity`\r\n* pressing the normal android \"back button\", quits the app.\r\n\r\n\r\n## TransferActivity\r\n\r\nThis activity will have the following UI elements:\r\n\r\n* A fixed text label \"Recipient\" \r\n* A drop-down list with names of your friends, e.g. \"Alice, Bob, Charlie, Dawn, Elvis, Frode\". The list must consist of at least 6 names. \r\n* A fixed label with text: \"Amount [EUR]\"\r\n* A single editable text field referred as: `txt_amount`\r\n* A single label with no text, referred as `lbl_amount_check`\r\n* A button with label \"Pay\", referred as: `btn_pay`, that is disabled by default\r\n\r\nBehaviour:\r\n* The `btn_pay` should not be enabled unless the recipient and amount are correctly selected. If the amount is provided and it is within correct bounds, and when the recipient is selected, then `btn_pay` is enabled.\r\n* Pressing `btn_pay` when enabled triggers the payment, which means, that the value from `txt_amount` is subtracted from the user total balance; a new transaction is made to the recipient selected from the drop-down. This new transaction is added to the list, subsequently visible in the `TransactionsActivity`. After the payment is triggered the user is moved back to the `MainActivity`.\r\n* When the amount is entered, and it is less or equal to 0, or, if the amount is greater then the current user balance, then `btn_pay` continues to be disabled, and the `lbl_amount_check` will display to the user the error, eg. \"Amount is outside of bounds\" or something similar. \r\n* pressing the normal android \"back button\" takes the user back to the `MainActivity` and no new transaction is made.\r\n\r\n\r\n## TransactionsActivity\r\n\r\nThis activity will communicate to the user all the historical transactions.  The first transaction should be the \"founding\" transaction with the total amount of the initial user balance. Therefore, for example, if the random number generated as the initial balance for the user account is 95.43 Euros, then, the very first transaction recorded should be the one on the very moment the number was generated, and it should be, for example: \r\n```\r\n15:30:21    |   Angel    |   95.43   |  95.43\r\n```\r\nwhich means, that at 15:30:21 (when the clock was at 3.30pm with 21 secs) the amount 95.43 Euros have been transferred to the user account, and the final balance on the account is 95.43 (the last column is the final balance on the account). Any subsequent payments, will be listed later as \"payments\" and the balance will be appropriately updated, eg:\r\n```\r\n15:30:21    |   Angel    |   95.43   |  95.43\r\n15:31:14    |   Alice    |   15.43   |  80.00\r\n15:32:18    |   Charlie  |   60.00   |  20.00\r\n```\r\n\r\nThe list items formatting is up to the implementation, but, it needs to communicate those 4 fields: exact time (hours:mins:sec of the time of payment), recipient, amount, final account balance.  Use your creativity to format the list items accordingly. \r\n\r\n\r\nBehaviour:\r\n* selecting (short-click) on the list items has no effect, but, user should be allowed to select a single item from the list\r\n* long-click on the list item displays a Toast that says \"XXX YYY\" where XXX is the name of the recipient from this particular transfer, and YYY is the amount.\r\n* pressing the normal android \"back button\" takes the user back to the `MainActivity` \r\n\r\n\r\n# Notes\r\n\r\n* You do not need to control for the app \"restarts\" - this should be left to Android. \r\n* You should not program the handling of the Android\'s \"back button\" presses - the behaviour should be the standard Android behaviour - in other words, to not code any back-button handling in your code. \r\n* You should use XML-based UI designs for your Activities, with the exception of programmatically populating the \"Recipients\" drop-down list. \r\n* You should also use XML-based view template for the Transaction list items.\r\n* The app should work in both, portrait and landscape modes.\r\n\r\n\r\n# Questions to consider\r\n\r\nWhen implementing the application with the above specification, not everything is precisely defined. Some things require you to think about the implementation details. This list of questions guides you through the process of clarifying the actual implementation details not covered in the SPEC. \r\n\r\n* Why do we represent the account balance as Integer? EURO has cents, and therefore it is a float with two decimal places, so why it is a bad idea to represent it as float for example? \r\n* How will you convert back-and-forth from Integer to a currency that has two decimal places in the UI for the user?\r\n* How will you represent the Transactions? Should it be a struct? Or class? \r\n* How will you represent the Recipients? \r\n* Where to store the application state in-memory, such as user account balance, and the list of transactions?\r\n* How will the Activities communicate data between themselves? \r\n* Who will be responsible for actually making the transfer: MainActivity or TransferActivity? Why?\r\n\r\n\r\n# Marking \r\n\r\nMarking/testing/scoring of the implementation is based on the form specified below. The total score is based on the sum of the individual items tested. The test either \"Fails\" or \"Passes\" - the task cannot be \"Partial\", ie. partial means \"Fail\". Each test case has a score associated with it - most are worth 1 point, but some are worth more. The marking schedule is based on ticking the appropriate test cases and then summing up those that have been ticked. \r\n\r\n* [ ] (1pt) The app has a custom icon (not the default Android one).\r\n* [ ] (1pt) The app `MainActivity` loads.\r\n* [ ] (1pt) The app\'s `MainActivity` contains all required UI elements as per SPEC.\r\n* [ ] (1pt) Pressing Android\'s \"back button\" on the `MainActivity` always quits the app, AND, this behaviour is not hardcoded in the code, ie. the code does not handle the Android back button presses, but instead, relies on the default Android behaviour.\r\n* [ ] (1pt) Pressing `btn_transactions` moves the user to TransactionsActivity.\r\n* [ ] (1pt) The default founding transaction from Angel is done correctly and visible in the `TransactionsActivity`. The user balance in `lbl_balance` matches the funding transaction.\r\n* [ ] (1pt) TransactionsActivity shows new payments correctly.\r\n* [ ] (1pt) TransactionsActivity moves the user back to MainActivity on \"back botton\" press.\r\n* [ ] (1pt) Pressing btn_transfer moves the user to TransferActivity.\r\n* [ ] (5pts) The `TransferActivity` has all required UI, the balance amount handling is working correctly with the appropriate lbl_amount_check error messages, and the btn_pay is enabled/disabled when appropriate. Pressing the `btn_pay` creates the new transaction, visible in TransactionsActivity, and updates the user balance accordingly.\r\n* [ ] (1pt) The app never crashes. In particular, changing from landscape to portrait and vice-versa does not crash the app.\r\n\r\n**Bonus tasks**\r\n\r\n* [ ] (2pts) The app UI elements are all well laid-out, look appealing and tidy, and the app has a solid/professional feel, ie. buttons properly centred, UI elements well-arranged, colour schemes appealing, and so on. The app works well in both, portrait and landscape modes. \r\n* [ ] (1pt) The app codebase is well-structured, readable, and follows professional Java and Android practices.\r\n* [ ] (1pt) The Recipient list for the drop-down is done programmatically such that the list source of names is provided as a string array from MainActivity, eg. `list = String[]{\"Alice\", \"Bob\", \"Charlie\", \"Dawn\", \"Elvis\", \"Frode\"};` This requires to dynamically update the TransferActivity UI based on the data passed to it from `MainActivity`. \r\n* [ ] (1pt) An extra bonus point for something outside of scope. \r\n\r\n\r\n**Total: 20**',
        '2019-02-28 14:25:10', '2019-02-28 16:24:00', '2019-03-24 23:59:00', 3, 1, 1, NULL, 2);
-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `course`
--

CREATE TABLE `course`
(
  `id`          int(11)                                       NOT NULL,
  `hash`        varchar(64) COLLATE utf8_danish_ci            NOT NULL,
  `coursecode`  varchar(10) COLLATE utf8_danish_ci            NOT NULL,
  `coursename`  varchar(64) COLLATE utf8_danish_ci            NOT NULL,
  `teacher`     int(11)                                       NOT NULL,
  `description` text                                          NOT NULL,
  `year`        int(11)                                       NOT NULL,
  `semester`    enum ('fall','spring') COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `course`
--

INSERT INTO `course` (`id`, `hash`, `coursecode`, `coursename`, `teacher`, `description`, `year`, `semester`)
VALUES (1, '3876438629b786', 'IMT1031', 'Grunnleggende Programmering', 2, 'Write hello, world in C++', 2019, 'fall'),
       (2, '12387teg817eg18', 'IMT1082', 'Objekt-orientert programmering', 2, 'Write Wazz up world in Python', 2019,
        'fall'),
       (3, '12e612eg1e17ge1', 'IMT3673', 'Mobile/Wearable Programming', 2,
        '# Lectures video recordings\r\n\r\nThe lecture videos are under the [Mobile Development playlist](https://www.youtube.com/playlist?list=PL17KQCa8hhvB1gKYcCH5MBi5Thc8gjnG5), on [GTL Youtube](https://www.youtube.com/channel/UCoChi3eU2ThtVkhIpq_F20Q) channel. [This year (2019) playlist](https://www.youtube.com/playlist?list=PL17KQCa8hhvA1kYaqG7Psxdx4RpHbWweQ) contains this year videos (for ease of access).\r\n\r\n# Dates\r\n\r\n* Labs submission deadline. Labs 1/2 **March 24**. Labs 3/4: **April 7th**.\r\n* Project submission deadline: **May 2nd**\r\n* Project presentations will start Thursday, 10:00, on May 2nd. \r\n* Written exam: **May 28**.\r\n\r\n\r\n# Lectures\r\n\r\n* Lecture 1, January 8: Introduction to the course. Logistics. Expectations.\r\n\r\n* Lecture 2, January 10: Background\r\n	* [lecture video](https://www.youtube.com/watch?v=1GVIUJESd20)\r\n   * Tutorial session, 2018, video lecture: [part1](https://youtu.be/c_qYJ1ovtjg) (Java, conventions) [part2](https://www.youtube.com/watch?v=bTeDFA-NU-8) (Android Studio, debugging)\r\n   * [web lecture: Java for C++ developers](http://pages.cs.wisc.edu/~hasti/cs368/JavaTutorial/)\r\n   * [pdf: Java programming for C/C++ developer](https://www.seas.upenn.edu/~cis1xx/resources/JavaForCppProgrammers/j-javac-cpp-ltr.pdf)\r\n   * [Java programming, code conventions and rules](https://docs.google.com/presentation/d/19shHLM6Co5vvxBQ43TH58LxjvhPY2p9IHreWBEIMvjQ/edit?usp=sharing)\r\n\r\n* Lecture 3, January 15: Background to mobile technologies\r\n   * [lecture slides](https://docs.google.com/presentation/d/1AEwES01hEf3-QY4YghovXsh4bprn8H_ry0e_8T1UXpo/edit?usp=sharing)\r\n   * 2019 [video lecture](https://www.youtube.com/watch?v=gHNlk6gZo-4&index=35&list=PL17KQCa8hhvB1gKYcCH5MBi5Thc8gjnG5)\r\n   * 2018 video lecture: [part1](https://youtu.be/dKzLDLmiG-4), [part2](https://youtu.be/0PPHHWIdhMA)\r\n\r\n* Lecture 4, January 16: Introduction to Android\r\n   * [lecture slides](https://docs.google.com/presentation/d/1vFysyTckSz8UzhEgN3NKB-tjbmLtEgCx8gT8uSDyMNU/edit?usp=sharing)\r\n   * 2019 [video lecture](https://www.youtube.com/watch?v=EGBqMsajRiY)\r\n   * 2018 video lecture: [part1](https://www.youtube.com/watch?v=cfuXPbTngdM), [part2](https://www.youtube.com/watch?v=nrzW8laKHTc)\r\n\r\n* [Practical session, January 17th](Android Studio Basics)\r\n\r\n* Lecture 5, January 22: Java and Android concurrency.\r\n   * [lecture video](https://youtu.be/9gke5NW951w)\r\n   * code example\r\n\r\n* Lecture 6, January 24: Android UI Development. XML format for UI. iOS and Apple ecosystem overview.\r\n   * [Projects, Course logistics, intro to Bundles, data passing](https://youtu.be/ARRVZzf-6YI), 2019 video\r\n   * [How to pass data between activities](https://youtu.be/v8BDxh7vxUg), 2019 video\r\n   * UI development [slides](https://docs.google.com/presentation/d/1Awc5h5IP37wyPETiUbEym4uYnYILF2B8tOxaY0k2WM0/edit#usp=sharing)\r\n   * [2018 video lecture](https://youtu.be/FxC61ktETrw)\r\n\r\n* Lecture 7, January 29: Background lecture on Mobility. Telephony systems and Mobile Communication.\r\n   * **Q&A session**\r\n   * [slides](https://docs.google.com/presentation/d/16xaadWxzgCd5-lXKypuF1nJQUYIM5nmOg3hhCkbmQhw/edit?usp=sharing)\r\n   * [2018 video lecture](https://www.youtube.com/watch?v=b7tiFXCRIM4)\r\n\r\n* Practical session, January 30, [Android Preferences (local, global)](https://developer.android.com/training/data-storage/shared-preferences)\r\n\r\n* Lecture 8, January 31\r\n   * [2019 lecture video](https://youtu.be/f7QjznjQQtY) Lab 1 and Lab 2. Discussion. Architecture and software design. [OO Design patterns](https://en.wikipedia.org/wiki/Design_Patterns) (Gang-of-Four). [SOLID principles](https://en.wikipedia.org/wiki/SOLID). MVC, [model-view-controller pattern](https://en.wikipedia.org/wiki/Model–view–controller).\r\n   * [2018 lecture video](https://youtu.be/fH0Ridp7M9o) Android concurrency. Android UI thread, main thread, and UI updating.\r\n   * [code example](https://gist.github.com/anonymous/a18004afd5a94f4130a1952092aa0e52)\r\n\r\n* Lecture 9, February 5: Android Persistence. Android file system. Internal storage. SQLite.\r\n  * [Slides](https://drive.google.com/file/d/1xB2_aWWRU8i2oy9VUMZ5-y4kkepX2e-s/view?usp=sharing)\r\n  * [Video lecture, part 1](https://youtu.be/v4LxCMWr7V8)\r\n  * [Video lecture, part 2](https://youtu.be/eSdZ0nFJEXs)\r\n\r\n\r\n* Lecture 10, February 6: Services, Service life-cycle.\r\n  * [Slides](https://drive.google.com/file/d/1_ABcudp0uI7xPfaoDOki_zExlwI6TwXB/view?usp=sharing)\r\n  * [Video lecture, part 1](https://youtu.be/Le8isG5xeiM)\r\n  * [Video lecture, part 2](https://youtu.be/Jdp34l16qK8)\r\n\r\n\r\n* Practical session, February 7: SQLite through [Room](https://developer.android.com/training/data-storage/room/)\r\n\r\n\r\n* Lecture 11, February 12: Notifications. Notification management\r\n  * [Slides](https://drive.google.com/file/d/1EJJyZgvA9um8LZlyp9dQyoXAGfWZPq47/view?usp=sharing)\r\n  * [Video lecture, part 1](https://youtu.be/FODyg5B8V1M)\r\n  * [Video lecture, part 2](https://youtu.be/W-ySUk4fOX4)\r\n  \r\n* Lecture 12, February 13: Intents and Broadcast Receivers\r\n  * [Slides](https://drive.google.com/file/d/1AnFElOctDlvC183GgK98pmowoiQtNo-e/view?usp=sharing)\r\n  * [Video lecture, part 1](https://youtu.be/VH7rGMDV4Kw)\r\n  * [Video lecture, part 2](https://youtu.be/JlCipnK20Bs)\r\n\r\n* Practical session, February 14: RecyclerView\r\n\r\n* Lecture 13, February 19: Sensors\r\n  * [Slides](https://docs.google.com/presentation/d/1COe-sOGtZBlMEeDgU6F7UeYRmw24Iq47t8ULnieoe9k/edit?usp=sharing)\r\n  * [Lecture video](https://youtu.be/E2asXLlsKo0)\r\n\r\n\r\n* Lecture 14, February 20: Android Networking: NFC & Wifi; Introduction to Volley\r\n   * [Slides](https://drive.google.com/file/d/12vHAqZXv6GhGHf3huObSIrMAelltz0eg/view?usp=sharing)\r\n   * [How to parse XML in Java](https://www.quora.com/What-is-the-best-way-to-parse-XML-in-Java-without-external-libraries?__nsrc__=4) - Quora answer.\r\n\r\n* Lecture 15, February 26: Build Tools\r\n   * [Slides](https://drive.google.com/file/d/1OdcyXQj53iik-oWVjrUa7bd1m2K9QrOa/view?usp=sharing)\r\n\r\n* Lecture 16, February 27: [Testing](https://drive.google.com/file/d/1a3UXdIV8rlVd0WBMX-QADjhYkDvHLYUH/view?usp=sharing)\r\n\r\n* Practical session, February 28: [Junit testing in android](https://developer.android.com/training/testing/junit-rules)\r\n\r\n* Lecture 17, March 5: [Advanced Testing](https://drive.google.com/file/d/1rdhFMuyxYpEtUiZhr1BJ3lsO5wn1Q6MR/view?usp=sharing)\r\n\r\n* Lecture 18, March 6: [Data Formats for Serialisation and Transmission](https://drive.google.com/file/d/1ird35rTS6H1OTOVa4iQWTzHGQ2XNLYl2/view?usp=sharing)\r\n\r\n# Labs\r\n\r\n* [Lab 1](lab_1)\r\n* [Lab 2](lab_2)\r\n* [Lab 3](lab_3)\r\n* Lab 4: TBA\r\n\r\n\r\n# Projects 2019\r\n\r\n* [Project IDEAS](Project Ideas) This is the list of potential projects that groups can pick from. We will have \"first-come-first-served\" policy on projects with external stakeholders, such that there is only one group per project stakeholder (this is to ease the communication with the idea providers).\r\n* [Projects 2019](Project Groups) This page shall contain the list of groups and associated projects.\r\n\r\n\r\n\r\n# Practical Exercises\r\n\r\n* [Basic Java](Exercise Basic Java)\r\n* [Advanced Java](Exercise Java)\r\n* [Playing with Basic Android UI](Exercise Basic Android UI)\r\n* [Android Sensors](Exercise Sensors)\r\n\r\n\r\n# [FAQ](faq)\r\n\r\n# Code examples\r\n\r\n* [Android project setup](project-setup)\r\n* [Code repository](https://bitbucket.org/gtl-hig/imt3662-mobile-labs/overview) \r\n* [Android project .gitignore](android-gitignore)\r\n* [Shared Preferences](shared-preferences)\r\n* [Using Camera2 API](https://github.com/googlesamples/android-Camera2Basic)\r\n\r\n\r\n\r\n# Tutorials\r\n\r\n* [Android Programming at Stanford](http://web.stanford.edu/class/cs193a/videos.shtml)\r\n* [Preferences - tutorial by Bjorn](https://docs.google.com/presentation/d/1OQHiPIzIjhAUS6Txl4QMjIeUNR_Xw-1mSBrr5ukSLZ4/edit?usp=sharing)\r\n* [Content Providers](https://developer.android.com/guide/topics/providers/content-providers.html)\r\n* [Setting up GoMobile](gomobile-setup)\r\n* [Kalman filter tutorial](https://home.wlu.edu/~levys/kalman_tutorial/)\r\n\r\n\r\n## News, articles, blog posts\r\n\r\n* [Material Motion](https://www.youtube.com/watch?v=aZ5V5e-phR8) youtube \r\n* [Angular 5.0](https://appinventiv.com/blog/angular-5-0-means-mobile-app-development/)\r\n* [10 years of Android](https://www.theverge.com/2011/12/7/2585779/android-10th-anniversary-google-history-pie-oreo-nougat-cupcake)\r\n\r\n\r\n## Security\r\n\r\n* [Forgetful Keystore](https://doridori.github.io/android-security-the-forgetful-keystore/)\r\n* [Pixels security, investments paying off](https://android-developers.googleblog.com/2018/01/android-security-ecosystem-investments.html)\r\n\r\n## Announcements\r\n\r\n* [Announcement 1, January 8th](announcement-1)\r\n* [Announcement 2, January 15th](announcement-2)\r\n* [Announcement 3, January 27th](announcement-3)\r\n\r\n\r\n\r\n## Reference group meetings\r\n\r\n## [[Course rules]]',
        2019, 'spring');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `fields`
--

CREATE TABLE `fields`
(
  `id`          int(11)     NOT NULL,
  `form_id`     int(11)     NOT NULL,
  `type`        varchar(64) NOT NULL,
  `name`        varchar(64) NOT NULL,
  `description` text        NOT NULL,
  `label`       varchar(64) DEFAULT NULL,
  `hasComment`  int(1)      NOT NULL,
  `priority`    int(11)     NOT NULL,
  `weight`      int(11)     DEFAULT NULL,
  `choices`     varchar(64) DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `fields`
--

INSERT INTO `fields` (`id`, `form_id`, `type`, `name`, `label`, `description`, `priority`, `weight`, `choices`)
VALUES (1, 1, 'text', 'github_form_text_0', 'Github handle', 'Username on github', 0, 0, ''),
       (2, 1, 'url', 'github_form_url_1', 'Github url', 'url to github', 1, 0, ''),
       (3, 1, 'textarea', 'github_form_textarea_2', 'Comments', 'Comments about your work', 2, 0, ''),
       (4, 2, 'radio', 'lab1_radio_0', 'Did the app have an MainActivity?', '', 0, 5, 'yes,no'),
       (5, 2, 'textarea', 'lab1_textarea_1', 'What could have been done better?', '', 1, 5, '');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `forms`
--

CREATE TABLE `forms`
(
  `id`          int(11)     NOT NULL,
  `prefix`      varchar(64) NOT NULL,
  `name`        varchar(64)      DEFAULT NULL,
  `created`     timestamp   NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `forms`
--


INSERT INTO `forms` (`id`, `prefix`, `name`, `created`)
VALUES (1, 'github_form', 'Github form', '2019-02-28 15:23:23'),
       (2, 'lab1', 'Lab1', '2019-03-12 11:00:43');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `logs`
--

CREATE TABLE `logs`
(
  `userid`       int(11)                            NOT NULL,
  `timestamp`    datetime                           NOT NULL,
  `activity`     varchar(32) COLLATE utf8_danish_ci NOT NULL,
  `assignmentid` int(11) DEFAULT NULL,
  `courseid`     int(11) DEFAULT NULL,
  `submissionid` int(11) DEFAULT NULL,
  `oldvalue`     text COLLATE utf8_danish_ci,
  `newValue`     text COLLATE utf8_danish_ci
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `peer_reviews`
--

CREATE TABLE `peer_reviews`
(
  `id`             int(11) NOT NULL,
  `submission_id`  int(11) NOT NULL,
  `assignment_id`  int(11) NOT NULL,
  `user_id`        int(11) NOT NULL,
  `review_user_id` int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `reviews`
--

CREATE TABLE `reviews`
(
  `id`      int(11) NOT NULL,
  `form_id` int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dumping data for table `reviews`
--

INSERT INTO `reviews` (`id`, `form_id`)
VALUES (1, 2);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `submissions`
--

CREATE TABLE `submissions`
(
  `id`      int(11) NOT NULL,
  `form_id` int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `submissions`
--

INSERT INTO `submissions` (`id`, `form_id`)
VALUES (1, 1);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `usercourse`
--

CREATE TABLE `usercourse`
(
  `userid`   int(11) NOT NULL,
  `courseid` int(11) NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `usercourse`
--

INSERT INTO `usercourse` (`userid`, `courseid`)
VALUES (1, 3),
       (2, 3),
       (3, 1),
       (3, 3),
       (4, 2),
       (4, 3),
       (5, 3),
       (6, 3),
       (7, 3),
       (8, 3),
       (9, 3),
       (10, 3);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `users`
--

CREATE TABLE `users`
(
  `id`            int(11)                            NOT NULL,
  `name`          varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `email_student` varchar(64) COLLATE utf8_danish_ci NOT NULL,
  `teacher`       tinyint(1)                         NOT NULL,
  `email_private` varchar(64) COLLATE utf8_danish_ci DEFAULT NULL,
  `password`      varchar(64) COLLATE utf8_danish_ci NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `users`
--

INSERT INTO `users` (`id`, `name`, `email_student`, `teacher`, `email_private`, `password`)
VALUES (1, 'Test User', 'hei@gmail.com', 1, 'test@yahoo.com',
        '$2a$14$MZj24p41j2NNGn6JDsQi0OsDb56.0LcfrIdgjE6WmZzp58O6V/VhK'),
       (2, 'Frode Haug', 'frodehg@teach.ntnu.no', 1, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (3, 'Ola Nordmann', 'olanor@stud.ntnu.no', 1, 'swag-meister69@ggmail.com',
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (4, 'Johan Klausen', 'johkl@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (5, 'Stian Fjerdingstad', 'stianfj@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (6, 'Svein Nilsen', 'sveini@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (7, 'Kjell Are-Kjelterud', 'kjellak@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (8, 'Marius Lillevik', 'mariuslil@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (9, 'Jorun Skaalnes', 'jorunska@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC'),
       (10, 'Klaus Aanesen', 'klausaa@stu.ntnu.no', 0, NULL,
        '$2a$14$vH/ibjwwXqBmOgJt8JCiK.S7D2r0VrBu46pYdCLs/dJMMk1aBV8RC');

-- --------------------------------------------------------

--
-- Table structure for table `user_reviews`
--

CREATE TABLE `user_reviews`
(
  `id`            int(11)     NOT NULL,
  `user_reviewer` int(11)     NOT NULL,
  `user_target`   int(11)     NOT NULL,
  `review_id`     int(11)     NOT NULL,
  `assignment_id` int(11)     NOT NULL,
  `type`          varchar(64) NOT NULL,
  `name`          varchar(64) NOT NULL,
  `label`         varchar(64) NOT NULL,
  `answer`        text        NOT NULL,
  `submitted`     datetime    NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dumping data for table `user_reviews`
--

INSERT INTO `user_reviews` (`id`, `user_reviewer`, `user_target`, `review_id`, `assignment_id`, `type`, `name`, `label`,
                            `answer`, `submitted`)
VALUES (3, 3, 4, 1, 1, 'text', 'test_text_0', 'A', 'good', '2019-03-08 01:08:43'),
       (4, 3, 4, 1, 1, 'text', 'test_text_1', 'B', 'bad', '2019-03-08 01:08:43');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `user_submissions`
--

CREATE TABLE `user_submissions`
(
  `id`            int(11)     NOT NULL,
  `user_id`       int(11)     NOT NULL,
  `assignment_id` int(11)     NOT NULL,
  `submission_id` int(11)     NOT NULL,
  `type`          varchar(64) NOT NULL,
  `answer`        text        NULL,
  `comment`       text        NULL,
  `submitted`     timestamp   NOT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Tabellstruktur for tabell `schedule_tasks`
--

CREATE TABLE `schedule_tasks`
(
  `id`             int(11)     NOT NULL,
  `submission_id`  int(11)     not null,
  `assignment_id`  int(11)     not null,
  `scheduled_time` datetime    not null,
  `task`           varchar(32) not null,
  `data`           blob        not null
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

--
-- Dataark for tabell `user_submissions`
--

INSERT INTO `user_submissions` (`id`, `user_id`, `assignment_id`, `submission_id`, `type`, `answer`, `submitted`)
VALUES (1, 4, 1, 1, 'text', 'JohanKlausen', '2019-03-13 10:40:26'),
       (2, 4, 1, 1, 'url', 'https://github.com/JohanKlausen/yeet', '2019-03-13 10:40:26'),
       (3, 4, 1, 1, 'textarea', 'I did good!', '2019-03-13 10:40:26'),
       (4, 5, 1, 1, 'text', 'StianFjerdingstad', '2019-03-01 23:59:59'),
       (5, 5, 1, 1, 'url', 'https://github.com/StianFjerdingstad/Sudoku', '2019-03-01 23:59:59'),
       (6, 5, 1, 1, 'textarea', 'I did sexy good!', '2019-03-01 23:59:59'),
       (7, 10, 1, 1, 'text', 'KlausAanesen', '2019-02-28 15:23:23'),
       (8, 10, 1, 1, 'url', 'https://github.com/KlausAanesen/1337yeet420', '2019-02-28 15:23:23'),
       (9, 10, 1, 1, 'textarea', 'I did bad :(', '2019-02-28 15:23:23');

INSERT INTO `peer_reviews` (`id`, `submission_id`, `assignment_id`, `user_id`, `review_user_id`)
VALUES (1, 1, 1, 3, 4),
       (2, 1, 1, 3, 5),
       (3, 1, 2, 9, 4);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `adminfaq`
--
ALTER TABLE `adminfaq`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `assignments`
--
ALTER TABLE `assignments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `assignments_courses_id_fk` (`course_id`),
  ADD KEY `assignments_reviews_id_fk` (`review_id`),
  ADD KEY `assignments_submissions_id_fk` (`submission_id`);

--
-- Indexes for table `course`
--
ALTER TABLE `course`
  ADD PRIMARY KEY (`id`),
  ADD KEY `teacher` (`teacher`);

--
-- Indexes for table `fields`
--
ALTER TABLE `fields`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fields_forms_id_fk` (`form_id`);

--
-- Indexes for table `forms`
--
ALTER TABLE `forms`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `logs`
--
ALTER TABLE `logs`
  ADD KEY `userid` (`userid`),
  ADD KEY `courseid` (`courseid`),
  ADD KEY `assignmentid` (`assignmentid`),
  ADD KEY `submissionid` (`submissionid`);


--
-- Indexes for table `reviews`
--
ALTER TABLE `reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `reviews_forms_id_fk` (`form_id`);

--
-- Indexes for table `submissions`
--
ALTER TABLE `submissions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `submissions_forms_id_fk` (`form_id`);

--
-- Indexes for table `usercourse`
--
ALTER TABLE `usercourse`
  ADD KEY `userid` (`userid`),
  ADD KEY `courseid` (`courseid`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email_student` (`email_student`),
  ADD UNIQUE KEY `email_private` (`email_private`);

--
-- Indexes for table `user_reviews`
--
ALTER TABLE `user_reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_reviews_reviews_id_fk` (`review_id`),
  ADD KEY `user_reviews_assignment_id_fk` (`assignment_id`),
  ADD KEY `user_reviews_user_reviewer_fk` (`user_reviewer`),
  ADD KEY `user_reviews_user_target_fk` (`user_target`);

--
-- Indexes for table `user_submissions`
--
ALTER TABLE `user_submissions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_submissions_submissions_id_fk` (`submission_id`),
  ADD KEY `user_submissions_users_id_fk` (`user_id`),
  ADD KEY `user_submission_assignment_id_fk` (`assignment_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `adminfaq`
--
ALTER TABLE `adminfaq`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 2;
--
-- AUTO_INCREMENT for table `assignments`
--
ALTER TABLE `assignments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 2;
--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 4;
--
-- AUTO_INCREMENT for table `fields`
--
ALTER TABLE `fields`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 4;
--
-- AUTO_INCREMENT for table `forms`
--
ALTER TABLE `forms`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 2;

--
-- AUTO_INCREMENT for table `reviews`
--
ALTER TABLE `reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `submissions`
--
ALTER TABLE `submissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 2;
--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 11;
--
-- AUTO_INCREMENT for table `user_reviews`
--
ALTER TABLE `user_reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `user_submissions`
--
ALTER TABLE `user_submissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 4;


ALTER TABLE `schedule_tasks`
  ADD PRIMARY KEY (`id`),
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- Begrensninger for dumpede tabeller
--

--
-- Indexes for table `peer_reviews`
--
ALTER TABLE `peer_reviews`
  ADD PRIMARY KEY (`id`),
  ADD KEY `peer_reviews_submissions_id_fk` (`submission_id`),
  ADD KEY `peer_reviews_assignment_id_fk` (`assignment_id`),
  ADD KEY `peer_reviews_user_id_fk` (`user_id`),
  ADD KEY `peer_reviews_review_user_id_fk` (`review_user_id`);

--
-- AUTO_INCREMENT for table `peer_reviews`
--
ALTER TABLE `peer_reviews`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;


--
-- Begrensninger for tabell `assignments`
--
ALTER TABLE `assignments`
  ADD CONSTRAINT `assignments_courses_id_fk` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`),
  ADD CONSTRAINT `assignments_reviews_id_fk` FOREIGN KEY (`review_id`) REFERENCES `reviews` (`id`),
  ADD CONSTRAINT `assignments_submissions_id_fk` FOREIGN KEY (`submission_id`) REFERENCES `submissions` (`id`);

--
-- Begrensninger for tabell `course`
--
ALTER TABLE `course`
  ADD CONSTRAINT `course_ibfk_1` FOREIGN KEY (`teacher`) REFERENCES `users` (`id`)
    ON UPDATE CASCADE;

--
-- Begrensninger for tabell `fields`
--
ALTER TABLE `fields`
  ADD CONSTRAINT `fields_forms_id_fk` FOREIGN KEY (`form_id`) REFERENCES `forms` (`id`);


--
-- Begrensninger for tabell `logs`
--
ALTER TABLE `logs`
  ADD CONSTRAINT `logs_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE CASCADE,
  ADD CONSTRAINT `logs_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`)
    ON DELETE NO ACTION
    ON UPDATE CASCADE;


--
-- Begrensninger for tabell `reviews`
--
ALTER TABLE `reviews`
  ADD CONSTRAINT `reviews_forms_id_fk` FOREIGN KEY (`form_id`) REFERENCES `forms` (`id`);


--
-- Begrensninger for tabell `submissions`
--
ALTER TABLE `submissions`
  ADD CONSTRAINT `submissions_forms_id_fk` FOREIGN KEY (`form_id`) REFERENCES `forms` (`id`);


--
-- Begrensninger for tabell `usercourse`
--
ALTER TABLE `usercourse`
  ADD CONSTRAINT `usercourse_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  ADD CONSTRAINT `usercourse_ibfk_2` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE;


--
-- Constraints for table `user_reviews`
--
ALTER TABLE `user_reviews`
  ADD CONSTRAINT `user_reviews_assignment_id_fk` FOREIGN KEY (`assignment_id`) REFERENCES `assignments` (`id`),
  ADD CONSTRAINT `user_reviews_reviews_id_fk` FOREIGN KEY (`review_id`) REFERENCES `reviews` (`id`),
  ADD CONSTRAINT `user_reviews_user_reviewer_fk` FOREIGN KEY (`user_reviewer`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `user_reviews_user_target_fk` FOREIGN KEY (`user_target`) REFERENCES `users` (`id`);

--
-- Begrensninger for tabell `user_submissions`
--
ALTER TABLE `user_submissions`
  ADD CONSTRAINT `user_submission_assignment_id_fk` FOREIGN KEY (`assignment_id`) REFERENCES `assignments` (`id`),
  ADD CONSTRAINT `user_submissions_submissions_id_fk` FOREIGN KEY (`submission_id`) REFERENCES `submissions` (`id`),
  ADD CONSTRAINT `user_submissions_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);


--
-- Begrensninger for tabell `schedule_tasks`
--
ALTER TABLE `schedule_tasks`
  ADD CONSTRAINT schedule_tasks_submission_id_fk FOREIGN KEY (submission_id) REFERENCES submissions (id),
  ADD CONSTRAINT schedule_tasks_assignment_id_fk FOREIGN KEY (assignment_id) REFERENCES assignments (id);

--
-- Begrensninger for tabell `peer_reviews`
--
ALTER TABLE `peer_reviews`
  ADD CONSTRAINT `peer_reviews_review_user_id_fk` FOREIGN KEY (`review_user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `peer_reviews_submissions_id_fk` FOREIGN KEY (`submission_id`) REFERENCES `submissions` (`id`),
  ADD CONSTRAINT `peer_reviews_assignment_id_fk` FOREIGN KEY (`assignment_id`) REFERENCES `assignments` (`id`),
  ADD CONSTRAINT `peer_reviews_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

