
public class Solution {

    private static final int MINUTES_IN_HOUR = 60;
    private static final int HOURS_FULL_CIRCLE = 12;
    private static final int ANGLE_DEGREES_FULL_CIRCLE = 360;

    private static final int ANGLE_DEGREES_PER_MINUTE = ANGLE_DEGREES_FULL_CIRCLE / MINUTES_IN_HOUR;
    private static final int MINUTES_SEPARATING_HOURS_ON_THE_DIAL = MINUTES_IN_HOUR / HOURS_FULL_CIRCLE;

    public double angleClock(int hour, int minutes) {
        /*
        minutesOnTheDialForInputHour 
        Note: not to be confused with minutes contained in the input hour
         */
        double minutesOnTheDialForInputHour
                = MINUTES_SEPARATING_HOURS_ON_THE_DIAL
                * (((double) minutes / MINUTES_IN_HOUR) + (hour % HOURS_FULL_CIRCLE));

        double angleDegrees = ANGLE_DEGREES_PER_MINUTE * Math.abs(minutes - minutesOnTheDialForInputHour);

        return Math.min(ANGLE_DEGREES_FULL_CIRCLE - angleDegrees, angleDegrees);
    }
}
