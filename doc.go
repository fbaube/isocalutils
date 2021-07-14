/*
package isocalutils works with ISO-8601 weeks.
There are two basic types used in the package.
The stdlib "time.Time" is referred to as "YMD".
OTOH "IsoYWD" refers to the type we prefer to
use. To prevent confusion, we redefine ints as
IsoYear, IsoMonth, and IsoDay.

About months being 0..11 or 1..12, ==TBS== 

https://en.wikipedia.org/wiki/ISO_week_date#Weeks_per_year

A "long year" of 53 weeks (17.75% of all cases)
is described by any of:
-  any year starting on Thursday
- leap year starting on Wensday
-  any year ending on Thursday
- leap year ending on Friday
- a year when 1 January and/or 31 December is a Thursday

Another way of stating the above rules is that in the ISO calendar,
- w01 is the earliest week with at least four days of January.
- the year's last week is the last with at least four days of December.
- w01 can begin as late as 4 January or as early as 29 December.
- the year's last week can end as early as 28 December and as late as 3 January.

A long ISO calendar year is always both preceded by and followed by
a short ISO calendar year.
*/
package isocalutils
