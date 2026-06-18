
using System;

public class Solution
{
    private static readonly int MINUTES_IN_HOUR = 60;
    private static readonly int HOURS_FULL_CIRCLE = 12;
    private static readonly int ANGLE_DEGREES_FULL_CIRCLE = 360;

    private static readonly int ANGLE_DEGREES_PER_MINUTE = ANGLE_DEGREES_FULL_CIRCLE / MINUTES_IN_HOUR;
    private static readonly int MINUTES_SEPARATING_HOURS_ON_THE_DIAL = MINUTES_IN_HOUR / HOURS_FULL_CIRCLE;

    public double AngleClock(int hour, int minutes)
    {
        /*
        minutesOnTheDialForInputHour 
        Note: not to be confused with minutes contained in the input hour
         */
        double minutesOnTheDialForInputHour
                = MINUTES_SEPARATING_HOURS_ON_THE_DIAL
                * (((double)minutes / MINUTES_IN_HOUR) + (hour % HOURS_FULL_CIRCLE));

        double angleDegrees = ANGLE_DEGREES_PER_MINUTE * Math.Abs(minutes - minutesOnTheDialForInputHour);

        return Math.Min(ANGLE_DEGREES_FULL_CIRCLE - angleDegrees, angleDegrees);
    }
}
