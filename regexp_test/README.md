## regexp test

When i using `regexp` to match wildcards, something wrong that i found.

```
+-------------+-----------+--------+
| test string | wildcards | result |
+-------------+-----------+--------+
|   test.js   |    *.js   |  false |
+-------------+-----------+--------+
|   test.js   |     js    |  true  |
+-------------+-----------+--------+
```