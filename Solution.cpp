
#include <cmath>
#include <algorithm>
using namespace std;

class Solution {

    static const int MINUTES_IN_HOUR = 60;
    static const int HOURS_FULL_CIRCLE = 12;
    static const int ANGLE_DEGREES_FULL_CIRCLE = 360;

    static const int ANGLE_DEGREES_PER_MINUTE = ANGLE_DEGREES_FULL_CIRCLE / MINUTES_IN_HOUR;
    static const int MINUTES_SEPARATING_HOURS_ON_THE_DIAL = MINUTES_IN_HOUR / HOURS_FULL_CIRCLE;

public:
    double angleClock(int hour, int minutes) {
        /*
        minutesOnTheDialForInputHour
        Note: not to be confused with minutes contained in the input hour
         */
        double minutesOnTheDialForInputHour
                = MINUTES_SEPARATING_HOURS_ON_THE_DIAL
                * ((static_cast<double>(minutes) / MINUTES_IN_HOUR) + (hour % HOURS_FULL_CIRCLE));

        double angleDegrees = ANGLE_DEGREES_PER_MINUTE * abs(minutes - minutesOnTheDialForInputHour);

        return min(ANGLE_DEGREES_FULL_CIRCLE - angleDegrees, angleDegrees);
    }
};
